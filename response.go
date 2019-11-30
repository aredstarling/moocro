package moocro

// Response describes an inteface to write to
type Response interface {
	IsFinished(path string) (bool, error)
	Write(body interface{}) error
	WritePath(path string, body interface{}) error
}
