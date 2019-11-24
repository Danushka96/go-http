package main

import (
	"./handlers"
	"./models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {

	router := mux.NewRouter()
	handlers.PersonMap = make(map[int64]models.Person)
	//mux := http.NewServeMux()

	server := http.Server{
		Addr:         ":8013",
		Handler:      router,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	router.Handle("/person", handlers.HandlePersonSave{}).
		Methods(http.MethodPost).
		Name("add-person")

	router.Handle("/persons", handlers.HandleGet{}).
		Methods(http.MethodGet).
		Name("get-person")

	router.Handle("/person/all", handlers.GetAll{}).
		Methods(http.MethodGet).
		Name("get-all")

	fmt.Println("server is starting")

	err := server.ListenAndServe()

	if err!=nil{
		panic(err)
	}

}
