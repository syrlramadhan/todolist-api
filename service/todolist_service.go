package service

import (
	"context"
	"go-restful/dto"
)

type TodoListService interface {
	CreateTodoList(ctx context.Context, todoListRequest dto.TodoListRequestDTO) dto.TodoListResponseDTO
	UpdateTodoList(ctx context.Context, todoListRequest dto.TodoListUpdateRequestDTO) dto.TodoListResponseDTO
	FindAll(ctx context.Context, limit int) []dto.TodoListResponseDTO
}
