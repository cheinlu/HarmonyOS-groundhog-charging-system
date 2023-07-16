//账号信息的ts类型
export interface ResponseData {
  code: number
  message: string
}
//单个权限列表展示的数据
export interface permission {
  key:string
  name:string
}
//权限列表集合
export type permissions = permission[]
//权限列表返回的数据
export interface permissionList extends ResponseData {
  data:permissions
}