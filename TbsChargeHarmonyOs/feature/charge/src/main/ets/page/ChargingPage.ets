import { StyleConst, TbsTitleBar, BottomPop, toast } from "@tbs/common"
import { router } from "@kit.ArkUI"
import { PileType, PriceType } from "../model/ChargeModel"
import { ChargePriceDialogView } from "../view/ChargePriceDialogView"
import { getChargePrice, getStartCharge } from "../viewModel/ChargeViewModel"
import './ChargeOrderPage'

//充电中
@Entry({ routeName: 'ChargingPage' })
@Component
export struct ChargingPage {
  @State priceList: PriceType[] = []
  @State pileId: number = 0
  @State pileCode: string = ''
  @State stationName: string = ''

  onPageShow(): void {
    const PileTypeParams = router.getParams() as Record<string, Object>
    if (PileTypeParams) {
      this.pileId = PileTypeParams.pileId as number
      this.pileCode = PileTypeParams.pileCode as string
      this.stationName = PileTypeParams.stationName as string
    }
  }

  chargePriceController: CustomDialogController | null = new CustomDialogController({
    builder: BottomPop({
      title: $r('app.string.rates'),
      bgColor: $r('app.color.color_f0f2f4'),
      titleBgColor: $r('app.color.color_f0f2f4'),
      contentBgColor: $r('app.color.color_f0f2f4'),
      keyboardAvoidMode: KeyboardAvoidMode.NONE,
      childBuilderParam: () => {
        this.chargePriceDialog()
      }
    }),
    customStyle: true,
    alignment: DialogAlignment.Bottom,
    maskColor: $r('app.color.color_80000000'),
  })

  @Builder
  chargePriceDialog() {
    ChargePriceDialogView({ priceList: this.priceList, pileCode: this.pileCode })
  }

  //获取不同时间段充电价格
  async getChargePriceData() {
    let data = await getChargePrice()
    if (data) {
      this.priceList = data
    }
  }

  //开始充电
  async getStartChargeData() {
    if (this.pileId) {
      let data = await getStartCharge(this.pileId)
      if (data == 0) {
        toast($r('app.string.charge_success'))
        router.pushNamedRoute({
          name: 'ChargeOrderPage'
        })
      }
    }

  }

  build() {
    Column({ space: 10 }) {
      //1.标题
      TbsTitleBar({ title: $r('app.string.charge_in'), showBackIcon: true, titleBarBgColor: $r('app.color.color_0aa671') })
      //2.主体
      Column({ space: 10 }) {
        //2.1 充电站
        Row() {
          Text($r('app.string.charge_name'))
          Text(this.stationName)
        }
        .height($r('app.float.size_50'))
        .justifyContent(FlexAlign.SpaceBetween)
        .chargingSty()

        //2.2 设备编号
        Row() {
          Text($r('app.string.pile_code'))
          Text(this.pileCode)
        }
        .height($r('app.float.size_50'))
        .justifyContent(FlexAlign.SpaceBetween)
        .chargingSty()

        //2.3 收费标准
        Column({ space: 20 }) {
          Row() {
            Text($r('app.string.rates'))
            Text($r('app.string.charge_rule'))
              .onClick(() => {
                this.chargePriceController?.open()
                this.getChargePriceData()
              })

          }
          .width(StyleConst.FULL_WIDTH)
          .justifyContent(FlexAlign.SpaceBetween)


          Row() {
            Text($r('app.string.charge_time'))
            Text($r('app.string.fill_and_stop'))
          }
          .width(StyleConst.FULL_WIDTH)
          .justifyContent(FlexAlign.SpaceBetween)


          Text('PS：此功能不涉及支付，只做演示使用')
            .fontColor(Color.Red)
        }
        .chargingSty()
        .alignItems(HorizontalAlign.Start)

        Button($r('app.string.charge'))
          .onClick(() => {
            this.getStartChargeData()
          })
          .fontColor(Color.White)
          .backgroundColor($r('app.color.color_0aa671'))
          .width(StyleConst.PERCENT_95)
      }
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .backgroundColor($r('app.color.color_efefef'))
  }

  //公共样式
  @Styles
  chargingSty(){
    .width(StyleConst.PERCENT_95)
    .padding($r('app.float.size_10'))
    .borderRadius($r('app.float.size_10'))
    .backgroundColor($r('app.color.color_fff'))
  }
}