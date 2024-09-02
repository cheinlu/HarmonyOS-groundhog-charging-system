
import useUserStore from '@/store/user.js'
export const setRequestConfig = () => {
  uni.$u.http.setConfig((config) => {
    /* config 为默认全局配置*/
    config.baseURL = 'http://101.132.109.52/' /* 本地根域名 */
    return config
  })
  // 请求拦截
  uni.$u.http.interceptors.request.use(
    (config) => {
		config.header['bypass-tunnel-reminder'] = '1234'  // 添加自定义头部
      let token = useUserStore().token
      if (token) {
        config.header.Authorization = `Bearer ${token}`
      }
      return config
    },
    (error) => {
      return Promise.reject(error)
    }
  )
  // 响应拦截
  uni.$u.http.interceptors.response.use(
    (response) => {
      if (response.data.code == 401) {
        // 提示重新登录
        uni.$showMsg('请登录')
        useUserStore().userLogout()
        setTimeout(() => {
          uni.switchTab({
            url: '/pages/my/my'
          })
        }, 1000);
        
      }
      return response
    },
    (error) => {
      return Promise.reject(error)
    }
  )
}
