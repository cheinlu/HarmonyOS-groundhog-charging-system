import requests from '../request'
// import mockRequests from '../mockIndex' //引入mock

//充电站接口   /api/station/list  get
export const chargeList = (params) => requests({ url: '/api/station/list', method: 'get', params })

//充电站mock接口
// export const charge = ()=>mockRequests({url:'/charge',method:'get'})

//删除充电站  /api/station/delete?id=5  delete
export const deletChargeList = (id) => requests({ url: `/api/station/delete?id=${id}`, method: 'delete' })

//添加充电站   /api/station/add     post
/*
{
    "name":"宝安西乡充电站",
    "address":"宝安西乡村17号",
    "coordinate":"128.333333,64.92222"
}
*/
// export const addChargeList = (data)=>requests({url:'/api/station/add',method:'post',data})
//修改充电站   /api/station/update  post
/*
{
    "id":1,
    "name":"宝安西乡充电站",
    "address":"宝安西乡村17号",
    "coordinate":"128.333333,64.92222"
}
*/
//添加和修改充电站的api  根据id来区分
export const addOrUpdateChargeList = (charge) => {
  if (charge.id) {
    return requests({ url: '/api/station/update', method: 'post', data: charge })
  } else {
    return requests({ url: '/api/station/add', method: 'post', data: charge })
  }
}

//------------------------------------------------------------------------

//充电桩接口  /api/pile/list get
export const pileList = (params) => requests({ url: '/api/pile/list', method: 'get', params })

//删除充电桩   /api/pile/delete?id=5  delete
export const deletePileList = (id) => requests({ url: `/api/pile/delete?id=${id}`, method: 'delete' })

//添加充电桩  /api/pile/add  post
/*
{
    "code":"02",
    "stationId":1
}  
*/

//修改充电桩 /api/pile/update  post
/*
{
    "id":1,
    "code":"02",
    "stationId":1,
    "state":0
}
*/

export const addOrUpdatePile = (pile) => {
  if (pile.id) {
    return requests({ url: '/api/pile/update', method: 'post', data: pile })
  } else {
    return requests({ url: '/api/pile/add', method: 'post', data: pile })
  }
}
