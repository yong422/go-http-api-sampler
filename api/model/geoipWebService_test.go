package model

import (
	"fmt"
	"sampler/config"
	"sampler/tool"
	"testing"
)

func TestSetGeoIpWebServiceAuthorizationInfo(t *testing.T) {
	SetGeoIpWebServiceAuthorizationInfo("testid", "testkey")
	if GeoIpWebService.licenseKey != "testkey" {
		t.Error("wrong result")
	}
	if GeoIpWebService.accountId != "testid" {
		t.Error("wrong result")
	}
	SetGeoIpWebServiceAuthorizationInfo("testid22", "test55key")
	if GeoIpWebService.licenseKey == "testkey" {
		t.Error("wrong result")
	}
	if GeoIpWebService.accountId == "testid" {
		t.Error("wrong result")
	}
}

// 	geo ip web service 테스트를 위한 테스트 결과값. 2020.08.21
//	host nexon.com test ip
//	nexon.com has address 52.78.153.209
//	nexon.com has address 52.78.145.30
/*
{
    "city": {
        "confidence": 0,
        "geoname_id": 1843564,
        "names": {
            "zh-CN": "仁川广域市",
            "de": "Incheon",
            "en": "Incheon",
            "es": "Inchon",
            "fr": "Incheon",
            "ja": "仁川広域市",
            "pt-BR": "Incheon",
            "ru": "Инчхон"
        }
    },
    "continent": {
        "code": "AS",
        "geoname_id": 6255147,
        "names": {
            "de": "Asien",
            "en": "Asia",
            "es": "Asia",
            "fr": "Asie",
            "ja": "アジア",
            "pt-BR": "Ásia",
            "ru": "Азия",
            "zh-CN": "亚洲"
        }
    },
    "country": {
        "confidence": 99,
        "iso_code": "KR",
        "geoname_id": 1835841,
        "names": {
            "ja": "大韓民国",
            "pt-BR": "Coreia do Sul",
            "ru": "Южная Корея",
            "zh-CN": "韩国",
            "de": "Südkorea",
            "en": "South Korea",
            "es": "Corea del Sur",
            "fr": "Corée du Sud"
        }
    },
    "location": {
        "accuracy_radius": 1000,
        "latitude": 37.4562,
        "longitude": 126.7288,
        "time_zone": "Asia/Seoul"
    },
    "maxmind": {
        "queries_remaining": 2476
    },
    "postal": {
        "confidence": 0,
        "code": "21539"
    },
    "registered_country": {
        "iso_code": "US",
        "geoname_id": 6252001,
        "names": {
            "ja": "アメリカ合衆国",
            "pt-BR": "Estados Unidos",
            "ru": "США",
            "zh-CN": "美国",
            "de": "USA",
            "en": "United States",
            "es": "Estados Unidos",
            "fr": "États-Unis"
        }
    },
    "subdivisions": [
        {
            "confidence": 0,
            "iso_code": "28",
            "geoname_id": 1843561,
            "names": {
                "en": "Incheon"
            }
        }
    ],
    "traits": {
        "is_anonymous": true,
        "is_hosting_provider": true,
        "user_type": "hosting",
        "autonomous_system_number": 16509,
        "autonomous_system_organization": "AMAZON-02",
        "domain": "amazonaws.com",
        "isp": "Amazon.com",
        "organization": "Amazon.com",
        "ip_address": "52.78.145.30",
        "network": "52.78.145.0/27"
    }
}
*/
func TestGetIpIntelligence(t *testing.T) {
	// test 를 위한 geoip authorization 정보를 가져온다.
	conf, err := config.RoadConfiugurationFile("../../config/sampler_config.json")
	if err != nil {
		t.Error(err)
	}
	SetGeoIpWebServiceAuthorizationInfo(conf.GeoIpWebServiceAccountId, conf.GeoIpWebServiceLicenseKey)
	// nexon.com ip 를 테스트 값으로 검색 실행.
	timer := tool.NewTimer()
	resultData, geoErr := GeoIpWebService.getIpIntelligence("52.78.145.30")
	fmt.Println(resultData)
	if geoErr != nil {
		t.Error(geoErr)
	} else {
		if resultData.Country.Names.En != "South Korea" {
			t.Error(fmt.Sprintf("%s != South Korea", resultData.Country.Names.En))
		}
		if resultData.Country.IsoCode != "KR" {
			t.Error(fmt.Sprintf("%s != KR", resultData.Country.IsoCode))
		}
		if resultData.Country.IsoCode != "KR" {
			t.Error(fmt.Sprintf("%s != KR", resultData.Country.IsoCode))
		}
		if resultData.Traits.UserType != "hosting" {
			t.Error(fmt.Sprintf("%s != hosting", resultData.Traits.UserType))
		}
	}
	if timer.ElapsedMilliseconds() >= 1000 {
		t.Error("Abnormal execution time")
	}
	fmt.Printf("execution time > %.4f ms", timer.ElapsedMilliseconds())
}
