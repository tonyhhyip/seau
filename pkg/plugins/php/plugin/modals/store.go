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
				WHERE php_repository.domain = $1`

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

func (s *Store) GetPackagesByVendor(domain, vendor string) (packages []*Package, err error) {
	conn, err := s.Opener.Open()
	if err != nil {
		return
	}

	defer conn.Close()

	query := `
			SELECT php_package.id, php_package.name, php_package.hash FROM php_package
				INNER JOIN php_vendor ON php_vendor.id = php_package.vendor
				INNER JOIN php_repository ON php_vendor.repository = php_repository.id
				WHERE php_repository.domain = $1 AND php_vendor.name = $2`

	rows, err := conn.Query(query, domain, vendor)
	if err != nil {
		return
	}

	defer rows.Close()

	packages = make([]*Package, 0)

	for rows.Next() {
		var pkg Package
		if err := rows.Scan(&pkg.ID, &pkg.Name, &pkg.Hash); err != nil {
			return nil, err
		}
		packages = append(packages, &pkg)
	}

	return
}

func (s *Store) GetPackageVersion(domain, vendor, pkg string) (versions []*PackageVersion, err error) {
	conn, err := s.Opener.Open()
	if err != nil {
		return
	}

	defer conn.Close()

	query := `
			SELECT
				php_package_version.id, php_package_version.version,
				php_package_version.release_time, php_package_version.shasum,
				php_package_version.composer_content
			FROM php_package_version
				INNER JOIN php_package ON php_package_version.package = php_package.id
				INNER JOIN php_vendor ON php_vendor.id = php_package.vendor
				INNER JOIN php_repository ON php_vendor.repository = php_repository.id
				WHERE php_repository.domain = $1 AND php_vendor.name = $2 AND php_package.name = $3`

	rows, err := conn.Query(query, domain, vendor, pkg)
	if err != nil {
		return
	}

	defer rows.Close()

	versions = make([]*PackageVersion, 0)

	for rows.Next() {
		var ver PackageVersion
		if err := rows.Scan(&ver.ID, &ver.Version, &ver.ReleaseTime, &ver.ShaSum, &ver.ComposerContent); err != nil {
			return nil, err
		}
		versions = append(versions, &ver)
	}

	return
}
