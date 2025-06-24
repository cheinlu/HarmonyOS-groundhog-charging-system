import { StyleConst } from "../../../../Index"

@Component
export struct BottomLoadingBuilder {
  @Prop haveNextPage: boolean
  @Prop loading: boolean
  @Prop dataLength: number

  build() {
    if (this.haveNextPage && this.loading) {
      Row() {
        LoadingProgress()
          .width($r('app.float.size_30'))
          .height($r('app.float.size_30'))
          .margin({ right: $r('app.float.size_10') })
        Text($r('app.string.loading'))
          .fontSize($r('app.float.size_text_14'))
          .fontColor($r('app.color.color_666'))
      }
      .width(StyleConst.FULL_WIDTH)
      .height($r('app.float.size_60'))
      .justifyContent(FlexAlign.Center)
    } else if (!this.haveNextPage && this.dataLength > 0) {
      Text($r('app.string.there_is_no_more_data'))
        .fontSize($r('app.float.size_text_14'))
        .fontColor($r('app.color.color_999'))
        .textAlign(TextAlign.Center)
        .width(StyleConst.FULL_WIDTH)
        .margin({ top: $r('app.float.size_10'), bottom: $r('app.float.size_20') })
    }
  }
}