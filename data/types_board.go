package types

// Board struct representing a board table entry
type Board struct {
	BoardID int `json:"boardID"`
	Name    string `json:"name"`
}

// BoardMember struct representing a boardmember table entry
type BoardMember struct {
	BoardFK int `json:"boardFK"`
	UserFK  string `json:"userFK"`
}

type NewBoard struct {
	BoardFK int    `json:"boardFK"`
	Code    string `json:"code"`
}

type BoardJoinCode struct {
	BoardFK int    `json:"boardFK"`
	Code   string `json:"code"`
}

type BoardJoin struct {
	UserFK int    `json:"userFK"`
	Code   string `json:"code"`
}