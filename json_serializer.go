package moocro

import (
	"encoding/json"
	"errors"
)

var (
	// ErrMarshal when the response can't be marshaled
	ErrMarshal = errors.New("could not marshal response")

	// ErrUnmarshal when the request can't be unmarshaled
	ErrUnmarshal = errors.New("could not unmarshal request")
)

const (
	jsonContentType = "application/json"
)

type jsonSerializer struct {
}

// CreateJSONSerializer initializes a new JSON Serializer
func CreateJSONSerializer() Serializer {
	return &jsonSerializer{}
}

// ContentType of the JSONSerializer.
func (s *jsonSerializer) ContentType() string {
	return jsonContentType
}

// Marshal returns the JSON encoding of v. On error it returns ErrMarshal.
func (s *jsonSerializer) Marshal(v interface{}) ([]byte, error) {
	value, err := json.Marshal(v)
	if err != nil {
		return nil, ErrMarshal
	}

	return value, nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. On error it returns ErrUnmarshal.
func (s *jsonSerializer) Unmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return ErrUnmarshal
	}

	return nil
}
