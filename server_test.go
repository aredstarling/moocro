// +build feature spec

package moocro

import (
	"sync"

	"gitlab.com/lyticaa-public/disclosure"
	"gitlab.com/lyticaa-public/golog"
)

const (
	greetingPath      = "Moocro/Greeting"
	findGreetingPath  = "Moocro/FindGreeting"
	greetingFoundPath = "Moocro/GreetingFound"
)

var (
	application *disclosure.Application
	logger      = golog.CreateNull()
	options     *Options
	serializer  = CreateJSONSerializer()
	waitGroup   = &sync.WaitGroup{}
)

func init() {
	var err error

	application, err = disclosure.CreateApplication("TEST_APP", logger)
	if err != nil {
		panic(err)
	}

	options = &Options{Application: application, Logger: logger, Serializer: serializer}
}
