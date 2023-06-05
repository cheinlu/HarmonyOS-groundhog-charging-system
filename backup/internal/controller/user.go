package controller

import (
	"context"
	v1 "login-demo/api/v1"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
)

type UserController struct {
}

func (UserController) List(ctx context.Context, req *v1.UserListReq) (pageRes model.PageRes, err error) {
	res := &[]v1.UserListRes{}
	// 初始化分页参数，设置默认值
	model.InitPageReq(&req.PageReq, 1, 10)
	userList, count, err := logic.User.UserList(ctx, req.Username, req.PageReq)
	for _, v := range userList {
		userRes := v1.UserListRes{
			Id:         v.Id,
			Username:   v.Passport,
			Nickname:   v.Nickname,
			Balance:    v.Balance,
			CreateTime: v.CreateAt,
		}
		tmp := append(*res, userRes)
		res = &tmp
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 添加用户
func (UserController) Add(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error) {
	var user do.User = do.User{
		Passport: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
	}
	err = logic.User.Add(ctx, &user)
	return
}

// 删除用户
func (UserController) Del(ctx context.Context, req *v1.UserDelReq) (res *v1.UserDelRes, err error) {
	err = logic.User.Del(ctx, req.Id)
	return
}

func (UserController) Update(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {

	user := do.User{
		Id:       req.Id,
		Passport: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
	}
	err = logic.User.Update(ctx, user)
	return
}
