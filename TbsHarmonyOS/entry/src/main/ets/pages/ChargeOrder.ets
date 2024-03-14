import HeaderComponent from '../common/components/HeaderComponent'
import ChargeModel from '../model/ChargeModel'
import OrderType from '../viewModel/OrderType';
import {formatDate} from '../common/utils/time'
@Entry
@Component
struct ChargeOrder {
  @State Orders:OrderType[] = []
  isLoading: boolean = false
  isMore: boolean = true

  aboutToAppear(){
    // 重新初始化页面状态
    this.Orders = [];
    ChargeModel.pageNo = 1;
    this.ChargeOrderInfo()
  }

  build() {
    Column({space:10}){
      //1.标题
      HeaderComponent({title:'充电订单'})
      //2.订单
      this.OrderCard()

    }
    .width('100%')
    .height('100%')
    .backgroundColor('#EFEFEF')
  }
  @Builder OrderCard(){
    List({space:10}){
      ForEach(this.Orders,(item:OrderType) => {
        ListItem(){
          Column({space:8}){
            Text(item.orderCode)
              .fontColor(Color.White)
              .width(200)
              .height(40)
              .padding(10)
              .borderRadius({ bottomRight:10,bottomLeft:10 })
              .margin({left:'50%'})
              .backgroundColor($r('app.color.charge_header'))
            if(item.state === 0){
              Text('充电中').fontWeight(FontWeight.Bold).fontColor(Color.Red)
            }else{
              Text('已完成').fontWeight(FontWeight.Bold).fontColor($r('app.color.charge_header'))
            }
            Row({space:10}){
              Text('充电桩编号:')
              Text(item.pileCode)
            }
            Row({space:10}){
              Text('充电站名:')
              Text(item.stationName)
            }
            Row({space:10}){
              Text('价格:')
              Text(`${item.price}`)
            }
            Row({space:10}){
              Text('订单状态:')
              Text(item.state == 0 ? '充电中' : item.state == 1 ? '已完成' : '未付账')
            }
            Row({space:10}){
              Text('开始时间:')
              Text(formatDate(item.startAt))
            }
            Row({space:10}){
              Text('结束时间:')
              if(item.stopAt !== undefined && item.stopAt !== null && item.state !== 0){
                Text(formatDate(item.stopAt))
              }
            }
            if(item.state === 0){
              Button('停止充电').width(90).backgroundColor(Color.Red)
                .onClick(() => {
                  this.StopChargeInfo(item)
                })
            }else{
              Text('')
            }
          }
          .width('94%')
          .padding({left:10,bottom:10})
          .borderRadius(10)
          .backgroundColor('#FFF')
          .alignItems(HorizontalAlign.Start)
          .margin({left:10})
        }
      })
    }
    .width('100%')
    .layoutWeight(1)
    .onReachEnd(() => {
      if(!this.isLoading && this.isMore){
        this.isLoading = true
        // 翻页
        ChargeModel.pageNo++
        this.ChargeOrderInfo()
      }
    })
  }

  //获取充电订单
  async ChargeOrderInfo(){
    let res:any = await ChargeModel.getChargeOrders()
    if(res.responseCode === 200){
      let chargeOrderRes = JSON.parse(res.result.toString())
      this.Orders = this.Orders.concat(chargeOrderRes.data.List)
      this.isLoading = false
      if(!chargeOrderRes.data.List || chargeOrderRes.data.List.length === 0){
        this.isMore = false
      }
    }
  }
  //停止充电
  async StopChargeInfo(item){
    let res:any = await ChargeModel.getStopCharge(item.id)
    if(res.responseCode === 200){
      item.state = 1
      this.ChargeOrderInfo()
    }
  }
}