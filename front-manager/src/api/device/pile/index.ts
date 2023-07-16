//充电桩相关的接口
import request from "@/utils/request";
//数据类型
import type {pileList} from './type'
enum API {
  //获取充电桩的数据
  PILE_URL = '/api/pile/list?',
  //添加充电桩的数据
  ADDPILE_URL = '/api/pile/add',
  //修改充电桩的数据
  UPDATEPILE_URL = '/api/pile/update',
  //删除充电桩的数据
  DELETEPILE_URL = '/api/pile/delete?'
}

//充电桩的数据接口
export const reqPile = (pageNo:number,pageSize:number,stationId:number|string,code:string)=>request.get<any,pileList>(API.PILE_URL+`pageNo=${pageNo}&pageSize=${pageSize}&stationId=${stationId}&code=${code}`)
//添加|修改充电桩数据接口
export const reqAddOrUpdatePile = (data:any)=>{
if(data.id){
  return request.post(API.UPDATEPILE_URL,data)
}else{
  return request.post(API.ADDPILE_URL,data)
}
}
//删除充电桩的接口
export const reqRemovePile = (id:number)=>request.delete(API.DELETEPILE_URL+`id=${id}`)