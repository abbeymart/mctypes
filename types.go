// @Author: abbeymart | Abi Akindele | @Created: 2020-12-23 | @Updated: 2020-12-23
// @Company: mConnect.biz | @License: MIT
// @Description: shared types

package mctypes

import "time"

type AuditStampType struct {
	IsActive  bool      `json:"is_active"` // => activate by modelOptionsType settings...
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInfoType struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Language  string `json:"language"`
	LoginName string `json:"login_name"`
	Token     string `json:"token"`
	Expire    uint   `json:"expire"`
}
