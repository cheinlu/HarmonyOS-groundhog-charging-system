import { StyleConst } from "@tbs/common"

@Component
export struct MinePrivacyTitleItemView {
  title?: Resource | string = '' //标题
  subTitleLeft?: Resource | string = '' //子标题
  @Prop rightTitle?: Resource | string = '' //右边标题
  showBottomLine: Boolean = false //底部分割线
  eventClick?: () => void //item点击
  linkClick?: () => void //超链接点击
  build() {
    Column() {
      Row() {
        Text(this.title)
          .margin({ left: $r('app.float.size_16'), right: $r('app.float.size_5') })
          .fontColor(Color.Black)
          .fontSize($r('app.float.size_text_15'))

        Row() {
          //右边文案
          Text(this.rightTitle)
            .margin({ right: $r('app.float.size_4') })
            .fontColor($r('app.color.color_898e92'))
            .fontSize($r('app.float.size_text_13'))

          Image($r("app.media.mine_min_arrow_gray"))
            .width($r('app.float.size_16'))
            .height($r('app.float.size_24'))
            .margin({ right: $r('app.float.size_8') })
            .objectFit(ImageFit.None)
        }
      }
      .width(StyleConst.FULL_WIDTH)
      .justifyContent(FlexAlign.SpaceBetween)
      .margin({ top: $r('app.float.size_15'), bottom: $r('app.float.size_6') })
      .onClick(()=>{
        if (this.eventClick) {
          this.eventClick()
        }
      })

      //子标题
      Row() {
        Text(this.subTitleLeft)
          .fontColor($r('app.color.color_91989E'))
          .fontSize($r('app.float.size_text_12'))
          .maxLines(2)
          .textOverflow({ overflow: TextOverflow.Ellipsis })
          .layoutWeight(1)
      }
      .width(StyleConst.FULL_WIDTH)
      .padding({ left: $r('app.float.size_16'), right: $r('app.float.size_12') })
      .margin({ bottom: $r('app.float.size_14') })
      .alignItems(VerticalAlign.Top)

      if (this.showBottomLine) {
        Row() {
          Blank().backgroundColor($r('app.color.color_dce0e4'))
            .height(0.5)
            .width(StyleConst.FULL_WIDTH)
        }.width(StyleConst.FULL_WIDTH)
        .padding({ left: $r('app.float.size_16') })
      }

    }
    .width(StyleConst.FULL_WIDTH)
    .backgroundColor(Color.White)
  }
}