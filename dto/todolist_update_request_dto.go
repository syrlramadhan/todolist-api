package dto

type TodoListUpdateRequestDTO struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}
