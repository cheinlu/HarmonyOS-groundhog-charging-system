//定义全部接口返回数据都拥有ts类型
export interface ResponseData{
  code :number
  message:string
  }

//已有的充电价格数据类型
export interface price {
  endHour: number|string
  id?: number
  price: number|string
  startHour: number|string
}
//保存全部时间段充电价格数据类型
export type changePrice = price[]
//充电价格列表返回的数据类型
export interface priceList extends ResponseData {
  data:{
    List:changePrice
    PageNo:number
    PageSize: number
    TotalCount: number
  }
}

//订单列表
//已有订单列表的数据类型
export interface charge {
  createAt: string
  id?: number
  nickname: string
  orderCode: string
  pileCode: string
  pileId: number|string
  price: number
  startAt: string
  state: number|string
  stationId: number|string
  stationName: string
  stopAt: string
  updateAt: string
  userId: number|string

}
//保存全部订单列表的数据
export type charges = charge[]
//订单列表返回的数据类型
export interface chargeList extends ResponseData {
  data:{
    List:charges
    PageNo:number
    PageSize: number
    TotalCount: number
  }
}
