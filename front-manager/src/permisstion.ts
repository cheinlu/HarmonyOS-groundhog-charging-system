//引入路由
import router from "./router";

//引入进度条
//@ts-expect-error
import nprogress from 'nprogress'
//引入进度条样式
import 'nprogress/nprogress.css'
//引入仓库
import useUserStore from '@/store/module/use'
import pinia from './store'
const userStore = useUserStore(pinia)
nprogress.configure({ showSpinner: false })
//全局前置守卫
router.beforeEach(async (to:any,_from:any,next:any)=>{
  //进度条打开
  nprogress.start()
  //获取token
  let token = userStore.token
  //获取用户名
  let name = userStore.username
  if(token){
    //有token登录成功，再去登录页是不行的，只能在首页
    if(to.path=='/login'){
      next('/')
    }else{
      //如果去的不是首页，是其他页面,判断有没name
      if(name){
        next()
      }else{
        try {
           //获取用户信息
           await userStore.setUsername(name as string);
           next();
        } catch (error) {
         userStore.userLogout()
          next({ path: '/login', query: { redirect: to.path } })
        }
      }
    }
  }else{
    //没有token，要去登录页，是可以的
    if(to.path=='/login'||to.path=='/register'){
      next()
    }else{
      //但如果去的不是登录页，让他只能去登录页，且传递确认样参数，重定向到退出的地址
      next({path:'/login',query:{redirect:to.path}})
    }
  }

})

//全局后置守卫
router.afterEach(()=>{
  //进度条关闭
  nprogress.done()
})