import http from '@ohos.net.http'

class ChargeModel {
  baseUrl: string = 'http://127.0.0.1:8000'
  pageNo:number = 1
  defaultHeaders: Record<string, string> = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${AppStorage.Get('TOKEN')}`
  }
  //定义sendRequest方法
  private async sendRequest(url: string, method: http.RequestMethod, extraData?: Record<string, any>) {
    try {
      const httpRequest = http.createHttp()
      const response = await httpRequest.request(url, { method, header: this.defaultHeaders, extraData })
      return response
    } catch (error) {
      console.log(`请求失败，URL: ${url}, Error: ${JSON.stringify(error)}`)
      throw new Error('请求失败')
    }
  }
  //获取充电站
  async getStations(name:string) {
    return this.sendRequest(`${this.baseUrl}/wx-api/station/list?name=${name}&pageNo=${this.pageNo}&pageSize=5`, http.RequestMethod.GET)
  }
  //获取充电桩
  async getPiles(stationId: number) {
    return this.sendRequest(`${this.baseUrl}/wx-api/pile/list?stationId=${stationId}`, http.RequestMethod.GET)
  }
  //获取充电价格
  async getChargePrice() {
    return this.sendRequest(`${this.baseUrl}/wx-api/price/list?pageNo=${this.pageNo}&pageSize=10`, http.RequestMethod.GET)
  }
  //开始充电
  async getStartCharge(pileId: number) {
    return this.sendRequest(`${this.baseUrl}/wx-api/charge/start`, http.RequestMethod.POST, { pileId })
  }
  //获取充电订单列表
  async getChargeOrders() {
    return this.sendRequest(`${this.baseUrl}/wx-api/charge/list?pageNo=${this.pageNo}&pageSize=4`, http.RequestMethod.GET)
  }
  //结束充电
  async getStopCharge(orderId: number) {
    return this.sendRequest(`${this.baseUrl}/wx-api/charge/stop`, http.RequestMethod.POST, { orderId })
  }
}

const chargeModel = new ChargeModel()
export default chargeModel as ChargeModel
