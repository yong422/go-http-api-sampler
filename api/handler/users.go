package handler

import (
	"net/http"
	"sampler/api/model"

	"github.com/gorilla/mux"
)

type UsersHandler struct {
	HandlerInterface
	Router *mux.Router
}

func (u UsersHandler) Route(path string, methodsToRegister ...string) {
	for _, method := range methodsToRegister {
		switch method {
		case "GET":
			u.Router.HandleFunc(path, u.Get).Methods(method)
		case "POST":
			u.Router.HandleFunc(path, u.Post).Methods(method)
		case "DELETE":
			u.Router.HandleFunc(path, u.Delete).Methods(method)
		case "PUT":
			u.Router.HandleFunc(path, u.Put).Methods(method)
		}
	}
}

func (u UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	users := model.UsersData{UserId: params["userId"]}

	if users.Get() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		users.ToJsonResponse(w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (u UsersHandler) Post(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u UsersHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
