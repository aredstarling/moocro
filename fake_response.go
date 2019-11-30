package moocro

// FakeResponse for testing
type FakeResponse struct {
	*FakePathContainer
}

// CreateFakeResponse for testing
func CreateFakeResponse(container *FakePathContainer) Response {
	return &FakeResponse{FakePathContainer: container}
}

// IsFinished processing messages in the path
func (r *FakeResponse) IsFinished(path string) (bool, error) {
	return len(r.writtenPaths[path]) == 0, nil
}

// Write a body for testing
func (r *FakeResponse) Write(body interface{}) error {
	return r.WritePath(defaultRPC, body)
}

// WritePath a body for testing
func (r *FakeResponse) WritePath(path string, body interface{}) error {
	r.writtenPaths[path] = append(r.writtenPaths[path], body)

	return nil
}
