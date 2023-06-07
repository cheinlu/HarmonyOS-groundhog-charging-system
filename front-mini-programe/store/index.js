import Vue from 'vue'
import Vuex from 'vuex'
//导入购物车的vuex模块
import moduleUser from './user.js'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules:{
    'm_user':moduleUser
  }
})

export default store