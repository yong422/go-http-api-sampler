package model

import (
	"fmt"
	"github.com/ip2location/ip2location-go"
	"net"
)

var (
	IpToLocation ipToLocation
)

type ipToLocation struct {
	ipToLocationModel
}

func (r ipToLocation) getAll(ip string) (*ip2location.IP2Locationrecord, error) {
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		return nil, fmt.Errorf("ip error")
	}
	db, err := r.getDB()
	if err != nil {
		return nil, err
	}
	allData, err := db.Get_all(ip)
	if err != nil {
		return nil, err
	}
	return &allData, nil
}

func (r ipToLocation) GetAll(ip string) (*ip2location.IP2Locationrecord, error) {
	return r.getAll(ip)
}

func (r ipToLocation) GetCountry(ip string) (string, error) {
	allData, err := r.getAll(ip)
	if err != nil {
		return "", err
	}
	return allData.Country_long, nil
}

func (r ipToLocation) GetProxyType(ip string) (string, error) {
	allData, err := r.getAll(ip)
	if err != nil {
		return "", err
	}
	return allData.Usagetype, nil
}

type ipToLocationModel struct {
	db *ip2location.DB
}

func (r *ipToLocationModel) setDB(db *ip2location.DB) {
	r.db = db
}

func (r ipToLocationModel) getDB() (*ip2location.DB, error) {
	if r.db == nil {
		return nil, fmt.Errorf("IpToLocation db is nil")
	}
	return r.db, nil
}

func SetIpToLocation(db *ip2location.DB) {
	if db == nil {
		return
	}
	IpToLocation.setDB(db)
}
