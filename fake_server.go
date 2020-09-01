package moocro

import (
	"fmt"

	"gitlab.com/lyticaa-public/disclosure"
)

// FakeServer for tetsing
type FakeServer struct {
	application *disclosure.Application
	*FakePathContainer
	RouteContainer
}

// CreateFakeServer for testing
func CreateFakeServer(container *FakePathContainer, application *disclosure.Application) Server {
	server := &FakeServer{
		FakePathContainer: container, RouteContainer: CreateRouteContainer(),
		application: application,
	}
	return server
}

// Start the server
func (s *FakeServer) Start() error {
	return nil
}

// Stop the server
func (s *FakeServer) Stop() error {
	return nil
}

// PerformAction that is configured
func (s *FakeServer) PerformAction(path string, body interface{}) error {
	route := s.Find(path)
	if route == nil {
		return fmt.Errorf("Could not find %s", path)
	}

	return s.application.Trace(path, func(tracePoint *disclosure.TracePoint) error {
		return route.CreateAction(tracePoint).Perform(&Request{Body: body, Path: path}, CreateFakeResponse(s.FakePathContainer))
	})
}
