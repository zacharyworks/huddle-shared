package types

// User represents a user
type User struct {
	UserID  int    `json:"userID"`
	OauthID string `json:"oauthID"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Name string `json:"name"`
	GivenName string `json:"givenName"`
	FamilyName string `json:"familyName"`
}
