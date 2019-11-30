// +build feature

package moocro

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

type amqpFeature struct {
	response *greetingResponse
	*feature
}

func (f *amqpFeature) sendGreeting(name string) error {
	response, err := f.client.Write(greetingPath, &nameRequest{Name: name}, f.response)
	if err != nil {
		return err
	}

	f.response = response.(*greetingResponse)

	return nil
}

func (f *amqpFeature) receiveGreeting(greeting string) error {
	if greeting != f.response.Greeting {
		return fmt.Errorf("Expected %s to equal %s", greeting, f.response.Greeting)
	}

	return nil
}

func (f *amqpFeature) sendFindGreeting(name string) error {
	waitGroup.Add(1)

	return f.client.WritePath(findGreetingPath, &nameRequest{Name: name})
}

func (f *amqpFeature) receiveGreetingFound(greeting string) error {
	waitGroup.Wait()

	if greeting != greetingFoundActionRequest.Greeting {
		return fmt.Errorf("Expected %s to equal %s", greeting, greetingFoundActionRequest.Greeting)
	}

	return nil
}

func AMQPContext(s *godog.Suite) {
	server, err := CreateAMQPServer(options)
	application.FailOnError(err)

	client, err := CreateAMQPClient(options)
	application.FailOnError(err)

	f := &amqpFeature{response: &greetingResponse{}, feature: createFeature(client, server)}

	s.BeforeScenario(f.setUp)

	s.Step(`^the AMQP system sends a name of "([^"]*)" to the greeting action$`, f.sendGreeting)
	s.Step(`^the AMQP system should send back "([^"]*)"$`, f.receiveGreeting)

	s.Step(`^the AMQP system sends a name of "([^"]*)" to the find greeting action$`, f.sendFindGreeting)
	s.Step(`^the AMQP system should respond "([^"]*)" as a greeting found$`, f.receiveGreetingFound)
}
