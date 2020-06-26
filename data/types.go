package types

// StringAction an action with a string as the payload
type StringAction struct {
	ActionSubset  string `json:"ActionSubset"`
	ActionType    string `json:"ActionType"`
	ActionPayload string `json:"ActionPayload"`
}
