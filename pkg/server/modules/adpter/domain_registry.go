package adpter

import "github.com/tonyhhyip/seau/pkg/server/repository"

type domainRegistry struct {
	name  string
	store repository.Store
}

func (dr *domainRegistry) Occupy(domain string) error {
	repo := &repository.Repository{
		Domain:          domain,
		AllowPublicRead: true,
		Handler:         dr.name,
	}

	return dr.store.Save(repo)
}

func (dr *domainRegistry) PrivateOccupy(domain string) error {
	repo := &repository.Repository{
		Domain:          domain,
		AllowPublicRead: false,
		Handler:         dr.name,
	}

	return dr.store.Save(repo)
}

func (dr *domainRegistry) Release(domain string) error {
	return dr.store.Delete(dr.name, domain)
}
