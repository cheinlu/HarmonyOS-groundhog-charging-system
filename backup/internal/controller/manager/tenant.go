package controller_manager

import (
	"context"
	"fmt"
	v2 "login-demo/api/manager"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 租户信息
type TenantController struct {
}

// 租户筛选框信息
func (TenantController) SelectList(ctx context.Context, req *v2.TenantSelectListReq) (res v2.TenantSelectListRes, err error) {
	model.InitPageReq(&req.PageReq, 1, 10000)
	tenants, _, err := logic.Tenant.TenantList(ctx, req.PageReq)
	res.Data = make([]v2.TenantSelectList, len(tenants))
	for i, tenant := range tenants {
		res.Data[i] = v2.TenantSelectList{
			Id:     tenant.Id,
			Name:   tenant.Name,
			Domain: tenant.Domain,
		}
	}
	return
}

// 租户注册
func (TenantController) Register(ctx context.Context, req *v2.TenantRegisterReq) (res v2.TenantRegisterRes, err error) {
	if req.ValidCode != "666666" {
		err = gerror.NewCode(gcode.New(1, "验证码错误", "验证码错误"))
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		tenantId, err := logic.Tenant.AddAndGetId(ctx, &do.Tenant{
			Name:   req.Name,
			Domain: req.Domain,
		})
		if err != nil {
			return err
		}
		// 生成随机数种子
		rand.Seed(time.Now().UnixNano())

		// 生成月日字符串
		now := time.Now()
		month := now.Month()
		day := now.Day()
		// 组合成新用户字符串
		newUserStr := fmt.Sprintf("新用户%s-%s", fmt.Sprintf("%02d%02d", month, day), fmt.Sprintf("%04d", rand.Intn(10000)))
		_, count, err := logic.User.UserList(ctx, req.AdminUsername, model.PageReq{
			PageNo:   1,
			PageSize: 10,
		})
		if count > 0 {
			err = gerror.NewCode(gcode.New(611, "该用户名已经注册，请换一个用户名", ""))
			return err
		}
		if err != nil {
			return err
		}
		roleId, err := logic.Role.InitTenantAdmin(ctx, tenantId)
		if err != nil {
			return err
		}
		err = logic.User.AddAndWallet(ctx, &do.User{
			Passport: req.AdminUsername,
			Password: req.AdminPassword,
			Nickname: newUserStr,
			TenantId: tenantId,
			RoleId:   roleId,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// 租户列表
func (TenantController) List(ctx context.Context, req *v2.TenantListReq) (pageRes model.PageRes, err error) {
	model.InitPageReq(&req.PageReq, 1, 10)
	tenants, count, err := logic.Tenant.TenantList(ctx, req.PageReq)
	res := make([]*v2.TenantListRes, len(tenants))
	for i, tenant := range tenants {
		res[i] = &v2.TenantListRes{
			Id:           tenant.Id,
			Name:         tenant.Name,
			Domain:       tenant.Domain,
			ContactName:  tenant.ContactName,
			ContactEmail: tenant.ContactEmail,
			ContactPhone: tenant.ContactPhone,
			Status:       tenant.Status,
			CreateAt:     tenant.CreateAt.Time,
			UpdateAt:     tenant.UpdateAt.Time,
		}
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 租户添加
func (TenantController) Add(ctx context.Context, req *v2.TenantAddReq) (res *v2.TenantAddRes, err error) {
	var tenant do.Tenant = do.Tenant{
		Name:         req.Name,
		Domain:       req.Domain,
		ContactName:  req.ContactName,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
		Status:       req.Status,
	}
	_, err = logic.Tenant.AddAndGetId(ctx, &tenant)
	return
}

// 删除租户
func (TenantController) Del(ctx context.Context, req *v2.TenantDelReq) (res *v2.TenantDelRes, err error) {
	err = logic.Tenant.Del(ctx, req.Id)
	return
}

// 修改租户
func (TenantController) Update(ctx context.Context, req *v2.TenantUpdateReq) (res *v2.TenantUpdateRes, err error) {
	tenant := do.Tenant{
		Id:           req.Id,
		Name:         req.Name,
		Domain:       req.Domain,
		ContactName:  req.ContactName,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
		Status:       req.Status,
	}
	err = logic.Tenant.Update(ctx, tenant)
	return
}
