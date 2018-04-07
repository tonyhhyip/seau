package adpter

import "github.com/tonyhhyip/seau/pkg/server/repository"

type DomainRegistryFactory struct {
	Store repository.Store
}

func (drf *DomainRegistryFactory) Create(name string) *domainRegistry {
	return &domainRegistry{
		name:  name,
		store: drf.Store,
	}
}
