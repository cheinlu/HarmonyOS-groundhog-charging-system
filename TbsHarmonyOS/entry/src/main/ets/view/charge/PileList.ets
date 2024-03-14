//表头
import router from '@ohos.router'
import PileType from '../../viewModel/PileType'
@Extend(Text) function TableHeaderSty(){
  .width('30%')
  .height(50)
  .textAlign(TextAlign.Center)
  .fontWeight(FontWeight.Bold)
  .fontSize(20)
}
//表行
@Extend(Text) function TableRootSty(){
  .width('30%')
  .height(45)
  .textAlign(TextAlign.Center)
  .borderRadius(10)
  .fontColor(Color.White)
}
@Component
export default struct PileList {

  @Link Piles:PileType[]


  build() {
    Column(){
      Flex({ direction: FlexDirection.Row }) {
        Text('编号').TableHeaderSty()
        Text('状态').TableHeaderSty().margin({left:20,right:20})
        Text('充电').TableHeaderSty()
      }
      .width('90%')
      ForEach(this.Piles,(pile:PileType) => {
        Flex({ direction: FlexDirection.Row }) {
          Row({space:18}){
            Text(pile.code).TableRootSty().backgroundColor(Color.Grey)
            if(pile.state == 0){
              Text('空闲中').TableRootSty().backgroundColor(Color.Orange)
            }else if(pile.state == 1){
              Text('充电中').TableRootSty().backgroundColor(Color.Green)
            }else{
              Text('故障中').TableRootSty().backgroundColor(Color.Red)
            }
            if(pile.state==0){
              Text('充电').TableRootSty().backgroundColor(Color.Green).onClick(()=>{
                router.pushUrl(
                  {
                    url:'pages/Charging',
                    params:{id:pile.id,code:pile.code,stationName:pile.stationName}
                  },
                  router.RouterMode.Single,
                  err => {
                    if(err){
                      console.log(`路由失败，errCode:${err.code},errMsg:${err.message}`)
                    }
                  }
                )
              })
            }else{
              Text('NO').TableRootSty().backgroundColor(Color.Brown)
            }

          }
        }
        .width('90%')
        .margin({top:10})
      })




    }
    .width('100%')
  }

}