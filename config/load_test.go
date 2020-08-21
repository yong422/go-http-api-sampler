package config

import "testing"

func TestRoadConfiugurationFile(t *testing.T) {
	samplerConfiguration, err := RoadConfiugurationFile("sampler_config.json")
	if err != nil {
		t.Error(err.Error())
	} else {
		if samplerConfiguration.AdminMail != "ykjo@nexon.co.kr" {
			t.Error("admin mail read failure")
		}
		if samplerConfiguration.GeoIpWebServiceAccountId != "377101" {
			t.Error("geoip account id failure")
		}
		if samplerConfiguration.GeoIpWebServiceLicenseKey != "nAWdDTEAUiCHGglX" {
			t.Error("geoip license key read failure")
		}
		if len(samplerConfiguration.GetRedisClusterAddresses()) != 3 {
			t.Error("redis cluster addresses read failure")
		}
	}
}
