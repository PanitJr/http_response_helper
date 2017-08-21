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
	case ErrDatabaseError:
		return http.StatusGatewayTimeout, ResultCodeDatabaseError
	case ErrConnectionTimeOut:
		return http.StatusGatewayTimeout, ResultCodeConnectionTimeOut
	case ErrConnectionError:
		return http.StatusServiceUnavailable, ResultCodeConnectionError
	case ErrServerBusy:
		return http.StatusServiceUnavailable, ResultCodeServerBusy
	case ErrRecordNotFound:
		return http.StatusNotFound, GormRecordNotFound
	case ErrInvalidSQL:
		return http.StatusInternalServerError, GormInvalidSQL
	case ErrInvalidTransaction:
		return http.StatusInternalServerError, GormInvalidTransaction
	case ErrCantStartTransaction:
		return http.StatusInternalServerError, GormCantStartTransaction
	case ErrUnaddressable:
		return http.StatusInternalServerError, GormUnaddressable
	case nil:
		return http.StatusOK, ResultCodeSuccess
	default:
		return http.StatusInternalServerError, ResultCodeUnknowError
	}
}
