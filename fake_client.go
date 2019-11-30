package moocro

// FakeClient for tetsing
type FakeClient struct {
	*FakePathContainer
}

// CreateFakeClient for testing
func CreateFakeClient(container *FakePathContainer) Client {
	return &FakeClient{FakePathContainer: container}
}

// Close the FakeClient
func (c *FakeClient) Close() error {
	return nil
}

// IsFinished processing messages in the path
func (c *FakeClient) IsFinished(path string) (bool, error) {
	return len(c.writtenPaths[path]) == 0, nil
}

// Write a body for testing
func (c *FakeClient) Write(path string, body interface{}, response interface{}) (interface{}, error) {
	return response, c.WritePath(path, body)
}

// WritePath a body for testing
func (c *FakeClient) WritePath(path string, body interface{}) error {
	c.writtenPaths[path] = append(c.writtenPaths[path], body)

	return nil
}
