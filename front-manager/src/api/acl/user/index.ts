//统一管理系统设置相关的接口
import request from "@/utils/request";
import type {UserResponseData,userAddOrUp,payRecord,userPay} from './type'
//统一管理接口
enum API {
  //获取全部已有用户账号信息
  ALLUSER_URL = '/saas-api/user/list?',
  //添加新用户
  ADDUSER_URL = '/saas-api/user/add',
  //更新已有的用户
  UPDATEUSER_URL = '/saas-api/user/update',
  //删除用户
  DELETEUSER_URL = ' /saas-api/user/delete?',
  //用户充值
  USERPAY = '/saas-api/user/pay',
  //充值列表数据
  PAYRECORD_URL = '/saas-api/user/pay/record?',
  //
}

//获取用户账号信息的接口
export const reqUserInfo = (pageNo:number,pageSize:number,username:string)=>request.get<any,UserResponseData>(API.ALLUSER_URL+`pageNo=${pageNo}&pageSize=${pageSize}&username=${username}`)
//添加|更新用户的接口
export const reqAddOrUpdateUser = (data:userAddOrUp)=>{
  if(data.id){
    return request.post<any,any>(API.UPDATEUSER_URL,data)
  }else{
    return request.post<any,any>(API.ADDUSER_URL,data)
  }
}
//删除用户的接口
export const reqRemoveUser = (id:number)=>request.delete(API.DELETEUSER_URL+`id=${id}`)
//用户充值
export const reqUserPay = (data:userPay)=>request.post(API.USERPAY,data)
//充值列表
export const reqPayList = (pageNo:number,pageSize:number)=>request.get<any,payRecord>(API.PAYRECORD_URL+`pageNo=${pageNo}&pageSize=${pageSize}`)
