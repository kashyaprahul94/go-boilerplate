package users

import (
	"errors"
	"fmt"

	"github.com/kashyaprahul94/go-boilerplate/internal/httperrors"
)

// Define all possible error types for user domain
var (
	ErrUserNotFound      = errors.New("USER_ERROR__NOT_FOUND")
	ErrUserAlreadyExists = errors.New("USER_ERROR__ALREADY_EXISTS")
	ErrUserMiscError     = errors.New("USER_ERROR__MISC")
)

// UserDomainError struct is the type representing a user domain error with human readable description
type UserDomainError struct {
	// Unique user domain error code
	code error

	// Human readable description of the error
	description string

	// Extra details specific to error
	details interface{}
}

// Implement the `error` interface for `UserDomainError` type
func (u *UserDomainError) Error() string {
	return fmt.Sprintf("%s: %s", u.code, u.description)
}

// ConvertToHttpError is the util method which helps to convert a user domain error to equivalent http error
// Useful when the domain functionalities are exposed via HTTP server
func (u *UserDomainError) ConvertToHttpError() *httperrors.HttpError {
	if errors.Is(u.code, ErrUserNotFound) {
		return httperrors.ResourceNotFoundError(u.description, u.details)
	}

	if errors.Is(u.code, ErrUserAlreadyExists) {
		return httperrors.ResourceAlreadyExistsError(u.description, u.details)
	}

	if errors.Is(u.code, ErrUserMiscError) {
		return httperrors.InternalServerError(u.description, u.details)
	}

	return httperrors.InternalServerError(u.description, u.details)
}

// Generate error type for each possbile user domain error type

// UserNotFoundError represent that the provided user attributes wasn't available in the system
func UserNotFoundError(reason string, details interface{}) *UserDomainError {
	return &UserDomainError{
		code:        ErrUserNotFound,
		description: reason,
		details:     details,
	}
}

// UserAlreadyExistsError represents that the provided user attributes, conflict with existing user in the system
func UserAlreadyExistsError(reason string, details interface{}) *UserDomainError {
	return &UserDomainError{
		code:        ErrUserAlreadyExists,
		description: reason,
		details:     details,
	}
}

// UserMiscError represents any other non-specialized error that could occur
func UserMiscError(reason string, details interface{}) *UserDomainError {
	return &UserDomainError{
		code:        ErrUserMiscError,
		description: reason,
		details:     details,
	}
}
