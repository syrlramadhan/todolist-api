package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoListController interface {
	CreateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
