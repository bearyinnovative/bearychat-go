package bearychat

import (
	"sync"
	"testing"
)

const (
	testRTMWSHost = "foobar"
)

func TestNewRTMLoop(t *testing.T) {
	l, err := NewRTMLoop(testRTMWSHost)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}
	if l.wsHost != testRTMWSHost {
		t.Errorf("unexpected wsHost: %s", l.wsHost)
	}
	if l.callId != 0 {
		t.Errorf("unexpected call id: %d", l.callId)
	}
}

func TestRTMLoop_ReadC_Closed(t *testing.T) {
	l, err := NewRTMLoop(testRTMWSHost)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if _, err := l.ReadC(); err != ErrRTMLoopClosed {
		t.Errorf("unexpected error: %+v", err)
	}
}

func TestRTMLoop_ErrC_Closed(t *testing.T) {
	l, err := NewRTMLoop(testRTMWSHost)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if c := l.ErrC(); c == nil {
		t.Errorf("expected error channel")
	}
}

func TestRTMLoop_advanceCallId(t *testing.T) {
	l, err := NewRTMLoop(testRTMWSHost)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	for i := uint64(1); i < 11; i = i + 1 {
		newCallId := l.advanceCallId()
		if l.callId != i || newCallId != l.callId {
			t.Errorf(
				"unexpected callId: %d, %d, %d",
				i,
				newCallId,
				l.callId,
			)
		}
	}
}

func TestRTMLoop_advanceCallId_Race(t *testing.T) {
	l, err := NewRTMLoop(testRTMWSHost)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	per := 15
	times := 10

	var wg sync.WaitGroup
	advance := func() {
		for i := 0; i < per; i = i + 1 {
			l.advanceCallId()
		}
		wg.Done()
	}

	for i := 0; i < times; i = i + 1 {
		wg.Add(1)
		go advance()
	}

	wg.Wait()
	if l.callId != uint64(per*times) {
		t.Errorf("unexepcted call id after data race: %d", l.callId)
	}
}
