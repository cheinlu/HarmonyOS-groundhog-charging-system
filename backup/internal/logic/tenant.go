package logic

import (
	"context"
	"login-demo/internal/consts"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type TenantLogic struct {
}

var (
	Tenant = TenantLogic{}
)

func (*TenantLogic) TenantList(ctx context.Context, page model.PageReq) (tenants []entity.Tenant, count int, err error) {
	err = dao.Tenant.Ctx(ctx).Page(page.PageNo, page.PageSize).OrderAsc("update_at").ScanAndCount(&tenants, &count, false)
	return
}

func (*TenantLogic) AddAndGetId(ctx context.Context, tenant *do.Tenant) (tenantId int, err error) {
	tenant.CreateAt, tenant.UpdateAt, tenant.Status = gtime.Now(), gtime.Now(), "active"
	id, err := dao.Tenant.Ctx(ctx).InsertAndGetId(tenant)
	tenantId = int(id)
	return
}

func (*TenantLogic) Del(ctx context.Context, id int) (err error) {
	_, err = dao.Tenant.Ctx(ctx).Delete("id = ?", id)
	return
}

func (*TenantLogic) Update(ctx context.Context, tenant do.Tenant) (err error) {
	tenant.UpdateAt = gtime.Now()
	_, err = dao.Tenant.Ctx(ctx).Update(tenant, "id = ?", tenant.Id)
	return
}

func SetCurrentTenantId(ctx context.Context, tenantId *interface{}) (err error) {
	id, ok := ctx.Value(consts.TenantIDKey).(int)
	if !ok {
		err = gerror.NewCode(gcode.New(1, "系统异常，没有找到租户信息", ""))
		return
	}
	*tenantId = id
	return
}

func (*TenantLogic) GetTenant(ctx context.Context, tenantId int) (tenant entity.Tenant, err error) {
	record, err := dao.Tenant.Ctx(ctx).Where("id", tenantId).One()
	if err != nil {
		err = gerror.WrapCode(gcode.New(1, "系统异常，获取租户信息失败", ""), err)
		return
	}
	if record.IsEmpty() {
		err = gerror.NewCode(gcode.New(404, "未找到该租户信息，请重新登录", ""))
		return
	}
	record.Struct(&tenant)
	return
}
