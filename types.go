// @Author: abbeymart | Abi Akindele | @Created: 2020-12-23 | @Updated: 2020-12-23
// @Company: mConnect.biz | @License: MIT
// @Description: shared types

package mctypes

import "time"

type BaseModelType struct {
	Id        string    `json:"id"`
	Language  string    `json:"language"`
	Desc      string    `json:"desc"`
	AppId     string    `json:"appId"`    // application-id in a multi-hosted apps environment (e.g. cloud-env)
	IsActive  bool      `json:"isActive"` // => activate by modelOptionsType settings...
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type AuditStampType struct {
	IsActive  bool      `json:"isActive"` // => activate by modelOptionsType settings...
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AppParamsType struct {
	AppId     string `json:"appId"`
	AccessKey string `json:"accessKey"`
	AppName   string `json:"appName"`
}

type AppType struct {
	BaseModelType
	AppName     string `json:"appName"`
	AccessKey   string `json:"accessKey"`
	AppCategory string `json:"appCategory"`
	OwnerId     string `json:"ownerId"`
}

type UserInfoType struct {
	UserId    string `json:"userId" form:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Language  string `json:"language"`
	LoginName string `json:"loginName" form:"loginName"`
	Token     string `json:"token"`
	Expire    int64  `json:"expire"`
	Email     string `json:"email" form:"email"`
	Group     string `json:"group"`
}

type ValueType struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}
