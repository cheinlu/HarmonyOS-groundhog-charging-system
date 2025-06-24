export class ChargeListData {
  PageNo?: number
  PageSize?: number
  TotalCount?: number
  List?: ChargeType[]
}

export class ChargeType {
  address?: string //地址
  coordinate?: string //坐标
  distance?: number //距离
  freePileTotal?: number //空闲数据
  id?: number //id
  imageUrl?: string //图片
  name?: string //站名
  pileTotal?: number //剩余数据
  pricePerHour?: number //多少钱一度
  tenantId?: number //租户id
  tenantName?: string //租户
}

export class PileListData {
  PageNo?: number
  PageSize?: number
  TotalCount?: number
  List?: PileType[]
}

export class PileType {
  code?: string //编号
  id?: number //id
  state?: number //充电桩状态
  stationId?: number //充电站站id
  stationName?: string //充电站站名
}

export class PriceListData {
  PageNo?: number
  PageSize?: number
  TotalCount?: number
  List?: PriceType[]
}

export class PriceType {
  endHour?: number|string  //结束时间
  id?: number  //id
  price?: string   //价格
  startHour?: number|string  //开始时间
  tenantId?: number  //租户id
}

export class OrderTypeData {
  PageNo?: number
  PageSize?: number
  TotalCount?: number
  List?: OrderType[]
}

export class OrderType{
  id?: number  //id
  pileCode?: string  //充电桩编号
  orderCode?: string  //订单号
  price?: number  //价格
  startAt?: string  //开始时间
  state?: number  //订单状态
  stationName?: string  //充电站名
  stopAt?: string  //结束时间
}

export class mapListItem {
  name: ResourceStr = ''
  link: string = ''
  isAppInstall: boolean = false
}

export class AddressData {
  latitude?: number = 0.0;
  longitude?: number = 0.0;
}

//桩参数
export const pileParam: Array<pileParamType> = [
  { text: $r('app.string.fast_charge'), type: '1' },
  { text: $r('app.string.180KW'), type: '2' },
  { text: $r('app.string.0V'), type: '3' },
  { text: $r('app.string.national_standard'), type: '4' }
]

export class pileParamType {
  text?: ResourceStr
  type?: string
}