package repository

import (
	"fmt"

	"github.com/gofiber/storage/s3"
)

type S3Repository struct {
	storage *s3.Storage
}

func NewS3Repository(storage *s3.Storage) *S3Repository {
	return &S3Repository{
		storage: storage,
	}
}

func (r *S3Repository) Upload(key string, value []byte) (string, error) {
	err := r.storage.Set(key, value, 0)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://flaiers.s3.amazonaws.com/flallet/%s", key), nil
}
