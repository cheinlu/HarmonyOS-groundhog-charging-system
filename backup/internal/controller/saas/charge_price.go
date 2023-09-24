package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
)

// 充电价格表
type ChargePriceController struct {
}

// 时段价格列表
func (ChargePriceController) List(ctx context.Context, req *v2.ChargePriceListReq) (pageRes model.PageRes, err error) {
	model.InitPageReq(&req.PageReq, 1, 10)
	chargePrices, count, err := logic.ChargePrice.ChargePriceList(ctx, req.PageReq)
	res := make([]v2.ChargePriceListRes, len(chargePrices))
	for i, chargePrice := range chargePrices {
		res[i] = v2.ChargePriceListRes{
			Id:        chargePrice.Id,
			TenantId:  chargePrice.TenantId,
			StartHour: chargePrice.StartHour,
			EndHour:   chargePrice.EndHour,
			Price:     chargePrice.Price,
			CreateAt:  chargePrice.CreateAt.Time,
			UpdateAt:  chargePrice.UpdateAt.Time,
		}
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 时段价格添加
func (ChargePriceController) Add(ctx context.Context, req *v2.ChargePriceAddReq) (res *v2.ChargePriceAddRes, err error) {
	var chargePrice do.ChargePrice = do.ChargePrice{
		StartHour: req.StartHour,
		EndHour:   req.EndHour,
		Price:     req.Price,
	}
	err = logic.ChargePrice.Add(ctx, chargePrice)
	return
}

// 删除时段价格
func (ChargePriceController) Del(ctx context.Context, req *v2.ChargePriceDelReq) (res *v2.ChargePriceDelRes, err error) {
	err = logic.ChargePrice.Del(ctx, req.Id)
	return
}

// 修改时段价格
func (ChargePriceController) Update(ctx context.Context, req *v2.ChargePriceUpdateReq) (res *v2.ChargePriceUpdateRes, err error) {
	chargePrice := do.ChargePrice{
		Id:        req.Id,
		StartHour: req.StartHour,
		EndHour:   req.EndHour,
		Price:     req.Price,
	}
	err = logic.ChargePrice.Update(ctx, chargePrice)
	return
}
