package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-restful/dto"
	"go-restful/service"
	"go-restful/util"
	"net/http"
)

type todoListControllerImpl struct {
	TodoListService service.TodoListService
}

func NewTodoListControllerImpl(todoListService service.TodoListService) TodoListController {
	return &todoListControllerImpl{
		TodoListService: todoListService,
	}
}

func (t todoListControllerImpl) CreateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestDTO := dto.TodoListRequestDTO{}
	util.ReadFromRequestBody(request, &requestDTO)
	responseDTO := t.TodoListService.CreateTodoList(request.Context(), requestDTO)
	util.WriteToResponseBody(writer, responseDTO)
}
