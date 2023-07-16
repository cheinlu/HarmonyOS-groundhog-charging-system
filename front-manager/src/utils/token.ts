//封装本地存储存储数据与读取数据方法
//存储TOKEN
export const SET_TOKEN = (token: string) => {
  localStorage.setItem('TOKEN', token)
}
//本地存储获取数据
export const GET_TOKEN = () => {
  return localStorage.getItem('TOKEN')
}
//本地存储删除数据方法
export const REMOVE_TOKEN = () => {
  localStorage.removeItem('TOKEN')
}

//本地存储用户名
export const SET_NAME = (username:string)=>{
  localStorage.setItem('name', username)
}
//本地存储获取名字
export const GET_NAME = () => {
  return localStorage.getItem('name')
}
//本地存储删除名字
export const REMOVE_NAME = () => {
  localStorage.removeItem('name')
}
//存储租户id
export const SET_TEID = (id:string)=>{
  localStorage.setItem('TENANT_ID',id)
}
//获取租户id
export const GET_TEID = () => {
  return localStorage.getItem('TENANT_ID')
}