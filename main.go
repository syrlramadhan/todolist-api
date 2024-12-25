package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-restful/config"
	"log"
	"net/http"
)

func main() {
	fmt.Print("coconut restfull api")

	config.GetEnv()
	postgresSql, err := config.OpenConnectionPostgres()
	if err != nil {
		panic(err)
	}

	log.Print(postgresSql)

	router := httprouter.New()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	errServer := server.ListenAndServe()
	if errServer != nil {
		panic(errServer)
	}
}
