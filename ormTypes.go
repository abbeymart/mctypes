// @Author: abbeymart | Abi Akindele | @Created: 2020-12-23 | @Updated: 2020-12-23
// @Company: mConnect.biz | @License: MIT
// @Description: orm types

package mctypes

import "github.com/jackc/pgx/v4/pgxpool"

// ORM types
type RecordValueType map[string]ValueParamType
type RecordDescType map[string]FieldDescType

type GetValueType func() interface{}
type SetValueType func(val interface{}) interface{}
type DefaultValueType func() interface{}
type ValidateMethodType func(val interface{}) bool
type ValidateMethodResponseType func(val interface{}) ValidateResponseType
type ComputedValueType func(val interface{}) interface{}

type ValidateMethodsType map[string]ValidateMethodResponseType
type ComputedMethodsType map[string]ComputedValueType

type FieldDescType struct {
	FieldType       string
	FieldLength     int   // default: 255 for DataType.STRING
	FieldPattern    string // "/^[0-9]{10}$/" => includes 10 digits, 0 to 9 | "/^[0-9]{6}.[0-9]{2}$/ => max 16 digits and 2 decimal places
	AllowNull       bool   // default: true
	Unique          bool
	Indexable       bool
	PrimaryKey      bool
	MinValue        int
	MaxValue        int
	SetValue        SetValueType       // set/transform fieldValue prior to save(create/insert), T=>fieldType
	DefaultValue    DefaultValueType   // result/T must be of fieldType
	Validate        ValidateMethodType // T=>fieldType, returns a bool (valid=true/invalid=false)
	ValidateMessage string
}

type ModelRelationType struct {
	SourceTable   string
	TargetTable   string
	SourceField   string
	TargetField   string
	RelationType  string
	SourceModel   ModelType
	TargetModel   ModelType
	ForeignField  string // source-to-targetField map
	RelationField string // relation-targetField, for many-to-many
	RelationTable string // optional tableName for many-to-many | default: sourceTable_targetTable
	OnDelete      string
	OnUpdate      string
}

type ModelType struct {
	AppDb           *pgxpool.Pool
	TableName       string
	RecordDesc      RecordDescType
	TimeStamp       bool // auto-add: createdAt and updatedAt | default: true
	ActorStamp      bool // auto-add: createdBy and updatedBy | default: true
	ActiveStamp     bool // record active status, isActive (true | false) | default: true
	Relations       []ModelRelationType
	ComputedMethods ComputedMethodsType // model-level functions, e.g fullName(a, b: T): T
	ValidateMethods ValidateMethodsType
	AlterSyncTable  bool // create / alter table/collection and sync existing data, if there was a change to the table structure | default: true
	// if alterSyncTable: false it will create/re-create the table, with no data sync
}
