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

func (t todoListControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// search := request.URL.Query().Get("search")
	limitStr := request.URL.Query().Get("limit")

	limit := util.GetLimitValue(limitStr)

	todolist := t.TodoListService.FindAll(request.Context(), limit)
	response := dto.ResponseList {
		Code: 200,
		Status: "OK",
		Data: todolist,
		Limit: limit,
	}

	writter.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writter, response)
}