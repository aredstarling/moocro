package moocro

type amqpResponse struct {
	channel       *amqpChannel
	correlationID string
	replyTo       string
	*Options
}

func (c *amqpChannel) createAMQPResponse(correlationID string, replyTo string) Response {
	response := &amqpResponse{channel: c, correlationID: correlationID, replyTo: replyTo, Options: c.Options}

	return response
}

// IsFinished processing messages in the path
func (r *amqpResponse) IsFinished(path string) (bool, error) {
	q, err := r.channel.inspect(path)
	if err != nil {
		return false, err
	}

	return q.Messages == 0, nil
}

// WritePath a body to amqp
func (r *amqpResponse) WritePath(path string, body interface{}) error {
	return r.channel.publish(path, body, r.correlationID, r.replyTo)
}

// Write a body to amqp
func (r *amqpResponse) Write(body interface{}) error {
	return r.WritePath(r.replyTo, body)
}
