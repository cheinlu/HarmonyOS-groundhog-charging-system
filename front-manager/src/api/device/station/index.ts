//充电站相关的接口
import request from "@/utils/request";
//数据类型
import type {stationList,stationAddOrUpdate} from './type'
enum API {
  //获取充电站的数据
  STATION_URL = '/api/station/list?',
  //添加充电站
  ADDSTATION_URL = '/api/station/add',
  //修改充电站
  UPDATESTATION_URL = '/api/station/update',
  //删除充电站
  DELETESTATION_URL = '/api/station/delete?'
}

//充电站的数据接口
export const reqStation = (pageNo:number,pageSize:number,name:string)=>request.get<any,stationList>(API.STATION_URL+`pageNo=${pageNo}&pageSize=${pageSize}&name=${name}`)
//添加|修改充电站接口
export const reqAddOrUpdateStation = (data:stationAddOrUpdate)=>{
  if(data.id){
    return request.post<any,any>(API.UPDATESTATION_URL,data)
  }else{
    return request.post<any,any>(API.ADDSTATION_URL,data)
  }
}

//删除充电站的接口
export const reqRemoveStation = (id:number)=>request.delete(API.DELETESTATION_URL+`id=${id}`)
