package blob

import (
	"fmt"
	"net/url"
	"time"

	"github.com/minio/minio-go"
	"github.com/tonyhhyip/seau/api"
	commonBlob "github.com/tonyhhyip/seau/pkg/common/blob"
)

type MinioManager struct {
	Client     *minio.Client
	BucketName string
}

func (m *MinioManager) objectName(handler, repo, filename string) string {
	return handler + "/" + repo + "/" + filename
}

func (m *MinioManager) GetFile(handler, repo, filename string) (blob api.Blob, err error) {
	objName := m.objectName(handler, repo, filename)
	obj, err := m.Client.GetObject(m.BucketName, objName, minio.GetObjectOptions{})
	if err != nil {
		return
	}

	stat, err := obj.Stat()
	if err != nil {
		return
	}

	blob = &commonBlob.Blob{
		BlobFilename: filename,
		BlobContent:  obj,
		BlobSize:     stat.Size,
		BlobMimeType: stat.ContentType,
	}

	return
}

func (m *MinioManager) GetFileWithUrl(handler, repo, filename string) (presignedURL string, err error) {
	values := make(url.Values)
	objName := m.objectName(handler, repo, filename)
	values.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	u, err := m.Client.PresignedGetObject(m.BucketName, objName, 5*time.Minute, values)
	if err != nil {
		return
	}

	presignedURL = u.String()
	return
}

func (m *MinioManager) PutFile(handler, repo string, blob api.Blob) (err error) {
	objName := m.objectName(handler, repo, blob.Filename())
	_, err = m.Client.PutObject(
		m.BucketName,
		objName,
		blob.Content(),
		blob.Size(),
		minio.PutObjectOptions{
			ContentType: blob.MimeType(),
		},
	)
	return
}

func (m *MinioManager) PutFileWithUrl(handler, repo, filename string) (presignedUrl string, err error) {
	objName := m.objectName(handler, repo, filename)
	u, err := m.Client.PresignedPutObject(m.BucketName, objName, 5*time.Minute)
	presignedUrl = u.String()
	return
}

func (m *MinioManager) RemoveFile(handler, repo, filename string) error {
	objName := m.objectName(handler, repo, filename)
	return m.Client.RemoveObject(m.BucketName, objName)
}
