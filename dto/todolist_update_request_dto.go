package dto

type TodoListUpdateRequestDTO struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}
