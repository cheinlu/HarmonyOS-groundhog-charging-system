package logic

import (
	"context"
	"fmt"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

var Station StationLogic

type StationLogic struct {
}

func (*StationLogic) StationList(ctx context.Context, query entity.Station, page model.PageReq) (stations []entity.Station, count int, err error) {
	err = dao.Station.Ctx(ctx).WhereLike("name", fmt.Sprintf("%%%s%%", query.Name)).WhereLike("address", fmt.Sprintf("%%%s%%", query.Address)).Page(page.PageNo, page.PageSize).ScanAndCount(&stations, &count, false)
	return
}

func (*StationLogic) Add(ctx context.Context, station do.Station) (err error) {
	station.CreateAt, station.UpdateAt = gtime.Now(), gtime.Now()
	_, err = dao.Station.Ctx(ctx).Insert(station)
	return
}

func (*StationLogic) Del(ctx context.Context, id int) (err error) {
	_, err = dao.Station.Ctx(ctx).Delete("id = ?", id)
	return
}

func (*StationLogic) Update(ctx context.Context, station do.Station) (err error) {
	station.UpdateAt = gtime.Now()
	_, err = dao.Station.Ctx(ctx).Update(station, "id = ?", station.Id)
	return
}
