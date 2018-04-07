package repository

type Store interface {
	GetByDomain(domain string) (*Repository, error)
	Save(repo *Repository) error
	Delete(handler, domain string) error
}
