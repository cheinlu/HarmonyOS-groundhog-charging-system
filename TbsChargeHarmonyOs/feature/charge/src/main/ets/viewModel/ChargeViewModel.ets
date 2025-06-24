import { requestApiCall, Apis, toast } from "@tbs/common";
import { http } from "@kit.NetworkKit";
import { ChargeListData, OrderTypeData, PileListData, PriceListData } from "../model/ChargeModel";
import { BusinessError } from "@kit.BasicServicesKit";

//获取充电站
export function getStations(name: string, pageNo: number = 1, pageSize: number = 20) {
  return requestApiCall(`${Apis.CHARGE_STATION_LIST}?name=${name}&pageNo=${pageNo}&pageSize=${pageSize}`,
    http.RequestMethod.GET).then(res => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    const data = res.data as ChargeListData
    return data
  }).catch((err: BusinessError) => {
    console.error(`lucy++ 获取充电站接口调用失败，${err.code}+${err.message}`)
  })
}

//获取充电桩
export function getPiles(stationId: number) {
  return requestApiCall(`${Apis.CHARGE_PILE_LIST}?stationId=${stationId}`, http.RequestMethod.GET).then(res => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    const data = res.data as PileListData
    return data.List
  }).catch((err: BusinessError) => {
    console.error(`lucy++ 获取充电桩接口调用失败，${err.code}+${err.message}`)
  })
}

//获取充电价格
export function getChargePrice(pageNo: number = 1, pageSize: number = 20) {
  return requestApiCall(`${Apis.CHARGE_PRICE}?pageNo=${pageNo}&pageSize=${pageSize}`, http.RequestMethod.GET)
    .then(res => {
      console.log('lucy++ msg',res.message)
      if (res.code !== 0 && res.message) {
        toast(res.message)
      }
      const data = res.data as PriceListData
      return data.List
    }).catch((err: BusinessError) => {
      console.error(`lucy++ 获取充电价格接口调用失败，${err.code}+${err.message}`)
    })
}

//开始充电
export function getStartCharge(pileId: number) {
  const params: Record<string, Object> = {
    "pileId": pileId
  }
  return requestApiCall(Apis.CHARGE_START, http.RequestMethod.POST, params).then(res => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    return res.code
  }).catch((err: BusinessError) => {
    console.error(`lucy++ 开始充电接口调用失败，${err.code}+${err.message}`)
  })
}

//获取充电订单列表
export function getChargeOrders(pageNo: number, pageSize: number) {
  return requestApiCall(`${Apis.CHARGE_ORDER_LIST}?pageNo=${pageNo}&pageSize=${pageSize}`, http.RequestMethod.GET)
    .then(res => {
      if (res.code !== 0 && res.message) {
        toast(res.message)
      }
      const data = res.data as OrderTypeData
      return data
    }).catch((err: BusinessError) => {
      console.error(`lucy++ 获取充电订单列表接口调用失败，${err.code}+${err.message}`)
    })
}

//结束充电
export function getStopCharge(orderId: number) {
  const params: Record<string, Object> = {
    "orderId": orderId
  }
  return requestApiCall(Apis.CHARGE_STOP, http.RequestMethod.POST, params).then(res => {
    if (res.code !== 0 && res.message) {
      toast(res.message)
    }
    return res.code
  }).catch((err: BusinessError) => {
    console.error(`lucy++ 结束充电接口调用失败，${err.code}+${err.message}`)
  })
}