package usecase

import (
	"sampler/api/model"
)

var (
	Lookup lookup
)

type lookupInterface interface {
	FindIpIntelligence(ip string)
	FindVpnIpIntelligence(ip string)
}

type lookup struct {
	lookupInterface
}

func (l *lookup) FindIpIntelligence(ip string) (*model.LookupIpData, error) {
	return model.Lookup.Get(ip)
}
func (l *lookup) FindIpIntelligenceFromWebService(ip string) (*model.LookupIpData, error) {
	return model.LookupFromWebService.Get(ip)
}

func (l *lookup) FindVpnIpIntelligence(ip string) (*model.LookupVpnIpData, error) {
	return model.LookupVpn.Get(ip)
}
