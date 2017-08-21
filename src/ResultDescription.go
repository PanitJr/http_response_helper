package http_response_helper

// ResultDescription string
var ResultDescription = map[int]string{
	ResultCodeSuccess:                "Success",
	ResultCodeCreated:                "Created Success",
	ResultCodeAccessDenied:           "Access denied",
	ResultCodeExceededDataAllowances: "Exceeded data allowances",
	ResultCodeDataExisted:            "Data existed",
	ResultCodeForbidden:              "Missing or invalid parameter",
	ResultCodeBadRequset:             "Bad Request",
	ResultCodeUnauthorized:           "Unauthorized",
	ResultCodeNotFound:               "Data not found",
	ResultCodeInternalServerError:    "Internal Server Error",
	ResultCodeNotImplemented:         "Not Implemented",
	ResultCodeServiceUnavailable:     "Service Unavailable",
	ResultCodeGatewayTimeout:         "Gateway Timeout",
}
