package moocro

import (
	"gitlab.com/getlytica/disclosure"
)

// Route defines an inteface for all routes
type Route interface {
	CreateAction(tracePoint *disclosure.TracePoint) Action
	CreateRequest() interface{}
	Path() string
}
