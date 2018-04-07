package blob

import "io"

type Blob struct {
	BlobFilename string
	BlobContent  io.Reader
	BlobSize     int64
	BlobMimeType string
}

func (b *Blob) Filename() string {
	return b.BlobFilename
}

func (b *Blob) Content() io.Reader {
	return b.BlobContent
}

func (b *Blob) Size() int64 {
	return b.BlobSize
}

func (b *Blob) MimeType() string {
	return b.BlobMimeType
}
