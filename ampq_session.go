package moocro

import (
	"errors"

	"gitlab.com/getlytica/disclosure"
	"gitlab.com/getlytica/golog"
	"github.com/streadway/amqp"
)

var (
	errStoppedListening = errors.New("stopped listening to path")
)

type amqpSession struct {
	channel *amqpChannel
	*Options
}

func (s *amqpServer) createAMQPSession() (*amqpSession, error) {
	channel, err := s.connection.createAMQPChannel()
	if err != nil {
		return nil, err
	}

	return &amqpSession{channel: channel, Options: s.Options}, nil
}

func (s *amqpSession) register(route Route) error {
	attributes := golog.Attributes{"path": route.Path()}

	s.Logger.Info("Starting listening to path.", attributes)

	_, messages, err := s.channel.consumeDurable(route.Path())
	if err != nil {
		return err
	}

	for message := range messages {
		messageAttributes := golog.Attributes{"request": string(message.Body)}.Merge(attributes)

		s.Logger.Info("Received a request.", messageAttributes)
		err := s.performAction(route, message)

		if err != nil {
			s.Logger.Warn("Could not perform action.", golog.Attributes{"error": err}.Merge(messageAttributes))
		}
	}

	s.Logger.Warn("Stopped listening to path.", attributes)

	return errStoppedListening
}

func (s *amqpSession) performAction(route Route, message amqp.Delivery) error {
	return s.Application.Trace(route.Path(), func(tracePoint *disclosure.TracePoint) error {
		response := s.channel.createAMQPResponse(message.CorrelationId, message.ReplyTo)
		request := route.CreateRequest()

		if err := s.Serializer.Unmarshal(message.Body, request); err != nil {
			_ = message.Nack(false, false)

			return err
		}

		action := route.CreateAction(tracePoint)

		if err := action.Perform(&Request{Body: request, Path: route.Path()}, response); err != nil {
			_ = message.Nack(false, true)

			return err
		}

		return message.Ack(false)
	})
}

// Close the amqpSession
func (s *amqpSession) Close() error {
	return s.channel.Close()
}
