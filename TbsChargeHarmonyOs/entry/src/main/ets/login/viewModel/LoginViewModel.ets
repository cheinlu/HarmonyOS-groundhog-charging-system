import { http } from "@kit.NetworkKit";
import { LoginData, TenantsData } from "../model/LoginModel";
import { Apis, StorageConst, commonStorage, requestApiCall, toast } from "@tbs/common";
import { UserInfoData } from "@tbs/mine/src/main/ets/model/MineModel";
import { BusinessError } from "@kit.BasicServicesKit";

//获取租户ID
export function getTenant() {
  return requestApiCall(Apis.USER_TENANT, http.RequestMethod.GET).then((res) => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    const data = res.data as TenantsData
    return data.Data
  }).catch((err: BusinessError) => {
    console.error(`lucy++ 获取租户ID接口调用失败，${err.code}+${err.message}`)
  })
}

//获取token-登录
export function getToken(tenant: string, username: string, password: string) {
  const headers: Record<string, string> = {
    'X-Tenant-ID': tenant,
  }
  const params: Record<string, Object> = {
    "username": username,
    "password": password
  }
  return requestApiCall(Apis.USER_LOGIN, http.RequestMethod.POST, params, headers).then(res => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    const data = res.data as LoginData
    let token = data.token
    console.log('lucy== token', token)
    if (token) {
      commonStorage.saveValueSync(StorageConst.TOKEN, token)
      //登录成功后，获取下用户信息
      getUserInfo()
    }
    return token
  }).catch((err: BusinessError) => {
    console.error(`lucy++ 获取登录接口调用失败，${err.code}+${err.message}`)
  })
}

//获取用户信息
export function getUserInfo() {
  return requestApiCall(Apis.USER_INFO, http.RequestMethod.GET).then(res => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    const data = res.data as UserInfoData
    if (data) {
      commonStorage.saveValueSync(StorageConst.USER_INFO, data)
    }
  }).catch((err: BusinessError) => {
    console.error(`lucy++ 获取用户信息接口调用失败，${err.code}+${err.message}`)
  })
}