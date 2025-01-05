package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restful/model"
	"go-restful/util"
)

type todoListRepositoryImpl struct {
}

func NewTodoListRepositoryImpl() TodoListRepository {
	return &todoListRepositoryImpl{}
}

func (t *todoListRepositoryImpl) CreateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) (model.MstTodoList, error) {

	query := `INSERT INTO mst_todolist(id, title, description, status, created_at) VALUES($1, $2, $3, $4, $5)`

	_, err := tx.ExecContext(ctx, query, todoList.ID, todoList.Title, todoList.Description, todoList.Status, todoList.CreatedAt)
	util.SendPanicIfError(err)

	return todoList, nil
}

func (t *todoListRepositoryImpl) UpdateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) model.MstTodoList {
	query := `UPDATE mst_todolist SET title = $1, description = $2, status = $3 WHERE id = $4`

	_, err := tx.ExecContext(ctx, query, todoList.Title, todoList.Description,todoList.Status, todoList.ID)
	util.SendPanicIfError(err)

	return todoList
}

func (t *todoListRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (model.MstTodoList, error) {
	query := `SELECT id, title, description, status FROM mst_todolist WHERE id = $1`
	rows, err := tx.QueryContext(ctx, query, id)
	util.SendPanicIfError(err)

	defer rows.Close()
	todoList := model.MstTodoList{}
	if rows.Next() {
		err := rows.Scan(&todoList.ID, &todoList.Title, &todoList.Description, &todoList.Status)
		util.SendPanicIfError(err)
		return todoList, err
	} else {
		return todoList, errors.New("id todolist not found")
	}
}

func (t *todoListRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, limit int) []model.MstTodoList {
	query := `SELECT id, title, description, status FROM mst_todolist LIMIT $1`
	rows, err := tx.Query(query, limit)
	util.SendPanicIfError(err)
	defer rows.Close()

	var todolist []model.MstTodoList
	for rows.Next() {
		list := model.MstTodoList{}
		err := rows.Scan(&list.ID, &list.Title, &list.Description, &list.Status)
		util.SendPanicIfError(err)
		todolist = append(todolist, list)
	}

	return todolist
}