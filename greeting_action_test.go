// +build feature spec

package moocro

import "fmt"

type nameRequest struct {
	Name string `json:"name"`
}

type greetingRequest struct {
	Greeting string `json:"greeting"`
}

type greetingResponse struct {
	Greeting string `json:"greeting"`
}

type greetingAction struct {
}

var (
	greetingFoundActionRequest *greetingRequest
)

func (a *greetingAction) Perform(request *Request, response Response) error {
	helloRequest := request.Body.(*nameRequest)
	greeting := fmt.Sprintf("Hello %s", helloRequest.Name)
	greetingResponse := &greetingResponse{Greeting: greeting}

	return response.Write(greetingResponse)
}

type findGreetingAction struct {
}

func (a *findGreetingAction) Perform(request *Request, response Response) error {
	helloRequest := request.Body.(*nameRequest)
	greeting := fmt.Sprintf("Hello %s", helloRequest.Name)
	greetingResponse := &greetingResponse{Greeting: greeting}

	return response.WritePath(greetingFoundPath, greetingResponse)
}

type greetingFoundAction struct {
}

func (a *greetingFoundAction) Perform(request *Request, response Response) error {
	defer waitGroup.Done()

	greetingFoundActionRequest = request.Body.(*greetingRequest)

	return nil
}
