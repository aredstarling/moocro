package moocro

// Serializer defines an interface that will be used to support diffrent serialization formats
type Serializer interface {
	ContentType() string
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}
