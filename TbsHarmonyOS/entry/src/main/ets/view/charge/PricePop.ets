import PriceType from '../../viewModel/PriceType'

//标题样式
@Extend(Text) function PriceHeaderSty(){
  .fontSize(18).fontWeight(FontWeight.Bold).fontColor(Color.White)
}
//行内样式
@Extend(Text) function PriceCellSty(){
  .fontSize(20).textAlign(TextAlign.Center)
}
@Component
export default struct PricePop {
  @Link prices:PriceType[]
  build() {
    Column({space:10}){
      Text('收费标准').textAlign(TextAlign.Center).fontWeight(FontWeight.Bold).fontSize(25)
      Column({space:10}){
        Row({space:50}){
          Text('收费时段').PriceHeaderSty()
          Text('收费(元/度)').PriceHeaderSty().margin({left:20})
        }
        .width('90%')
        .padding(15)
        .backgroundColor('#0aa671')
        .borderRadius(5)
        Column({space:10}){
          List({space:10}){
            ForEach(this.prices,(item:PriceType) => {
              ListItem(){
                Column(){
                  Row({space:100}){
                    Text(`${item.startHour} - ${item.endHour}`).width(80).PriceCellSty()
                    Text(`${item.price}`).width(60).PriceCellSty()
                  }
                  .width('90%')
                  .height(40)
                  Divider().width('95%').opacity(0.6)
                }
              }
            })
          }
          .width('90%')
          .layoutWeight(1)
        }
        .width('90%')
        .height(180)
        .padding(10)
        .borderWidth(1).borderColor(Color.Gray)
      }
      .width('90%')
    }
    .width('100%')
    .padding(10)
  }
}