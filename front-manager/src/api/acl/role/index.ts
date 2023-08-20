//统一管理角色相关的接口
import request from "@/utils/request";
//数据类型
import type {roleList,roleAddOrUp} from './type'
enum API {
  //获取全部的职位接口
  ALLROLE_URL = '/saas-api/role/list?',
  //新增角色的接口地址
  ADDROLE_URL = '/saas-api/role/add',
  //更新已有的职位
  UPDATEROLE_URL = '/saas-api/role/update',
  //删除已有的职位
  REMOVEROLE_URL = '/saas-api/role/delete?'
}

//获取全部的角色
export const reqAllRoleList = (pageNo:number,pageSize:number,name:string)=>request.get<any,roleList>(API.ALLROLE_URL+`pageNo=${pageNo}&pageSize=${pageSize}&name=${name}`)
//添加职位与更新已有职位接口
export const reqAddOrUpdateRole = (data:roleAddOrUp) => {
  if (data.id) {
    return request.post<any, any>(API.UPDATEROLE_URL, data)
  } else {
    return request.post<any, any>(API.ADDROLE_URL, data)
  }
}
//删除已有的职位
export const reqRemoveRole = (id: number) =>
  request.delete<any, any>(API.REMOVEROLE_URL + `id=${id}`)