import { createApp } from 'vue'
//引入i18n
import i18n from '@/lang/index'
//引入ElementPlus
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css'
//国际化
//@ts-expect-error忽略当前文件ts类型的检测否则有红色提示(打包会失败)
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
//暗黑模式
import 'element-plus/theme-chalk/dark/css-vars.css'
import App from '@/App.vue'
//引入路由
import router from './router/index'
// 引入svg
import 'virtual:svg-icons-register'
import gloablComponent from '@/components/index'
//引入全局样式
import '@/styles/index.scss'
//引入仓库
import pinia from './store'
//引入七牛
import * as qiniu from 'qiniu-js';
//引入百度地图
import BaiduMap from 'vue-baidu-map-3x';
const app = createApp(App)
//引入自定义指令
import {isHasPermission} from '@/directive/has'
isHasPermission(app)

//安装i18n
app.use(i18n)
//安装百度地图
app.use(BaiduMap, {ak: "124wXbeEy3ngDIaCEzUMlu6DA7gPwGsF"})
//安装仓库
app.use(pinia)
//安装element-plus插件
app.use(ElementPlus, {
    locale: zhCn
})
app.config.globalProperties.$qiniu = qiniu;
//注册模板路由
app.use(router)
//引入路由鉴权文件
import './permisstion'
app.mount('#app')
//安装自定义插件
app.use(gloablComponent)

