package types

// Todo item
type Todo struct {
	TodoID   int    `json:"todoID"`
	Status   int    `json:"status"`
	Value    string `json:"value"`
	ParentFK int    `json:"parentFK"`
	BoardFK  int    `json:"boardFK"`
}