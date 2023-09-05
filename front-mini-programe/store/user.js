import { defineStore } from 'pinia'
//引入本地存储
import {SET_TOKEN,GET_TOKEN,REMOVE_TOKEN,SET_INFO,GET_INFO,REMOVE_INFO} from '@/utils/token'
//创建用户小仓库
const useUserStore = defineStore('User', {
  state: () => {
    return {
      token: GET_TOKEN() || '',
      info: GET_INFO() || '{}',
      activeTab: 0, // 默认选中的索引
    }
  },
  actions: {
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
})

//暴露用户小仓库
export default useUserStore
