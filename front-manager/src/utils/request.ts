import axios from "axios";
//引入token
import {SET_TOKEN} from '@/utils/token'
//引入jwt
import jwtDecode from 'jwt-decode';
//引入token续期的api接口
import {reqRefresh} from '@/api/user/index.ts'
//引入用户相关的仓库
import useUserStore from '@/store/module/use.ts'
let request = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API,
  timeout: 10000
})

// 添加请求拦截器
request.interceptors.request.use((config)=>{
//获取用户相关的小仓库:获取仓库内部token和租户id,登录成功以后携带给服务器
const userStore = useUserStore()
let token = userStore.token
let id = userStore.id
if (token) {
  const decodedHrader:any = jwtDecode(token)
  //现在时间
  const now = Math.floor(Date.now() / 1000)
  
  //只要token有效期小于1小时，就会请求refresh接口，重新获取token    3600是60秒*60=1小时
  if(config.url !== '/saas-api/token/refresh' && decodedHrader.exp > now && decodedHrader.exp - 3600 < now){
    reqRefresh().then(res=>{
      SET_TOKEN(res.data.token)
    }).finally()
  }
  //header头中携带token
  config.headers.Authorization = `Bearer ${token}`
}

if(userStore.id){
  config.headers['X-Tenant-ID'] = id;
}
// 在发送请求之前做的
return config
})

//响应拦截器
request.interceptors.response.use((response)=>{
  return response.data
},
(error) => {
  return Promise.reject(error)
}
)

// 对外导出
export default request