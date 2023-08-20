//租户管理相关的数据接口，包含租户注册，选择租户登录和租户列表数据
import request from "@/utils/request";
//引入数据类型
import type {registerFormData,tenantSelect,tenantList,tenant} from './type'
//统一管理接口,
enum API {
  //租户注册
  TE_REGISTER_URL = '/saas-api/tenant/register',
  //租户登录，选择租户
  TE_LOGIN_URL = '/saas-api/tenant/selects',
  //租户列表数据
  TE_LIST_URL = '/saas-api/tenant/list',
  //添加租户数据
  TE_ADDTE_URL = '/saas-api/tenant/add',
  //修改租户数据
  TE_UPDATE_URL = '/saas-api/tenant/update',
  //删除租户数据
  TE_DELETE_URL = '/saas-api/tenant/delete?'
}
//暴露请求函数
//租户注册接口
export const reqRegister = (data:registerFormData)=>request.post<any,any>(API.TE_REGISTER_URL,data)
//租户登录
export const reqTeLogin =()=>request.get<any,tenantSelect>(API.TE_LOGIN_URL)
//租户列表数据的接口
export const reqTeList = ()=>request.get<any,tenantList>(API.TE_LIST_URL)
//添加|修改租户数据的接口
export const reqAddOrUpdateTe = (data:tenant)=>{
  if(data.id){
    return request.post(API.TE_UPDATE_URL,data)
  }else{
    return request.post(API.TE_ADDTE_URL,data)
  }
}
//删除租户
export const reqDeleteTe = (id:number)=>request.delete<any,any>(API.TE_DELETE_URL+`id=${id}`)