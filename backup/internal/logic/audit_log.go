package logic

import (
	"context"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/entity"
)

type AuditLogLogic struct {
}

var AuditLog AuditLogLogic

func (AuditLogLogic) List(ctx context.Context, page model.PageReq, auditLogQuery entity.AuditLog) (auditLogs []entity.AuditLog, count int, err error) {
	auditLogModel := dao.AuditLog.Ctx(ctx).Page(page.PageNo, page.PageSize)
	if auditLogQuery.OperationType != "" {
		auditLogModel = auditLogModel.WhereLike("operation_type", "%"+auditLogQuery.OperationType+"%")
	}
	if auditLogQuery.ChangedUser != "" {
		auditLogModel = auditLogModel.WhereLike("changed_user", "%"+auditLogQuery.ChangedUser+"%")
	}
	err = auditLogModel.ScanAndCount(&auditLogs, &count, false)
	return
}
