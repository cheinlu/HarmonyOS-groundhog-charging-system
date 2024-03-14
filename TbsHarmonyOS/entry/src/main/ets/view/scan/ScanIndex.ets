import HeaderComponent from '../../common/components/HeaderComponent'
@Component
export default struct ScanIndex {
  build() {
    Column(){
      Row({space: 5}){
        Text('扫码')
          .fontColor(Color.White)
          .fontSize(20)
      }
      .width('100%')
      .height('10%')
      .justifyContent(FlexAlign.Center)
      .backgroundColor($r('app.color.charge_header'))
      //2.主体
      Text('扫码').margin({top:200})
    }
    .width('100%')
    .height('100%')
    .backgroundColor('#EFEFEF')
  }
}