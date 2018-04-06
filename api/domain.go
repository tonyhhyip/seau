package api

type DomainRegistry interface {
	Occupy(domain string) error
	PrivateOccupy(domain string) error
}
