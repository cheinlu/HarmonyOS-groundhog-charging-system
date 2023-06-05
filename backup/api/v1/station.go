package v1

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type StationListReq struct {
	g.Meta  `path:"/api/station/list" tags:"station" method:"get" summary:"充电站点列表"`
	Name    string `json:"name"`
	Address string `json:"address"`
	model.PageReq
}

type StationListRes struct {
	g.Meta     `mime:"application/json" example:"string"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Coordinate string    `json:"coordinate" `  //
	CreateAt   time.Time `json:"createAt"    ` //
	UpdateAt   time.Time `json:"updateAt"    ` //
}

type StationAddReq struct {
	g.Meta     `path:"/api/station/add" tags:"stationAdd" method:"POST" summary:"添加充电站"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Coordinate string `json:"coordinate" ` //
}

type StationAddRes struct {
	g.Meta `mime:"application/json" `
}

type StationDelReq struct {
	g.Meta `path:"/api/station/delete" tags:"stationDelete" method:"DELETE" summary:"删除充电桩"`
	Id     int `json:"id"`
}

type StationDelRes struct {
	g.Meta `mime:"application/json" `
}

type StationUpdateReq struct {
	g.Meta     `path:"/api/station/update" tags:"stationDelete" method:"post" summary:"修改充电桩"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Coordinate string `json:"coordinate" ` //
}

type StationUpdateRes struct {
	g.Meta `mime:"application/json" `
}
