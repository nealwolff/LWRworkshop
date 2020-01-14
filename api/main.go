package main

import (
	"LWRworkshop/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	theRouter := mux.NewRouter()

	theRouter.HandleFunc("/route", handlers.Route).Methods(http.MethodGet)
	theRouter.HandleFunc("/user", handlers.UserHandler).Methods(http.MethodPost, http.MethodGet)
	theRouter.HandleFunc("/user/{id}/loan/{loanId}", handlers.UserHandler).Methods(http.MethodGet)

	log.Println("The Api is listening")

	http.ListenAndServe(":8080", theRouter)
}
