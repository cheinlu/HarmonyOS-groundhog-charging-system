import store from '@/store/index.js'
module.exports = () => {
  // 初始化请求配置
  uni.$u.http.setConfig(config => {
    // config.baseURL = "http://127.0.0.1:8000"; /* 根域名 */
    config.baseURL = 'https://7483f88q46.goho.co' /* 根域名 */
    return config
  })

  // 请求拦截
  uni.$u.http.interceptors.request.use(
    async config => {
      const token = await getToken() // 获取 token，异步方式
      // 如果存在 token，则将其添加到请求头中
      if (token) {
        config.header.Authorization = `Bearer ${token}`
      }
      return config
    },
    config => {
      return Promise.reject(config)
    }
  )

  // 响应拦截，主要用于处理 token 过期的情况
  uni.$u.http.interceptors.response.use(
    async response => {
      // 判断响应头和响应数据中的状态码是否为 50008，避免误判
      if (response.header['X-Status-Code'] === 50008 || response.data.code === 50008) {
        // 提示用户重新登录
        await showReLoginDialog()
        // 尝试重新发起之前的请求并获取响应数据
        const newResponse = await retryRequest(response.config)
        return newResponse
      } else {
        // 如果返回状态正常，则直接返回响应数据
        return response
      }
    },
    error => {
      return Promise.reject(error)
    }
  )
}

// 封装获取 token 的函数，异步方式
function getToken() {
  return new Promise(resolve => {
    uni.getStorage({
      key: 'token',
      success: res => {
        resolve(res.data)
      },
      fail: () => {
        resolve(null)
      }
    })
  })
}

// // 封装弹窗提示用户重新登录的函数，异步方式
function showReLoginDialog() {
  return new Promise(resolve => {
    uni.showToast({
      icon: 'none',
      title: '登录已过期，请重新登录',
      success: () => {
        store.commit('m_user/clearUserInfoAndToken')
        uni.switchTab({
          url: '/pages/my/my'
        })
        resolve()
      }
    })
  })
}
