package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	LookupFromWebService lookupFromWebService
	Lookup               lookup
	LookupVpn            lookupVpn
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
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		fmt.Println(err)
	}
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

type LookupVpnIpData struct {
	//ModelInterface
	Ip     string `json:"ip"`
	Region string `json:"region,string"`
	City   string `json:"city,string"`
}

func (u *LookupVpnIpData) ToJsonResponse(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		fmt.Println(err)
	}
}

type lookupVpn struct {
}

func (u *lookupVpn) Get(ip string) (*LookupVpnIpData, error) {
	// find user data by user id from redis
	vpnData, err := IpToLocation.GetAll(ip)
	if err != nil {
		return nil, errors.New("Not found ip data")
	}
	ipData := LookupVpnIpData{
		Ip:     ip,
		Region: vpnData.Region,
		City:   vpnData.City}
	return &ipData, nil
}

type lookupFromWebService struct {
}

func (u *lookupFromWebService) Get(ip string) (*LookupIpData, error) {
	// find user data by user id from redis
	//city, err := GeoIp.GetCity(ip)
	ipIntelligence, err := GeoIpWebService.getIpIntelligence(ip)
	if err != nil {
		return nil, errors.New("Not found intelligence data")
	}
	ipData := LookupIpData{
		Ip:          ip,
		CountryName: ipIntelligence.Country.Names.En,
		CountryCode: ipIntelligence.Country.IsoCode,
		TimeZone:    ipIntelligence.Location.TimeZone}
	return &ipData, nil
}
