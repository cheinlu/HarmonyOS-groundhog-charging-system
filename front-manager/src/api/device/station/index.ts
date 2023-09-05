//引入二次封装的axios
import request from "@/utils/request";
//数据类型
import type {stationList,stationAddOrUpdate} from './type'
enum API {
  //获取充电站的数据
  STATION_URL = '/saas-api/station/list?',
  //添加充电站
  ADDSTATION_URL = '/saas-api/station/add',
  //修改充电站
  UPDATESTATION_URL = '/saas-api/station/update',
  //删除充电站
  DELETESTATION_URL = '/saas-api/station/delete?',
  //上传图片 
  UPLOAD_URL = '/saas-api/station/upload'
  
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
// 上传图片
export const reqUpload = (file:any) => {
  // 创建了一个新的 FormData 对象，用于构建表单数据,并将file添加到FormData对象中
  const formData = new FormData();
  formData.append('file', file);

  return request.post(API.UPLOAD_URL, formData);
};