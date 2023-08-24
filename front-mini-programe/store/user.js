import { defineStore } from 'pinia'
//引入接口
import { requestLogin } from '@/utils/api/user.js'
//引入本地存储
import {SET_TOKEN,GET_TOKEN,REMOVE_TOKEN,SET_INFO,GET_INFO,REMOVE_INFO} from '@/utils/token'
//创建用户小仓库
const useUserStore = defineStore('User', {
  state: () => {
    return {
      token: GET_TOKEN() || '',
      info: GET_INFO() || '{}',
      activeTab: 1 // 默认选中的索引
    }
  },
  actions: {
    //用户登录
    async userLogin(lcode) {
      const code = { code: lcode }
      let { data: res } = await requestLogin(code)
      if (res.code == 0) {
        this.token = res.data.token
        SET_TOKEN(res.data.token)
      } else {
        uni.$showMsg(res.message)
        uni.showModal({
          title: '提示',
          content: '请注册用户',
          success: function (res) {
            if (res.confirm) {
              uni.navigateTo({
                url: '/subpkg/register/register'
              })
            } else if (res.cancel) {
              console.log('用户点击取消')
            }
          }
        })
      }
    },
    //获取用户信息
    updateUserInfo(info) {
      SET_INFO(JSON.stringify(info))
    },
    //获取token
    updateToken(token) {
      SET_TOKEN(token)
    },
    //退出登录
    userLogout() {
      this.token = ''
      this.info = {}
      REMOVE_TOKEN()
      REMOVE_INFO()
    },
    //设置active的值
    setActive(active) {
      this.activeTab = active;
    },
  },
  getters: {}
})

//暴露用户小仓库
export default useUserStore
