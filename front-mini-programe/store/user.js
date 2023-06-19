export default {
  // 开启命名空间
  namespaced: true,
  //数据
  state: () => ({
    token: uni.getStorageSync('token') || '',
    userInfo: uni.getStorageSync('userInfo') || '{}'
  }),
  //方法
  mutations: {
    //更改userInfo
    updateUserInfo(state, userInfo) {
      state.userInfo = userInfo
      this.commit('m_user/saveUserInfoToStorage')
    },
    //持久化存储用户信息
    saveUserInfoToStorage(state) {
      uni.setStorageSync('userInfo', JSON.stringify(state.userInfo))
    },
    //更改token
    updateToken(state, token) {
      state.token = token
      this.commit('m_user/saveTokenToStorage')
    },
    //持久化存储token
    saveTokenToStorage(state) {
      uni.setStorageSync('token', state.token)
    },
    //清除token和userInfo
    clearUserInfoAndToken(state){
      state.token = ''
      state.userInfo = {}
      uni.removeStorageSync('token')
      uni.removeStorageSync('userInfo')
    }
  },
  getters: {}
}
