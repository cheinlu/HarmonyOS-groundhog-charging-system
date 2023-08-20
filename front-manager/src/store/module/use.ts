//创建用户相关的小仓库
import {defineStore} from 'pinia'
//引入接口
import {reqLogin} from '@/api/user/index'
import {reqTeLogin} from '@/api/acl/tenant/index'
import {reqCurrentPermi} from '@/api/acl/authority/index'
import { ElMessage } from 'element-plus'
//引入ts类型
import type {loginFormData,loginResponseData} from '@/api/user/type'
import type {tenantSelect} from '@/api/acl/tenant/type'
import type { UserState } from './types/type'
//引入本地存储
import {SET_TOKEN,GET_TOKEN,REMOVE_TOKEN,SET_NAME,GET_NAME,REMOVE_NAME,SET_TEID,GET_TEID,SET_PERMISSION,SET_TENAME,GET_TENAME} from '@/utils/token'
//引入路由（常量路由）
import { constantRoute,asyncRoute,anyRoute} from '@/router/routes'
//引入路由
import router from '@/router'
//引入深拷贝方法
//@ts-expect-error
import cloneDeep from 'lodash/cloneDeep'
//用于过滤当前用户需要展示的异步路由
function filterRoutes(routes:any,names:any, depth = 0){
  const MAX_DEPTH = 1000; // 设置递归深度限制
  if (depth > MAX_DEPTH) {
    throw new Error('Exceeded maximum recursion depth');
  }
  
  let filteredRoutes:any = [];
  routes.forEach((route:any) => {
    let filteredRoute = { ...route };
    if (route.children) {
      // 过滤子路由
      filteredRoute.children = filterRoutes(route.children, names, depth + 1);
      // 如果子路由数量大于0，则将当前路由添加到过滤后的路由数组中
      if (filteredRoute.children.length > 0) {
        filteredRoutes.push(filteredRoute);
      }
    } else {
      // 如果当前路由在名字数组中，则将其添加到过滤后的路由数组中
      if (names.includes(route.name)) {
        filteredRoutes.push(filteredRoute);
      }
    }
  });
  return filteredRoutes;
}

const useUserStore = defineStore('User',{
  state:():UserState=>{
    return {
      token:GET_TOKEN(),
      menuRoutes: constantRoute, //仓库存储生成菜单需要数组(路由)
      username:GET_NAME(),
      id:GET_TEID(),//租户id
      tenantName:GET_TENAME(),
      teArr:[],//选择租户数据
      permissionList:[]//当前用户的权限
    }
  },
  actions:{
    //用户登录
    async userLogin(data:loginFormData){
    let res:loginResponseData = await reqLogin(data)
    if(res.code==0){
      this.token = res.data.token
      //本地持久化存储
      SET_TOKEN(res.data.token)
        //调用setCurrentPermission方法更新权限
     await this.setCurrentPermission();
      //能保证当前async函数返回一个成功的promise
      return 'ok'
    } else {
      return Promise.reject(new Error(res.message))
    }
    },
    //将输入的用户名保存到仓库中
    setUsername(username:string){
      this.username = username
      //本地持久化存储
      SET_NAME(username)
    },
    //退出登录
    userLogout(){
      this.token = ''
      this.username = ''
      REMOVE_TOKEN()
      REMOVE_NAME()
    },
     // 设置租户ID
    setTenantId(id:any){
      this.id = id
      SET_TEID(id)
    },
    setTenantName(tenantName:any){
      this.tenantName = tenantName
      SET_TENAME(tenantName)
    },
      // 获取租户信息
    async tenantSelect(){
     let res:tenantSelect = await reqTeLogin()
      if(res.code==0){
        this.teArr = res.data.Data
      }
    },
    //获得当前用户登录的权限
    async setCurrentPermission(){
     let res:any = await reqCurrentPermi()
     if(res.code==0){
      this.permissionList = res.data.permissionList

      SET_PERMISSION(res.data.permissionList)
      //当前用户展示的异步路由
      let filteredAsyncRoute = filterRoutes(cloneDeep(asyncRoute), res.data.permissionList)
      //菜单需要的路由数据整理
      this.menuRoutes = [...constantRoute,...filteredAsyncRoute,anyRoute]
      //目前路由器管理的只有常量路由:用户计算完毕异步路由、任意路由动态追加
      ;[...filteredAsyncRoute,anyRoute].forEach((route:any)=>{
        router.addRoute(route)
      })
     }else{
      ElMessage({type:'error',message:res.message})
     }
     
    }
  },
  getters:{
  }
})

//对外暴露
export default useUserStore