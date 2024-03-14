import router from '@ohos.router'
import prompt from '@ohos.prompt';
import HeaderComponent from '../common/components/HeaderComponent'
import ChargeModel from '../model/ChargeModel'
import PricePop from '../view/charge/PricePop'
import PriceType from '../viewModel/PileType'
import Prompt from '@system.prompt';
@Entry
@Component
struct Charging {
  @State params: any = router.getParams()
  @State showPricePanel:boolean = false
  @State prices:PriceType[] = []

  build() {
    Column({space:10}){
      //1.标题
      HeaderComponent({title:'充电ing'})
      //2.主体
      Column({space:10}){
        //2.1 充电站
        Row(){
          Text('充电站')
          // Text('宝安新安充电站')
          Text(this.params.stationName)
        }
        .height(50)
        .justifyContent(FlexAlign.SpaceBetween)
        .chargingSty()
        //2.2 设备编号
        Row(){
          Text('设备编号')
          // Text('03')
          Text(this.params.code)
        }
        .height(50)
        .justifyContent(FlexAlign.SpaceBetween)
        .chargingSty()
        //2.3 收费标准
        Column({space:20}){
          Row({space:170}){
            Text('收费标准')
            Text('收费规则 >>')
              .onClick(()=>{
                this.showPricePanel = true
                this.PriceInfo()
              })

          }
          Row({space:170}){
            Text('充电时长')
            Text('充满即止 >>')
          }
          Text('PS：此功能不涉及支付，只做演示使用')
            .fontColor(Color.Red)
        }
        .height(120)
        .chargingSty()
        .alignItems(HorizontalAlign.Start)
        Button('充电')
          .onClick(() => {
              this.StartChargeInfo()
          })
          .fontColor(Color.White)
          .backgroundColor($r('app.color.charge_header'))
          .width('95%')
      }
      //面板
      Panel(this.showPricePanel){
        Button('知道了').onClick(() => this.showPricePanel = false)
        PricePop({prices:$prices})
      }
      .mode(PanelMode.Full)
      .dragBar(false)
      .backgroundMask($r('app.color.light_gray'))
      .backgroundColor(Color.White)
    }
    .width('100%')
    .height('100%')
    .backgroundColor('#EFEFEF')
  }
  //公共样式
  @Styles chargingSty(){
    .width('95%')
    .padding(10)
    .borderRadius(10)
    .backgroundColor('#FFF')
  }

  //获取不同时间段充电价格
  async PriceInfo(){
    let res:any = await ChargeModel.getChargePrice()
    if(res.responseCode === 200){
      let priceRes = JSON.parse(res.result.toString())
      this.prices = priceRes.data.List
    }
  }
  //获取开始充电
  async StartChargeInfo(){
    let res:any = await ChargeModel.getStartCharge(this.params.id)
    if(res.responseCode === 200){
      let startChargeRes = JSON.parse(res.result.toString())
      if(startChargeRes.code == 0){
        Prompt.showToast({ message:'充电成功'})
        router.pushUrl(
          {
            url:'pages/ChargeOrder',
          },
          router.RouterMode.Single,
          err => {
            if(err){
              console.log(`路由失败，errCode:${err.code},errMsg:${err.message}`)
            }
          }
        )
      }
    }
  }
}