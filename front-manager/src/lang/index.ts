//引入i18n
import { createI18n } from 'vue-i18n'
import enLocale from './en/index'
import zhLocale from './zh/index'

//创建i18n对象
const i18n = createI18n({
  locale:localStorage.getItem('lang') || "zh", //默认显示的语言 
  fallbackLocale: localStorage.getItem('lang') || "zh",//预设语言环境
  globalInjection: true, //全局生效$t
  legacy:false,
  messages:{
    zh: zhLocale,
    en: enLocale,
  }
})

export default i18n