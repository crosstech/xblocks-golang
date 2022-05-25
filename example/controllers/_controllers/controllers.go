package controllers

import (
	"net/http"

	"github.com/crosstech/xblocks-golang/example/controllers/todoController"
	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/todos/{id}", todoController.GetById).Methods("GET")

	http.ListenAndServe(":7000", router)
}
