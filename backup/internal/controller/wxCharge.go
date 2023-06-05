package controller

import (
	"context"
	v1 "login-demo/api/v1"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
	"sort"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type WxChargeController struct {
}

// StationList 获取充电站列表，包含充电桩数量和当前价格
func (c WxChargeController) StationList(ctx context.Context, req *v1.WxStationListReq) (pageRes *model.PageRes, err error) {
	// 初始化分页参数
	pageRes = &model.PageRes{}
	model.InitPageReq(&req.PageReq, 1, 10)

	// 构造查询条件
	query := entity.Station{
		Name:    req.Name,
		Address: req.Address,
	}

	// 查询充电站列表
	stations, count, err := logic.Station.StationList(ctx, query, req.PageReq)
	if err != nil {
		return nil, err
	}
	// 获得当前位置
	myCoordinate := req.Coordinate
	if myCoordinate == "" {
		// 默认采用延安的位置
		myCoordinate = "36.5853,109.4898"
	}

	// 构造响应数据
	wxStationListRes := make([]*v1.WxStationListRes, len(stations))
	for i, station := range stations {

		// 计算当前位置与充电站的距离
		distance := logic.Geo.GetDistance(myCoordinate, station.Coordinate)

		// 填充充电站基本信息
		wxStationListRes[i] = &v1.WxStationListRes{
			Id:         station.Id,
			Name:       station.Name,
			Address:    station.Address,
			Coordinate: station.Coordinate,
			Distance:   distance,
			CreateAt:   station.CreateAt.Time,
			UpdateAt:   station.UpdateAt.Time,
		}

		// 查询充电桩数量
		totalQuery := do.Pile{
			StationId: station.Id,
		}
		pileTotal, _ := logic.Pile.PileTotal(ctx, totalQuery)
		freeTotalQuery := do.Pile{
			StationId: station.Id,
			State:     0,
		}
		freePileTotal, _ := logic.Pile.PileTotal(ctx, freeTotalQuery)
		wxStationListRes[i].PileTotal, wxStationListRes[i].FreePileTotal = pileTotal, freePileTotal

		// 计算价格
		hour := time.Now().Hour()
		chargePricesByHour, err := logic.ChargePrice.GetChargePricesByHour(ctx, hour)
		if err != nil {
			err = gerror.NewCode(gcode.New(1, "获取充电价格失败", err.Error()))
			return nil, err
		}
		if len(chargePricesByHour) > 0 {
			wxStationListRes[i].PricePerHour = chargePricesByHour[0].Price
		} else {
			// 默认价格 1.5
			wxStationListRes[i].PricePerHour = 1.5
		}

	}

	// 对充电站按距离远近排序
	sort.Slice(wxStationListRes, func(i, j int) bool {
		return wxStationListRes[i].Distance < wxStationListRes[j].Distance
	})

	// 填充分页信息
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = wxStationListRes, req.PageNo, req.PageSize, count
	return
}

// 充电桩列表
func (c WxChargeController) PileList(ctx context.Context, req *v1.WxPileListReq) (pageRes model.PageRes, err error) {
	model.InitPageReq(&req.PageReq, 1, 10)
	query := entity.Pile{
		StationId: req.StationId,
	}
	piles, count, err := logic.Pile.PileList(ctx, query, req.PageReq)
	res := make([]*v1.WxPileListRes, len(piles))
	for i, pile := range piles {
		res[i] = &v1.WxPileListRes{
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

// 开始充电
func (c WxChargeController) StartCharge(ctx context.Context, req *v1.WXStartChargeReq) (res v1.WXStartChargeRes, err error) {
	err = logic.ChargeOrder.StartCharge(ctx, req.PileId, req.ChargeHours)
	return
}

// 结束充电
func (c WxChargeController) StopCharge(ctx context.Context, req *v1.WXStopChargeReq) (res v1.WXStopChargeRes, err error) {
	err = logic.ChargeOrder.StopCharge(ctx, req.OrderId)
	return
}

// 我的充电订单
func (c WxChargeController) MyChargeOrders(ctx context.Context, req *v1.WXChargeListReq) (pageRes model.PageRes, err error) {
	res := &[]v1.ChargeOrderListRes{}
	// 初始化分页参数，设置默认值
	model.InitPageReq(&req.PageReq, 1, 10)
	// 获得当前用户
	currentUser, err := logic.Ctx.GetCurrentUser(ctx)
	if err != nil {
		err = gerror.NewCode(gcode.New(1, "系统异常", err.Error()))
		return
	}
	query := model.ChargeOrderQuery{
		UserId:  currentUser.Id,
		PageReq: req.PageReq,
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

// 我的个人信息
func (c WxChargeController) AboutMe(ctx context.Context, req *v1.WXMeInfoReq) (res v1.WXMeInfoRes, err error) {
	// 获得当前用户
	currentUser, err := logic.Ctx.GetCurrentUser(ctx)
	if err != nil {
		err = gerror.NewCode(gcode.New(1, "系统异常", err.Error()))
		return
	}
	res.Nickname, res.Balance, res.AvatarUrl = currentUser.Nickname, currentUser.Balance, ""
	return
}
