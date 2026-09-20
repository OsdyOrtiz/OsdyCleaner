package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/core"
	"github.com/osdy/OsdyCleaner/internal/scan"
)

type fakeScan struct {
	snapshot core.Snapshot
	err      error
	calls    int
}

func (f *fakeScan) run(context.Context) (core.Snapshot, error) { f.calls++; return f.snapshot, f.err }

type terminalBuffer struct{ bytes.Buffer }

func TestCommandIdentityAndHelp(t *testing.T) {
	command := NewCommand(context.Background(), Dependencies{})
	if command.Use != "osdy-cleaner" {
		t.Fatalf("Use=%q", command.Use)
	}

	var help bytes.Buffer
	command.SetOut(&help)
	command.SetArgs([]string{"--help"})
	if err := command.Execute(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(help.String(), "osdy-cleaner") || strings.Contains(help.String(), "osdy [command]") {
		t.Fatalf("help=%q", help.String())
	}
}

func TestCommandDefaults(t *testing.T) {
	for _, tt := range []struct {
		name string
		tty  bool
		want string
	}{{"terminals use tui", true, "tui"}, {"pipes use text", false, "text"}} {
		t.Run(tt.name, func(t *testing.T) {
			_, out, _, formats, code := execute(t, completeSnapshot(t), []string{"scan"}, tt.tty)
			if code != ExitComplete || formats[0] != tt.want || out.String() != "" && tt.want == "tui" {
				t.Fatalf("code=%d formats=%v stdout=%q", code, formats, out.String())
			}
		})
	}
}

func TestCommandFormats(t *testing.T) {
	for _, format := range []string{"text", "json"} {
		t.Run(format, func(t *testing.T) {
			_, out, errOut, formats, code := execute(t, completeSnapshot(t), []string{"scan", "--format", format}, false)
			if code != ExitComplete || formats[0] != format || errOut.Len() != 0 {
				t.Fatalf("code=%d formats=%v stderr=%q", code, formats, errOut.String())
			}
			if format == "json" {
				var document map[string]bool
				if json.Unmarshal(out.Bytes(), &document) != nil || !document["scan"] || strings.Count(out.String(), "\n") != 1 {
					t.Fatalf("stdout is not one JSON document: %q", out.String())
				}
			} else if out.String() != "text\n" {
				t.Fatalf("text stdout=%q", out.String())
			}
		})
	}
}

func TestCommandStreams(t *testing.T) {
	for _, tt := range []struct{ format, want string }{{"text", "text\n"}, {"json", "{\"scan\":true}\n"}} {
		t.Run(tt.format, func(t *testing.T) {
			_, out, errOut, _, code := execute(t, partialSnapshot(t), []string{"scan", "--format", tt.format}, false)
			if code != ExitPartial || errOut.Len() != 0 || out.String() != tt.want {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out.String(), errOut.String())
			}
		})
	}
}

func TestCommandExitClasses(t *testing.T) {
	for _, tt := range []struct {
		name    string
		snap    core.Snapshot
		err     error
		want    int
		wantErr string
	}{{"complete", completeSnapshot(t), nil, ExitComplete, ""}, {"partial", partialSnapshot(t), nil, ExitPartial, ""}, {"cancelled", cancelledSnapshot(t), nil, ExitCancelled, ""}, {"unsupported", core.Snapshot{}, scan.ErrUnsupported, ExitUnsupported, "osdy-cleaner: scan platform is unsupported\n"}, {"runtime", core.Snapshot{}, errors.New("broken"), ExitFailure, "osdy-cleaner: broken\n"}} {
		t.Run(tt.name, func(t *testing.T) {
			f, out, errOut, _, code := execute(t, tt.snap, []string{"scan", "--format", "text"}, false, tt.err)
			if code != tt.want || f.calls != 1 || errOut.String() != tt.wantErr || (tt.err == nil && out.String() != "text\n") {
				t.Fatalf("code=%d calls=%d stdout=%q stderr=%q; want %d", code, f.calls, out.String(), errOut.String(), tt.want)
			}
		})
	}
}

func TestCommandTerminalValidation(t *testing.T) {
	f, _, _, _, code := execute(t, completeSnapshot(t), []string{"scan", "--format", "tui"}, false)
	if code != ExitInput || f.calls != 0 {
		t.Fatalf("code=%d calls=%d", code, f.calls)
	}
}

func TestCommandFailures(t *testing.T) {
	for _, tt := range []struct {
		name   string
		args   []string
		render error
		write  bool
		want   int
	}{{"invalid format", []string{"scan", "--format", "xml"}, nil, false, ExitInput}, {"render", []string{"scan", "--format", "text"}, errors.New("render"), false, ExitFailure}, {"write", []string{"scan", "--format", "text"}, nil, true, ExitFailure}} {
		t.Run(tt.name, func(t *testing.T) {
			f, out, _, _, code := executeWith(t, completeSnapshot(t), tt.args, false, nil, tt.render, tt.write)
			if code != tt.want || (tt.want == ExitInput && f.calls != 0) || (tt.render != nil && out.Len() != 0) {
				t.Fatalf("code=%d calls=%d stdout=%q", code, f.calls, out.String())
			}
		})
	}
}

func TestCommandViewerFailure(t *testing.T) {
	f, out, errOut, formats, code := executeCustom(t, completeSnapshot(t), []string{"scan", "--format", "tui"}, true, nil, nil, nil, errors.New("viewer"), nil)
	if code != ExitFailure || f.calls != 1 || out.Len() != 0 || errOut.String() != "osdy-cleaner: viewer\n" || strings.Join(formats, ",") != "tui" {
		t.Fatalf("code=%d calls=%d stdout=%q stderr=%q formats=%v", code, f.calls, out.String(), errOut.String(), formats)
	}
}

func TestCommandCancelledStreams(t *testing.T) {
	for _, tt := range []struct {
		name, format, want string
	}{{"text", "text", "text\n"}, {"json", "json", "{\"scan\":true}\n"}} {
		t.Run(tt.name, func(t *testing.T) {
			_, out, errOut, _, code := executeCustom(t, cancelledSnapshot(t), []string{"scan", "--format", tt.format}, false, nil, nil, nil, nil, nil)
			if code != ExitCancelled || out.String() != tt.want || errOut.Len() != 0 {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out.String(), errOut.String())
			}
			if tt.format == "json" {
				var document map[string]bool
				decoder := json.NewDecoder(out)
				if err := decoder.Decode(&document); err != nil || !document["scan"] || decoder.Decode(&struct{}{}) != io.EOF {
					t.Fatalf("stdout is not exactly one JSON document: %q", out.String())
				}
			}
		})
	}
}

func TestCommandCancellationPrecedence(t *testing.T) {
	snapshot := cancelledSnapshot(t)
	if snapshot.Outcome() != core.OutcomeCancelled {
		t.Fatalf("outcome=%q", snapshot.Outcome())
	}
	_, out, errOut, _, code := executeCustom(t, snapshot, []string{"scan", "--format", "text"}, false, nil, nil, nil, nil, nil)
	if code != ExitCancelled || out.String() != "text\n" || errOut.Len() != 0 {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, out.String(), errOut.String())
	}
}

func TestCommandFailureDiagnostics(t *testing.T) {
	for _, tt := range []struct {
		name                               string
		args                               []string
		tty                                bool
		scanErr, textErr, jsonErr, viewErr error
		writer                             io.Writer
		wantCode                           int
		wantOut, wantErr                   string
		wantCalls                          int
	}{
		{"scan", []string{"scan", "--format", "text"}, false, errors.New("scan"), nil, nil, nil, nil, ExitFailure, "", "osdy-cleaner: scan\n", 1},
		{"text render", []string{"scan", "--format", "text"}, false, nil, errors.New("render"), nil, nil, nil, ExitFailure, "", "osdy-cleaner: render\n", 1},
		{"JSON render", []string{"scan", "--format", "json"}, false, nil, nil, errors.New("render"), nil, nil, ExitFailure, "", "osdy-cleaner: render\n", 1},
		{"viewer", []string{"scan", "--format", "tui"}, true, nil, nil, nil, errors.New("viewer"), nil, ExitFailure, "", "osdy-cleaner: viewer\n", 1},
		{"write partial", []string{"scan", "--format", "text"}, false, nil, nil, nil, nil, &partialWriter{n: 3, err: errors.New("write")}, ExitFailure, "tex", "osdy-cleaner: write\n", 1},
		{"invalid input", []string{"scan", "--format", "xml"}, false, nil, nil, nil, nil, nil, ExitInput, "", "osdy-cleaner: unsupported format \"xml\"\n", 0},
		{"unsupported", []string{"scan", "--format", "text"}, false, scan.ErrUnsupported, nil, nil, nil, nil, ExitUnsupported, "", "osdy-cleaner: scan platform is unsupported\n", 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			f, out, errOut, _, code := executeCustom(t, completeSnapshot(t), tt.args, tt.tty, tt.scanErr, tt.textErr, tt.jsonErr, tt.viewErr, tt.writer)
			gotOut := out.String()
			if writer, ok := tt.writer.(*partialWriter); ok {
				gotOut = writer.String()
			}
			if code != tt.wantCode || f.calls != tt.wantCalls || gotOut != tt.wantOut || errOut.String() != tt.wantErr {
				t.Fatalf("code=%d calls=%d stdout=%q stderr=%q", code, f.calls, gotOut, errOut.String())
			}
		})
	}
}

func TestCommandNoCleanupPath(t *testing.T) {
	for _, args := range [][]string{{"scan", "/tmp/x"}, {"scan", "--path", "/tmp/x"}, {"scan", "--cleanup", "/tmp/x"}} {
		f, _, _, _, code := execute(t, completeSnapshot(t), args, false)
		if code != ExitInput || f.calls != 0 {
			t.Fatalf("args=%v code=%d calls=%d", args, code, f.calls)
		}
	}
}

func TestRunDisposableFixtureJSON(t *testing.T) {
	_, out, errOut, _, code := execute(t, completeSnapshot(t), []string{"scan", "--format", "json"}, false)
	var document map[string]bool
	if code != ExitComplete || errOut.Len() != 0 || json.Unmarshal(out.Bytes(), &document) != nil || !document["scan"] {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, out.String(), errOut.String())
	}
}

func execute(t *testing.T, snapshot core.Snapshot, args []string, tty bool, scanErr ...error) (*fakeScan, *bytes.Buffer, *bytes.Buffer, []string, int) {
	var err error
	if len(scanErr) != 0 {
		err = scanErr[0]
	}
	return executeWith(t, snapshot, args, tty, err, nil, false)
}
func executeWith(t *testing.T, snapshot core.Snapshot, args []string, tty bool, scanErr, renderErr error, failWrite bool) (*fakeScan, *bytes.Buffer, *bytes.Buffer, []string, int) {
	t.Helper()
	f, out, errOut := &fakeScan{snapshot: snapshot, err: scanErr}, &bytes.Buffer{}, &bytes.Buffer{}
	var formats []string
	writer := io.Writer(out)
	if failWrite {
		writer = failingWriter{}
	}
	code := Execute(context.Background(), Dependencies{Scan: f.run, RenderText: renderer("text\n", renderErr, &formats, "text"), RenderJSON: renderer("{\"scan\":true}\n", renderErr, &formats, "json"), View: func(core.Snapshot) error { formats = append(formats, "tui"); return nil }, Input: strings.NewReader(""), Output: writer, Error: errOut, IsTerminal: func(io.Reader, io.Writer) bool { return tty }}, args)
	return f, out, errOut, formats, code
}
func renderer(value string, err error, formats *[]string, format string) func(core.Snapshot) ([]byte, error) {
	return func(core.Snapshot) ([]byte, error) { *formats = append(*formats, format); return []byte(value), err }
}

func executeCustom(t *testing.T, snapshot core.Snapshot, args []string, tty bool, scanErr, textErr, jsonErr, viewErr error, writer io.Writer) (*fakeScan, *bytes.Buffer, *bytes.Buffer, []string, int) {
	t.Helper()
	f, out, errOut := &fakeScan{snapshot: snapshot, err: scanErr}, &bytes.Buffer{}, &bytes.Buffer{}
	var formats []string
	if writer == nil {
		writer = out
	}
	code := Execute(context.Background(), Dependencies{Scan: f.run, RenderText: renderer("text\n", textErr, &formats, "text"), RenderJSON: renderer("{\"scan\":true}\n", jsonErr, &formats, "json"), View: func(core.Snapshot) error { formats = append(formats, "tui"); return viewErr }, Input: strings.NewReader(""), Output: writer, Error: errOut, IsTerminal: func(io.Reader, io.Writer) bool { return tty }}, args)
	return f, out, errOut, formats, code
}

type partialWriter struct {
	bytes.Buffer
	n   int
	err error
}

func (w *partialWriter) Write(p []byte) (int, error) {
	if w.n > len(p) {
		w.n = len(p)
	}
	_, _ = w.Buffer.Write(p[:w.n])
	return w.n, w.err
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errors.New("write") }

func completeSnapshot(t *testing.T) core.Snapshot  { return snapshot(t, core.RootScanned) }
func partialSnapshot(t *testing.T) core.Snapshot   { return snapshot(t, core.RootPartial) }
func cancelledSnapshot(t *testing.T) core.Snapshot { return snapshot(t, core.RootCancelled) }
func snapshot(t *testing.T, state core.RootStatus) core.Snapshot {
	t.Helper()
	policy, _ := core.NewScanPolicy(true, true, 1, 1)
	areas := []core.AreaID{core.AreaNPMCache, core.AreaHomebrewCache, core.AreaGradleCaches, core.AreaXcodeDerivedData, core.AreaCoreSimulator}
	names := []string{"npm cache", "Homebrew cache", "Gradle caches", "Xcode DerivedData", "CoreSimulator"}
	paths := []string{"~/.npm", "~/Library/Caches/Homebrew", "~/.gradle/caches", "~/Library/Developer/Xcode/DerivedData", "~/Library/Developer/CoreSimulator"}
	roots := make([]core.RootObservation, 0, 5)
	for i := range areas {
		status, reason, completeness := core.RootScanned, core.ReasonCompleted, core.CompletenessComplete
		if i == 0 && state == core.RootPartial {
			status, reason, completeness = core.RootPartial, core.ReasonEntryVisibilityGap, core.CompletenessIncomplete
		}
		if state == core.RootCancelled && i == 0 {
			status, reason, completeness = core.RootPartial, core.ReasonEntryVisibilityGap, core.CompletenessIncomplete
		}
		if state == core.RootCancelled && i == 1 {
			status, reason, completeness = core.RootCancelled, core.ReasonCancelledDuringScan, core.CompletenessIncomplete
		}
		logical, _ := core.NewValueEstimate(0, core.BasisRegularFileStatSize, completeness)
		allocated, _ := core.NewValueEstimate(0, core.BasisStatBlocks512, completeness)
		estimate, _ := core.NewEstimate(logical, allocated, true)
		root, err := core.NewRootObservation(areas[i], names[i], paths[i], status, reason, estimate, 0, nil)
		if err != nil {
			t.Fatal(err)
		}
		roots = append(roots, root)
	}
	snapshot, err := core.NewSnapshot("test", policy, roots, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	return snapshot
}
