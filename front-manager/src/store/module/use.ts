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
import {SET_TOKEN,GET_TOKEN,REMOVE_TOKEN,SET_NAME,GET_NAME,REMOVE_NAME,SET_TEID,GET_TEID} from '@/utils/token'
//引入路由（常量路由）
import { constantRoute} from '@/router/routes'
const useUserStore = defineStore('User',{
  state:():UserState=>{
    return {
      token:GET_TOKEN(),
      menuRoutes: constantRoute, //仓库存储生成菜单需要数组(路由)
      username:GET_NAME(),
      id:GET_TEID(),//租户id
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
      // 获取租户信息
    async tenantSelect(){
     let res:tenantSelect = await reqTeLogin()
      if(res.code==0){
        this.teArr = res.data
      }
    },
    //获得当前用户登录的权限
    async setCurrentPermission(){
     let res:any = await reqCurrentPermi()
     if(res.code==0){
      this.permissionList = res.data.permissionList
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