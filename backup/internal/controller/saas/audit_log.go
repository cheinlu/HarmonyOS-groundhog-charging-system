package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/entity"
)

type AuditLogController struct {
}

func (AuditLogController) List(ctx context.Context, req *v2.AuditLogListReq) (pageRes model.PageRes, err error) {
	// 初始化分页参数，设置默认值
	model.InitPageReq(&req.PageReq, 1, 10)
	auditLogs, count, err := logic.AuditLog.List(ctx, req.PageReq, entity.AuditLog{
		OperationType: req.OperationType,
		ChangedUser:   req.ChangedUser,
	})
	pageRes.List, pageRes.TotalCount, pageRes.PageNo, pageRes.PageSize = auditLogs, count, req.PageNo, req.PageSize
	return
}
