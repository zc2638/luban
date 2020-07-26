/**
 * Created by zc on 2020/6/6.
 */
package errs

import "github.com/pkg/errors"

func New(message string) error {
	return errors.New(message)
}

type Error string

func (e Error) Error() string {
	return string(e)
}

func (e Error) With(err error) error {
	return errors.Wrap(err, e.Error())
}

const (
	// ErrBodyParse is returned when body parse error
	ErrBodyParse = Error("Body parameter format error")

	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = Error("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = Error("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = Error("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = Error("Not Found")

	// ErrNotImplemented is returned when an endpoint is not implemented.
	ErrNotImplemented = Error("Not Implemented")

	// ErrInvalidSpace is returned when the api request space is invalid.
	ErrInvalidSpace = Error("Invalid or missing space")

	// ErrInvalidConfig is returned when the api request config is invalid.
	ErrInvalidConfig = Error("Invalid or missing config")
)
