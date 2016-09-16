package bearychat

import (
	"errors"
	"time"
)

type RTMLoopState string

const (
	RTMLoopStateClosed RTMLoopState = "closed"
	RTMLoopStateOpen                = "open"
)

var (
	ErrRTMLoopClosed = errors.New("rtm loop is closed")
)

// RTMLoop is used to interactive with BearyChat's RTM websocket message protocol.
type RTMLoop interface {
	// Connect to RTM, returns after connected
	Start() error
	// Stop the connection
	Stop() error
	// Get current state
	State() RTMLoopState
	// Send a ping message
	Ping() error
	// Keep connection alive. Closes ticker before return
	Keepalive(interval *time.Ticker) error
	// Send a message
	Send(m RTMMessage) error
	// Get message receiving channel
	ReadC() (chan RTMMessage, error)
	// Get error channel
	ErrC() chan error
}
