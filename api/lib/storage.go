package lib

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

const (
	bucketName = "statsbomb-open-data-api-data"
)

func ReadGCSFile(path string, client *storage.Client) ([]byte, error) {
	rc, err := client.Bucket(bucketName).Object(path).NewReader(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to read file from GCS: %w", err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	return data, nil
}
