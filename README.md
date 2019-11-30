# moocro

Our take to building distributed/asynchronous architectures

## Setup

The application has the following environment variables

* `export AMQP_PROVIDER=` - The AMQP provider
* `export MOOCRO_CONCURRENCY=` - The number of go routines per path

## Server and Client

The best way to get going is to look at the examples presented in following files:

* `server_test.go`
* `greeting_action_test.go`
* `amqp_steps_test.go`
* `fake_steps_test.go`
