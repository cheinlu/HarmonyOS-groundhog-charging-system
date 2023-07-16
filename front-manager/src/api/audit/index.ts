//审计的接口
import request from "@/utils/request";
//数据类型
enum API {
AUDITLOG_URL = '/api/audit/log/list?'
}
//审计的接口
export const reqAuditList = (pageNo:number,pageSize:number,operationTypeL:string,ChangedUser:string)=>request.get(API.AUDITLOG_URL+`pageNo=${pageNo}&pageSize=${pageSize}&operationTypeL=${operationTypeL}&ChangedUser=${ChangedUser}`)
// export const reqAuditList = (params:any)=>request.get(API.AUDITLOG_URL,params)