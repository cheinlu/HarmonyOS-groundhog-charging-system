import { StyleConst } from "@tbs/common"
import { PriceType } from "../model/ChargeModel"

@Component
export struct ChargePriceDialogView {
  @Prop priceList: PriceType[]
  @Prop pileCode: string

  build() {
    Column() {
      // 顶部标题
      this.buildHeader()

      // 价格表格区域
      this.buildPriceTable()

      // 分隔线
      Divider()
        .strokeWidth(1)
        .color($r('app.color.color_e0e0e0'))
        .margin({ top: $r('app.float.size_25'), bottom: $r('app.float.size_20') })

      // 底部提示
      this.buildFooter()
    }
    .padding($r('app.float.size_20'))
    .backgroundColor($r('app.color.color_f5f7fa'))
    .borderRadius($r('app.float.size_12'))
    .width(StyleConst.FULL_WIDTH)
    .height('auto')
  }

  // 顶部标题组件
  @Builder
  buildHeader() {
    Text() {
      Span($r('app.string.pile_number_1'))
      Span(this.pileCode)
    }
    .fontSize($r('app.float.size_text_14'))
    .fontColor($r('app.color.color_666'))
    .margin({ bottom: $r('app.float.size_15') })
  }

  // 价格表格组件 - 简化为两列
  @Builder
  buildPriceTable() {
    Column() {
      // 表头 - 左右分布
      Row() {
        Text($r('app.string.charging_period'))
          .fontSize($r('app.float.size_text_16'))
          .fontWeight(FontWeight.Bold)
          .layoutWeight(1) // 左边占据1份空间

        Text($r('app.string.price_yuan'))
          .fontSize($r('app.float.size_text_16'))
          .fontWeight(FontWeight.Bold)
          .layoutWeight(1) // 右边占据1份空间
      }
      .width(StyleConst.FULL_WIDTH)
      .padding({ bottom: $r('app.float.size_10') })

      // 表格内容 - 左右分布
      ForEach(this.priceList,
        (item: PriceType, index: number) => {
          Row() {
            Text(`${item.startHour} - ${item.endHour}`)
              .fontSize($r('app.float.size_text_16'))
              .fontColor($r('app.color.color_333'))
              .layoutWeight(1) // 左边占据1份空间

            if (item.price) {
              Text(`${item.price}`)
                .fontSize($r('app.float.size_text_16'))
                .fontColor($r('app.color.color_e91e63'))// 价格使用红色强调
                .fontWeight(FontWeight.Medium)
                .layoutWeight(1) // 右边占据1份空间
            }

          }
          .width(StyleConst.FULL_WIDTH)
          .padding({ top: $r('app.float.size_10'), bottom: $r('app.float.size_10') })
          .backgroundColor(index % 2 === 0 ? $r('app.color.color_fff') : $r('app.color.color_f8f9fa'))
        })
    }
    .width(StyleConst.FULL_WIDTH)
    .padding($r('app.float.size_12'))
    .backgroundColor(Color.White)
    .borderRadius($r('app.float.size_8'))
    .shadow({
      radius: $r('app.float.size_4'),
      color: $r('app.color.color_20000'),
      offsetX: 1,
      offsetY: 1
    })
  }

  // 底部提示组件
  @Builder
  buildFooter() {
    Text('*充电费用仅供参考，请以实际收取费用为准')
      .fontSize($r('app.float.size_text_12'))
      .fontColor($r('app.color.color_999'))
      .fontStyle(FontStyle.Italic)
      .alignSelf(ItemAlign.Start)
  }
}

