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

type RTMMessageType string

const (
	RTMMessageTypeUnknown              RTMMessageType = "unknown"
	RTMMessageTypePing                                = "ping"
	RTMMessageTypePong                                = "pong"
	RTMMessageTypeReply                               = "reply"
	RTMMessageTypeOk                                  = "ok"
	RTMMessageTypeP2PMessage                          = "message"
	RTMMessageTypeP2PTyping                           = "typing"
	RTMMessageTypeChannelMessage                      = "channel_message"
	RTMMessageTypeChannelTyping                       = "channel_typing"
	RTMMessageTypeUpdateUserConnection                = "update_user_connection"
)

// RTMMessage represents a message entity send over RTM protocol.
type RTMMessage map[string]interface{}

func (m RTMMessage) Type() RTMMessageType {
	if t, present := m["type"]; present {
		if mtype, ok := t.(string); ok {
			return RTMMessageType(mtype)
		}
	}

	return RTMMessageTypeUnknown
}
