// 引入axios
import axios from 'axios'
//引入jwt
import jwtDecode from 'jwt-decode'
//引入token续期的api接口
import { refreshToken } from './user'

const requests = axios.create({
  // baseURL: ' http://localhost:8080',
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 5000
})

// 添加请求拦截器
requests.interceptors.request.use(
  (config) => {
    let token = localStorage.getItem('token')

    if (token) {
      const decodedHrader = jwtDecode(token)
      //现在时间
      const now = Date.parse(new Date().toString()) / 1000
      //只要token有效期小于1小时，就会请求refresh接口，重新获取token    3600是60秒*60=1小时
      if (config.url !== '/api/token/refresh' && decodedHrader.exp > now && decodedHrader.exp - 3600 < now) {
        refreshToken()
          .then((res) => {
            console.log('请求', res)
            localStorage.setItem('token', res.data.data.token)
          })
          .finally()
      }
      if (!config.headers) {
        config.headers = {}
      }
      //header头中携带token
      config.headers.Authorization = `Bearer ${token}`
    }

    // 在发送请求之前做的
    return config
  },
  (error) => {
    // 对请求错误做些什么
    return Promise.reject(error)
  }
)

// 验证响应中是否包含错误信息
requests.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    console.log(error)
    return Promise.reject(error)
  }
)

// 对外导出
export default requests
