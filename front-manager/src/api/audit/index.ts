//审计的接口
import request from "@/utils/request";
//数据类型
enum API {
AUDITLOG_URL = '/saas-api/audit/log/list?'
}
//审计的接口
export const reqAuditList = (params:any)=>request.get(API.AUDITLOG_URL,params)