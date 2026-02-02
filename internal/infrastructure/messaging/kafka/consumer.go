package kafka

// Consumer interface for Kafka message consumption
type Consumer interface {
	Consume(topic string, handler func(message []byte) error) error
	Close() error
}

// TODO: Implement Kafka consumer

