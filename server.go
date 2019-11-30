package moocro

import "errors"

var (
	// ErrStartServer indicates the server had an issue starting.
	ErrStartServer = errors.New("could not start server")
)

// Server defines an inteface for all servers
type Server interface {
	Route(route Route)
	Start() error
	Stop() error
}
