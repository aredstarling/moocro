package moocro

// Client defines an inteface for all clients
type Client interface {
	Close() error
	IsFinished(path string) (bool, error)
	Write(path string, body interface{}, response interface{}) (interface{}, error)
	WritePath(path string, body interface{}) error
}
