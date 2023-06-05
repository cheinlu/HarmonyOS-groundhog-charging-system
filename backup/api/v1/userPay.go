package v1

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type UserPayRecordReq struct {
	g.Meta `path:"/api/user/pay/record" tags:"userPayRecord" method:"get" summary:"用户充值列表"`
	model.PageReq
}

type UserPayRecordRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Id       int       `json:"id"`
	PayCode  string    `json:"payCode"  `   //
	UserId   int       `json:"userId"   `   //
	Username string    `json:"username"   ` //
	State    int       `json:"state"    `   //
	Price    float64   `json:"price"    `   //
	PayAt    time.Time `json:"payAt"    `   //
	CreateAt time.Time `json:"createAt" `   //
	UpdateAt time.Time `json:"updateAt" `   //
}

type UserPayReq struct {
	g.Meta `path:"/api/user/pay" tags:"userPay" method:"POST" summary:"用户充值"`
	UserId int     `json:"userId" ` //
	Price  float64 `json:"price" `  //
}

type UserPayRes struct {
	g.Meta `mime:"application/json" `
}
