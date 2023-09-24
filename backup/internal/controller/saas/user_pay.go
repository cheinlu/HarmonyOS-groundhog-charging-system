package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
)

// 用户充值相关

type UserPayController struct {
}

// 查看用户充值列表
func (UserPayController) List(ctx context.Context, req *v2.UserPayRecordReq) (pageRes model.PageRes, err error) {
	// 初始化分页参数，设置默认值
	model.InitPageReq(&req.PageReq, 1, 10)
	userPayList, count, err := logic.UserPay.UserPayList(ctx, req.PageReq)
	res := make([]v2.UserPayRecordRes, len(userPayList))
	for i, v := range userPayList {
		res[i] = v2.UserPayRecordRes{
			Id:       v.Id,
			PayCode:  v.PayCode,
			UserId:   v.UserId,
			Username: v.Username,
			State:    v.State,
			Price:    v.Price,
			PayAt:    v.PayAt,
			CreateAt: v.CreateAt,
			UpdateAt: v.UpdateAt,
		}
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 用户充值
func (UserPayController) Pay(ctx context.Context, req *v2.UserPayReq) (res *v2.UserPayRes, err error) {
	var userPay do.PayRecord = do.PayRecord{
		UserId: req.UserId,
		Price:  req.Price,
	}
	err = logic.UserPay.UserPay(ctx, userPay)
	return
}
