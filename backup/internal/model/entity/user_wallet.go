// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWallet is the golang structure for table user_wallet.
type UserWallet struct {
	UserId   int         `json:"userId"   ` //
	TenantId int         `json:"tenantId" ` //
	Balance  float64     `json:"balance"  ` //
	CreateAt *gtime.Time `json:"createAt" ` //
	UpdateAt *gtime.Time `json:"updateAt" ` //
}
