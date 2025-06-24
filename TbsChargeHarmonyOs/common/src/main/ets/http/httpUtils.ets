import http from '@ohos.net.http';
import { commonStorage } from '../../../../Index';
import { RoutePageConst, StorageConst } from '../manager';
import { ResponseResult } from './response';
import { router } from '@kit.ArkUI';
import { BusinessError } from '@kit.BasicServicesKit';

export function requestApiCall(url: string,
  method: http.RequestMethod,
  data?: Record<string, number | boolean | string | Object | undefined>,
  headers?: Record<string, string>,
): Promise<ResponseResult> {
  const BASE_URL = 'https://charging.lusson.xyz'
  const httpRequest = http.createHttp();
  const defaultHeaders: Record<string, string> = {
    'Content-Type': 'application/json',
  }
  const token = commonStorage.getStringValueSync(StorageConst.TOKEN, "")
  if (token) {
    defaultHeaders['Authorization'] = `Bearer ${token}`
  }

  const finalHeaders: Record<string, string> = {};
  Object.keys(defaultHeaders).forEach(key => {
    finalHeaders[key] = defaultHeaders[key];
  });
  if (headers) {
    Object.keys(headers).forEach(key => {
      finalHeaders[key] = String(headers[key]);
    });
  }

  let responseResult = httpRequest.request(BASE_URL + url, {
    method: method,
    extraData: data ? JSON.stringify(data) : undefined,
    header: finalHeaders,
    connectTimeout: 15000
  });
  let response: ResponseResult = new ResponseResult();
  return responseResult.then((value: http.HttpResponse) => {
    if (value.responseCode === 200) {
      let res: ResponseResult = JSON.parse(`${value.result}`);
      console.log('lucy++ res呀呀呀', res.code, res.message)
      response.data = res.data;
      response.code = res.code;
      response.message = res.message;
      if (res.code === 401 || (res.message && res.message.includes('token is expired'))) {
        commonStorage.deleteValueSync(StorageConst.TOKEN)
        commonStorage.deleteValueSync(StorageConst.USER_INFO)

        // 跳转到登录页面
        router.clear();
        router.replaceUrl({ url: RoutePageConst.LoginPage })
      }
    } else {
      response.message = '请求错误';
      response.code = 400;
    }
    return response;
  }).catch((error: BusinessError) => {
    console.log('lucy++ httpUtils error', JSON.stringify(error))
    response.message = error.message;
    response.code = error.code;
    return response;
  })
}