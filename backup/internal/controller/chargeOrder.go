package controller

import (
	"context"
	v1 "login-demo/api/v1"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
)

// 充电订单

type ChargeOrderController struct {
}

func (ChargeOrderController) List(ctx context.Context, req *v1.ChargeOrderListReq) (pageRes model.PageRes, err error) {
	res := &[]v1.ChargeOrderListRes{}
	// 初始化分页参数，设置默认值
	model.InitPageReq(&req.PageReq, 1, 10)
	query := model.ChargeOrderQuery{
		OrderCode: req.OrderCode,
		UserId:    req.UserId,
		StationId: req.StationId,
		PileId:    req.PileId,
		BeforeAt:  req.BeforeAt,
		AfterAt:   req.AfterAt,
		State:     req.State,
		PageReq:   req.PageReq,
	}
	chargeOrderList, count, err := logic.ChargeOrder.ChargeOrderList(ctx, query)
	for _, v := range chargeOrderList {
		chargeOrderRes := v1.ChargeOrderListRes{
			Id:          v.Id,
			OrderCode:   v.OrderCode,
			UserId:      v.UserId,
			Nickname:    v.Nickname,
			StationId:   v.StationId,
			StationName: v.StationName,
			PileId:      v.PileId,
			PileCode:    v.PileCode,
			StartAt:     v.StartAt,
			StopAt:      v.StopAt,
			State:       v.State,
			Price:       v.Price,
			CreateAt:    v.CreateAt,
			UpdateAt:    v.UpdateAt,
		}
		tmp := append(*res, chargeOrderRes)
		res = &tmp
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 修改充电订单状态
func (ChargeOrderController) Update(ctx context.Context, req *v1.ChargeOrderUpdateReq) (res *v1.ChargeOrderUpdateRes, err error) {
	chargeOrder := do.ChargeOrder{
		Id:    req.Id,
		State: req.State,
		Price: req.Price,
	}
	err = logic.ChargeOrder.Update(ctx, chargeOrder)
	return
}
