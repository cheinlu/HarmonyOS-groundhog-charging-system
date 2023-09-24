package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
)

type UserController struct {
}

func (UserController) List(ctx context.Context, req *v2.UserListReq) (pageRes model.PageRes, err error) {
	// 初始化分页参数，设置默认值
	model.InitPageReq(&req.PageReq, 1, 10)
	userList, count, err := logic.User.UserList(ctx, req.Username, req.PageReq)
	res := make([]v2.UserListRes, len(userList))
	for i, v := range userList {
		userRes := v2.UserListRes{
			Id:         v.Id,
			Username:   v.Passport,
			Nickname:   v.Nickname,
			Balance:    v.Balance,
			RoleName:   v.RoleName,
			CreateTime: v.CreateAt,
		}
		res[i] = userRes
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 添加用户
func (UserController) Add(ctx context.Context, req *v2.UserAddReq) (res *v2.UserAddRes, err error) {
	var user do.User = do.User{
		Passport: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
	}
	err = logic.SetCurrentTenantId(ctx, &user.TenantId)
	if err != nil {
		return
	}
	err = logic.User.AddAndWallet(ctx, &user)
	return
}

// 删除用户
func (UserController) Del(ctx context.Context, req *v2.UserDelReq) (res *v2.UserDelRes, err error) {
	err = logic.User.Del(ctx, req.Id)
	return
}

func (UserController) Update(ctx context.Context, req *v2.UserUpdateReq) (res *v2.UserUpdateRes, err error) {

	user := do.User{
		Id:       req.Id,
		Passport: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
	}
	err = logic.User.Update(ctx, user)
	return
}
