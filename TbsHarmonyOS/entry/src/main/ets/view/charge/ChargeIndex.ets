import ChargeModel from '../../model/ChargeModel'
import ChargeType from '../../viewModel/ChargeType'
import ChargeList from './ChargeList'
import PileInfo from '../../viewModel/PileType'
import PileList from './PileList'
@Component
export default struct ChargeIndex {
  @State Piles:PileInfo[] = []
  @State PileList:PileInfo[] = []
  @State showPanel:boolean = false
  @State stationId:number = 0
  @State searchValue:string = ''
  onPanelShow(){
    this.showPanel = true
    setTimeout(()=>{
      this.PileList = this.Piles
    },100)
  }
  build() {
    Column(){
      //1、搜索
      Column({space:15}){
        Text('充电')
          .fontColor(Color.White)
          .fontSize(20)
          .margin({top:30})
        Search({placeholder:'请输入搜索内容'})
          .placeholderColor(Color.Gray)
          .backgroundColor('#F5F5F5')
          .textFont({size:18})
          .onChange((value: string) => {
            this.searchValue = value
          })
      }
      .width('100%')
      .height('20%')
      .padding({left:20,right:20})
      .backgroundColor($r('app.color.charge_header'))
      //2.充电列表
      ChargeList({searchValue:$searchValue,stationId:$stationId,showPanel:this.onPanelShow.bind(this),Piles:$Piles})
        .layoutWeight(1)
      //3.底部面板
      Panel(this.showPanel){
        Button('关闭').onClick(() => this.showPanel = false)
        PileList({Piles:$Piles})
      }
      .mode(PanelMode.Full)
      .dragBar(false)
      .backgroundMask($r('app.color.light_gray'))
      .backgroundColor(Color.White)
    }
    .width('100%')
    .height('100%')
    .backgroundColor('#EFEFEF')
  }
}