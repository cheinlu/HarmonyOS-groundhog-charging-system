import { window } from "@kit.ArkUI"
import { BusinessError } from "@kit.BasicServicesKit"
import { StyleConst } from "../../../../Index"

@CustomDialog
@Component
export struct ImagePreviewView {
  controller: CustomDialogController
  @Prop imageUrl: ResourceStr
  @State windowWidth: number = 0;
  @State windowHeight: number = 0;

  aboutToAppear(): void {
    // 获取应用主窗口的宽高
    window.getLastWindow(getContext(this), (err: BusinessError, data: window.Window) => {
      let rect: window.Rect = data.getWindowProperties().windowRect;
      this.windowWidth = px2vp(rect.width);
      this.windowHeight = px2vp(rect.height);
      data.on("windowSizeChange", this.windowSizeChangeCallback)
    })
  }

  // 窗口尺寸变化的监听回调
  private windowSizeChangeCallback = (size: window.Size) => {
    this.windowWidth = px2vp(size.width);
    this.windowHeight = px2vp(size.height);
  }

  build() {
    Stack({ alignContent: Alignment.Top }) {
      Image(this.imageUrl)
        .alt($r('app.media.common_preview_img'))//接口有返回图片就使用接口的，否则使用本地固定图片展示
        .width(this.windowWidth)
        .height(this.windowHeight)
        .objectFit(ImageFit.Contain)

      Image($r('app.media.common_close_icon_white'))
        .width($r('app.float.size_30'))
        .height($r('app.float.size_30'))
        .onClick(() => {
          this.controller.close()
        })
        .position({
          x: 12,
          y: -5
        })
        .margin({ top: $r('app.float.size_63') })
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .backgroundColor(Color.Black)
    .expandSafeArea([SafeAreaType.SYSTEM, SafeAreaType.CUTOUT], [SafeAreaEdge.TOP, SafeAreaEdge.BOTTOM])
  }
}