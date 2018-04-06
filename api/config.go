package api

type Config interface {
	Opener() Opener
	DomainRegistry() DomainRegistry
}
