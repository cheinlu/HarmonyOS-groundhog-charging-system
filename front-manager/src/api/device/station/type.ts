//定义全部接口返回数据都拥有ts类型
export interface ResponseData{
  code :number
  message:string
  }

//已有的充电站列表数据类型
export interface station {
  address: string
  coordinate: string
  createAt: string
  id?: number
  name: string
  tenantName: string
  updateAt: string
}
//保存全部充电站列表数据类型
export type stations = station[]
//充电站列表返回的数据类型
export interface stationList extends ResponseData {
  data:{
    List:stations
    PageNo:number
    PageSize: number
    TotalCount: number
  }
}
//添加|修改充电站参数的类型
export interface stationAddOrUpdate { 
  name: string
  address: string
  coordinate: string
  imageUrl:string
  id?:number
}