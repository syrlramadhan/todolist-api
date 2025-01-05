package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TodoListController interface {
	CreateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
