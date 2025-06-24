export class ChargingQrData {
  type: string = ''
  pileId: string = ''
  pileCode: string = ''
  stationName: string = ''
}

//草料二维码网站地址：https://cli.im/text
//草料二维码选择‘网址’，输入http/https格式的网址，App扫描二维码跳转对应外部网站，网址如：https://www.baidu.com/
//草料二维码选择‘文本’，输入以下内容格式接口生成二维码，App扫描二维码跳转至充电ing页面
// {
//   "type": "charging",
//   "pileId": "2",
//   "pileCode": "02",
//   "stationName": "宝安新安充电站"
// }
