package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
)

type PileController struct {
}

// 充电站列表
func (PileController) List(ctx context.Context, req *v2.PileListReq) (pageRes model.PageRes, err error) {
	model.InitPageReq(&req.PageReq, 1, 10)
	query := entity.Pile{
		StationId: req.StationId,
		Code:      req.Code,
	}
	piles, count, err := logic.Pile.PileList(ctx, query, req.PageReq)
	res := make([]*v2.PileListRes, len(piles))
	for i, pile := range piles {
		res[i] = &v2.PileListRes{
			Id:          pile.Id,
			Code:        pile.Code,
			StationId:   pile.StationId,
			StationName: pile.Name,
			State:       pile.State,
			CreateAt:    pile.CreateAt.Time,
			UpdateAt:    pile.UpdateAt.Time,
		}
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 充电站添加
func (PileController) Add(ctx context.Context, req *v2.PileAddReq) (res *v2.PileAddRes, err error) {
	var pile do.Pile = do.Pile{
		Code:      req.Code,
		StationId: req.StationId,
		State:     0,
	}
	err = logic.Pile.Add(ctx, pile)
	return
}

// 删除充电站
func (PileController) Del(ctx context.Context, req *v2.PileDelReq) (res *v2.PileDelRes, err error) {
	err = logic.Pile.Del(ctx, req.Id)
	return
}

func (PileController) Update(ctx context.Context, req *v2.PileUpdateReq) (res *v2.PileUpdateRes, err error) {

	pile := do.Pile{
		Id:        req.Id,
		Code:      req.Code,
		StationId: req.StationId,
		State:     req.State,
	}
	err = logic.Pile.Update(ctx, pile)
	return
}
