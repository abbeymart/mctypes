// @Author: abbeymart | Abi Akindele | @Created: 2020-12-23 | @Updated: 2020-12-23
// @Company: mConnect.biz | @License: MIT
// @Description: app messages

package appMessages

// Standard/Default messages for task-responses and/or errors
const (
	LoginError            = "Login audit error"
	LogoutError           = "Logout audit error"
	AuthErrorMessage      = "Authorization error"
	ConfirmCashPayment    = "Confirm cash-payment completed / received, to complete order - It cannot be undone"
	ConfirmInvoicePayment = "Confirm invoice-payment (pay-on-delivery, cheque, wire transfer), to complete order"
	ConfirmDelete         = "Do you really want to permanently delete this record? It cannot be undone"
	ConfirmMsgDelete      = "Do you really want to delete this message?"
	ConfirmMsgUndelete    = "Do you really want to un-delete/recover this message?"
	ItemMissing           = "Item not selected or missing, please try again"
	UpdateDenied          = "You are not authorized / permitted to update the record"
	DocDuplicate          = "Record exist, record must be unique"
	DocExists             = "No record created/changed, record with the same name exists"
	InsertSuccess         = "New record saved successfully"
	UpdateStatus          = "Record updated successfully"
	ValidateInputs        = "Required information must be provided in the required format"
	SubItems              = "Record contains sub-items and may not be removed"
	RecordDeleted         = "Record removed successfully"
	RemoveDenied          = "You are not authorized / permitted to remove this record / item"
	NoRecord              = "Record does not exist"
	LangCodeFormat        = "Language code should be of the format en-US"
	NameFormat            = "Name should contains letters, words and spaces only"
	PostalCodeFormat      = "Postal code contains letters, numbers and spaces only"
	NumberFormat          = "Only numbers (0-9) are permitted"
	FloatFormat           = "Currency / Float: only numbers (0-9) and . or , are permitted"
	PhoneFormat           = "Only numbers (0-9), -, (, and ) are permitted"
	BnFormat              = "Only numbers (0-9) and - are permitted"
	TitleFormat           = "Title should contains words/letters, spaces and hyphens only"
	DescFormat            = "Description contains unacceptable inputs... include permitted formats only"
	CodeFormat            = "Code should contain letters and numbers only"
	CurrencyFormat        = "Currency format is required= $, #, CA$, US$..."
	PathFormat            = "Path (tag) should contain letters and numbers only"
	UrlFormat             = "URL should contain alphanumeric only"
	UsernameFormat        = "Username should contains character/letters only"
	LangRequired          = "Language is required"
	NameRequired          = "Name ( is required - cannot be empty"
	DescRequired          = "Description is required - cannot be empty"
	CatRequired           = "Category is required - cannot be empty"
	CodeRequired          = "Code is required - cannot be empty"
	InfoRequired          = "Information item is required"
	TypeRequired          = "Type is required"
	GroupRequired         = "Group is required"
	NumberRequired        = "Number (0-9) is required - cannot be empty"
	TitleRequired         = "Title is required - cannot be empty"
	EmailRequired         = "Valid Email is required"
	EmailMessage          = "Valid Email should be in the correct format"
	LoginRequired         = "Valid Email or Username is required"
	PasswordRequired      = "Password is required"
	PasswordMessage       = "Password should include at least an upper-case and special character(@, #, $...)"
	PasswordConfirm       = "Password and confirm-Password must be the same"
	ConfirmPassRequired   = "Confirmed password is required"
	StdCodeRequired       = "Standard code name may contains letters, numbers, spaces and hyphens only"
	ContentErrorMessage   = "Content information should include valid title, category type and language"
	UpdateErrorMessage    = "Unable to update record - try again or contact system administrator"
	InsertErrorMessage    = "Unable to create new record  - try again or contact system administrator"
	RemoveErrorMessage    = "Unable to remove record  - try again or contact system administrator"
	SearchErrorMessage    = "Unable to retrieve record(s)  - try again or contact system administrator"
	ParamErrorMessage     = "Error validating the inputs - information required in the acceptable formats"
	UpdateNotAuthorized   = "You are not authorized / permitted to update the record"
	NoUpdate              = "Update not completed - current record ID is unknown or not provided"
	NoInfo                = "Information not available... no data/records available"
	AccessKeyMissing      = "Access-key missing or expired. Please login again."
	OrderMsgSubject       = "Order Details"
	CancelMsgSubject      = "Order / Order-Item(s) Cancelled"
	ReturnMsgSubject      = "Order-Item(s) Returned"
	PayMsgSubject         = "Payment Made"
	ShipMsgSubject        = "Order/Order-Items Shipped"
	RefundMsgSubject      = "Order Payment Refund"
	GetRequired           = "GET method is acceptable/required for this action/task"
	PostRequired          = "POST method is acceptable/required for this action/task"
	GetPostRequired       = "GET or POST method is acceptable/required for this action/task"
	CancelTypeOrder       = "Order - All Items"
	CancelTypeItem        = "Item(s) - Part of Order"
	IsStringAlpha         = "Parameter should contains alphanumeric (a-z, A-Z, 0-9, -, _) only"
	IsObject              = "Parameter should be object ({})"
	IsArray               = "Parameter should be array ([])"
	UsernameExists        = "Username exists, a unique username is required"
	EmailExists           = "Email exists, a unique email is required"
	NoPageInfo            = "Information is not available for this page"
)
