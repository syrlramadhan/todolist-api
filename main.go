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

	//create
	router.POST("/api/v1/todolist/create", todoListController.CreateTodoList)

	//read
	router.GET("/api/v1/todolist", todoListController.FindAll)

	//update
	router.PUT("/api/v1/todolist/update/:todoListId", todoListController.UpdateTodoList)

	//delete
	router.DELETE("/api/v1/todolist/delete/:todoListId", todoListController.DeleteTodoList)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	errServer := server.ListenAndServe()
	util.SendPanicIfError(errServer)
}
