package s3

// Client interface for S3 operations
type Client interface {
	Upload(bucket, key string, data []byte) error
	Download(bucket, key string) ([]byte, error)
	Delete(bucket, key string) error
}

// TODO: Implement S3 client

