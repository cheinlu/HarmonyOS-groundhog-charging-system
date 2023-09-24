package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type RoleController struct {
}

// 充电站列表
func (RoleController) List(ctx context.Context, req *v2.RoleListReq) (pageRes model.PageRes, err error) {
	model.InitPageReq(&req.PageReq, 1, 10)
	roles, count, err := logic.Role.RoleList(ctx, req.Name, req.PageReq)
	res := make([]*v2.RoleListRes, len(roles))
	for i, role := range roles {
		userCount, _ := logic.User.UserCount(ctx, do.User{
			RoleId: role.Id,
		})
		res[i] = &v2.RoleListRes{
			Id:        role.Id,
			Name:      role.Name,
			BindUsers: userCount,
			CreateAt:  role.CreateAt.Time,
			UpdateAt:  role.UpdateAt.Time,
		}
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 充电站添加
func (RoleController) Add(ctx context.Context, req *v2.RoleAddReq) (res *v2.RoleAddRes, err error) {
	if req.Name == "" {
		err = gerror.NewCode(gcode.New(401, "角色名不为空", ""))
		return
	}
	tenantId, ok := ctx.Value(consts.TenantIDKey).(int)
	if !ok {
		err = gerror.NewCode(gcode.New(401, "tenantId 不为空", ""))
		return
	}
	var role do.Role = do.Role{
		Name:     req.Name,
		TenantId: tenantId,
	}
	err = logic.Role.Add(ctx, &role)
	return
}

// 删除充电站
func (RoleController) Del(ctx context.Context, req *v2.RoleDelReq) (res *v2.RoleDelRes, err error) {
	if req.Id == 0 {
		err = gerror.NewCode(gcode.New(401, "删除角色失败，id 不能为空", ""))
		return
	}
	err = logic.Role.Del(ctx, req.Id)
	return
}

func (RoleController) Update(ctx context.Context, req *v2.RoleUpdateReq) (res *v2.RoleUpdateRes, err error) {
	if req.Id == 0 {
		err = gerror.NewCode(gcode.New(401, "修改角色失败，id 不能为空", ""))
		return
	}
	role := do.Role{
		Id:   req.Id,
		Name: req.Name,
	}
	err = logic.Role.Update(ctx, role)
	return
}
