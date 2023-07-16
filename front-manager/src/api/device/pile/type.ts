//定义全部接口返回数据都拥有ts类型
export interface ResponseData{
  code :number
  message:string
  }
//已有的充电桩列表数据类型
export interface pile {
  code: string
  id: number
  state: number
  stationId: number
  stationName: string

}
//保存全部充电桩列表数据类型
export type piles = pile[]
//充电桩列表返回的数据类型
export interface pileList extends ResponseData {
  data:{
    List:piles
    PageNo:number
    PageSize: number
    TotalCount: number
  }
}
//添加|修改充电桩参数的类型
export interface pileAddOrUpdate { 
  code: string
  stationId: number|string
  id?:number
  state?:number
}