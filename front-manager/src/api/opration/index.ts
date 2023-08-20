//运营管理相关的接口----包含订单和规则
import request from "@/utils/request";
//数据类型
import type {priceList,price,chargeList} from './type'
enum API {
  //获取价格区间的数据
  PRICE_URL = '/saas-api/charge/price/list?',
   //添加充电价格
   ADDPRICE_URL = '/saas-api/charge/price/add',
   //更新已有充电价格
   UPDATEPRICE_URL = '/saas-api/charge/price/update',
 //删除已有充电价格
  DELETEPRICE_URL = '/saas-api/charge/price/delete?',
  //充电订单的数据
  ORDER_URL = '/saas-api/charge/list?',
  //充电订单修改
  UPDATEORDER_URL = '/saas-api/charge/update?'
}

//价格区间的接口
export const reqChargePrice = (pageNo:number,pageSize:number)=>request.get<any,priceList>(API.PRICE_URL+`pageNo=${pageNo}&pageSize=${pageSize}`)
//添加|更改充电的接口
export const reqAddOrUpdatePrice = (data:price)=>{
  if(data.id){
    return request.post<any,any>(API.UPDATEPRICE_URL,data)
  }else{
    return request.post<any,any>(API.ADDPRICE_URL,data)
  }
}
//删除已有充电价格的接口
export const reqRemovePrice = (id:number)=>request.delete(API.DELETEPRICE_URL+`id=${id}`)

//充电订单的接口
export const reqOrder = (pageNo:number,pageSize:number,orderCode:string,userId:number,stationId:number,pileId:number,beforeAt:string,afterAt:string,state:string)=>request.get<any,chargeList>(API.ORDER_URL+`pageNo=${pageNo}&pageSize=${pageSize}&orderCode=${orderCode}&userId=${userId}&stationId=${stationId}&pileId=${pileId}&beforeAt=${beforeAt}&afterAt=${afterAt}&state=${state}`)
//充电订单的修改
export const reqUpdateOrder = (id:number,price:number|null,state:any)=>request.post(API.UPDATEORDER_URL+`id=${id}&price=${price}&state=${state}`)