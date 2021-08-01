package httperrors

import (
	"net/http"
)

// ConvertToHttpError is a util function which helps to convert an error to HTTP error if possible,
// fallbacks to internal server error when can't
func ConvertToHttpError(err error) *HttpError {
	if errVal, ok := err.(HttpErrorType); ok {
		return errVal.ConvertToHttpError()
	} else {
		return InternalServerError(err.Error(), nil)
	}
}

// BadRequestError represents HTTP 400 error
func BadRequestError(description string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode:  http.StatusBadRequest,
		Description: description,
		Details:     details,
	}
}

// UnauthenticatedError represents HTTP 401 error
func UnauthenticatedError(description string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode:  http.StatusUnauthorized,
		Description: description,
		Details:     details,
	}
}

// PermissionDeniedError represents HTTP 403 error
func PermissionDeniedError(description string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode:  http.StatusForbidden,
		Description: description,
		Details:     details,
	}
}

// ResourceNotFoundError represents HTTP 404 error
func ResourceNotFoundError(description string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode:  http.StatusNotFound,
		Description: description,
		Details:     details,
	}
}

// MethodNotImplementedError represents HTTP 405 error
func MethodNotImplementedError(description string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode:  http.StatusNotImplemented,
		Description: description,
		Details:     details,
	}
}

// ResourceAlreadyExistsError represents HTTP 409 error
func ResourceAlreadyExistsError(description string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode:  http.StatusConflict,
		Description: description,
		Details:     details,
	}
}

// InternalServerError represents HTTP 500 error
func InternalServerError(description string, details interface{}) *HttpError {
	return &HttpError{
		StatusCode:  http.StatusInternalServerError,
		Description: description,
		Details:     details,
	}
}
