import { ScreenUtils, StyleConst } from "@tbs/common";

@Component
export struct CustomScanView {
  @State dividerOffset: number = 0
  @Link isFlashlightOn: boolean // 手电筒开启状态
  @Prop isShowFlashlight: boolean = false // 是否隐藏灯光
  @Prop isScanned: boolean = false // 是否已扫出结果
  onBack?: () => void
  handlerLamplightCompletion?: (isOpen: boolean) => void
  animation_offset_end: number =
    (ScreenUtils.getInstance().getScreenHeightVp() - ScreenUtils.getInstance().getStatusBarHeight() -
      190 - 136)

  @Builder
  topBuilder() {
    Column() {
      Flex({ justifyContent: FlexAlign.SpaceBetween, alignItems: ItemAlign.Center }) {
        Image($r('app.media.scan_back')).width($r('app.float.size_36')).aspectRatio(1)
          .onClick(() => {
            if (this.onBack) {
              this.onBack()
            }
          })
      }
      .height($r('app.float.size_44'))
      .margin({
        top: ScreenUtils.getInstance().getStatusBarHeight(),
        left: $r('app.float.size_16'),
        right: $r('app.float.size_16')
      })
    }
    .backgroundColor(Color.Transparent)
  }

  @Builder
  scanBuilder() {
    Column() {
      Image($r('app.media.scan_strip'))
        .width(StyleConst.FULL_WIDTH)
        .height($r('app.float.size_24'))
        .objectFit(ImageFit.Fill)
        .position({ x: 0, y: $r('app.float.size_80') })
        .translate({ x: 0, y: this.dividerOffset })
    }
    .onAppear(() => {
      this.setQRCodeScanAnimation()
    })
    .width(StyleConst.FULL_WIDTH)
    .height(ScreenUtils.getInstance().getScreenHeightVp() - ScreenUtils.getInstance().getStatusBarHeight() - 325)
  }

  @Builder
  bottomBuilder() {
    RelativeContainer(){
      // 闪光灯按钮，启动相机流后才能使用
      Image(this.isFlashlightOn ? $r('app.media.scan_light_open') : $r('app.media.scan_light_close'))
        .width($r('app.float.size_40'))
        .aspectRatio(1)
        .objectFit(ImageFit.Contain)
        .visibility(this.isShowFlashlight ? Visibility.Visible : Visibility.Hidden)
        .onClick(()=>{
          this.isFlashlightOn = !this.isFlashlightOn
          this.handlerLamplightCompletion && this.handlerLamplightCompletion(this.isFlashlightOn)
        })
        .id('flashLight')
        .alignRules({
          top: { anchor: '__container__', align: VerticalAlign.Top },
          middle: { anchor: '__container__', align: HorizontalAlign.Center },
        })

      // 扫描二维码文案
      Text($r('app.string.scan_barQR_code'))
        .fontSize($r('app.float.size_text_16'))
        .fontColor(Color.White)
        .textAlign(TextAlign.Center)
        .maxLines(2)
        .textOverflow({ overflow: TextOverflow.Ellipsis })
        .backgroundColor(Color.Transparent)
        .borderRadius(0)
        .padding(0)
        .margin({
          left: $r('app.float.size_32'),
          right: $r('app.float.size_32'),
          top: $r('app.float.size_16'),
          bottom: $r('app.float.size_190')
        })
        .id('tips')
        .alignRules({
          top: { anchor: 'flashLight', align: VerticalAlign.Bottom },
          middle: { anchor: '__container__', align: HorizontalAlign.Center },
        })
    }
    .width(StyleConst.FULL_WIDTH)
    .layoutWeight(1)
    .backgroundColor(Color.Transparent)
  }

  setQRCodeScanAnimation() {
    let ANI_CONFIG: AnimateParam = {
      duration: 1500, // 动画时间
      tempo: 0.3, // 动画速率
      curve: Curve.EaseInOut,
      iterations: -1, // 动画是否重复播放
      playMode: PlayMode.Normal,
    }
    animateTo(ANI_CONFIG, () => {
      // 扫描动画结束Y坐标
      this.dividerOffset = this.animation_offset_end - 140
    })
  }

  build() {
    Column() {
      // 顶部按钮
      this.topBuilder()
      // 扫描动画组件
      this.scanBuilder()
      // 底部视图 (提示文案、灯光等)
      this.bottomBuilder()
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
  }
}