package service

import (
	"context"
	"go-restful/dto"
)

type TodoListService interface {
	CreateTodoList(ctx context.Context, todoListRequest dto.TodoListRequestDTO) dto.TodoListResponseDTO
}
