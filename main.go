package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-restful/config"
	"go-restful/controller"
	"go-restful/repository"
	"go-restful/service"
	"go-restful/util"
	"net/http"
)

func main() {
	fmt.Print("coconut restfull api")

	config.GetEnv()
	postgresSql, err := config.OpenConnectionPostgres()
	util.SendPanicIfError(err)

	todoListRepository := repository.NewTodoListRepositoryImpl()
	todoListService := service.NewTodoListServiceImpl(todoListRepository, postgresSql)
	todoListController := controller.NewTodoListControllerImpl(todoListService)

	router := httprouter.New()

	router.POST("/api/v1/todolist/create", todoListController.CreateTodoList)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	errServer := server.ListenAndServe()
	util.SendPanicIfError(errServer)
}
