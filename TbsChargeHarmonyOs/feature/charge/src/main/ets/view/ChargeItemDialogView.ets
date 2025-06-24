import { StorageConst, ScreenUtils, StyleConst } from "@tbs/common"
import { pileParam, pileParamType, PileType } from "../model/ChargeModel"

@Component
export struct ChargeItemDialogView {
  @StorageProp(StorageConst.SCREEN_WIDTH_VP) screenWidth: number =
    ScreenUtils.getInstance().getScreenWidthVp()
  @Prop Piles: PileType[]
  clickPileItemCallBack?: (item: PileType) => void

  @Builder
  pileItem(item: PileType) {
    Row() {
      //左
      this.pileItemLeftBuilder(item)

      //右
      this.pileItemRightBuilder(item)
    }
    .padding({ top: $r('app.float.size_12'), bottom: $r('app.float.size_12') })
    .width(this.screenWidth - 12 * 2)
    .margin({ bottom: $r('app.float.size_12') })
    .borderRadius($r('app.float.size_12'))
    .backgroundColor(Color.White)
    .onClick(() => {
      if (this.clickPileItemCallBack) {
        this.clickPileItemCallBack(item)
      }
    })
  }

  //桩左侧UI
  @Builder
  pileItemLeftBuilder(item: PileType) {
    Column({ space: 10 }) {
      Image(this.pileImage(item.state ?? 0))
        .height($r('app.float.size_30'))
      Text(this.pileStateText(item.state ?? 0))
        .textAlign(TextAlign.Center)
        .fontSize($r('app.float.size_text_10'))
        .fontColor(this.pileStateTextColor(item.state ?? 0))
        .width($r('app.float.size_40'))
        .height($r('app.float.size_20'))
        .border({
          width: $r('app.float.size_1'),
          color: this.pileStateTextColor(item.state ?? 0),
          style: BorderStyle.Solid,
          radius: $r('app.float.size_20')
        })
    }
    .margin({ left: $r('app.float.size_20'), right: $r('app.float.size_20') })
  }

  //桩右侧UI
  @Builder
  pileItemRightBuilder(item: PileType) {
    Column({ space: 10 }) {
      Text() {
        Span($r('app.string.pile_number')).fontColor($r('app.color.color_898e92'))
        Span(item.code).fontColor(Color.Black)
      }.fontSize($r('app.float.size_text_14'))

      Text() {
        Span($r('app.string.station_name')).fontColor($r('app.color.color_898e92'))
        Span(item.stationName).fontColor(Color.Black)
      }.fontSize($r('app.float.size_text_14'))

      List({ space: 10 }) {
        ForEach(pileParam, (item: pileParamType) => {
          ListItem() {
            Row() {
              Text(item.text)
                .textAlign(TextAlign.Center)
                .fontSize($r('app.float.size_text_11'))
                .fontColor(Color.Black)
            }
            .backgroundColor($r('app.color.color_f5f5f5'))
            .padding({
              top: $r('app.float.size_5'),
              bottom: $r('app.float.size_5'),
              left: $r('app.float.size_8'),
              right: $r('app.float.size_8')
            })
          }
        })
      }
      .width(StyleConst.FULL_WIDTH)
      .listDirection(Axis.Horizontal)
      .scrollBar(BarState.Off)
      .edgeEffect(EdgeEffect.None)
      .padding({right:$r('app.float.size_10')})
    }
    .layoutWeight(1)
    .alignItems(HorizontalAlign.Start)
  }

  //区分不同状态的桩图片
  pileImage(state: number) {
    switch (state) {
      case 0:
        return $r('app.media.pile_idle')
      case 1:
        return $r('app.media.pile_charging_in_progress')
      case 2:
        return $r('app.media.pile_faults')
      default:
        return $r('app.media.pile_idle')
    }
  }

  //区分不同状态的文本
  pileStateText(state: number) {
    switch (state) {
      case 0:
        return $r('app.string.idle')
      case 1:
        return $r('app.string.charging_in_progress')
      case 2:
        return $r('app.string.faults')
      default:
        return $r('app.string.idle')
    }
  }

  //区分不同状态的文本颜色
  pileStateTextColor(state: number) {
    switch (state) {
      case 0:
        return $r('app.color.color_009688')
      case 1:
        return $r('app.color.color_d68751')
      case 2:
        return $r('app.color.color_F44336')
      default:
        return $r('app.color.color_009688')
    }
  }

  build() {
    List() {
      ForEach(this.Piles, (item: PileType) => {
        ListItem() {
          this.pileItem(item)
        }
        .width(StyleConst.FULL_WIDTH)
      })
    }
    .width(StyleConst.FULL_WIDTH)
  }
}