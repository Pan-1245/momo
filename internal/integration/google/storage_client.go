package google

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// StorageClient is a wrapper around the Google Cloud Storage client.
type StorageClient struct {
	Client *storage.Client
}

// NewStorageClient creates a new instance of StorageClient.
func NewStorageClient(credentialsFilePath string) (*StorageClient, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create storage client: %v", err)
	}
	return &StorageClient{Client: client}, nil
}

// Close closes the storage client connection.
func (sc *StorageClient) Close() error {
	return sc.Client.Close()
}
