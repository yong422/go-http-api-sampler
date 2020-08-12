package config

import (
	"github.com/tkanos/gonfig"
)

type SamplerConfiguration struct {
	AdminMail             string
	GeoIp2CityDatabase    string
	RedisClusterAddresses []string
}

func (r *SamplerConfiguration) GetRedisClusterAddresses() []string {
	return r.RedisClusterAddresses
}

func RoadConfiugurationFile(fileName string) (*SamplerConfiguration, error) {
	configuration := SamplerConfiguration{}
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}
