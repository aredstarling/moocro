// +build feature

package moocro

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

type fakeFeature struct {
	container *FakePathContainer
	*feature
}

func (f *fakeFeature) sendGreeting(name string) error {
	server := f.server.(*FakeServer)

	return server.PerformAction(greetingPath, &nameRequest{Name: name})
}

func (f *fakeFeature) receiveGreeting(greeting string) error {
	server := f.server.(*FakeServer)

	responses, err := server.WrittenRPC()
	if err != nil {
		return err
	}

	return f.compareGreetingResponse(greeting, responses)
}

func (f *fakeFeature) sendFindGreeting(name string) error {
	server := f.server.(*FakeServer)

	return server.PerformAction(findGreetingPath, &nameRequest{Name: name})
}

func (f *fakeFeature) receiveGreetingFound(greeting string) error {
	server := f.server.(*FakeServer)

	responses, err := server.WrittenPath(greetingFoundPath)
	if err != nil {
		return err
	}

	return f.compareGreetingResponse(greeting, responses)
}

func (f *fakeFeature) compareGreetingResponse(greeting string, responses []interface{}) error {
	if len(responses) > 1 {
		return fmt.Errorf("Only expected one")
	}

	response := responses[0].(*greetingResponse)

	if greeting != response.Greeting {
		return fmt.Errorf("Expected %s to equal %s", greeting, response.Greeting)
	}

	return nil
}

func FakeContext(s *godog.Suite) {
	container := CreateFakePathContainer()
	server := CreateFakeServer(container, application)
	client := CreateFakeClient(container)

	f := &fakeFeature{container: container, feature: createFeature(client, server)}

	s.BeforeScenario(func(interface{}) { container.Clear() })
	s.BeforeScenario(f.setUp)

	s.Step(`^the fake system sends a name of "([^"]*)" to the greeting action$`, f.sendGreeting)
	s.Step(`^the fake system should send back "([^"]*)"$`, f.receiveGreeting)

	s.Step(`^the fake system sends a name of "([^"]*)" to the find greeting action$`, f.sendFindGreeting)
	s.Step(`^the fake system should respond "([^"]*)" as a greeting found$`, f.receiveGreetingFound)
}
