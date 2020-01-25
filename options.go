package moocro

import (
	"errors"

	"gitlab.com/getlytica/disclosure"
	"gitlab.com/getlytica/golog"
)

var (
	errEmptyOptions            = errors.New("empty options")
	errEmptyOptionsApplication = errors.New("empty options application")
	errEmptyOptionsLogger      = errors.New("empty options logger")
	errEmptyOptionsSerializer  = errors.New("empty options serializer")
)

// Options used for the server and client
type Options struct {
	Application *disclosure.Application
	Logger      golog.Logger
	Serializer  Serializer
}

func (o *Options) valid() error {
	switch {
	case o == nil:
		return errEmptyOptions
	case o.Application == nil:
		return errEmptyOptionsApplication
	case o.Logger == nil:
		return errEmptyOptionsLogger
	case o.Serializer == nil:
		return errEmptyOptionsSerializer
	default:
		return nil
	}
}
