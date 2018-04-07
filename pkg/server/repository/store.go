package repository

import (
	"database/sql"

	"github.com/tonyhhyip/seau/api"
)

type Store struct {
	Opener api.Opener
}

func (s *Store) GetByDomain(domain string) (r *Repository, err error) {
	conn, err := s.Opener.Open()
	if err != nil {
		return
	}

	defer conn.Close()

	query := "SELECT domain, allow_public_read, handler FROM repositories WHERE domain = $1"
	row := conn.QueryRow(query, domain)

	var repo Repository
	r = &repo
	err = row.Scan(&repo.Domain, &repo.AllowPublicRead, &repo.Handler)

	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

func (s *Store) Save(repo *Repository) (err error) {
	conn, err := s.Opener.Open()
	if err != nil {
		return
	}

	defer conn.Close()

	query := "INSERT INTO repositories (domain, allow_public_read, handler) VALUES ($1, $2, $3)"
	_, err = conn.Exec(query, repo.Domain, repo.AllowPublicRead, repo.Handler)
	return
}

func (s *Store) Delete(handler, domain string) (err error) {
	conn, err := s.Opener.Open()
	if err != nil {
		return
	}

	defer conn.Close()

	query := "DELETE FROM repositories WHERE domain = $1 AND handler = $2"
	_, err = conn.Exec(query, domain, handler)
	return
}
