import requests from '../request'
// import mockRequests from '../mockIndex'

//------------登录与token续期接口
//登录接口 /api/user/login  post
export const login = (data) => requests({ url: '/api/user/login', method: 'post', data })

//token续期的接口 /api/token/refresh
export const refreshToken = () => requests({ url: '/api/token/refresh', method: 'post'})

//------------用户列表

//用户列表  /api/user/list  get
export const userList = (params) => requests({ url: '/api/user/list', method: 'get', params })

//添加用户   /api/user/add  post
/*
参数
{
    "username":"guest",
    "nickname":"访客",
    "password":"guest"
}
*/

//修改用户   /api/user/update  post
/*
{
    "id":4,
    "username":"lucy1",
    "nickname":"lucy1",
    "password":"lucy1"
}
*/
//添加和修改用户的api  根据id来区分
export const arrOrUpdateUserList = (user)=>{
    if(user.id){
      return  requests({url:'/api/user/update',method:'post',data:user})
    }else{
      return  requests({url:'/api/user/add',method:'post',data:user})
    }
}

//删除用户  /api/user/delete?id=5  delete
export const deletUserList = (id) => requests({ url: `/api/user/delete?id=${id}`, method: 'delete' })

//------------用户充值

//  用户充值列表 /api/user/pay/record  get
export const recordList = (params) => requests({url:'/api/user/pay/record',method:'get',params})

//用户充值  /api/user/pay  post
export const recordPay = (data)=>requests({url:'/api/user/pay',method:'post',data})