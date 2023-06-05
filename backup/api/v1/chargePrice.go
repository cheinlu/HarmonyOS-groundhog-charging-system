package v1

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type ChargePriceListReq struct {
	g.Meta `path:"/api/charge/price/list" tags:"chargePrice" method:"get" summary:"时段价格列表"`
	model.PageReq
}

type ChargePriceListRes struct {
	g.Meta    `mime:"application/json" example:"string"`
	Id        int       `json:"id"`
	StartHour int       `json:"startHour" ` //
	EndHour   int       `json:"endHour"   ` //
	Price     float64   `json:"price"     ` //
	CreateAt  time.Time `json:"createAt"  ` //
	UpdateAt  time.Time `json:"updateAt"  ` //
}

type ChargePriceAddReq struct {
	g.Meta    `path:"/api/charge/price/add" tags:"chargePriceAdd" method:"POST" summary:"添加时段价格表"`
	StartHour int     `json:"startHour" ` //
	EndHour   int     `json:"endHour"   ` //
	Price     float64 `json:"price"     ` //
}

type ChargePriceAddRes struct {
	g.Meta `mime:"application/json" `
}

type ChargePriceDelReq struct {
	g.Meta `path:"/api/charge/price/delete" tags:"chargePriceDelete" method:"DELETE" summary:"删除时段价格表"`
	Id     int `json:"id"`
}

type ChargePriceDelRes struct {
	g.Meta `mime:"application/json" `
}

type ChargePriceUpdateReq struct {
	g.Meta    `path:"/api/charge/price/update" tags:"chargePriceDelete" method:"post" summary:"修改时段价格表"`
	Id        int     `json:"id"`
	StartHour int     `json:"startHour" ` //
	EndHour   int     `json:"endHour"   ` //
	Price     float64 `json:"price"     ` //
}

type ChargePriceUpdateRes struct {
	g.Meta `mime:"application/json" `
}
