package http_response_helper

import "errors"

var (
	ErrNotFound            = errors.New("Not Found")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrBadRequset          = errors.New("Bad Requset")
	ErrForbidden           = errors.New("Forbidden")
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotImplemented      = errors.New("Not Implemented")
	ErrServiceUnavailable  = errors.New("Service Unavailable")
	ErrGatewayTimeout      = errors.New("Gateway Timeout")
)
