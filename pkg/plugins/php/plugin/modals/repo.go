package modals

type Repository struct {
	ID     string
	Domain string
	Owner  string

	store *Store
}

type RepositoryStore struct {
	*Store
}

func (s *RepositoryStore) Get(domain string) (r *Repository, err error) {
	conn, err := s.Opener.Open()

	if err != nil {
		return
	}
	defer conn.Close()

	query := "SELECT id, domain, owner FROM php_repository WHERE domain = $1"
	row := conn.QueryRow(query, domain)
	var repo Repository
	if err := row.Scan(&repo.ID, &repo.Domain, &repo.Owner); err != nil {
		return nil, err
	}
	r = &repo

	return
}
