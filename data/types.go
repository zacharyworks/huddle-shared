package types

// StringAction an action with a string as the payload
type StringAction struct {
	Subset  string `json:"subset"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

// StringAction an action with a string as the payload
type Action struct {
	Subset  string `json:"subset"`
	Type    string `json:"type"`
	Payload interface{} `json:"payload"`
}
