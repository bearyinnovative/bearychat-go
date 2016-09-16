package bearychat

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
		if mtype, ok := t.(RTMMessageType); ok {
			return mtype
		}
	}

	return RTMMessageTypeUnknown
}

// Reply a message (with copying type, vchannel_id)
func (m RTMMessage) Reply(text string) RTMMessage {
	reply := RTMMessage{
		"text":        text,
		"vchannel_id": m["vchannel_id"],
	}

	if m.IsP2P() {
		reply["type"] = RTMMessageTypeP2PMessage
		reply["to_uid"] = m["uid"]
	} else {
		reply["type"] = RTMMessageTypeChannelMessage
		reply["channel_id"] = m["channel_id"]
	}

	return reply
}

// Refer a message
func (m RTMMessage) Refer(text string) RTMMessage {
	refer := m.Reply(text)
	refer["refer_key"] = m["key"]

	return refer
}

func (m RTMMessage) IsP2P() bool {
	mt := m.Type()
	if mt == RTMMessageTypeP2PMessage || mt == RTMMessageTypeP2PTyping {
		return true
	}

	return false
}

func (m RTMMessage) IsChatMessage() bool {
	mt := m.Type()
	if mt == RTMMessageTypeP2PMessage || mt == RTMMessageTypeChannelMessage {
		return true
	}

	return false
}

func (m RTMMessage) IsFromMe(u User) bool {
	return m["uid"] == u.Id
}
