package service

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go-restful/dto"
	"go-restful/model"
	"go-restful/repository"
	"go-restful/util"
	"time"
)

type todoListServiceImpl struct {
	TodoListRepository repository.TodoListRepository
	DB                 *sql.DB
}

func NewTodoListServiceImpl(todoListRepository repository.TodoListRepository, DB *sql.DB) TodoListService {
	return &todoListServiceImpl{
		TodoListRepository: todoListRepository,
		DB:                 DB,
	}
}

func (t *todoListServiceImpl) CreateTodoList(ctx context.Context, todoListRequest dto.TodoListRequestDTO) dto.TodoListResponseDTO {

	tx, err := t.DB.Begin()
	util.SendPanicIfError(err)
	defer util.CommitOrRollBack(tx)

	todoList := model.MstTodoList{
		ID:          uuid.New().String(),
		Title:       todoListRequest.Title,
		Description: todoListRequest.Description,
		Status:      todoListRequest.Status,
		CreatedAt:   time.Now(),
	}

	createTodoList, errSave := t.TodoListRepository.CreateTodoList(ctx, tx, todoList)
	util.SendPanicIfError(errSave)

	return convertToResponseDTO(createTodoList)
}

func convertToResponseDTO(mstTodoList model.MstTodoList) dto.TodoListResponseDTO {
	return dto.TodoListResponseDTO{
		Title:       mstTodoList.Title,
		Description: mstTodoList.Description,
		Status:      mstTodoList.Status,
		CreatedAt:   time.Now(),
	}
}
