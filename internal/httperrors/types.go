package httperrors

import (
	"encoding/json"
	"errors"
)

// HttpError type represents an error type useful for HTTP based flows
type HttpError struct {
	// Status code useful for HTTP status code
	StatusCode int `json:"statusCode"`

	// Description is the human readable explanation of the error
	Description string `json:"description"`

	// Details represents error specific extra fields
	Details interface{} `json:"details,omitempty"`
}

// HttpErrorType interface specifies tha protocol to identify the errors which can be treated as HTTP errors
type HttpErrorType interface {
	// A method which should convert error to its equivalent HTTP based error
	ConvertToHttpError() *HttpError
}

// Implement the `error` interface for `HttpError` type
func (e *HttpError) Error() string {
	bytes, _ := json.Marshal(e.Description)
	return errors.New(string(bytes)).Error()
}

// Implement the `stringer` interface for `HttpError` type
func (e *HttpError) String() string {
	bytes, _ := json.Marshal(e)
	return string(bytes)
}
