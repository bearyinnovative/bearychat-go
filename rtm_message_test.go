package bearychat

import (
	"fmt"
	"testing"
)

func TestRTMMessage_Type(t *testing.T) {
	cases := [][]RTMMessageType{
		{RTMMessageTypeUnknown, RTMMessageTypeUnknown},
		{RTMMessageTypePing, RTMMessageTypePing},
		{RTMMessageTypePong, RTMMessageTypePong},
		{RTMMessageTypeReply, RTMMessageTypeReply},
		{RTMMessageTypeOk, RTMMessageTypeOk},
		{RTMMessageTypeP2PMessage, RTMMessageTypeP2PMessage},
		{RTMMessageTypeP2PTyping, RTMMessageTypeP2PTyping},
		{RTMMessageTypeChannelMessage, RTMMessageTypeChannelMessage},
		{RTMMessageTypeChannelTyping, RTMMessageTypeChannelTyping},
		{RTMMessageTypeUpdateUserConnection, RTMMessageTypeUpdateUserConnection},
	}

	for _, c := range cases {
		m := RTMMessage{"type": c[0]}
		if m.Type() != c[1] {
			t.Errorf("expected type: %s, got: %s", c[0], m.Type())
		}
	}
}

func TestRTMMessage_IsP2P(t *testing.T) {
	cases := []struct {
		mt       RTMMessageType
		expected bool
	}{
		{RTMMessageTypeUnknown, false},
		{RTMMessageTypePing, false},
		{RTMMessageTypePong, false},
		{RTMMessageTypeReply, false},
		{RTMMessageTypeOk, false},
		{RTMMessageTypeP2PMessage, true},
		{RTMMessageTypeP2PTyping, true},
		{RTMMessageTypeChannelMessage, false},
		{RTMMessageTypeChannelTyping, false},
		{RTMMessageTypeUpdateUserConnection, false},
	}

	for _, c := range cases {
		m := RTMMessage{"type": c.mt}
		if m.IsP2P() != c.expected {
			t.Errorf("expected: %+v, got: %+v", c.expected, m.IsP2P())
		}
	}
}

func TestRTMMessage_IsChatMessage(t *testing.T) {
	cases := []struct {
		mt       RTMMessageType
		expected bool
	}{
		{RTMMessageTypeUnknown, false},
		{RTMMessageTypePing, false},
		{RTMMessageTypePong, false},
		{RTMMessageTypeReply, false},
		{RTMMessageTypeOk, false},
		{RTMMessageTypeP2PMessage, true},
		{RTMMessageTypeP2PTyping, false},
		{RTMMessageTypeChannelMessage, true},
		{RTMMessageTypeChannelTyping, false},
		{RTMMessageTypeUpdateUserConnection, false},
	}

	for _, c := range cases {
		m := RTMMessage{"type": c.mt}
		if m.IsChatMessage() != c.expected {
			t.Errorf("expected: %+v, got: %+v", c.expected, m.IsChatMessage())
		}
	}
}

func TestRTMMessage_IsFrom(t *testing.T) {
	uid := "1"
	user := User{Id: uid}
	var m RTMMessage

	m = RTMMessage{"uid": uid}
	if !m.IsFromUser(user) {
		t.Errorf("expected from user: %+v", m)
	}
	if !m.IsFromUID(uid) {
		t.Errorf("expected from uid: %+v", m)
	}

	m = RTMMessage{"uid": uid + "1"}
	if m.IsFromUser(user) {
		t.Errorf("unexpected from user: %+v", m)
	}
	if m.IsFromUID(uid) {
		t.Errorf("expected from uid: %+v", m)
	}
}

func TestRTMMessage_Refer_ChannelMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeChannelMessage,
		"channel_id":  "foobar",
		"vchannel_id": "foobar",
		"key":         "foobar",
	}

	referText := "foobar"
	refer := m.Refer(referText)
	if refer["text"] != referText {
		t.Errorf("unexpected %s", refer["text"])
	}
	if refer.Type() != RTMMessageTypeChannelMessage {
		t.Errorf("unexpected %s", refer.Type())
	}
	if refer["channel_id"] != m["channel_id"] {
		t.Errorf("unexpected %s", refer["channel_id"])
	}
	if refer["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", refer["vchannel_id"])
	}
	if refer["refer_key"] != m["key"] {
		t.Errorf("unexpected %s", refer["refer_key"])
	}
}

func TestRTMMessage_Refer_P2PMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeP2PMessage,
		"uid":         "foobar",
		"vchannel_id": "foobar",
		"key":         "foobar",
	}

	referText := "foobar"
	refer := m.Refer(referText)
	if refer["text"] != referText {
		t.Errorf("unexpected %s", refer["text"])
	}
	if refer.Type() != RTMMessageTypeP2PMessage {
		t.Errorf("unexpected %s", refer.Type())
	}
	if refer["to_uid"] != m["uid"] {
		t.Errorf("unexpected %s", refer["to_uid"])
	}
	if refer["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", refer["vchannel_id"])
	}
	if refer["refer_key"] != m["key"] {
		t.Errorf("unexpected %s", refer["refer_key"])
	}
}

func TestRTMMessage_Reply_ChannelMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeChannelMessage,
		"channel_id":  "foobar",
		"vchannel_id": "foobar",
	}

	replyText := "foobar"
	reply := m.Reply(replyText)
	if reply["text"] != replyText {
		t.Errorf("unexpected %s", reply["text"])
	}
	if reply.Type() != RTMMessageTypeChannelMessage {
		t.Errorf("unexpected %s", reply.Type())
	}
	if reply["channel_id"] != m["channel_id"] {
		t.Errorf("unexpected %s", reply["channel_id"])
	}
	if reply["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", reply["vchannel_id"])
	}
}

func TestRTMMessage_Reply_P2PMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeChannelMessage,
		"channel_id":  "foobar",
		"vchannel_id": "foobar",
	}

	replyText := "foobar"
	reply := m.Reply(replyText)
	if reply["text"] != replyText {
		t.Errorf("unexpected %s", reply["text"])
	}
	if reply.Type() != RTMMessageTypeChannelMessage {
		t.Errorf("unexpected %s", reply.Type())
	}
	if reply["to_uid"] != m["uid"] {
		t.Errorf("unexpected %s", reply["to_uid"])
	}
	if reply["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", reply["vchannel_id"])
	}
}

func TestRTMMessage_ParseMention(t *testing.T) {
	uid := "=1"
	text := "abc"
	user := User{Id: uid}

	m := RTMMessage{}
	var mentioned bool
	var content string

	expect := func(expectMentioned bool, expectContent string) {
		if mentioned != expectMentioned {
			t.Errorf("expected mentioned: '%v', got '%v', m: %+v", expectMentioned, mentioned, m)
		}
		if content != expectContent {
			t.Errorf("expected content: '%v', got '%v'", expectContent, content)
		}
	}

	m["text"] = text
	m["type"] = RTMMessageTypeP2PMessage
	mentioned, content = m.ParseMentionUser(user)
	expect(true, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, text)

	m["type"] = RTMMessageTypeChannelMessage
	mentioned, content = m.ParseMentionUser(user)
	expect(false, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(false, text)

	m["text"] = fmt.Sprintf("@<=%s=> %s", uid, text)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, text)

	m["text"] = fmt.Sprintf("123123123 12312 123@<=%s=> %s", uid, text)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, text)

	m["text"] = fmt.Sprintf("@<=%s=>", uid)
	mentioned, content = m.ParseMentionUser(user)
	expect(false, m.Text())
	mentioned, content = m.ParseMentionUID(uid)
	expect(false, m.Text())

	m["text"] = fmt.Sprintf("@<=%s=> ", uid)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, "")
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, "")

	m["text"] = fmt.Sprintf("@<=%s=> 你和 @<==bwOwr=> 谁聪明", uid)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, "你和 @<==bwOwr=> 谁聪明")
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, "你和 @<==bwOwr=> 谁聪明")
}
