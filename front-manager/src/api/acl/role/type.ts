//定义全部接口返回数据都拥有ts类型
export interface ResponseData{
  code :number
  message:string
  }
//已有的角色数据类型
export interface role {
  bindUsers: number
  createAt: string
  id: number
  name: string
  updateAt: string
}
//保存全部角色列表数据类型
export type roles = role[]
//角色列表返回的数据类型
export interface roleList extends ResponseData {
  data:{
    List:roles
    PageNo:number
    PageSize: number
    TotalCount: number
  }
}
//添加|修改角色参数的类型
export interface roleAddOrUp { 
  name: string
  id?:number
}