package usecase

import (
	"sampler/api/model"
)

var (
	Lookup lookup
)

type lookupInterface interface {
	FindIpIntelligence(ip string)
}

type lookup struct {
	lookupInterface
}

func (l *lookup) FindIpIntelligence(ip string) (*model.LookupIpData, error) {
	return model.Lookup.Get(ip)
}
