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
	response := dto.ResponseList {
		Code: 200,
		Status: "OK",
		Data: responseDTO,
	}
	util.WriteToResponseBody(writer, response)
}

func (t todoListControllerImpl) UpdateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListRequest := dto.TodoListUpdateRequestDTO{}
	util.ReadFromRequestBody(request, &todoListRequest)

	todoListID := params.ByName("todoListId")
	todoListRequest.Id = todoListID

	responseDTO := t.TodoListService.UpdateTodoList(request.Context(), todoListRequest)
	response := dto.ResponseList {
		Code: 200,
		Status: "OK",
		Data: responseDTO,
	}
	util.WriteToResponseBody(writer, response)
}

func (t todoListControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

func (t todoListControllerImpl) DeleteTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListId := params.ByName("todoListId")

	responseDTO := t.TodoListService.DeleteTodoList(request.Context(), todoListId)
	response := dto.ResponseList {
		Code: 200,
		Status: "OK",
		Data: responseDTO,
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}