package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
)

var GeoIpWebService geoIpWebService

type GeoIpWebServiceData struct {
	Continent struct {
		Code  string `json:"code"`
		Names struct {
			En string `json:"en"`
		}
	}
	City struct {
		Confidence int `json:"confidence"`
		Names      struct {
			En string `json:"en"`
		}
	}
	Country struct {
		IsoCode string `json:"iso_code"`
		Names   struct {
			En string `json:"en"`
		}
	}
	Location struct {
		TimeZone string `json:"time_zone"`
	}
	Traits struct {
		UserType string `json:"user_type"`
		Domain   string `json:"domain"`
	}
}

type geoIpWebService struct {
	geoIpWebServiceModel
}

type geoIpWebServiceModel struct {
	accountId  string
	licenseKey string
}

func (g *geoIpWebServiceModel) getIpIntelligence(ip string) (*GeoIpWebServiceData, error) {
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		return nil, errors.New("ip error")
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://geoip.maxmind.com/geoip/v2.1/insights/"+ip, nil)
	req.SetBasicAuth(g.accountId, g.licenseKey)
	req.Header.Add("Accept", "application/vnd.maxmind.com-country+json; charset=UTF-8;")
	resp, reqErr := client.Do(req)
	resultData := GeoIpWebServiceData{}
	if reqErr != nil {
		return nil, reqErr
	} else {
		responseData, _ := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		jsonErr := json.Unmarshal(responseData, &resultData)
		if jsonErr != nil {
			return nil, jsonErr
		}
		//fmt.Println(responseString)
	}

	return &resultData, nil
}

func SetGeoIpWebServiceAuthorizationInfo(accountId string, licenseKey string) {
	GeoIpWebService.accountId = accountId
	GeoIpWebService.licenseKey = licenseKey
}
