package util

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	// Success status
	Success codes.Code = 200
	//SuccessCreated status
	SuccessCreated codes.Code = 201
	// SuccessNoContent status
	SuccessNoContent codes.Code = 204
	// InvalidArgument status
	InvalidArgument codes.Code = 400
	// Unauthorized status
	Unauthorized codes.Code = 401
	// Forbidden status
	Forbidden codes.Code = 403
	// NotFound status
	NotFound codes.Code = 404
	// Cancelled status
	Cancelled codes.Code = 405
	// RequestTimeout status
	RequestTimeout codes.Code = 408

	// InactiveAccount status
	InactiveAccount codes.Code = 410
	// InvalidToken status
	InvalidToken codes.Code = 411
	// InvalidAPIKey status
	InvalidAPIKey codes.Code = 412
	// InvalidSession status
	InvalidSession codes.Code = 413
	// ResourceExhausted status
	ResourceExhausted codes.Code = 414

	// InvalidSubdomain status
	InvalidSubdomain codes.Code = 420
	// InactiveSubdomain status
	InactiveSubdomain codes.Code = 421
	// SuspendedSubdomain status
	SuspendedSubdomain codes.Code = 422

	// InvalidTransaction status
	InvalidTransaction codes.Code = 430
	// DuplicateTransaction status
	DuplicateTransaction codes.Code = 431

	// InternalError status
	InternalError codes.Code = 500
	// ProcessingError status
	ProcessingError codes.Code = 502
)

// HTTPStatusFromCode return HTTP Status for each code
func HTTPStatusFromCode(c codes.Code) int {
	switch c {
	case Success:
		return http.StatusOK
	case SuccessCreated:
		return http.StatusCreated
	case SuccessNoContent:
		return http.StatusOK
	case InvalidArgument:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case Cancelled:
		return http.StatusRequestTimeout
	case RequestTimeout:
		return http.StatusRequestTimeout
	case InactiveAccount:
		return http.StatusUnauthorized
	case InvalidToken:
		return http.StatusUnauthorized
	case InvalidAPIKey:
		return http.StatusUnauthorized
	case InvalidSession:
		return http.StatusUnauthorized
	case ResourceExhausted:
		return http.StatusTooManyRequests
	case InvalidSubdomain:
		return http.StatusNotFound
	case InactiveSubdomain:
		return http.StatusNotFound
	case SuspendedSubdomain:
		return http.StatusForbidden
	case InvalidTransaction:
		return http.StatusBadRequest
	case DuplicateTransaction:
		return http.StatusConflict
	case ProcessingError:
		return http.StatusInternalServerError
	case InternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
