package bearychat

import "time"

type RTMLoopState string

const (
	RTMLoopStateClosed RTMLoopState = "closed"
	RTMLoopStateOpen                = "open"
)

// RTMLoop is used to interactive with BearyChat's RTM websocket message protocol.
type RTMLoop interface {
	// Connect to RTM
	Start() error
	// Stop the connection
	Stop() error
	// Get current state
	State() RTMLoopState
	// Send a ping message
	Ping() error
	// Keep connection alive
	Keepalive(interval *time.Ticker) error
	// Send a message
	Send(m RTMMessage) error
	// Get message receiving channel
	ReadC() (chan *RTMMessage, error)
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
		if mtype, ok := t.(RTMMessageType); ok {
			return mtype
		}
	}

	return RTMMessageTypeUnknown
}
