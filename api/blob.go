package api

import "io"

type BlobManager interface {
	Put(repo string, blob Blob) error
	Get(repo, filename string) (Blob, error)
	Remove(repo, filename string) error

	PutByUrl(repo, filename string) (string, error)
	GetByUrl(repo, filename string) (string, error)
}

type Blob interface {
	Filename() string
	Content() io.Reader
	Size() int64
	MimeType() string
}
