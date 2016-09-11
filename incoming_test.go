package bearychat

import "testing"

func TestValidateIncoming(t *testing.T) {
	var m Incoming

	m = Incoming{}
	if err := m.Validate(); err == nil {
		t.Errorf("text should not be empty: %+v", err)
	}

	m = Incoming{
		Text:        "test",
		Attachments: []IncomingAttachment{{}},
	}
	if err := m.Validate(); err == nil {
		t.Errorf("title or text should not be empty: %+v", err)
	}

	m = Incoming{
		Text: "test",
		Attachments: []IncomingAttachment{
			{
				Text:   "test",
				Images: []IncomingAttachmentImage{{}},
			},
		},
	}
	if err := m.Validate(); err == nil {
		t.Errorf("image url should not be empty: %+v", err)
	}
}

func ExampleIncoming() {
	m := Incoming{
		Text:         "Hello, **BearyChat",
		Notification: "Hello, BearyChat in notification",
		Markdown:     true,
		Channel:      "#所有人",
		User:         "@bearybot",
		Attachments: []IncomingAttachment{
			{Text: "attachment 1", Color: "#cb3f20"},
			{Title: "attachment 2", Color: "#ffa500"},
			{
				Text: "愿原力与你同在",
				Images: []IncomingAttachmentImage{
					{Url: "http://img3.douban.com/icon/ul15067564-30.jpg"},
				},
			},
		},
	}

	m.Build()
}
