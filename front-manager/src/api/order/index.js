import requests from '../request'

//充电订单  /api/charge/list  get
/*
参数
OrderCode：订单编号
UserId：用户id，使用下拉框
StationId：充电桩id，使用下拉框
PileId：充电桩编号，使用下拉框，下拉框与 充电站下拉框联动
BeforeAt：充电开始时间（大于）
AfterAt：充电开始时间（小于）
State：订单状态（0:充电中,1:已完成,2:未付账）
*/
export const chargeOrderList = (params)=>requests({url:'/api/charge/list',method:'get',params})

//修改订单 /api/charge/update  post
/*
参数
{
"Id":1,
"State":1,  
"Price":1.3
}

*/
export const updateChargeOrder = ()=>requests({url:'/api/charge/update',method:'post'})