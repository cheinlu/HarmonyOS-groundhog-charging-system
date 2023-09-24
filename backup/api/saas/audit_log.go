package v2_saas

import (
	"login-demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AuditLogListReq struct {
	g.Meta        `path:"/audit/log/list" tags:"审计" method:"get" summary:"审计操作列表"`
	OperationType string `json:"operationType" `   //
	ChangedUser   string `json:"ChangedUser"     ` //
	model.PageReq
}

type AuditLogListRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
