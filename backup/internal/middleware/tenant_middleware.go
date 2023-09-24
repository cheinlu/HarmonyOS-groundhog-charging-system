package middleware

import (
	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"strconv"

	"github.com/gogf/gf/v2/net/ghttp"
)

func TenantMiddleware(req *ghttp.Request) {
	// 提取租户信息，例如从请求头或URL参数中获取
	tenantID := req.GetHeader("X-Tenant-ID")

	if tenantID == "" {
		handleError(req, 401, nil, "缺少租户ID")
		return
	}
	id, err := strconv.Atoi(tenantID)
	if err != nil {
		handleError(req, 401, nil, "租户ID格式错误，不是数字")
		return
	}
	// 将租户ID存储在请求的上下文中
	req.SetCtxVar(consts.TenantIDKey, id)

	_, err = logic.Tenant.GetTenant(req.Context(), id)
	if err != nil {
		handleError(req, 404, err, "未找到该租户信息，请重新登录")
		return
	}

	// 继续处理请求
	req.Middleware.Next()
}
