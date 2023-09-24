package logic

import (
	"context"
	"fmt"
	"login-demo/internal/consts"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type LogicRole struct {
}

var (
	Role = LogicRole{}
)

func (*LogicRole) RoleList(ctx context.Context, rolename string, page model.PageReq) (roles []entity.Role, count int, err error) {
	model := dao.Role.Ctx(ctx)
	if rolename != "" {
		model = model.WhereLike("name", fmt.Sprintf("%%%s%%", rolename))
	}
	err = model.OrderDesc("role.update_at").Page(page.PageNo, page.PageSize).ScanAndCount(&roles, &count, false)
	return
}

func (*LogicRole) Add(ctx context.Context, role *do.Role) (err error) {
	role.CreateAt, role.UpdateAt = gtime.Now(), gtime.Now()
	err = SetCurrentTenantId(ctx, &role.TenantId)
	if err != nil {
		return
	}
	_, err = dao.Role.Ctx(ctx).Insert(role)
	return
}

func (*LogicRole) Del(ctx context.Context, id int) (err error) {
	_, err = dao.Role.Ctx(ctx).Delete("id = ?", id)
	return
}

func (*LogicRole) Update(ctx context.Context, role do.Role) (err error) {
	role.UpdateAt = gtime.Now()
	_, err = dao.Role.Ctx(ctx).Update(role, "id = ?", role.Id)
	return
}

// 从 context 取出 Role 信息
func (*LogicRole) GetCurrentRole(ctx context.Context) (role entity.Role, err error) {
	user, err := User.GetCurrentUser(ctx)
	if err != nil {
		err = gerror.NewCode(gcode.New(1, "系统异常，未找到登录用户信息", ""))
		return
	}
	record, err := dao.Role.Ctx(ctx).Where("id", user.RoleId).One()
	if err != nil {
		err = gerror.NewCode(gcode.New(1, "系统异常，未找到登录角色信息", ""))
		return
	}
	record.Struct(&role)
	return
}

// 初始化租户管理员权限
func (*LogicRole) InitTenantAdmin(ctx context.Context, tenantId int) (roleId int, err error) {
	insertRole := do.Role{
		Name:     consts.AdminRoleName,
		TenantId: tenantId,
	}
	id, err := dao.Role.Ctx(ctx).InsertAndGetId(insertRole)
	roleId = int(id)
	if err != nil {
		err = gerror.WrapCode(gcode.New(1, "系统异常，添加角色失败", ""), err)
		return
	}
	insertRolePermission := do.RolePermission{
		RoleId:      roleId,
		TenantId:    tenantId,
		Permissions: consts.AdminRolePermission,
		CreateAt:    gtime.Now(),
		UpdateAt:    gtime.Now(),
	}
	_, err = dao.RolePermission.Ctx(ctx).Insert(insertRolePermission)
	if err != nil {
		err = gerror.WrapCode(gcode.New(1, "系统异常，添加权限失败", ""), err)
		return
	}
	return
}
