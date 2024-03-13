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
//存储权限列表permission
export const SET_PERMISSION = (permission:string[])=>{
  localStorage.setItem('PERMISSION',JSON.stringify(permission))
}

//存储租户名字
export const SET_TENAME = (tenantName:string)=>{
  localStorage.setItem('TENANT_NAME',tenantName)
}
//获取租户名字
export const GET_TENAME = () => {
  return localStorage.getItem('TENANT_NAME')
}



[
  {
    address: "深圳市南山南头街道5区12号",
    id: 3,
    name: "南山南头充电站",
    tenantName: "钢铁侠充电公司",
  },
  {
    address: "深圳市宝安区新安街道17区3号",
    id: 1,
    name: "宝安新安充电站",
    tenantName: "土拨鼠充电公司",
  },
  {
    address: "深圳市宝安区西乡街道盐田村5号",
    id: 2,
    name: "宝安西乡充电站",
    tenantName: "土拨鼠充电公司",
  }
]