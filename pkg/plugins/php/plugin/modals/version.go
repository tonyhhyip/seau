package modals

import "time"

type PackageVersion struct {
	ID              string
	Version         string
	ReleaseTime     time.Time
	ShaSum          string
	ComposerContent string
}
