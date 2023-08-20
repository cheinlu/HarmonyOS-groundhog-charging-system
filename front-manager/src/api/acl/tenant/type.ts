//租户相关的数据类型
//定义全部接口返回数据都拥有ts类型
export interface ResponseData{
  code :number
  message:string
  }

//租户注册需要携带的参数类型
export interface registerFormData {
  name: string
  domain: string
  adminUsername: string
  adminPassword: string
  validCode: string
}

export interface data{
  Data:{
    domain: string
    id: number
    name:string
  }
}
//租户登录时返回的数据类型
export interface tenantSelect extends ResponseData {
  data:data
}

//已有的租户ts类型
export interface tenant {
  contactEmail: string
  contactName: string
  contactPhone: string
  domain: string
  id?: number
  name: string
  status: string
}
//保存全部租户的ts数据类型
export type teList = tenant[]
//租户列表返回的数据类型
export interface tenantList extends ResponseData {
  data:{
    List:teList
    PageNo:number
    PageSize: number
    TotalCount: number
  }
}