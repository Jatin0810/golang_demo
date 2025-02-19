package routes

import (
	"github.com/gorilla/mux"
	"main.go/api/handler"
)

func Routes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", handler.ServerStart).Methods("GET")
	router.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/UserUpdate", handler.UserUpdate).Methods("POST")
	router.HandleFunc("/DeleteUser", handler.DeleteUser).Methods("DELETE")

	return router

}
