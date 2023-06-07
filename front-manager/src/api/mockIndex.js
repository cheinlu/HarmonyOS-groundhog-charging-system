// 引入axios
import axios from 'axios'

const requests = axios.create({
  baseURL: '/mock',
  timeout: 5000
})

// 添加请求拦截器
requests.interceptors.request.use(
  (config) => {
    let token = localStorage.getItem('token')
    config.headers.Authorization = `Bearer ${token}`
    // 在发送请求之前做些什么
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
