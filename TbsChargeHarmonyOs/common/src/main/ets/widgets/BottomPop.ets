import { StyleConst } from "../common";
import { ScreenUtils, WindowModel,ComponentUtils } from "../utils";

@CustomDialog
@Component
export struct BottomPop {
  controller: CustomDialogController;
  title: ResourceStr = '';
  titlePadding: Padding = { top: $r('app.float.size_16'), bottom: $r('app.float.size_16') }
  titleFontSize: ResourceStr = $r('app.float.size_text_16')
  titleTextColor: ResourceColor = Color.Black
  titleBgColor: ResourceColor = Color.Transparent
  contentBgColor: ResourceColor = Color.Transparent
  bgColor?: ResourceColor = Color.White
  // 软键盘是否显示
  @State keyBoardVisible: boolean = false
  keyboardAvoidMode: KeyboardAvoidMode = KeyboardAvoidMode.DEFAULT
  minHeight: Length = ScreenUtils.getInstance().getScreenHeightVp() * 0.3 // 弹窗最低高度
  maxHeight: Length = ScreenUtils.getInstance().getScreenHeightVp() * 0.8 // 弹窗最大高度
  @State dialogMaxHeight: number = 0 // 通过 maxHeight 计算出来的内容的高度

  @Builder
  child() {
  };

  @BuilderParam childBuilderParam: () => void = this.child;

  aboutToAppear(): void {
    if (this.isKeyboardAvoidNone()) {
      this.getKeyBoardVisible()
    }
    this.getTitleHeight()
  }

  // 软键盘是否拉起
  getKeyBoardVisible() {
    WindowModel.getInstance().getKeyBoardVisible((visible) => {
      this.keyBoardVisible = visible
    })
  }

  // 弹窗是否在拉起软键盘时进行自动避让
  isKeyboardAvoidNone() {
    return this.keyboardAvoidMode === KeyboardAvoidMode.NONE // 弹窗不避让软键盘
  }

  getTitleHeight() {
    ComponentUtils.getComponentHeightSync('BottomPop_title', (componentHeight) => {
      const safeBottomHeight = ScreenUtils.getInstance().getSafeBottomHeight() // 底部安全区高度
      // 内容最大高度 = 屏幕80%高度 - title高度 - 底部安全区高度
      this.dialogMaxHeight = Number(this.maxHeight) - componentHeight - safeBottomHeight
    })
  }

  build() {
    Stack({ alignContent: Alignment.TopEnd }) {
      Column() {
        Text(this.title)
          .padding(this.titlePadding)
          .fontColor(this.titleTextColor)
          .fontSize(this.titleFontSize)
          .fontWeight(FontWeight.Medium)
          .textAlign(TextAlign.Center)
          .width(StyleConst.FULL_WIDTH)
          .height($r('app.float.size_50'))
          .backgroundColor(this.titleBgColor)
          .id('BottomPop_title')
          .borderRadius({
            topLeft: $r('app.float.size_16'),
            topRight: $r('app.float.size_16')
          })
        Column() {
          this.childBuilderParam()
        }
        .backgroundColor(this.contentBgColor)
        .constraintSize({
          maxHeight: this.dialogMaxHeight,
        })
      }
      .padding({
        bottom: ScreenUtils.getInstance().getSafeBottomHeight(), // 设置padding避开底部安全区
      })
      .width(StyleConst.FULL_WIDTH)
      .backgroundColor(this.bgColor)
      .borderRadius({
        topLeft: $r('app.float.size_16'),
        topRight: $r('app.float.size_16')
      })
      .constraintSize({
        maxHeight: this.maxHeight,
        minHeight: this.minHeight
      })

      Image($r("app.media.common_bottom_sheet_close"))
        .size({ width: $r('app.float.size_20'), height: $r('app.float.size_20') })
        .margin({ top: $r('app.float.size_16'), right: $r('app.float.size_16') })
        .responseRegion({
          x: '-25%',
          y: '-25%',
          width: '150%',
          height: '150%'
        })
        .onClick(() => {
          this.controller.close();
        })
    }
  }
}