package logic

import (
	"context"
	"fmt"
	"login-demo/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

type LogicContext struct {
}

var Ctx LogicContext = LogicContext{}

const (
	HTTPContextKey = "paramsCtx"
)

type HTTPContextParams struct {
	Username string
}

func (LogicContext) SetHTTPContextParams(r *ghttp.Request, username string) {
	r.SetCtxVar(HTTPContextKey, HTTPContextParams{
		Username: username,
	})
}

func (LogicContext) GetHTTPContextParams(ctx context.Context) *HTTPContextParams {
	v, ok := ctx.Value(HTTPContextKey).(HTTPContextParams)
	if ok {
		return &v
	} else {
		return nil
	}
}

func (l LogicContext) GetCurrentUser(ctx context.Context) (user model.UserMore, err error) {
	username := l.GetHTTPContextParams(ctx).Username
	users, count, err := User.UserList(ctx, username, model.PageReq{})
	if count != 1 || err != nil {
		err = gerror.NewCode(gcode.New(1, "系统异常", fmt.Sprintf("%s 用户不存在", username)))
		return
	}
	user = users[0]
	return
}
