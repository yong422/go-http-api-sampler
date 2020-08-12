package model

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	Lookup lookup
)

type LookupIpData struct {
	//ModelInterface
	Ip          string `json:"ip"`
	CountryName string `json:"country_name,string"`
	CountryCode string `json:"country_code,string"`
	City        string `json:"city,string"`
	TimeZone    string `json:"time_zone,string"`
}

type lookup struct {
}

func (u *LookupIpData) ToJsonResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(u)
}

func (u *lookup) Get(ip string) (*LookupIpData, error) {
	// find user data by user id from redis
	city, err := GeoIp.GetCity(ip)
	if err != nil {
		return nil, errors.New("Not found city data")
	}
	//fmt.Println(city)

	ipData := LookupIpData{
		Ip:          ip,
		CountryName: city.Country.Names["en"],
		CountryCode: city.Country.IsoCode,
		TimeZone:    city.Location.TimeZone}
	return &ipData, nil
}
