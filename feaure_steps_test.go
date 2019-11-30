// +build feature

package moocro

import "github.com/sellernomics/disclosure"

type feature struct {
	client Client
	server Server
}

func createFeature(client Client, server Server) *feature {
	return &feature{client, server}
}

func (f *feature) setUp(interface{}) {
	routes := []Route{
		CreateSimpleRoute(greetingPath,
			func() interface{} { return &nameRequest{} },
			func(tracePoint *disclosure.TracePoint) Action { return &greetingAction{} }),
		CreateSimpleRoute(findGreetingPath,
			func() interface{} { return &nameRequest{} },
			func(tracePoint *disclosure.TracePoint) Action { return &findGreetingAction{} }),
		CreateSimpleRoute(greetingFoundPath,
			func() interface{} { return &greetingRequest{} },
			func(tracePoint *disclosure.TracePoint) Action { return &greetingFoundAction{} }),
	}

	for _, route := range routes {
		f.server.Route(route)
	}

	go func() {
		if err := f.server.Start(); err != nil {
			logger.FatalError("Could not start server!", err)
		}
	}()
}
