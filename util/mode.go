package util

import (
	"go-restful/dto"
	"go-restful/model"
)

func ToTodoListResponse(todolist model.MstTodoList) dto.TodoListResponseDTO {
	return dto.TodoListResponseDTO{
		Title: todolist.Title,
		Description: todolist.Description,
		Status: todolist.Status,
	}
}