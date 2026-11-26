package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/core"
)

func TestCommandScanCancellationErrors(t *testing.T) {
	for _, format := range []string{"text", "json", "tui"} {
		t.Run(format, func(t *testing.T) {
			f, out, errOut, formats, code := execute(t, core.Snapshot{}, []string{"scan", "--format", format}, format == "tui", context.Canceled)
			wantFormats := ""
			if format == "tui" {
				wantFormats = "tui"
			}
			if code != ExitCancelled || f.calls != 1 || out.Len() != 0 || errOut.Len() != 0 || strings.Join(formats, ",") != wantFormats {
				t.Fatalf("code=%d calls=%d stdout=%q stderr=%q formats=%v", code, f.calls, out.String(), errOut.String(), formats)
			}
		})
	}
}

func TestCommandShortWrites(t *testing.T) {
	for _, format := range []struct{ name, payload string }{{"text", "text\n"}, {"json", "{\"scan\":true}\n"}} {
		t.Run(format.name, func(t *testing.T) {
			writer := &countedPartialWriter{partialWriter: &partialWriter{n: 3}}
			f, _, errOut, _, code := executeCustom(t, completeSnapshot(t), []string{"scan", "--format", format.name}, false, nil, nil, nil, nil, writer)
			if code != ExitFailure || f.calls != 1 || writer.calls != 1 || writer.String() != format.payload[:3] || errOut.String() != "osdy-cleaner: "+io.ErrShortWrite.Error()+"\n" {
				t.Fatalf("code=%d scan_calls=%d write_calls=%d stdout=%q stderr=%q", code, f.calls, writer.calls, writer.String(), errOut.String())
			}
		})
	}
}

func TestCommandScanErrorClassification(t *testing.T) {
	for _, format := range []string{"text", "json", "tui"} {
		for _, tt := range []struct {
			name string
			err  error
			want int
		}{
			{"wrapped cancellation", fmt.Errorf("scan stopped: %w", context.Canceled), ExitCancelled},
			{"deadline", context.DeadlineExceeded, ExitFailure},
			{"wrapped deadline", fmt.Errorf("scan stopped: %w", context.DeadlineExceeded), ExitFailure},
		} {
			t.Run(format+"/"+tt.name, func(t *testing.T) {
				f, out, errOut, formats, code := execute(t, core.Snapshot{}, []string{"scan", "--format", format}, format == "tui", tt.err)
				wantErr, wantFormats := "", ""
				if tt.want == ExitFailure {
					wantErr = "osdy-cleaner: " + tt.err.Error() + "\n"
				}
				if format == "tui" {
					wantFormats = "tui"
				}
				if code != tt.want || f.calls != 1 || out.Len() != 0 || errOut.String() != wantErr || strings.Join(formats, ",") != wantFormats {
					t.Fatalf("code=%d calls=%d stdout=%q stderr=%q formats=%v", code, f.calls, out.String(), errOut.String(), formats)
				}
			})
		}
	}
}

func TestCommandWriteResultPrecedence(t *testing.T) {
	writeErr := errors.New("device write failed")
	for _, format := range []struct{ name, payload string }{{"text", "text\n"}, {"json", "{\"scan\":true}\n"}} {
		for _, tt := range []struct {
			name    string
			snap    core.Snapshot
			n       int
			err     error
			want    int
			wantErr error
		}{
			{"zero", completeSnapshot(t), 0, nil, ExitFailure, io.ErrShortWrite},
			{"full", completeSnapshot(t), len(format.payload), nil, ExitComplete, nil},
			{"short with error", completeSnapshot(t), 3, writeErr, ExitFailure, writeErr},
			{"full with error", completeSnapshot(t), len(format.payload), writeErr, ExitFailure, writeErr},
			{"writer cancellation", completeSnapshot(t), 0, context.Canceled, ExitFailure, context.Canceled},
			{"partial snapshot", partialSnapshot(t), 3, nil, ExitFailure, io.ErrShortWrite},
			{"cancelled snapshot", cancelledSnapshot(t), 3, nil, ExitFailure, io.ErrShortWrite},
		} {
			t.Run(format.name+"/"+tt.name, func(t *testing.T) {
				writer := &countedPartialWriter{partialWriter: &partialWriter{n: tt.n, err: tt.err}}
				f, _, errOut, _, code := executeCustom(t, tt.snap, []string{"scan", "--format", format.name}, false, nil, nil, nil, nil, writer)
				wantErr := ""
				if tt.wantErr != nil {
					wantErr = "osdy-cleaner: " + tt.wantErr.Error() + "\n"
				}
				if code != tt.want || f.calls != 1 || writer.calls != 1 || writer.String() != format.payload[:tt.n] || errOut.String() != wantErr {
					t.Fatalf("code=%d scan_calls=%d write_calls=%d stdout=%q stderr=%q", code, f.calls, writer.calls, writer.String(), errOut.String())
				}
			})
		}
	}
}

func TestCommandRenderErrorClassification(t *testing.T) {
	for _, format := range []string{"text", "json"} {
		for _, renderErr := range []error{context.Canceled, context.DeadlineExceeded} {
			t.Run(format+"/"+renderErr.Error(), func(t *testing.T) {
				f, out, errOut, formats, code := executeCustom(t, completeSnapshot(t), []string{"scan", "--format", format}, false, nil, renderErr, renderErr, nil, nil)
				if code != ExitFailure || f.calls != 1 || out.Len() != 0 || errOut.String() != "osdy-cleaner: "+renderErr.Error()+"\n" || strings.Join(formats, ",") != format {
					t.Fatalf("code=%d calls=%d stdout=%q stderr=%q formats=%v", code, f.calls, out.String(), errOut.String(), formats)
				}
			})
		}
	}
}

type countedPartialWriter struct {
	*partialWriter
	calls int
}

func (w *countedPartialWriter) Write(p []byte) (int, error) {
	w.calls++
	return w.partialWriter.Write(p)
}
