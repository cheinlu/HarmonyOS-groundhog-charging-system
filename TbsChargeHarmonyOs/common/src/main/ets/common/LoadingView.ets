import { StyleConst } from "."

@Component
export struct LoadingView {
  build() {
    Row() {
      LoadingProgress()
        .width($r('app.float.size_50'))
        .height($r('app.float.size_50'))
        .margin({ right: $r('app.float.size_10') })
      Text($r('app.string.loading'))
        .fontSize($r('app.float.size_text_14'))
        .fontColor($r('app.color.color_666'))
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .justifyContent(FlexAlign.Center)
  }
}