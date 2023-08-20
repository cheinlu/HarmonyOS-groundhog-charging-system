//统一管理权限相关的接口
import request from "@/utils/request";
//数据类型
import type {permissionList} from './type'
enum API {
//权限列表
ALLPERMI_URL = '/saas-api/permission/list',
//角色绑定权限
BINDPERMI_URL = '/saas-api/permission/bind?',
//当前用户权限
CURRENTPERMI_URL = '/saas-api/permission/current'
}

//权限列表的接口
export const reqAllPermiList = ()=>request.get<any,permissionList>(API.ALLPERMI_URL)
//角色绑定权限的接口
export const reqBindPermi = (roleId:number,permissionList:string[]) => request.post<any,any>(API.BINDPERMI_URL+`roleId=${roleId}&permissionList=${permissionList}`)
//当前用户权限的接口
export const reqCurrentPermi = ()=>request.get<any,any>(API.CURRENTPERMI_URL)