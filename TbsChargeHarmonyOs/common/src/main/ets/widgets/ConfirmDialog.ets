
// 确认弹窗
@CustomDialog
@Component
export struct ConfirmDialog {
  controller: CustomDialogController
  title: ResourceStr = '标题';
  content: ResourceStr = '内容'
  leftBtn: ResourceStr = '取消';
  rightBtn: ResourceStr = '确定';
  isShowInputBox: boolean = false
  theInputContent: string = '' //输入框内容
  onTapLeft?: () => void;
  onTapRight?: (value?: string) => void;
  @BuilderParam contentBuilder: () => void;
  build() {
    Column() {
      Text(this.title)
        .fontSize($r('app.float.size_text_16'))
        .fontColor(Color.Black)
        .fontWeight(FontWeight.Bold)
        .margin({ top: $r('app.float.size_24'), left: $r('app.float.size_12'), right: $r('app.float.size_12') })
        .textAlign(TextAlign.Center)
      if (this.contentBuilder) {
        this.contentBuilder()
      } else {
        Text(this.content)
          .fontSize($r('app.float.size_text_14'))
          .fontColor($r('app.color.color_898e92'))
          .margin({ top: $r('app.float.size_12'), left: $r('app.float.size_24'), right: $r('app.float.size_24') })
          .textAlign(TextAlign.Center)
      }

      Blank()
      Divider().color('#dce0e4').width(311)
        .margin({ top: $r('app.float.size_24') })
      Row() {
        Row() {
          Text(this.leftBtn)
            .fontColor(Color.Black)
        }
        .layoutWeight(1)
        .justifyContent(FlexAlign.Center)
        .onClick(()=>{
          if (this.onTapLeft) {
            this.onTapLeft();
          } else {
            this.controller.close();
          }
        })

        Divider()
          .vertical(true)
          .color('#dce0e4')
        Row() {
          Text(this.rightBtn)
            .fontColor('#0165b8')
            .textAlign(TextAlign.Center)
        }
        .layoutWeight(1)
        .justifyContent(FlexAlign.Center)
        .onClick(()=>{
          if (this.onTapRight) {
            this.onTapRight();
          } else {
            this.controller.close();
          }
        })
      }.height($r('app.float.size_48'))
    }
    .justifyContent(FlexAlign.SpaceBetween)
    .width(311)
    .constraintSize({ maxHeight: '60%', minHeight: $r('app.float.size_149') })
    .backgroundColor(Color.White)
    .borderRadius($r('app.float.size_12'))
  }
}