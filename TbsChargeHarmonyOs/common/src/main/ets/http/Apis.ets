export class Apis {
  //注册
  static readonly USER_REGISTER = "/saas-api/tenant/register"
  //获取token-登录
  static readonly USER_LOGIN = "/saas-api/user/login"
  //获取租户ID
  static readonly USER_TENANT = "/saas-api/tenant/selects"
  //获取用户信息
  static readonly USER_INFO = "/wx-api/me/info"
  //获取充电站
  static readonly CHARGE_STATION_LIST = "/wx-api/station/list"
  //获取充电桩
  static readonly CHARGE_PILE_LIST = "/wx-api/pile/list"
  //获取充电价格
  static readonly CHARGE_PRICE = "/wx-api/price/list"
  //开始充电
  static readonly CHARGE_START = "/wx-api/charge/start"
  //结束充电
  static readonly CHARGE_STOP = "/wx-api/charge/stop"
  //获取充电订单列表
  static readonly CHARGE_ORDER_LIST = "/wx-api/charge/list"
}