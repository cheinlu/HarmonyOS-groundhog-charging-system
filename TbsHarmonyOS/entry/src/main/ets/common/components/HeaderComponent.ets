import router from '@ohos.router'
@Component
export default struct HeaderComponent {
  private  title:ResourceStr
  build() {
    Row({space: 5}){
      Image($r('app.media.back'))
        .width(45)
        .onClick(() => {
          router.showAlertBeforeBackPage({
            message: '确定返回吗？'
          })
          router.back()
        })
      Text(this.title)
        .fontColor(Color.White)
        .fontSize(20)
        .margin({left:80})
    }
    .width('100%')
    .height('10%')
    .padding(10)
    .backgroundColor($r('app.color.charge_header'))
  }
}