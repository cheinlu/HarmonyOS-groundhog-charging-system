import ChargeType from '../../viewModel/ChargeType'
import PileType from '../../viewModel/PileType'
import ChargeModel from '../../model/ChargeModel'

@Component
export default struct ChargeList {
  @State charges:ChargeType[] = []
  @Link stationId:number
  @Link Piles:PileType[]
  @Link @Watch('onSearch')searchValue:string
  isLoading: boolean = false
  isMore: boolean = true
  showPanel:()=>void

  //监听搜索框输入值的函数
  onSearch():void{
    if(this.searchValue){
      this.chargeInfo(this.searchValue, true); // 传入true来重置页面编号和清空原有数据
    } else {
      this.chargeInfo('', true); // 传入true来重置页面编号和清空原有数据
    }
  }

  aboutToAppear(){
    //加载充电站数据
    this.chargeInfo('')
  }

  build() {
    Column({space:10}){
      //列表
      List({space:10}){
        ForEach(this.charges,(item:ChargeType) => {
          ListItem(){
            this.chargeItem(item)
          }
          .onClick(() =>{
            this.PileInfo(item.id)
            this.stationId = item.id
            this.showPanel()
          })
        })
      }
      .width('100%')
      .padding(10)
      .layoutWeight(1)
      .onReachEnd(() => {
        if(!this.isLoading && this.isMore){
          this.isLoading = true
          // 翻页
          ChargeModel.pageNo++
          this.chargeInfo('')
        }
      })
    }
    .width('100%')
  }

  @Builder chargeItem(item:ChargeType){
    Column({space:12}){
      Text(item.name).fontWeight(FontWeight.Bold)
      Row({space:8}){
        Text(`￥ ${item.pricePerHour}`).fontColor($r('app.color.pile_tanan'))
        Text('/度').fontColor($r('app.color.limit'))
        Blank()
        Text(item.tenantName).fontSize(18).fontColor($r('app.color.pile_tanan'))
        Blank()
        Image($r('app.media.fast')).width(20)
        Text(`闲${item.freePileTotal}`).fontWeight(FontWeight.Bold)
        Text(`/${item.pileTotal}`)
      }
      .width('100%')
      Row({space:8}){
        Text(`${item.distance.toFixed(2)}km`)
        Image($r('app.media.location')).width(20)
        Image($r('app.media.imgUrl')).width(20)
      }

    }
    .backgroundColor('#FFF')
    .height(120)
    .padding(10)
    .borderRadius(20)
    .alignItems(HorizontalAlign.Start)
  }

  //充电站数据
  async chargeInfo(name, resetPage = false){
    if(resetPage){
      ChargeModel.pageNo = 1; // 重置页面编号为1
      this.charges = []; // 清空原有数据
      this.isMore = true; // 重置isMore为true，表示还有更多数据可加载
    }
    let res:any = await ChargeModel.getStations(name)
    if(res.responseCode === 200){
      let chargeRes = JSON.parse(res.result.toString())
      if(chargeRes.code == 0){
        this.charges = this.charges.concat(chargeRes.data.List);
        this.isLoading = false;
        if(!chargeRes.data.List || chargeRes.data.List.length === 0){
          this.isMore = false;
        }
      }
    }
  }

  //充电桩数据
  async PileInfo(stationId){
    let res:any = await ChargeModel.getPiles(stationId)
    if(res.responseCode === 200){
      let result = res.result.toString()
      let pileRes = JSON.parse(result.toString())
      if(pileRes.code == 0){
        this.Piles = pileRes.data.List
      }
    }
  }
}