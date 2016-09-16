package bearychat

// Team information
type Team struct {
	Id          string `json:"id"`
	Subdomain   string `json:"subdomain"`
	Name        string `json:"name"`
	UserId      string `json:"uid"`
	Description string `json:"description"`
	EmailDomain string `json:"email_domain"`
	Inactive    bool   `json:"inactive"`
	CreatedAt   string `json:"created"` // TODO parse date
	UpdatedAt   string `json:"updated"` // TODO parse date
}

const (
	UserRoleOwner   = "owner"
	UserRoleAdmin   = "admin"
	UserRoleNormal  = "normal"
	UserRoleVisitor = "visitor"
)

const (
	UserTypeNormal    = "normal"
	UserTypeAssistant = "assistant"
	UserTypeHubot     = "hubot"
)

// User information
type User struct {
	Id         string `json:"id"`
	TeamId     string `json:"team_id"`
	VChannelId string `json:"vchannel_id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	AvatarUrl  string `json:"avatar_url"`
	Role       string `json:"role"`
	Type       string `json:"type"`
	Conn       string `json:"conn"`
	CreatedAt  string `json:"created"` // TODO parse date
	UpdatedAt  string `json:"updated"` // TODO parse date
}

// IsOnline tells user connection status.
func (u User) IsOnline() bool {
	return u.Conn == "connected"
}

// IsNormal tells if this user a normal user (owner, admin or normal)
func (u User) IsNormal() bool {
	return u.Type == UserTypeNormal && u.Role != UserRoleVisitor
}

// Channel information.
type Channel struct {
	Id         string `json:"id"`
	TeamId     string `json:"team_id"`
	UserId     string `json:"uid"`
	VChannelId string `json:"vchannel_id"`
	Name       string `json:"name"`
	IsPrivate  bool   `json:"private"`
	IsGeneral  bool   `json:"general"`
	Topic      string `json:"topic"`
	CreatedAt  string `json:"created"` // TODO parse date
	UpdatedAt  string `json:"updated"` // TODO parse date
}
