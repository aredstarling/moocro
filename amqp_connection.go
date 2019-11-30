package moocro

import (
	"os"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/streadway/amqp"
)

type amqpConnection struct {
	consumer  *amqp.Connection
	publisher *amqp.Connection
	*Options
}

func createAMQPConnection(options *Options) (*amqpConnection, error) {
	consumer, err := connect()
	if err != nil {
		return nil, err
	}

	publisher, err := connect()
	if err != nil {
		options.Application.Close(consumer)

		return nil, err
	}

	newConnection := &amqpConnection{consumer: consumer, publisher: publisher, Options: options}

	return newConnection, nil
}

func (c *amqpConnection) consumerChannel() (*amqp.Channel, error) {
	return c.consumer.Channel()
}

func (c *amqpConnection) publisherChannel() (*amqp.Channel, error) {
	return c.publisher.Channel()
}

// Close the amqpConnection
func (c *amqpConnection) Close() error {
	var result error

	if err := c.publisher.Close(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := c.consumer.Close(); err != nil {
		result = multierror.Append(result, err)
	}

	return result
}

func connect() (*amqp.Connection, error) {
	return amqp.Dial(os.Getenv(os.Getenv("AMQP_PROVIDER")))
}
