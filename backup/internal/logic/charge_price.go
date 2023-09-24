package logic

import (
	"context"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// 时段价格表

var ChargePrice ChargePriceLogic

type ChargePriceLogic struct {
}

func (*ChargePriceLogic) ChargePriceList(ctx context.Context, page model.PageReq) (chargePrices []entity.ChargePrice, count int, err error) {
	err = dao.ChargePrice.Ctx(ctx).Page(page.PageNo, page.PageSize).OrderAsc("start_hour").ScanAndCount(&chargePrices, &count, false)
	return
}

func (*ChargePriceLogic) Add(ctx context.Context, chargePrice do.ChargePrice) (err error) {
	chargePrice.CreateAt, chargePrice.UpdateAt = gtime.Now(), gtime.Now()
	err = SetCurrentTenantId(ctx, &(chargePrice.TenantId))
	if err != nil {
		return
	}
	_, err = dao.ChargePrice.Ctx(ctx).Insert(chargePrice)
	return
}

func (*ChargePriceLogic) Del(ctx context.Context, id int) (err error) {
	rs, err := dao.ChargePrice.Ctx(ctx).Delete("id = ?", id)
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

func (*ChargePriceLogic) Update(ctx context.Context, chargePrice do.ChargePrice) (err error) {
	chargePrice.UpdateAt = gtime.Now()
	rs, err := dao.ChargePrice.Ctx(ctx).Update(chargePrice, "id = ?", chargePrice.Id)
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

// GetChargePricesByHour 根据小时数获取对应的收费价格
func (*ChargePriceLogic) GetChargePricesByHour(ctx context.Context, targetHour int) (matchingPrices []entity.ChargePrice, err error) {
	err = dao.ChargePrice.Ctx(ctx).WhereLTE("start_hour", targetHour).WhereGT("end_hour", targetHour).Scan(&matchingPrices)
	return
}
