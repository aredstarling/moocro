package moocro

import (
	"gitlab.com/lyticaa-public/disclosure"
)

// Route defines an inteface for all routes
type Route interface {
	CreateAction(tracePoint *disclosure.TracePoint) Action
	CreateRequest() interface{}
	Path() string
}
