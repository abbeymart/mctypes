// @Author: abbeymart | Abi Akindele | @Created: 2020-11-30 | @Updated: 2020-11-30
// @Company: mConnect.biz | @License: MIT
// @Description: mcconfig: standard network codes and texts

package netStatus

const (
	/** RFC 7231, 6.2.1 */
	Continue = 100
	/** RFC 7231, 6.2.2 */
	SwitchingProtocols = 101
	/** RFC 2518, 10.1 */
	Processing = 102
	/** RFC 8297 **/
	EarlyHints = 103
	/** RFC 7231, 6.3.1 */
	OK = 200
	/** RFC 7231, 6.3.2 */
	Created = 201
	/** RFC 7231, 6.3.3 */
	Accepted = 202
	/** RFC 7231, 6.3.4 */
	NonAuthoritativeInfo = 203
	/** RFC 7231, 6.3.5 */
	NoContent = 204
	/** RFC 7231, 6.3.6 */
	ResetContent = 205
	/** RFC 7233, 4.1 */
	PartialContent = 206
	/** RFC 4918, 11.1 */
	MultiStatus = 207
	/** RFC 5842, 7.1 */
	AlreadyReported = 208
	/** RFC 3229, 10.4.1 */
	IMUsed = 226
	/** RFC 7231, 6.4.1 */
	MultipleChoices = 300
	/** RFC 7231, 6.4.2 */
	MovedPermanently = 301
	/** RFC 7231, 6.4.3 */
	Found = 302
	/** RFC 7231, 6.4.4 */
	SeeOther = 303
	/** RFC 7232, 4.1 */
	NotModified = 304
	/** RFC 7231, 6.4.5 */
	UseProxy = 305
	/** RFC 7231, 6.4.7 */
	TemporaryRedirect = 307
	/** RFC 7538, 3 */
	PermanentRedirect = 308
	/** RFC 7231, 6.5.1 */
	BadRequest = 400
	/** RFC 7235, 3.1 */
	Unauthorized = 401
	/** RFC 7231, 6.5.2 */
	PaymentRequired = 402
	/** RFC 7231, 6.5.3 */
	Forbidden = 403
	/** RFC 7231, 6.5.4 */
	NotFound = 404
	/** RFC 7231, 6.5.5 */
	MethodNotAllowed = 405
	/** RFC 7231, 6.5.6 */
	NotAcceptable = 406
	/** RFC 7235, 3.2 */
	ProxyAuthRequired = 407
	/** RFC 7231, 6.5.7 */
	RequestTimeout = 408
	/** RFC 7231, 6.5.8 */
	Conflict = 409
	/** RFC 7231, 6.5.9 */
	Gone = 410
	/** RFC 7231, 6.5.10 */
	LengthRequired = 411
	/** RFC 7232, 4.2 */
	PreconditionFailed = 412
	/** RFC 7231, 6.5.11 */
	RequestEntityTooLarge = 413
	/** RFC 7231, 6.5.12 */
	RequestURITooLong = 414
	/** RFC 7231, 6.5.13 */
	UnsupportedMediaType = 415
	/** RFC 7233, 4.4 */
	RequestedRangeNotSatisfiable = 416
	/** RFC 7231, 6.5.14 */
	ExpectationFailed = 417
	/** RFC 7168, 2.3.3 */
	Teapot = 418
	/** RFC 7540, 9.1.2 */
	MisdirectedRequest = 421
	/** RFC 4918, 11.2 */
	UnprocessableEntity = 422
	/** RFC 4918, 11.3 */
	Locked = 423
	/** RFC 4918, 11.4 */
	FailedDependency = 424
	/** RFC 8470, 5.2 */
	TooEarly = 425
	/** RFC 7231, 6.5.15 */
	UpgradeRequired = 426
	/** RFC 6585, 3 */
	PreconditionRequired = 428
	/** RFC 6585, 4 */
	TooManyRequests = 429
	/** RFC 6585, 5 */
	RequestHeaderFieldsTooLarge = 431
	/** RFC 7725, 3 */
	UnavailableForLegalReasons = 451

	/** RFC 7231, 6.6.1 */
	InternalServerError = 500
	/** RFC 7231, 6.6.2 */
	NotImplemented = 501
	/** RFC 7231, 6.6.3 */
	BadGateway = 502
	/** RFC 7231, 6.6.4 */
	ServiceUnavailable = 503
	/** RFC 7231, 6.6.5 */
	GatewayTimeout = 504
	/** RFC 7231, 6.6.6 */
	HTTPVersionNotSupported = 505
	/** RFC 2295, 8.1 */
	VariantAlsoNegotiates = 506
	/** RFC 4918, 11.5 */
	InsufficientStorage = 507
	/** RFC 5842, 7.2 */
	LoopDetected = 508
	/** RFC 2774, 7 */
	NotExtended = 510
	/** RFC 6585, 6 */
	NetworkAuthenticationRequired = 511
)

var StatusText = map[uint]string{
	Continue:                      "Continue",
	SwitchingProtocols:            "Switching Protocols",
	Processing:                    "Processing",
	EarlyHints:                    "Early Hints",
	OK:                            "OK",
	Created:                       "Created",
	Accepted:                      "Accepted",
	NonAuthoritativeInfo:          "Non-Authoritative Information",
	NoContent:                     "No Content",
	ResetContent:                  "Reset Content",
	PartialContent:                "Partial Content",
	MultiStatus:                   "Multi-Status",
	AlreadyReported:               "Already Reported",
	IMUsed:                        "IM Used",
	MultipleChoices:               "Multiple Choices",
	MovedPermanently:              "Moved Permanently",
	Found:                         "Found",
	SeeOther:                      "See Other",
	NotModified:                   "Not Modified",
	UseProxy:                      "Use Proxy",
	TemporaryRedirect:             "Temporary Redirect",
	PermanentRedirect:             "Permanent Redirect",
	BadRequest:                    "Bad Request",
	Unauthorized:                  "Unauthorized",
	PaymentRequired:               "Payment Required",
	Forbidden:                     "Forbidden",
	NotFound:                      "Not Found",
	MethodNotAllowed:              "Method Not Allowed",
	NotAcceptable:                 "Not Acceptable",
	ProxyAuthRequired:             "Proxy Authentication Required",
	RequestTimeout:                "Request Timeout",
	Conflict:                      "Conflict",
	Gone:                          "Gone",
	LengthRequired:                "Length Required",
	PreconditionFailed:            "Precondition Failed",
	RequestEntityTooLarge:         "Request Entity Too Large",
	RequestURITooLong:             "Request URI Too Long",
	UnsupportedMediaType:          "Unsupported Media Type",
	RequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	ExpectationFailed:             "Expectation Failed",
	Teapot:                        "I'm a teapot",
	MisdirectedRequest:            "Misdirected Request",
	UnprocessableEntity:           "Unprocessable Entity",
	Locked:                        "Locked",
	FailedDependency:              "Failed Dependency",
	TooEarly:                      "Too Early",
	UpgradeRequired:               "Upgrade Required",
	PreconditionRequired:          "Precondition Required",
	TooManyRequests:               "Too Many Requests",
	RequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	UnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	InternalServerError:           "Internal Server Error",
	NotImplemented:                "Not Implemented",
	BadGateway:                    "Bad Gateway",
	ServiceUnavailable:            "Service Unavailable",
	GatewayTimeout:                "Gateway Timeout",
	HTTPVersionNotSupported:       "HTTP Version Not Supported",
	VariantAlsoNegotiates:         "Variant Also Negotiates",
	InsufficientStorage:           "Insufficient Storage",
	LoopDetected:                  "Loop Detected",
	NotExtended:                   "Not Extended",
	NetworkAuthenticationRequired: "Network Authentication Required",
}
