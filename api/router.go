package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"sampler/api/handler"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, "Hello World!"); err != nil {
		fmt.Println(err)
	}

}

//	@return *mux.Router
func CreateApiRouter() *mux.Router {
	routerHandler := mux.NewRouter()
	routerHandler.HandleFunc("/", HelloWorld).Methods("GET")
	handler.UsersHandler{Router: routerHandler}.Route("/users/{userId}", "GET", "POST")

	lookupRouterHandler := routerHandler.PathPrefix("/lookup").Subrouter()

	//handler.LookupHandler{Router: lookupRouterHandler}.Route("/ip/{ip}", "GET")
	handler.LookupFromWebServiceHandler{Router: lookupRouterHandler}.Route("/ip/{ip}", "GET")
	handler.VpnLookupHandler{Router: lookupRouterHandler}.Route("/vpn/{ip}", "GET")

	return routerHandler
}
