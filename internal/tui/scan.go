package tui

import (
	"context"
	"sync"
	"sync/atomic"

	tea "charm.land/bubbletea/v2"

	"github.com/osdy/OsdyCleaner/internal/core"
	"github.com/osdy/OsdyCleaner/internal/scan"
)

// ScanFunc performs one read-only scan and reports its transient progress.
type ScanFunc func(context.Context, scan.ProgressObserver) (core.Snapshot, error)

// V1 emits at most twelve lifecycle events, followed by one result.
const scanMessageCapacity = 13

type scanSession struct {
	ctx      context.Context
	cancel   context.CancelFunc
	scan     ScanFunc
	messages chan tea.Msg
	done     chan struct{}
	once     sync.Once
	began    atomic.Bool

	snapshot core.Snapshot
	err      error
}

type scanStartedMsg struct{}
type scanCancelMsg struct{}
type scanProgressMsg struct {
	processed, total int
	category         string
}
type scanResultMsg struct {
	snapshot core.Snapshot
	err      error
}

func newScanSession(parent context.Context, fn ScanFunc) *scanSession {
	ctx, cancel := context.WithCancel(parent)
	return &scanSession{ctx: ctx, cancel: cancel, scan: fn, messages: make(chan tea.Msg, scanMessageCapacity), done: make(chan struct{})}
}

func newScanModel(parent context.Context, fn ScanFunc) Model {
	m := NewModel(core.Snapshot{})
	m.session, m.running = newScanSession(parent, fn), true
	return m
}

func (s *scanSession) start() tea.Cmd {
	return func() tea.Msg {
		s.once.Do(func() {
			s.began.Store(true)
			go func() {
				defer close(s.done)
				s.snapshot, s.err = s.scan(s.ctx, func(event scan.ProgressEvent) {
					msg := scanProgressMsg{processed: event.Processed(), total: event.Total(), category: event.CategoryDisplayName()}
					select {
					case s.messages <- msg:
					case <-s.ctx.Done():
					}
				})
				s.messages <- scanResultMsg{snapshot: s.snapshot, err: s.err}
			}()
		})
		return scanStartedMsg{}
	}
}

func (s *scanSession) wait() tea.Cmd {
	return func() tea.Msg { return <-s.messages }
}

func (s *scanSession) requestCancel() tea.Cmd {
	return func() tea.Msg {
		s.cancel()
		return scanCancelMsg{}
	}
}

func (s *scanSession) waitShutdown() (core.Snapshot, error) {
	if !s.began.Load() {
		return core.Snapshot{}, nil
	}
	<-s.done
	return s.snapshot, s.err
}
