package handler

import (
	"net/http"
	"sampler/api/usecase"

	"github.com/gorilla/mux"
)

type LookupHandler struct {
	HandlerInterface
	Router *mux.Router
}

func (u LookupHandler) Route(path string, methodsToRegister ...string) {
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

func (u LookupHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ipIntelligence, err := usecase.Lookup.FindIpIntelligence(params["ip"])

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		ipIntelligence.ToJsonResponse(w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (u LookupHandler) Post(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u LookupHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u LookupHandler) Put(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

type VpnLookupHandler struct {
	HandlerInterface
	Router *mux.Router
}

func (u VpnLookupHandler) Route(path string, methodsToRegister ...string) {
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

func (u VpnLookupHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ipIntelligence, err := usecase.Lookup.FindVpnIpIntelligence(params["ip"])

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		ipIntelligence.ToJsonResponse(w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (u VpnLookupHandler) Post(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u VpnLookupHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u VpnLookupHandler) Put(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

type LookupFromWebServiceHandler struct {
	HandlerInterface
	Router *mux.Router
}

func (u LookupFromWebServiceHandler) Route(path string, methodsToRegister ...string) {
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

func (u LookupFromWebServiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ipIntelligence, err := usecase.Lookup.FindIpIntelligenceFromWebService(params["ip"])

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		ipIntelligence.ToJsonResponse(w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (u LookupFromWebServiceHandler) Post(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u LookupFromWebServiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u LookupFromWebServiceHandler) Put(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusMethodNotAllowed)
}
