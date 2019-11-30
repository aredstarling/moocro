package moocro

import (
	"fmt"
)

const (
	defaultRPC = "Test/Write"
)

// FakePathContainer for testing
type FakePathContainer struct {
	writtenPaths map[string][]interface{}
}

// CreateFakePathContainer for testing
func CreateFakePathContainer() *FakePathContainer {
	return &FakePathContainer{writtenPaths: make(map[string][]interface{})}
}

// WrittenPath in the container
func (c *FakePathContainer) WrittenPath(path string) ([]interface{}, error) {
	data, ok := c.writtenPaths[path]
	if !ok {
		return nil, fmt.Errorf("Could not find %s", path)
	}

	return data, nil
}

// WrittenRPC in the container
func (c *FakePathContainer) WrittenRPC() ([]interface{}, error) {
	return c.WrittenPath(defaultRPC)
}

// Clear the container
func (c *FakePathContainer) Clear() {
	c.writtenPaths = make(map[string][]interface{})
}
