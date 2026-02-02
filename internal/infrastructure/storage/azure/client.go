package azure

// Client interface for Azure Storage operations
type Client interface {
	Upload(container, filename string, data []byte) error
	Download(container, filename string) ([]byte, error)
	Delete(container, filename string) error
}

// TODO: Implement Azure Storage client

