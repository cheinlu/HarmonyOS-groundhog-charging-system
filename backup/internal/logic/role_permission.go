package logic

import (
	"context"
	"login-demo/internal/dao"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gtime"
)

type LogicRolePermission struct {
}

var (
	RolePermission = LogicRolePermission{}
)

func ConfigScan(a func(*gcfg.Config) error) error {
	Config, _ := gcfg.New()
	return a(Config)
}

func (*LogicRolePermission) Bind(ctx context.Context, rolePermission *do.RolePermission) (err error) {
	rolePermission.UpdateAt = gtime.Now()
	_, err = dao.RolePermission.Ctx(ctx).Where("role_id", rolePermission.RoleId.(int)).One()
	if err != nil {
		return
	}
	err = SetCurrentTenantId(ctx, &rolePermission.TenantId)
	if err != nil {
		return
	}
	_, err = dao.RolePermission.Ctx(ctx).Replace(rolePermission)
	return
}

func (*LogicRolePermission) GetByRoleId(ctx context.Context, roleId int) (rolePermission entity.RolePermission, err error) {
	rolePermission.UpdateAt = gtime.Now()
	userRecord, err := dao.RolePermission.Ctx(ctx).Where("role_id", roleId).One()
	if err != nil {
		return
	}
	if userRecord.IsEmpty() {
		err = gerror.NewCode(gcode.New(404, "没有找到权限信息，请检查角色是否绑定权限", ""))
		return
	}
	userRecord.Struct(&rolePermission)
	return
}
