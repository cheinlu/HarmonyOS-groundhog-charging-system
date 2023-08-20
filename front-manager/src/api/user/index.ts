//统一管理用户相关的接口,
import request from "@/utils/request";
//引入数据类型
import type {loginResponseData,loginFormData} from './type'
//统一管理接口,
enum API {
  //登录
  LOGIN_URL = '/saas-api/user/login',
  //续期登录
  REFRESH_URL = '/saas-api/token/refresh',
  
}

//登录接口方法
export const reqLogin = (data:loginFormData) =>
  request.post<any, loginResponseData>(API.LOGIN_URL, data)
//token续期的接口 
export const reqRefresh = () => request.post<any,loginResponseData>(API.REFRESH_URL)
