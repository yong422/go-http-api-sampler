package model

import "net/http"

type ModelInterface interface {
	ToJsonResponse(w http.ResponseWriter)
}
