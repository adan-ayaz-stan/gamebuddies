package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const AvatarBucket = "gamebuddies-avatars"

// GamesBucket is the public bucket that stores game cover images.
const GamesBucket = "gm-public"

// Client wraps the MinIO client with app-specific helpers.
type Client struct {
	mc     *minio.Client
	bucket string
}

// NewClient creates a MinIO client from environment variables:
//
//	MINIO_ENDPOINT  (default: localhost:9000)
//	MINIO_ACCESS_KEY
//	MINIO_SECRET_KEY
//	MINIO_USE_SSL   (default: false)
func NewClient() (*Client, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:9000"
	}
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"

	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("minio: failed to create client: %w", err)
	}

	c := &Client{mc: mc, bucket: AvatarBucket}

	// Ensure the bucket exists
	ctx := context.Background()
	exists, err := mc.BucketExists(ctx, AvatarBucket)
	if err != nil {
		return nil, fmt.Errorf("minio: bucket check failed: %w", err)
	}
	if !exists {
		if err := mc.MakeBucket(ctx, AvatarBucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("minio: failed to create bucket: %w", err)
		}
		// Set bucket policy to public-read so avatar URLs are directly accessible
		policy := fmt.Sprintf(`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::%s/*"]}]}`, AvatarBucket)
		if err := mc.SetBucketPolicy(ctx, AvatarBucket, policy); err != nil {
			return nil, fmt.Errorf("minio: failed to set bucket policy: %w", err)
		}
	}

	return c, nil
}

// UploadAvatar uploads a file to MinIO and returns the public object URL.
// objectName should be a unique identifier (e.g., "avatars/user-{id}.webp").
func (c *Client) UploadAvatar(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) (string, error) {
	_, err := c.mc.PutObject(ctx, c.bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("minio: upload failed: %w", err)
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:9000"
	}

	scheme := "http"
	if os.Getenv("MINIO_USE_SSL") == "true" {
		scheme = "https"
	}

	url := fmt.Sprintf("%s://%s/%s/%s", scheme, endpoint, c.bucket, objectName)
	return url, nil
}

// DeleteAvatar removes an avatar object from the bucket.
func (c *Client) DeleteAvatar(ctx context.Context, objectName string) error {
	return c.mc.RemoveObject(ctx, c.bucket, objectName, minio.RemoveObjectOptions{})
}
