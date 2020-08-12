package api

import (
	"io"
	"net/http"
	"sampler/api/handler"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello World!")
}

func CreateApiRouter() *mux.Router {
	routerHandler := mux.NewRouter()
	routerHandler.HandleFunc("/", HelloWorld).Methods("GET")
	//usersHandler := handler.UsersHandler{}
	//routerHandler.HandleFunc("/users/{userId}", usersHandler.Get).Methods("GET", "POST")
	handler.UsersHandler{Router: routerHandler}.Route("/users/{userId}", "GET", "POST")

	lookupRouterHandler := routerHandler.PathPrefix("/lookup").Subrouter()

	handler.LookupHandler{Router: lookupRouterHandler}.Route("/ip/{ip}", "GET")

	return routerHandler
}
