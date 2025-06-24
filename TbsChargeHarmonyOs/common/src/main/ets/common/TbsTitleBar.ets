import { ScreenUtils, StyleConst } from "../../../../Index"
import { router } from "@kit.ArkUI"

@Component
export struct TbsTitleBar {
  @Prop title: ResourceStr = ''
  @Prop titleColor: ResourceColor = Color.Black
  @Prop showBackIcon: boolean = false
  @Prop titleBarBgColor: ResourceColor = Color.White

  build() {
    Column() {
      Blank().height(ScreenUtils.getInstance().getStatusBarHeight())
      RelativeContainer() {
        Text(this.title)
          .fontColor(this.titleColor)
          .fontSize($r('app.float.size_text_16'))
          .textAlign(TextAlign.Center)
          .maxLines(1)
          .textOverflow({ overflow: TextOverflow.Ellipsis })
          .id('text')
          .alignRules({
            center: { anchor: "__container__", align: VerticalAlign.Center },
            middle: { anchor: "__container__", align: HorizontalAlign.Center }
          })
        if (this.showBackIcon) {
          Image($r('app.media.common_back_black'))
            .width($r('app.float.size_20'))
            .margin({ left: $r('app.float.size_5') })
            .id('image')
            .alignRules({
              left: { anchor: "__container__", align: HorizontalAlign.Start },
              center: { anchor: "__container__", align: VerticalAlign.Center },
            })
            .onClick(() => {
              router.back()
            })
        }
      }
      .width(StyleConst.FULL_WIDTH)
      .height($r('app.float.size_45'))
    }
    .width(StyleConst.FULL_WIDTH)
    .backgroundColor(this.titleBarBgColor)
  }
}