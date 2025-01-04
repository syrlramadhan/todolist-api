package controller

import (
	"go-restful/dto"
	"go-restful/service"
	"go-restful/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

func (t todoListControllerImpl) UpdateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListRequest := dto.TodoListUpdateRequestDTO{}
	util.ReadFromRequestBody(request, &todoListRequest)

	todoListID := params.ByName("todoListId")
	todoListRequest.Id = todoListID

	responseDTO := t.TodoListService.UpdateTodoList(request.Context(), todoListRequest)
	util.WriteToResponseBody(writer, responseDTO)
}
