package repository

import (
	"context"
	"database/sql"
	"go-restful/model"
)

type TodoListRepository interface {
	CreateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) (model.MstTodoList, error)
	UpdateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) model.MstTodoList
	FindById(ctx context.Context, tx *sql.Tx, id string) (model.MstTodoList, error)
	FindAll(ctx context.Context, tx *sql.Tx, limit int) []model.MstTodoList
	DeleteTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) model.MstTodoList
}
