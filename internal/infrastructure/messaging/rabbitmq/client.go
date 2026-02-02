package rabbitmq

// Client interface for RabbitMQ operations
type Client interface {
	Publish(exchange, routingKey string, message []byte) error
	Consume(queue string, handler func(message []byte) error) error
	Close() error
}

// TODO: Implement RabbitMQ client

