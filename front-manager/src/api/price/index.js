import requests from '../request'

//区间充电价格  /api/charge/price/list  get
export const priceList = (data)=>requests({url:'/api/charge/price/list',method:'get',data})

//区间充电价格删除 /api/charge/price/delete?id=5 delete
export const priceDelete = (id)=>requests({url:`/api/charge/price/delete?id=${id}`,method:'delete'})

//添加  /api/charge/price/add  post
/*
参数{
    "startHour":0,
    "endHour":1,
    "price":1.5
} 
*/

//修改  /api/charge/price/update  post
/*
参{
    "id":1,
    "startHour":0,
    "endHour":1,
    "price":1.5
}
*/

export const addOrUpdatePrice = (price)=>{
  if(price.id){
   return requests({url:'/api/charge/price/update',method:'post',data:price})
  }else{
    return requests({url:'/api/charge/price/add',method:'post',data:price})
  }
}




