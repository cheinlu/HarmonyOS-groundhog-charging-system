package logic

import (
	"context"
	"login-demo/internal/dao"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
)

type WxUserLogic struct{}

var WxUser WxUserLogic

func (*WxUserLogic) GetUserByOpenID(ctx context.Context, openID string) (*entity.User, error) {
	var user entity.User
	err := dao.WxUser.Ctx(ctx).InnerJoin("user", "user.id=wx_user.user_id").Fields("user.*").Where("open_id=?", openID).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*WxUserLogic) Update(ctx context.Context, wxUser do.WxUser) error {
	_, err := dao.WxUser.Ctx(ctx).Update(wxUser, "wx_user.user_id=?", wxUser.UserId)
	return err
}
