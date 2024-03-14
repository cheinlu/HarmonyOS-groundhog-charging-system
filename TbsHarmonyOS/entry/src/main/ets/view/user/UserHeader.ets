@Component
export default struct UserHeader {
  @State account:string = AppStorage.Get('account')

  build() {
    Column({space:15}){
      Text('我的')
        .fontColor(Color.White)
        .fontSize(20)
        .margin({top:40})
      Image($r('app.media.Groundhog')).width(60).borderRadius(50)
      Text(this.account).fontColor(Color.White)
    }
    .width('100%')
    .height('30%')
    .padding({left:20,right:20})
    .backgroundColor($r('app.color.charge_header'))
  }
}