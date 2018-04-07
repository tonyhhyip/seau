package adpter

import (
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/blob"
)

type BlobManager struct {
	Name   string
	Manger blob.Manager
}

func (b *BlobManager) Put(repo string, blob api.Blob) error {
	return b.Manger.PutFile(b.Name, repo, blob)
}

func (b *BlobManager) Get(repo, filename string) (api.Blob, error) {
	return b.Manger.GetFile(b.Name, repo, filename)
}

func (b *BlobManager) Remove(repo, filename string) error {
	return b.Manger.RemoveFile(b.Name, repo, filename)
}

func (b *BlobManager) PutByUrl(repo, filename string) (string, error) {
	return b.Manger.PutFileWithUrl(b.Name, repo, filename)
}

func (b *BlobManager) GetByUrl(repo, filename string) (string, error) {
	return b.Manger.GetFileWithUrl(b.Name, repo, filename)
}
