package http_response_helper

const (
	ResultCodeSuccess                = 20000
	ResultCodeCreated                = 20100
	ResultCodeBadRequset             = 40000
	ResultCodeUnauthorized           = 40100
	ResultCodeAccessDenied           = 40101
	ResultCodeExceededDataAllowances = 40113
	ResultCodeForbidden              = 40300
	ResultCodeDataExisted            = 40301
	ResultCodeNotFound               = 40401
	ResultCodeInternalServerError    = 50000
	ResultCodeUnknowError            = 50060
	ResultCodeDatabaseError          = 50001
	ResultCodeConnectionTimeOut      = 50002
	ResultCodeConnectionError        = 50003
	ResultCodeNotImplemented         = 50100
	ResultCodeServerBusy             = 50300
	ResultCodeServiceUnavailable     = 50301
	ResultCodeGatewayTimeout         = 50400

	GormRecordNotFound       = 151
	GormInvalidSQL           = 101
	GormInvalidTransaction   = 178
	GormCantStartTransaction = 179
	GormUnaddressable        = 141

	LdapSuccess                     = 0
	LdapOperationsError             = 1
	LdapProtocolError               = 2
	LdapTimeLimitExceeded           = 3
	LdapSizeLimitExceeded           = 4
	LdapAuthMethodNotSupported      = 7
	LdapStrongerAuthRequired        = 8
	LdapAdminLimitExceeded          = 11
	LdapNoSuchAttribute             = 16
	LdapUndefinedAttributeType      = 17
	LdapInappropriateMatching       = 18
	LdapConstraintViolation         = 19
	LdapAttributeOrValueExists      = 20
	LdapInvalidAttributeSyntax      = 21
	LdapNoSuchObject                = 32
	LdapAliasProblem                = 33
	LdapInvalidDNSyntax             = 34
	LdapAliasDereferencingProblem   = 36
	LdapInappropriateAuthentication = 48
	LdapInvalidCredentials          = 49
	LdapInsufficientAccessRights    = 50
	LdapBusy                        = 51
	LdapUnavailable                 = 52
	LdapUnwillingToPerform          = 53
	LdapLoopDetect                  = 54
	LdapNamingViolation             = 64
	LdapObjectClassViolation        = 65
	LdapNotAllowedOnNonLeaf         = 66
	LdapNotAllowedOnRDN             = 67
	LdapEntryAlreadyExists          = 68
	LdapObjectClassModsProhibited   = 69
	LdapOther                       = 80
)
