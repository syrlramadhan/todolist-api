package dto

import "time"

type TodoListResponseDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
