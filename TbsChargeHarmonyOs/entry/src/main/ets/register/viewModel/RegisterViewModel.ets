import { http } from "@kit.NetworkKit";
import { Apis, requestApiCall, toast } from "@tbs/common";

//注册
export function getRegister(name: string, domain: string, adminUsername: string, adminPassword: string,
  validCode: string) {
  const params: Record<string, Object> = {
    "name": name,
    "domain": domain,
    "adminUsername": adminUsername,
    "adminPassword": adminPassword,
    "validCode": validCode
  }
  return requestApiCall(Apis.USER_REGISTER, http.RequestMethod.POST, params).then(res => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    return res.code
  })
}