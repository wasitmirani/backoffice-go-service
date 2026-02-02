package redis

// Client interface for Redis operations
type Client interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration int) error
	Delete(key string) error
}

// TODO: Implement Redis client

