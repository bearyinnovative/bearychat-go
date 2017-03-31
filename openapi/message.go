package openapi

type MessageKey string

type MessageSubtype string

const (
	MessageSubtypeNormal MessageSubtype = "normal"
	MessageSubtypeInfo                  = "info"
)

type MessageAttachmentImage struct {
	Url *string `json:"url,omitempty"`
}

type MessageAttachment struct {
	Title  *string                  `json:"title,omitempty"`
	Text   *string                  `json:"text,omitempty"`
	Color  *string                  `json:"color,omitempty"`
	Images []MessageAttachmentImage `json:"images,omitempty"`
}

type Message struct {
	Key         *MessageKey         `json:"key,omitempty"`
	TeamID      *string             `json:"team_id,omitempty"`
	UID         *string             `json:"uid,omitempty"`
	RobotID     *string             `json:"robot_id,omitempty"`
	VChannelID  *string             `json:"vchannel_id,omitempty"`
	ReferKey    *MessageKey         `json:"refer_key,omitempty"`
	Subtype     *MessageSubtype     `json:"subtype,omitempty"`
	Text        *string             `json:"text,omitempty"`
	Fallback    *string             `json:"fallback,omitempty"`
	Attachments []MessageAttachment `json:"attachments,omitempty"`
	Created     *Time               `json:"created,omitempty"`
	Updated     *Time               `json:"updated,omitempty"`
	CreatedTS   *VChannelTS         `json:"created_ts,omitempty"`
	IsChannel   *bool               `json:"is_channel,omitempty"`
}

type MessageService service
