// @Author: abbeymart | Abi Akindele | @Created: 2020-12-23 | @Updated: 2020-12-23
// @Company: mConnect.biz | @License: MIT
// @Description: application [system] constants

// Package appConstants publish default application/system constants
package appConstants

// Default application[system] constants
const (
	ApiTimeout         = 10 * 1000               // 10 seconds, in milliseconds
	LoginTimeout       = 24 * 60 * 60 * 1000     // 24 hours, in milliseconds
	StateTimeout       = 30 * 60 * 1000          // 30 minutes, in milliseconds
	RememberTimeout    = 5 * 24 * 60 * 60 * 1000 // 5 days, in milliseconds
	PageCacheExpire    = 30 * 60 * 1000          // 30 minutes, in milliseconds
	DbCacheExpire      = 5 * 60 * 1000           // 5 minutes, in milliseconds
	ShortDesc          = 20                      // length/slice of a long string to return
	MaxFileCount       = 20
	MaxProductQty      = 20
	MaxQueryLimit      = 100000
	MaxFileSize        = 10485760 // in bytes
	DefaultLanguage    = "en-US"
	DefaultCurrency    = "usd"
	ReadLogType        = "read"
	CreateLogType      = "create"
	SearchLogType      = "search"
	UpdateLogType      = "update"
	RemoveLogType      = "remove"
	DeleteLogType      = "delete"
	LoginLogType       = "login"
	LogoutLogType      = "logout"
	LogAll             = false
	LogCrud            = false
	LogCreate          = true
	LogUpdate          = true
	LogDelete          = true
	LogRead            = false
	LogLogin           = true
	LogLogout          = false
	LogSearch          = false
	DefaultCart        = "Cart"
	DefaultWish        = "Wish List"
	LoginRequired      = false
	RecursiveDelete    = false
	PasswordMinLength  = 6
	LoginNameMinLength = 6
	LoginMaxRetry      = 5
	LoginLockoutTime   = 15 * 60 * 1000 // 15 minutes, in milliseconds
)

// LoginServices function returns supported third-party login services
func LoginServices() []string {
	return []string{"google", "facebook", "twitter"}
}

// AllowedDocTypes function returns the supported document types for upload
func AllowedDocTypes() []string {
	return []string{
		"text/plain", "image/jpeg", "image/jpg", "image/png", "image/gif",
		"audio/wav", "audio/mpeg", "application/pdf", "application/msword",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.ms-powerpoint",
	}
}

func PermittedUserChangeRequests() []string {
	return []string{
		"verifyRegistration",
		"resetPassword",
		"changeEmail",
		"changeUsername",
		"getUsername",
	}
}

func AppCurrencies() []string {
	return []string{
		"USD",
		"CAD",
		"FFR",
		"NGN",
	}
}

func ContentAccess() []string {
	return []string{
		"internal",
		"private",
		"public",
	}
}

func ContentTypes() []string {
	return []string{
		"document",
		"news",
		"page",
		"post",
	}
}
