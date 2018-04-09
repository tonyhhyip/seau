package repository

import "github.com/dgrijalva/lfu-go"

type CachedStore struct {
	Store
	Cache *lfu.Cache
}

func (s *CachedStore) GetByDomain(domain string) (r *Repository, err error) {
	val := s.Cache.Get(domain)
	if val != nil {
		return val.(*Repository), nil
	}
	r, err = s.Store.GetByDomain(domain)
	s.Cache.Set(domain, r)
	return
}
