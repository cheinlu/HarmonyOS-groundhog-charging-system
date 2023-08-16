import App from './App'
import uviewPlus from '@/uni_modules/uview-plus'
import * as Pinia from 'pinia'

// 封装弹框的方法
uni.$showMsg = function (title = '数据请求失败！', duration = 1500) {
  uni.showToast({
    title,
    duration,
    icon: 'none'
  })
}
// #ifndef VUE3
import Vue from 'vue'
import './uni.promisify.adaptor'
Vue.config.productionTip = false
App.mpType = 'app'
const app = new Vue({
  ...App
})
app.$mount()
// #endif

// #ifdef VUE3
import { createSSRApp } from 'vue'
export function createApp() {
  const app = createSSRApp(App)
  app.use(uviewPlus),
  app.use(Pinia.createPinia())
  return {
    app,
    Pinia
  }
}
// #endif
