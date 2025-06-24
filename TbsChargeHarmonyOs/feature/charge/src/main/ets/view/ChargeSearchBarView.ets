import { ScreenUtils, StyleConst, TbsTitleBar } from "@tbs/common"

@Extend(TextInput)
function orderInputStyle(focus: boolean) {
  .textAlign(TextAlign.Start)
  .fontColor(Color.Black)
  .placeholderColor($r('app.color.color_898e92'))
  .placeholderFont({ size: $r('app.float.size_text_13') })
  .fontSize($r('app.float.size_text_14'))
  .backgroundColor(Color.Transparent)
  .enterKeyType(EnterKeyType.Search) //设置输入法回车键类型 显示为搜索样式
  .defaultFocus(focus)
}

@Component
export struct ChargeSearchBarView {
  @Link name: string
  @State cancelMode: CancelButtonStyle = CancelButtonStyle.INVISIBLE // 输入框清除按钮模式
  inputController?: TextInputController
  keyboardFocus: boolean = false // 控制键盘首次是否自动弹出
  maxLength: number = Infinity
  searchButtonCallback?: (value: string) => void // 点击搜索按钮回调

  //输入框
  @Builder
  inputBuilder() {
    Row() {
      Image($r('app.media.search_icon'))
        .width($r('app.float.size_18'))
        .height($r('app.float.size_18'))
        .margin({ left: $r('app.float.size_7') })
      TextInput({
        text: this.name,
        placeholder: $r('app.string.search_content'),
        controller: this.inputController
      })
        .orderInputStyle(this.keyboardFocus)
        .maxLines(this.maxLength)
        .layoutWeight(1)
        .cancelButton({
          style: this.cancelMode,
          icon: { src: $r('app.media.search_delete'), size: $r('app.float.size_14') }
        })
        .padding({ left: $r('app.float.size_6') })
        .onChange((value: string) => {
          this.name = value
          if (value) {
            //输入框有值时，清除按钮显示
            this.cancelMode = CancelButtonStyle.INPUT
          }
        })
        .onSubmit(() => {
          this.lostKeyboard()
          if (this.searchButtonCallback) {
            this.searchButtonCallback(this.name)
          }
        })
    }
    .height($r('app.float.size_34'))
    .clip(true)
    .layoutWeight(1)
    .backgroundColor($r('app.color.color_fff'))
    .borderRadius($r('app.float.size_8'))
    .margin({ left: $r('app.float.size_6'), right: $r('app.float.size_8') })
  }

  //搜索按钮
  @Builder
  searchButtonBuilder() {
    Text($r('app.string.search'))
      .fontSize($r('app.float.size_text_14'))
      .fontColor($r('app.color.color_0aa671'))
      .fontWeight(FontWeight.Medium)
      .height($r('app.float.size_30'))
      .padding({ left: $r('app.float.size_12'), right: $r('app.float.size_12') })
      .backgroundColor(Color.White)
      .borderRadius($r('app.float.size_6'))
      .onClick(() => {
        if (this.searchButtonCallback) {
          this.searchButtonCallback(this.name)
        }
        if (this.name) {
          this.lostKeyboard()
        }
      })
  }

  lostKeyboard() {
    if (this.inputController) {
      this.inputController.stopEditing()
    }
  }

  build() {
    Column() {
      TbsTitleBar({title:$r('app.string.charge'), titleBarBgColor: $r('app.color.color_0aa671') })
      Row() {
        //输入框
        this.inputBuilder()
        //搜索按钮
        this.searchButtonBuilder()
      }
      .width(StyleConst.FULL_WIDTH)
      .padding({ left: $r('app.float.size_12'), right: $r('app.float.size_12') })
      .margin({ bottom: $r('app.float.size_15') })
    }
    .width(StyleConst.FULL_WIDTH)
    .backgroundColor($r('app.color.color_0aa671'))
  }
}