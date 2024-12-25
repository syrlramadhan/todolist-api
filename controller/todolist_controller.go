package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TodoListController interface {
	CreateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
