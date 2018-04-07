package modals

import "github.com/tonyhhyip/seau/api"

type Store struct {
	Opener api.Opener
}

func (s *Store) GetVendorByDomain(domain string) (vendors []*Vendor, err error) {
	conn, err := s.Opener.Open()
	if err != nil {
		return
	}

	defer conn.Close()

	query := `
			SELECT php_vendor.id, php_vendor.name, php_vendor.hash FROM php_vendor
				INNER JOIN php_repository ON php_vendor.repository = php_repository.id
				WHERE php_repository.domain = $1
`

	rows, err := conn.Query(query, domain)
	if err != nil {
		return
	}

	defer rows.Close()

	vendors = make([]*Vendor, 0)

	for rows.Next() {
		var vendor Vendor
		if err := rows.Scan(&vendor.ID, &vendor.Name, &vendor.Hash); err != nil {
			return nil, err
		}
		vendors = append(vendors, &vendor)
	}

	return
}
