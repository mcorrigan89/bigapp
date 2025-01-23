package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/mcorrigan89/simple_auth/server/internal/common"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/xid"
)

var useSSL = false

type blobStorageService struct {
	config *common.Config
}

func NewBlobStorageService(cfg *common.Config) *blobStorageService {
	return &blobStorageService{
		config: cfg,
	}
}

func (s *blobStorageService) GetObject(ctx context.Context, bucketName, objectKey string) ([]byte, error) {
	minioClient, err := minio.New(s.config.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.config.Storage.AccessKey, s.config.Storage.SecretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	objResponse, err := minioClient.GetObject(ctx, bucketName, objectKey, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer objResponse.Close()

	respByte, err := io.ReadAll(objResponse)
	if err != nil {
		return nil, err
	}

	return respByte, nil
}

func (s *blobStorageService) UploadObject(ctx context.Context, bucketName, objectKey string, object io.Reader, size int64) (*string, error) {
	minioClient, err := minio.New(s.config.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.config.Storage.AccessKey, s.config.Storage.SecretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	uniqueId := xid.New().String()
	uniqueObjectKey := fmt.Sprintf("%s-%s", uniqueId, objectKey)

	_, err = minioClient.PutObject(ctx, bucketName, uniqueObjectKey, object, size, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &uniqueObjectKey, nil
}
