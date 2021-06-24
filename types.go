// @Author: abbeymart | Abi Akindele | @Created: 2020-12-23 | @Updated: 2020-12-23
// @Company: mConnect.biz | @License: MIT
// @Description: shared types

package mctypes

import "time"

type AuditStampType struct {
	IsActive  bool      `json:"isActive"` // => activate by modelOptionsType settings...
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserInfoType struct {
	UserId    string `json:"userId" form:"userId" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Language  string `json:"language"`
	LoginName string `json:"loginName" form:"loginName" binding:"required"`
	Token     string `json:"token"`
	Expire    int64   `json:"expire"`
	Group     string `json:"group"`
	Email     string `json:"email" form:"email" binding:"required"`
}

type ValueType struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}
