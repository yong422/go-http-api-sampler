package model

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	Users users
)

type UsersData struct {
	//ModelInterface
	UserId      string `json:"userId"`
	LastUpdated uint32 `json:"lastUpdated,int"`
	Status      uint32 `json:"status,int"`
}

type users struct {
}

func (u *UsersData) ToJsonResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(u)
}

func (u *UsersData) Get() bool {
	// find user data by user id from redis
	u.LastUpdated = uint32(time.Now().Unix())
	u.Status = 1
	return true
}
