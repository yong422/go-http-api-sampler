package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerInterface interface {
	Route(r *mux.Route)
	//	get resource
	Get(w http.ResponseWriter, r *http.Request)
	// 	create resource
	Post(w http.ResponseWriter, r *http.Request)
	// 	delete resource
	Delete(w http.ResponseWriter, r *http.Request)
	// 	update resource if exists
	Put(w http.ResponseWriter, r *http.Request)
}
