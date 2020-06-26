package types

// Todo item
type Todo struct {
	TodoID   int    `json:"todoID"`
	Status   int    `json:"status"`
	Value    string `json:"value"`
	ParentFK int    `json:"parentFK"`
	BoardFK  int    `json:"boardFK"`
}

// TodoAction - a command to be broadcast / consumed
type TodoAction struct {
	ActionSubset  string `json:"ActionSubset"`
	ActionType    string `json:"ActionType"`
	ActionPayload []Todo `json:"ActionPayload"`
}

// TodoActionSingle stuff
type TodoActionSingle struct {
	ActionSubset  string `json:"ActionSubset"`
	ActionType    string `json:"ActionType"`
	ActionPayload Todo   `json:"ActionPayload"`
}

// TodoActionID stuff
type TodoActionID struct {
	ActionSubset  string `json:"ActionSubset"`
	ActionType    string `json:"ActionType"`
	ActionPayload int    `json:"ActionPayload"`
}
