package moocro

import (
	"time"
)

var (
	emptyString = ""
)

type amqpServer struct {
	connection *amqpConnection
	RouteContainer
	*Options
}

// CreateAMQPServer for messaging
func CreateAMQPServer(options *Options) (Server, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	connection, err := createAMQPConnection(options)
	if err != nil {
		return nil, err
	}

	server := &amqpServer{connection: connection, Options: options, RouteContainer: CreateRouteContainer()}

	return server, nil
}

// Start the server
func (s *amqpServer) Start() error {
	for _, route := range s.All() {
		for i := 1; i <= concurrency(); i++ {
			go func(route Route) {
				if err := s.register(route); err != nil {
					s.Application.FailOnError(err)
				}
			}(route)
		}
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

// Stop the server
func (s *amqpServer) Stop() error {
	return s.connection.Close()
}

func (s *amqpServer) register(route Route) error {
	session, err := s.createAMQPSession()
	if err != nil {
		return err
	}

	defer s.Application.Close(session)

	err = session.register(route)
	if err == errStoppedListening {
		return s.register(route)
	}

	return err
}
