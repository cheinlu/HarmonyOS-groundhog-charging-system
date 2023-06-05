package logic

import (
	"context"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// 时段价格表

var ChargePrice ChargePriceLogic

type ChargePriceLogic struct {
}

func (*ChargePriceLogic) ChargePriceList(ctx context.Context, page model.PageReq) (chargePrices []entity.ChargePrice, count int, err error) {
	err = dao.ChargePrice.Ctx(ctx).Page(page.PageNo, page.PageSize).ScanAndCount(&chargePrices, &count, false)
	return
}

func (*ChargePriceLogic) Add(ctx context.Context, chargePrice do.ChargePrice) (err error) {
	chargePrice.CreateAt, chargePrice.UpdateAt = gtime.Now(), gtime.Now()
	_, err = dao.ChargePrice.Ctx(ctx).Insert(chargePrice)
	return
}

func (*ChargePriceLogic) Del(ctx context.Context, id int) (err error) {
	_, err = dao.ChargePrice.Ctx(ctx).Delete("id = ?", id)
	return
}

func (*ChargePriceLogic) Update(ctx context.Context, chargePrice do.ChargePrice) (err error) {
	chargePrice.UpdateAt = gtime.Now()
	_, err = dao.ChargePrice.Ctx(ctx).Update(chargePrice, "id = ?", chargePrice.Id)
	return
}

// GetChargePricesByHour 根据小时数获取对应的收费价格
func (*ChargePriceLogic) GetChargePricesByHour(ctx context.Context, targetHour int) (matchingPrices []entity.ChargePrice, err error) {
	err = dao.ChargePrice.Ctx(ctx).WhereLTE("start_hour", targetHour).WhereGT("end_hour", targetHour).Scan(&matchingPrices)
	return
}
