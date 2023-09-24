package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"

	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"login-demo/internal/model/do"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
)

type PermissionController struct {
}

func (PermissionController) List(ctx context.Context, req *v2.PermissionListReq) (rep v2.PermissionListRes, err error) {
	Config, _ := gcfg.New()
	err = Config.MustGet(ctx, "permissionList").Scan(&rep.Data)
	// err = logic.ConfigScan(func(c *gcfg.Config) error {
	// 	return c.MustGet(ctx, "permissionList").Scan(&rep.Data)
	// })
	tenantId, ok := ctx.Value(consts.TenantIDKey).(int)
	if !ok {
		err = gerror.NewCode(gcode.New(401, "tenantId 不为空", ""))
	}
	if tenantId != 1 {
		tenantIndex := 0
		for i, val := range rep.Data {
			if val.Model == "租户管理" {
				tenantIndex = i
			}
		}
		tmp := make([]v2.PermissionList, len(rep.Data)-1)
		copy(tmp, rep.Data[:tenantIndex])
		copy(tmp, rep.Data[tenantIndex+1:])
		rep.Data = tmp
	}
	return
}

func (PermissionController) Bind(ctx context.Context, req *v2.BindPermissionReq) (rep *v2.BindPermissionRes, err error) {
	err = logic.RolePermission.Bind(ctx, &do.RolePermission{
		RoleId:      req.RoleId,
		Permissions: strings.Join(req.PermissionList, ","),
	})
	return
}

func (PermissionController) CurrentPermissions(ctx context.Context, req *v2.CurrentPermissionReq) (rep *v2.CurrentPermissionRes, err error) {
	user, err := logic.User.GetCurrentUser(ctx)
	if err != nil {
		return
	}
	rolePermission, err := logic.RolePermission.GetByRoleId(ctx, user.RoleId)
	if err != nil {
		return
	}
	rep = &v2.CurrentPermissionRes{
		PermissionList: strings.Split(rolePermission.Permissions, ","),
	}
	return
}
