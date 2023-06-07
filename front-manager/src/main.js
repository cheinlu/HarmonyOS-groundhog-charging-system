import Vue from 'vue'
import App from './App.vue'
import router from './router'
//引入全局样式
import './assets/css/global.css'
//引入element
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)
// 组件实例的原型的原型指向的是Vue.prototype，任意组件可以使用API相关的接口
import API from '@/api'
Vue.prototype.$API = API
//引入百度地图
import BaiduMap from 'vue-baidu-map'
Vue.use(BaiduMap, {
  // ak 是在百度地图开发者平台申请的密钥 详见 http://lbsyun.baidu.com/apiconsole/key */
  ak: '124wXbeEy3ngDIaCEzUMlu6DA7gPwGsF'
})

//引入MockServer.js---mock数据
import './mock/mockServe'

Vue.config.productionTip = false

new Vue({
  router,
  render: (h) => h(App),
  beforeCreate() {
    //安装全局事件总线
    Vue.prototype.$bus = this
  }
}).$mount('#app')
