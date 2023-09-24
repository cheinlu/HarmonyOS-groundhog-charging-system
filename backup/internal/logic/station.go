package logic

import (
	"context"
	"fmt"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

var Station StationLogic

type StationLogic struct {
}

func (*StationLogic) StationList(ctx context.Context, query entity.Station, page model.PageReq) (stations []model.Station, count int, err error) {
	err = dao.Station.Ctx(ctx).InnerJoin("tenant", "station.tenant_id = tenant.id").WhereLike("station.name", fmt.Sprintf("%%%s%%", query.Name)).WhereLike("station.address", fmt.Sprintf("%%%s%%", query.Address)).OrderDesc("station.update_at").Fields("station.*,tenant.name as tenantName").Page(page.PageNo, page.PageSize).ScanAndCount(&stations, &count, false)
	return
}

func (*StationLogic) Add(ctx context.Context, station do.Station) (err error) {
	station.CreateAt, station.UpdateAt = gtime.Now(), gtime.Now()
	err = SetCurrentTenantId(ctx, &station.TenantId)
	if err != nil {
		return
	}
	_, err = dao.Station.Ctx(ctx).Insert(station)
	return
}

func (*StationLogic) Del(ctx context.Context, id int) (err error) {
	rs, err := dao.Station.Ctx(ctx).Delete("id = ?", id)
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，删除失败", ""), err)
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，删除失败", ""), err)
	}
	if rowsAffected == 0 {
		err = gerror.NewCode(gcode.New(1, "删除失败，未找到原数据，可能已被其他人删除", ""))
		return err
	}
	return
}

func (*StationLogic) Update(ctx context.Context, station do.Station) (err error) {
	station.UpdateAt = gtime.Now()
	rs, err := dao.Station.Ctx(ctx).Update(station, "id = ?", station.Id)
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，修改失败", ""), err)
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，修改失败", ""), err)
	}
	if rowsAffected == 0 {
		err = gerror.NewCode(gcode.New(1, "修改失败，未找到原数据，可能已被其他人删除", ""))
		return err
	}
	return
}
