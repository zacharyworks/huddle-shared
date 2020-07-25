package types

// Board struct representing a board table entry
type Board struct {
	BoardID int `json:"boardID"`
	BoardType int `json:"boardType"`
	Name    string `json:"name"`
}

// BoardMember struct representing a boardmember table entry
type BoardMember struct {
	BoardFK int `json:"boardFK"`
	UserFK  string `json:"userFK"`
}

type NewBoard struct {
	Board 	Board    `json:"board"`
	UserFK  string `json:"userFK"`
}

type BoardJoinCode struct {
	BoardFK int    `json:"boardFK"`
	Code   string `json:"code"`
}

type BoardJoin struct {
	UserFK string    `json:"userFK"`
	Code   string `json:"code"`
}