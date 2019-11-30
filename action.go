package moocro

const (
	// SuccessStatus of the action response
	SuccessStatus = "success"

	// ErrorStatus of the action response
	ErrorStatus = "error"
)

// Action gives us the ability to respond to a request.
type Action interface {
	Perform(request *Request, response Response) error
}

// ActionResponse is a common response between services.
type ActionResponse struct {
	Status string    `json:"status"`
	Errors *[]string `json:"errors,omitempty"`
}

// CreateSuccessActionResponse for communication between services.
func CreateSuccessActionResponse() *ActionResponse {
	return &ActionResponse{Status: SuccessStatus}
}

// CreateErrorActionResponse for communication between services.
func CreateErrorActionResponse(err error) *ActionResponse {
	return &ActionResponse{Status: ErrorStatus, Errors: &[]string{err.Error()}}
}

// IsSuccess the action response
func (a *ActionResponse) IsSuccess() bool {
	if a == nil {
		return false
	}

	return a.Status == SuccessStatus
}

// IsError the action response
func (a *ActionResponse) IsError() bool {
	if a == nil {
		return true
	}

	return a.Status == ErrorStatus
}
