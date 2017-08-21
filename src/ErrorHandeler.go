package http_response_helper

import "net/http"

func handlerError(err error) (int, int) {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound, ResultCodeNotFound
	case ErrBadRequset:
		return http.StatusBadRequest, ResultCodeBadRequset
	case ErrForbidden:
		return http.StatusForbidden, ResultCodeForbidden
	case ErrUnauthorized:
		return http.StatusUnauthorized, ResultCodeUnauthorized
	case ErrNotImplemented:
		return http.StatusNotImplemented, ResultCodeNotImplemented
	case ErrServiceUnavailable:
		return http.StatusServiceUnavailable, ResultCodeServiceUnavailable
	case ErrGatewayTimeout:
		return http.StatusGatewayTimeout, ResultCodeGatewayTimeout
	case nil:
		return http.StatusOK, ResultCodeSuccess
	default:
		return http.StatusInternalServerError, ResultCodeInternalServerError
	}
}
