package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"textgopher/middleware"
	routes "textgopher/routes"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user/login", routes.SignIn).Methods("POST")
	router.HandleFunc("/api/user/registration", routes.SignUp).Methods("POST")
	router.HandleFunc("/api/upload", middleware.IsAuthorized(routes.Upload)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
