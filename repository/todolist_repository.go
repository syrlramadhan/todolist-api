package repository

import (
	"context"
	"database/sql"
	"go-restful/model"
)

type TodoListRepository interface {
	CreateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) (model.MstTodoList, error)
}
