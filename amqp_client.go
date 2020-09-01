 package moocro

import (
	"errors"
	"time"

	"gitlab.com/getlytica/golog"
	"github.com/streadway/amqp"
)

var (
	// ErrClientTimeout is returned when we have not received a response in time.
	ErrClientTimeout = errors.New("received a timeout")
)

type amqpClient struct {
	connection *amqpConnection
	*Options
}

type responseResult struct {
	response interface{}
	err      error
}

// CreateAMQPClient for messaging
func CreateAMQPClient(options *Options) (Client, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	connection, err := createAMQPConnection(options)
	if err != nil {
		return nil, err
	}

	client := &amqpClient{connection: connection, Options: options}

	return client, nil
}

// Close the AMQPClient
func (c *amqpClient) Close() error {
	return c.connection.Close()
}

// IsFinished processing messages in the path
func (c *amqpClient) IsFinished(path string) (bool, error) {
	channel, err := c.connection.createAMQPChannel()
	if err != nil {
		return false, err
	}

	defer c.Application.Close(channel)

	q, err := channel.inspect(path)
	if err != nil {
		return false, err
	}

	return q.Messages == 0, nil
}

func (c *amqpClient) Write(path string, body interface{}, response interface{}) (interface{}, error) {
	channel, err := c.connection.createAMQPChannel()
	if err != nil {
		return nil, err
	}

	defer c.Application.Close(channel)

	q, msgs, err := channel.consumeTransient()
	if err != nil {
		return response, err
	}

	correlationID, err := generateCorrelationID()
	if err != nil {
		return response, err
	}

	if err := channel.publish(path, body, correlationID, q.Name); err != nil {
		return response, err
	}

	responseChannel := make(chan responseResult, 1)

	go c.receiveMessage(path, correlationID, response, msgs, responseChannel)

	select {
	case res := <-responseChannel:
		return res.response, res.err
	case <-time.After(10 * time.Second):
		return response, ErrClientTimeout
	}
}

// WritePath a body to amqp
func (c *amqpClient) WritePath(path string, body interface{}) error {
	channel, err := c.connection.createAMQPChannel()
	if err != nil {
		return err
	}

	defer c.Application.Close(channel)

	correlationID, err := generateCorrelationID()
	if err != nil {
		return err
	}

	return channel.publish(path, body, correlationID, emptyString)
}

func (c *amqpClient) receiveMessage(path string, correlationID string, response interface{}, messages <-chan amqp.Delivery, responseChannel chan responseResult) {
	for message := range messages {
		if correlationID == message.CorrelationId {
			c.Logger.Info("Received message from path", golog.Attributes{"path": path, "body": string(message.Body), "correlationID": correlationID})

			if err := c.Serializer.Unmarshal(message.Body, response); err != nil {
				responseChannel <- responseResult{response: response, err: err}
			} else {
				responseChannel <- responseResult{response: response, err: nil}
			}

			break
		}
	}
}
