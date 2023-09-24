package v2_wx

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 微信小程序充电接口

// 小程序：开始充电
type WXStartChargeReq struct {
	g.Meta      `path:"/wx-api/charge/start" tags:"微信小程序业务接口" method:"POST" summary:"小程序：开始充电"`
	PileId      int `json:"pileId"    ` //
	ChargeHours int `json:"chargeHours"`
}
type WXStartChargeRes struct {
	g.Meta `mime:"application/json" `
}

// 小程序：结束充电
type WXStopChargeReq struct {
	g.Meta  `path:"/wx-api/charge/stop" tags:"微信小程序业务接口" method:"POST" summary:"小程序：结束充电"`
	OrderId int `json:"orderId" ` //
}
type WXStopChargeRes struct {
	g.Meta `mime:"application/json" `
}

// 小程序：充电详情
type WXChargeDetailReq struct {
	g.Meta    `path:"/wx-api/charge/detail" tags:"微信小程序业务接口" method:"get" summary:"小程序：充电详情"`
	OrderCode string `json:"chargename"`
}
type WXChargeDetailRes struct {
	g.Meta      `mime:"application/json" `
	Id          int       `json:"id"`
	OrderCode   string    `json:"orderCode" `   //
	StationId   int       `json:"stationId" `   //
	StationName int       `json:"stationName" ` //
	PileId      int       `json:"pileId"    `   //
	PileCode    int       `json:"pileCode"    ` //
	StartAt     time.Time `json:"startAt"   `   //
	StopAt      time.Time `json:"stopAt"    `   //
	State       int       `json:"state"     `   //
	Price       float64   `json:"price"     `   //
	CreateAt    time.Time `json:"createAt"  `   //
	UpdateAt    time.Time `json:"updateAt"  `   //
}

// 小程序：充电列表
type WXChargeListReq struct {
	g.Meta `path:"/wx-api/charge/list" tags:"微信小程序业务接口" method:"get" summary:"小程序：充电列表"`
	model.PageReq
}

type WXChargeListRes struct {
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

// 充电站列表
type WxStationListReq struct {
	g.Meta     `path:"/wx-api/station/list" tags:"微信小程序业务接口" method:"get" summary:"充电站点列表"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Coordinate string `json:"coordinate"`
	TenantName string `json:"tenantName"`
	model.PageReq
}

type WxStationListRes struct {
	g.Meta        `mime:"application/json" example:"string"`
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	TenantId      int       `json:"tenantId"`
	TenantName    string    `json:"tenantName"`
	Address       string    `json:"address"`
	Coordinate    string    `json:"coordinate" ` //
	PileTotal     int       `json:"pileTotal" `
	FreePileTotal int       `json:"freePileTotal" `
	PricePerHour  float64   `json:"pricePerHour" `
	Distance      float64   `json:"distance" `
	ImageUrl      string    `json:"imageUrl"`     //
	CreateAt      time.Time `json:"createAt"    ` //
	UpdateAt      time.Time `json:"updateAt"    ` //
}

type WxPileListReq struct {
	g.Meta    `path:"/wx-api/pile/list" tags:"微信小程序业务接口" method:"get" summary:"充电桩点列表"`
	StationId int `json:"stationId"`
	model.PageReq
}

type WxPileListRes struct {
	g.Meta      `mime:"application/json" example:"string"`
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	StationId   int       `json:"stationId"`
	StationName string    `json:"stationName"`
	State       int       `json:"state"`        // 0:空闲中,1:充电中,2:故障
	CreateAt    time.Time `json:"createAt"    ` //
	UpdateAt    time.Time `json:"updateAt"    ` //
}

type WXMeInfoReq struct {
	g.Meta `path:"/wx-api/me/info" tags:"微信小程序业务接口" method:"get" summary:"小程序：个人信息"`
}

type WXMeInfoRes struct {
	g.Meta    `mime:"application/json" `
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Nickname  string  `json:"nickname"`
	Balance   float64 `json:"balance"  `  //
	AvatarUrl string  `json:"avatarUrl" ` //
}

// 价格区间
type WXPriceListReq struct {
	g.Meta `path:"/wx-api/price/list" tags:"微信小程序业务接口" method:"get" summary:"时段价格列表"`
	model.PageReq
}

type WXPriceListRes struct {
	g.Meta    `mime:"application/json" example:"string"`
	Id        int       `json:"id"`
	StartHour int       `json:"startHour" `    //
	EndHour   int       `json:"endHour"   `    //
	Price     float64   `json:"price"     `    //
	TenantId  int       `json:"tenantId"     ` //
	CreateAt  time.Time `json:"createAt"  `    //
	UpdateAt  time.Time `json:"updateAt"  `    //
}

// 上传个人头像
type WXUploadAvatarReq struct {
	g.Meta `path:"/wx-api/upload/avatar" tags:"微信小程序业务接口" method:"post" summary:"上传个人头像"`
	File   *ghttp.UploadFile `json:"file" type:"json" dc:"选择上传文件"`
}

type WXUploadAvatarRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"  ` //
}

// 修改个人昵称
type WXUpdateNicknameReq struct {
	g.Meta      `path:"/wx-api/update/nickname" tags:"微信小程序业务接口" method:"post" summary:"修改个人昵称"`
	NewNickname string `json:"newNickname"`
}

type WXUpdateNicknameRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"  ` //
}
