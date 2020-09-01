package moocro

import (
	multierror "github.com/hashicorp/go-multierror"
	"gitlab.com/lyticaa-public/golog"
	"github.com/streadway/amqp"
)

type amqpChannel struct {
	consumer  *amqp.Channel
	publisher *amqp.Channel
	*Options
}

func (c *amqpConnection) createAMQPChannel() (*amqpChannel, error) {
	consumer, err := c.consumerChannel()
	if err != nil {
		return nil, err
	}

	publisher, err := c.publisherChannel()
	if err != nil {
		c.Application.Close(consumer)

		return nil, err
	}

	newChannel := &amqpChannel{consumer: consumer, publisher: publisher, Options: c.Options}

	return newChannel, nil
}

func (c *amqpChannel) publish(path string, body interface{}, correlationID string, replyTo string) error {
	value, err := c.Serializer.Marshal(body)
	if err != nil {
		return err
	}

	publishing := amqp.Publishing{
		ContentType: c.Serializer.ContentType(), CorrelationId: correlationID,
		Body: value, ReplyTo: replyTo, DeliveryMode: amqp.Persistent,
	}

	attributes := golog.Attributes{"path": path, "body": string(value), "correlation_id": correlationID, "reply_to": replyTo}

	if err := c.publisher.Publish(emptyString, path, false, false, publishing); err != nil {
		c.Logger.Warn("Could not write to path", golog.Attributes{"error": err}.Merge(attributes))

		return err
	}

	c.Logger.Info("Wrote message to path", attributes)

	return nil
}

func (c *amqpChannel) inspect(path string) (amqp.Queue, error) {
	return c.publisher.QueueInspect(path)
}

func (c *amqpChannel) declareDurable(path string) (amqp.Queue, error) {
	return c.publisher.QueueDeclare(path, true, false, false, false, nil)
}

func (c *amqpChannel) declareTransient() (amqp.Queue, error) {
	return c.publisher.QueueDeclare(emptyString, false, true, false, false, nil)
}

func (c *amqpChannel) consumeDurable(path string) (*amqp.Queue, <-chan amqp.Delivery, error) {
	q, err := c.declareDurable(path)
	if err != nil {
		return nil, nil, err
	}

	messages, err := c.consumeQueue(&q, true, false)
	if err != nil {
		return nil, nil, err
	}

	return &q, messages, nil
}

func (c *amqpChannel) consumeTransient() (*amqp.Queue, <-chan amqp.Delivery, error) {
	q, err := c.declareTransient()
	if err != nil {
		return nil, nil, err
	}

	messages, err := c.consumeQueue(&q, false, true)
	if err != nil {
		return nil, nil, err
	}

	return &q, messages, nil
}

func (c *amqpChannel) consumeQueue(queue *amqp.Queue, durable, autoAck bool) (<-chan amqp.Delivery, error) {
	if err := c.consumer.Qos(1, 0, false); err != nil {
		return nil, err
	}

	messages, err := c.consumer.Consume(queue.Name, emptyString, autoAck, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Close the amqpChannel
func (c *amqpChannel) Close() error {
	var result error

	if err := c.publisher.Close(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := c.consumer.Close(); err != nil {
		result = multierror.Append(result, err)
	}

	return result
}
