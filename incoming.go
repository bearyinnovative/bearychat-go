package bearychat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Incoming message builder.
//
//      m := Incoming{
//              Text: "Hello, **BearyChat**",
//              Markdown: true,
//              Notification: "Hello, BearyChat in Notification",
//      }
//      output, _ := m.Build()
//      http.Post("YOUR INCOMING HOOK URI HERE", "application/json", output)
//
// For full documentation, visit https://bearychat.com/integrations/incoming .
type Incoming struct {
	Text         string               `json:"text"`
	Notification string               `json:"notification,omitempty"`
	Markdown     bool                 `json:"markdown,omitempty"`
	Channel      string               `json:"channel,omitempty"`
	User         string               `json:"user,omitempty"`
	Attachments  []IncomingAttachment `json:"attachments,omitempty"`
}

// Build an incoming message.
func (m Incoming) Build() (io.Reader, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}

	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

// Validate fields.
func (m Incoming) Validate() error {
	if m.Text == "" {
		return fmt.Errorf("`text` is required for incoming message")
	}

	for i, a := range m.Attachments {
		if err := a.Validate(); err != nil {
			return errors.Wrapf(
				err,
				"#%d incoming attachment validate failed",
				i,
			)
		}
	}

	return nil
}

// IncomingAttachment contains incoming attachment fields.
type IncomingAttachment struct {
	Title  string                    `json:"title,omitempty"`
	Text   string                    `json:"text,omitempty"`
	Color  string                    `json:"color,omitempty"`
	Images []IncomingAttachmentImage `json:"images,omitempty"`
}

// Validate fields.
func (a IncomingAttachment) Validate() error {
	if a.Title == "" && a.Text == "" {
		return fmt.Errorf("`title`/`text` is required for incoming attachment")
	}

	for i, im := range a.Images {
		if err := im.Validate(); err != nil {
			return errors.Wrapf(
				err,
				"#%d incoming image validate failed",
				i,
			)
		}
	}

	return nil
}

// IncomingAttachmentImage contains attachment image fields.
type IncomingAttachmentImage struct {
	URL string `json:"url"`
}

// Validate fields.
func (i IncomingAttachmentImage) Validate() error {
	if i.URL == "" {
		return fmt.Errorf("`url` is required for incoming image")
	}

	return nil
}
