package http_response_helper

import "errors"

var (
	// ErrNotFound Api Not found error
	ErrNotFound = errors.New("Not Found")
	// ErrUnauthorized ..
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrBadRequset ..
	ErrBadRequset = errors.New("Bad Requset")
	// ErrForbidden ..
	ErrForbidden = errors.New("Forbidden")
	// ErrInternalServerError ..
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotImplemented ..
	ErrNotImplemented = errors.New("Not Implemented")
	// ErrServiceUnavailable ..
	ErrServiceUnavailable = errors.New("Service Unavailable")
	// ErrGatewayTimeout ..
	ErrGatewayTimeout = errors.New("Gateway Timeout")
	// ErrUnknowError ..
	ErrUnknowError = errors.New("Unknow Error")
	// ErrDatabaseError ..
	ErrDatabaseError = errors.New("Database Error")
	// ErrConnectionTimeOut ..
	ErrConnectionTimeOut = errors.New("Connection Timeout")
	// ErrConnectionError ..
	ErrConnectionError = errors.New("Connection Error")
	// ErrServerBusy ..
	ErrServerBusy = errors.New("Server Busy")

	// ErrRecordNotFound record not found error, happens when haven't find any matched data when looking up with a struct
	ErrRecordNotFound = errors.New("record not found")
	// ErrInvalidSQL invalid SQL error, happens when you passed invalid SQL
	ErrInvalidSQL = errors.New("invalid SQL")
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("no valid transaction")
	// ErrCantStartTransaction can't start transaction when you are trying to start one with `Begin`
	ErrCantStartTransaction = errors.New("can't start transaction")
	// ErrUnaddressable unaddressable value
	ErrUnaddressable = errors.New("using unaddressable value")
)
