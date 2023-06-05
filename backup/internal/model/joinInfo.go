package model

import (
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// 充电相关
type ChargeOrderQuery struct {
	OrderCode string    `json:"qrderCode"`
	UserId    int       `json:"userId"`
	StationId int       `json:"stationId"`
	PileId    int       `json:"pileId"`
	BeforeAt  time.Time `json:"beforeAt"` // 订单开始时间（起始）
	AfterAt   time.Time `json:"afterAt"`  // 订单开始时间（结束）
	State     int       `json:"state"`    //0:充电中,1:已完成,2:未付账
	PageReq
}

// 充电订单聚合查询
type ChargeOrderResult struct {
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

// Pile is the golang structure for table pile.
type Pile struct {
	Id        int         `json:"id"        ` //
	Code      string      `json:"code"      ` //
	Name      string      `json:"name"      ` //
	StationId int         `json:"stationId" ` //
	State     int         `json:"state"     ` //
	CreateAt  *gtime.Time `json:"createAt"  ` //
	UpdateAt  *gtime.Time `json:"updateAt"  ` //
}

// user 聚合信息，聚合余额
type UserMore struct {
	Id       int       `json:"id"       ` //
	Passport string    `json:"passport" ` //
	Password string    `json:"password" ` //
	Nickname string    `json:"nickname" ` //
	Balance  float64   `json:"balance"  ` //
	CreateAt time.Time `json:"createAt" ` //
	UpdateAt time.Time `json:"updateAt" ` //
}

// payRecord 聚合信息，聚合用户名
type PayRecordInfo struct {
	Id       int       `json:"id"       `   //
	PayCode  string    `json:"payCode"  `   //
	UserId   int       `json:"userId"   `   //
	State    int       `json:"state"    `   //
	Price    float64   `json:"price"    `   //
	Username string    `json:"userName"   ` //
	PayAt    time.Time `json:"payAt"    `   //
	CreateAt time.Time `json:"createAt" `   //
	UpdateAt time.Time `json:"updateAt" `   //
}
