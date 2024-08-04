package google

import (
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

// GenerateSignedURL generates a v4 signed URL for a specified object in a bucket.
func (sc *GoogleStorageClient) GenerateSignedURL(bucketName, objectName string, expiry time.Duration) (string, error) {
	url, err := sc.Client.Bucket(bucketName).SignedURL(objectName, &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(expiry),
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %v", err)
	}
	return url, nil
}
