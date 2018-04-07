package blob

import "github.com/tonyhhyip/seau/api"

type Manager interface {
	GetFile(handler, repo, filename string) (api.Blob, error)
	GetFileWithUrl(handler, repo, filename string) (string, error)
	PutFile(handler, repo string, blob api.Blob) error
	PutFileWithUrl(handler, repo, filename string) (string, error)
	RemoveFile(handler, repo, filename string) error
}
