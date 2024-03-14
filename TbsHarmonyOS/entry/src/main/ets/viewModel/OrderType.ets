export default class OrderType{
  id: number  //id
  pileCode: string  //充电桩编号
  orderCode: string  //订单号
  price: number  //价格
  startAt: string  //开始时间
  state: number  //订单状态
  stationName: string  //充电站名
  stopAt: string  //结束时间
}