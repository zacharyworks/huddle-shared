package types

// Board struct representing a board table entry
type Board struct {
	BoardID string `json:"boardID"`
	Name    string `json:"name"`
}

// BoardMember struct representing a boardmember table entry
type BoardMember struct {
	BoardFK string `json:"boardFK"`
	UserFK  string `json:"userFK"`
}

type NewBoard struct {
	BoardFK int `json:"boardFK"`
	Code string `json:"code"`
}

type JoinBoard struct {
	UserFK int `json:"userFK"`
	Code string `json:"code"`
}