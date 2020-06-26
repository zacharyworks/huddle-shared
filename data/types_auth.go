package types

// Session struct representing a session
type Session struct {
	SessionID       string `json:"sessionID"`
	Token           string `json:"token"`
	State           string `json:"state"`
	UserFK          string `json:"userFK"`
	CreatedDateTime string `json:"createdDateTime"`
}

// OauthResponse represents a response from oAuth
type OauthResponse struct {
	Response Response `json:"Response"`
}

// Response from oAuth
type Response struct {
	ID      string `json:"iD"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}
