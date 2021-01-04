// @Author: abbeymart | Abi Akindele | @Created: 2020-12-22 | @Updated: 2020-12-22
// @Company: mConnect.biz | @License: MIT
// @Description: crud operations' types

package mctypes

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleServiceType struct {
	ServiceId            string `json:"service_id"`
	RoleId               string `json:"role_id"`
	ServiceCategory      string `json:"service_category"`
	CanRead              bool   `json:"can_read"`
	CanCreate            bool   `json:"can_create"`
	CanUpdate            bool   `json:"can_update"`
	CanDelete            bool   `json:"can_delete"`
	TableAccessPermitted bool   `json:"table_access_permitted"`
}

type CheckAccessType struct {
	UserId       string            `json:"user_id" mcorm:"user_id"`
	Group        string            `json:"group" mcorm:"group"`
	Groups       []string          `json:"groups" mcorm:"groups"`
	IsActive     bool              `json:"is_active" mcorm:"is_active"`
	IsAdmin      bool              `json:"is_admin" mcorm:"is_admin"`
	RoleServices []RoleServiceType `json:"role_services" mcorm:"role_services"`
	TableId      string            `json:"table_id" mcorm:"table_id"`
}

type CheckAccessParamsType struct {
	AccessDb     *pgxpool.Pool `json:"access_db"`
	UserInfo     UserInfoType  `json:"user_info"`
	TableName    string        `json:"table_name"`
	RecordIds    []string      `json:"record_ids"` // for update, delete and read tasks
	AccessTable  string        `json:"access_table"`
	UserTable    string        `json:"user_table"`
	RoleTable    string        `json:"role_table"`
	ServiceTable string        `json:"service_table"`
}

type RoleFuncType func(it1 string, it2 RoleServiceType) bool
type FieldValueType interface{}
type ValueParamType map[string]interface{}
type ValueToDataType map[string]interface{}
type ActionParamsType []ValueParamType
type ExistParamType map[string]interface{}
type ExistParamsType []ExistParamType
type SortParamType map[string]int     // 1 for "asc", -1 for "desc"
type ProjectParamType map[string]bool // 1 or true for inclusion, 0 or false for exclusion

type QueryItemType struct {
	GroupItem      map[string]map[string]interface{} `json:"group_item"`       // key1 => fieldName, key2 => fieldOperator, interface{}=> value(s)
	GroupItemOrder int                               `json:"group_item_order"` // item/field order within the group
	GroupItemOp    string                            `json:"group_item_op"`    // group-item relationship to the next item (AND, OR), the last item groupItemOp should be "" or will be ignored
}

type QueryGroupType struct {
	GroupName   string          `json:"group_name"`    // for group-items(fields) categorization
	GroupItems  []QueryItemType `json:"group_items"`   // group items to be composed by category
	GroupOrder  int             `json:"group_order"`   // group order
	GroupLinkOp string          `json:"group_link_op"` // group relationship to the next group (AND, OR), the last group groupLinkOp should be "" or will be ignored
}

type QueryParamType []QueryGroupType
type WhereParamType []QueryGroupType

// CrudParamsType is the struct type for receiving, composing and passing CRUD inputs
type CrudParamsType struct {
	AppDb         *pgxpool.Pool    `json:"-"`
	TableName     string           `json:"-"`
	UserInfo      UserInfoType     `json:"user_info"`
	ActionParams  ActionParamsType `json:"action_params"`
	ExistParams   ExistParamsType  `json:"exist_params"`
	QueryParams   WhereParamType   `json:"query_params"`
	RecordIds     []string         `json:"record_ids"`
	ProjectParams ProjectParamType `json:"project_params"`
	SortParams    SortParamType    `json:"sort_params"`
	Token         string           `json:"token"`
	Skip          int              `json:"skip"`
	Limit         int              `json:"limit"`
	TaskName      string           `json:"-"`
}

type CrudOptionsType struct {
	ParentTables          []string
	ChildTables           []string
	RecursiveDelete       bool
	CheckAccess           bool
	AccessDb              *pgxpool.Pool
	AuditDb               *pgxpool.Pool
	ServiceDb             *pgxpool.Pool
	AuditTable            string
	ServiceTable          string
	UserTable             string
	RoleTable             string
	AccessTable           string
	VerifyTable           string
	MaxQueryLimit         int
	LogAll                bool
	LogCreate             bool
	LogUpdate             bool
	LogRead               bool
	LogDelete             bool
	LogLogin              bool
	LogLogout             bool
	UnAuthorizedMessage   string
	RecExistMessage       string
	CacheExpire           int
	LoginTimeout          int
	UsernameExistsMessage string
	EmailExistsMessage    string
	MsgFrom               string
}

type CrudParamType struct {
	AppDb           *pgxpool.Pool // use *pgxpool.Pool, preferred || *pgx.Conn
	TableName       string
	Token           string
	UserInfo        UserInfoType
	UserId          string
	Group           string
	Groups          []string
	RecordIds       []string
	ActionParams    ActionParamsType
	QueryParams     QueryParamType
	ExistParams     ExistParamsType
	ProjectParams   ProjectParamType
	SortParams      SortParamType
	Skip            int
	Limit           int
	ParentTables    []string
	ChildTables     []string
	RecursiveDelete bool
	CheckAccess     bool
	AccessDb        *pgxpool.Pool
	AuditDb         *pgxpool.Pool
	AuditTable      string
	ServiceTable    string
	UserTable       string
	RoleTable       string
	AccessTable     string
	MaxQueryLimit   int
	LogAll          bool
	LogCreate       bool
	LogUpdate       bool
	LogRead         bool
	LogDelete       bool
	//transLog AuditLog
	HashKey             string
	IsRecExist          bool
	ActionAuthorized    bool
	UnAuthorizedMessage string
	RecExistMessage     string
	IsAdmin             bool
	CreateItems         ActionParamsType
	UpdateItems         ActionParamsType
	CurrentRecs         ActionParamsType
	RoleServices        []RoleServiceType
	SubItems            []bool
	CacheExpire         int
	Params              CrudParamsType
}

// MongoDB specific types
type MongoCrudTaskType struct {
	AppDb         *mongo.Client
	TableName     string
	UserInfo      UserInfoType
	ActionParams  ActionParamsType
	ExistParams   ExistParamsType
	QueryParams   QueryParamType
	RecordIds     []string
	ProjectParams ProjectParamType
	SortParams    SortParamType
	Token         string
	Options       MongoCrudOptionsType
	TaskName      string
}

type MongoCrudOptionsType struct {
	Skip                  int
	Limit                 int
	ParentTables          []string
	ChildTables           []string
	RecursiveDelete       bool
	CheckAccess           bool
	AccessDb              *mongo.Client
	AuditDb               *mongo.Client
	ServiceDb             *mongo.Client
	AuditTable            string
	ServiceTable          string
	UserTable             string
	RoleTable             string
	AccessTable           string
	VerifyTable           string
	MaxQueryLimit         int
	LogAll                bool
	LogCreate             bool
	LogUpdate             bool
	LogRead               bool
	LogDelete             bool
	LogLogin              bool
	LogLogout             bool
	UnAuthorizedMessage   string
	RecExistMessage       string
	CacheExpire           int
	LoginTimeout          int
	UsernameExistsMessage string
	EmailExistsMessage    string
	MsgFrom               string
}

type MongoCrudParamType struct {
	AppDb           *mongo.Client
	TableName       string
	Token           string
	UserInfo        UserInfoType
	UserId          string
	Group           string
	Groups          []string
	RecordIds       []string
	ActionParams    ActionParamsType
	QueryParams     QueryParamType
	ExistParams     ExistParamsType
	ProjectParams   ProjectParamType
	SortParams      SortParamType
	Skip            int
	Limit           int
	ParentTables    []string
	ChildTables     []string
	RecursiveDelete bool
	CheckAccess     bool
	AccessDb        *mongo.Client
	AuditDb         *mongo.Client
	AuditTable      string
	ServiceTable    string
	UserTable       string
	RoleTable       string
	AccessTable     string
	MaxQueryLimit   int
	LogAll          bool
	LogCreate       bool
	LogUpdate       bool
	LogRead         bool
	LogDelete       bool
	//transLog AuditLog
	HashKey             string
	IsRecExist          bool
	ActionAuthorized    bool
	UnAuthorizedMessage string
	RecExistMessage     string
	IsAdmin             bool
	CreateItems         ActionParamsType
	UpdateItems         ActionParamsType
	CurrentRecs         ActionParamsType
	RoleServices        []RoleServiceType
	SubItems            []bool
	CacheExpire         int
	Params              MongoCrudTaskType
}

type MessageObject map[string]string

type ValidateResponseType struct {
	Ok     bool          `json:"ok"`
	Errors MessageObject `json:"errors"`
}
type OkResponse struct {
	Ok bool `json:"ok"`
}

// CRUD operations
type CreateQueryResponseType struct {
	CreateQuery string
	FieldNames  []string
	FieldValues [][]interface{}
}

type UpdateQueryResponseType struct {
	UpdateQuery string
	WhereQuery  string
	FieldValues []interface{}
}

type WhereQueryResponseType struct {
	WhereQuery  string
	FieldValues []interface{}
}

type DeleteQueryResponseType struct {
	DeleteQuery string
	WhereQuery  string
	FieldValues []interface{}
}

type SelectQueryResponseType struct {
	SelectQuery string
	WhereQuery  string
	FieldValues []interface{}
}

type SaveParamsType struct {
	UserInfo     UserInfoType     `json:"user_info"`
	QueryParams  QueryParamType   `json:"query_params"`
	RecordIds    []string         `json:"record_ids"`
	//ActionParams ActionParamsType `json:"action_params"`
}

type DeleteParamsType struct {
	UserInfo    UserInfoType   `json:"user_info"`
	RecordIds   []string       `json:"record_ids"`
	QueryParams QueryParamType `json:"query_params"`
}

type GetParamsType struct {
	UserInfo     UserInfoType     `json:"user_info"`
	Skip         int              `json:"skip"`
	Limit        int              `json:"limit"`
	RecordIds    []string         `json:"record_ids"`
	QueryParams  QueryParamType   `json:"query_params"`
	SortParam    SortParamType    `json:"sort_params"`
	ProjectParam ProjectParamType `json:"project_param"`
}

// ErrorType provides the structure for error reporting
type ErrorType struct {
	Code    string
	Message string
}

type SaveError ErrorType
type CreateError ErrorType
type UpdateError ErrorType
type DeleteError ErrorType
type ReadError ErrorType
type AuthError ErrorType
type ConnectError ErrorType
type SelectQueryError ErrorType
type WhereQueryError ErrorType
type CreateQueryError ErrorType
type UpdateQueryError ErrorType
type DeleteQueryError ErrorType

// sample Error() implementation
func (err ErrorType) Error() string {
	return fmt.Sprintf("Error-code: %v | Error-message: %v", err.Code, err.Message)
}
