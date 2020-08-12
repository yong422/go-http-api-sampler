package model

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var (
	GeoIp geoIp
)

type geoIp struct {
	geoIpModel
}

func (r geoIp) getCity(ip string) (*geoip2.City, error) {
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		return nil, fmt.Errorf("ip error")
	}
	reader, err := r.getReader()
	if err != nil {
		return nil, err
	}
	city, err := reader.City(ipAddress)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (r geoIp) GetCity(ip string) (*geoip2.City, error) {
	city, err := r.getCity(ip)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (r geoIp) GetCountryIsoCode(ip string) (string, error) {
	city, err := r.getCity(ip)
	if err != nil {
		return "", err
	}
	return city.Country.IsoCode, nil
}

// read-only model
type geoIpModel struct {
	reader *geoip2.Reader
}

func (r *geoIpModel) setReader(reader *geoip2.Reader) {
	r.reader = reader
}

func (r geoIpModel) getReader() (*geoip2.Reader, error) {
	if r.reader == nil {
		return nil, fmt.Errorf("geoip reader is nil")
	}
	return r.reader, nil
}

func SetGeoIp(reader *geoip2.Reader) {
	if reader == nil {
		return
	}
	GeoIp.setReader(reader)
}
