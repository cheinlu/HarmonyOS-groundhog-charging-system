//账号信息的ts类型
export interface ResponseData {
  code: number
  message: string
}
//添加|修改用户携带的参数
export interface userAddOrUp {
  username: string
  password: string
  nickname: string
  id?:number
}
//代表一个账号信息的ts类型
export interface User {
  balance?:number
  createTime?: string
  id?: number
  nickname?: string
  username?: string
  password?: string
}
//数组包含全部的用户信息
export type Records = User[]

//获取全部用户信息接口返回的数据ts类型
export interface UserResponseData extends ResponseData {
  data:{
    List:Records
    PageNo:number
    PageSize:number
    TotalCount:number
  }
}
//已有的用户充值列表包含的ts类型
export interface pay {
  id: number
  payCode: string
  userId: number
  username: string
  state: number
  price: number
  payAt: string
  createAt: string
  updateAt: string
}
//数组包含全部的充值数据
export type pays = pay[]
//获取全部用户信息接口返回的数据ts类型
export interface payRecord extends ResponseData {
  data:{
    List:pays
    PageNo:number
    PageSize:number
    TotalCount:number
  }
}
//用户充值携带的参数类型
export interface userPay {
userId: number|string
price: number
}