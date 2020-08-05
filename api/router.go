package api

import (
	"sampler/api/handler"

	"github.com/gorilla/mux"
)

func CreateApiRouter() *mux.Router {
	routerHandler := mux.NewRouter()
	//usersHandler := handler.UsersHandler{}
	//routerHandler.HandleFunc("/users/{userId}", usersHandler.Get).Methods("GET", "POST")
	handler.UsersHandler{Router: routerHandler}.Route("/users/{userId}", "GET", "POST")
	return routerHandler
}
