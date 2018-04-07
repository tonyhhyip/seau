package adpter

import "github.com/tonyhhyip/seau/pkg/server/blob"

type BlobFactory struct {
	Manager blob.Manager
}

func (bf *BlobFactory) Create(name string) *BlobManager {
	return &BlobManager{
		Name:   name,
		Manger: bf.Manager,
	}
}
