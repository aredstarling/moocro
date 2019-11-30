package moocro

import "github.com/sellernomics/disclosure"

// CreateRequestFun a factory for requests.
type CreateRequestFun func() interface{}

// CreateActionFunc a factory for actions.
type CreateActionFunc func(tracePoint *disclosure.TracePoint) Action

type simpleRoute struct {
	request CreateRequestFun
	action  CreateActionFunc
	path    string
}

// CreateSimpleRoute for a quick route
func CreateSimpleRoute(path string, request CreateRequestFun, action CreateActionFunc) Route {
	return &simpleRoute{path: path, request: request, action: action}
}

// CreateAction for the simple route
func (r *simpleRoute) CreateAction(tracePoint *disclosure.TracePoint) Action {
	if r.action == nil {
		return nil
	}

	return r.action(tracePoint)
}

// CreateRequest for the simple route
func (r *simpleRoute) CreateRequest() interface{} {
	if r.request == nil {
		return nil
	}

	return r.request()
}

// Path of the simple route
func (r *simpleRoute) Path() string {
	return r.path
}
