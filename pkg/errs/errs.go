/**
 * Created by zc on 2020/6/6.
 */
package errs

import "github.com/pkg/errors"

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = errors.New("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = errors.New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = errors.New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = errors.New("Not Found")

	// ErrNotImplemented is returned when an endpoint is not implemented.
	ErrNotImplemented = errors.New("Not Implemented")
)
