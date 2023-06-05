package v1

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// 管理台：充电订单列表
type ChargeOrderListReq struct {
	g.Meta    `path:"/api/charge/list" tags:"chargeList" method:"get" summary:"充电订单列表"`
	OrderCode string    `json:"orderCode"`
	UserId    int       `json:"userId"`
	StationId int       `json:"stationId"`
	PileId    int       `json:"pileId"`
	BeforeAt  time.Time `json:"startAt"` // 订单开始时间（起始）
	AfterAt   time.Time `json:"stopAt"`  // 订单开始时间（结束）
	State     int       `json:"state"`   // 0:充电中,1:已完成,2:未付账
	model.PageReq
}

type ChargeOrderListRes struct {
	g.Meta      `mime:"application/json" `
	Id          int       `json:"id"`
	OrderCode   string    `json:"orderCode" `   //
	UserId      int       `json:"userId"    `   //
	Nickname    string    `json:"nickname"    ` //
	StationId   int       `json:"stationId" `   //
	StationName string    `json:"stationName" ` //
	PileId      int       `json:"pileId"    `   //
	PileCode    string    `json:"pileCode"    ` //
	StartAt     time.Time `json:"startAt"   `   //
	StopAt      time.Time `json:"stopAt"    `   //
	State       int       `json:"state"     `   //
	Price       float64   `json:"price"     `   //
	CreateAt    time.Time `json:"createAt"  `   //
	UpdateAt    time.Time `json:"updateAt"  `   //
}

// 修改订单
type ChargeOrderUpdateReq struct {
	g.Meta `path:"/api/charge/update" tags:"chargeUpdate" method:"post" summary:"修改订单状态"`
	Id     int     `json:"id"`
	State  int     `json:"state"     ` //
	Price  float64 `json:"price"     ` //
}

type ChargeOrderUpdateRes struct {
	g.Meta `mime:"application/json" `
}
