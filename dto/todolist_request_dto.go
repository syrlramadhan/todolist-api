package dto

type TodoListRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}
