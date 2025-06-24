import { LoadingView, ScreenUtils, StyleConst, toast } from "@tbs/common";
import { ChargeType } from "../model/ChargeModel";
import { ChargeListView } from "../view/ChargeListView";
import { ChargeSearchBarView } from "../view/ChargeSearchBarView";
import { getStations } from "../viewModel/ChargeViewModel";

@Component
export struct ChargePage {
  @State @Watch('onChangeName') name: string = ''
  @State chargeList: ChargeType[] = []
  @State isLoading: boolean = false
  inputController: TextInputController = new TextInputController()

  aboutToAppear(): void {
    this.getStationData()
  }

  //获取充电站数据
  async getStationData() {
    this.isLoading = true
    let data = await getStations(this.name)
    if (data && data.List) {
      this.chargeList = data.List
      this.isLoading = false
    }
  }

  //监听输入框的输入内容的变化，若为空时再调接口
  onChangeName() {
    if (this.name == '') {
      this.getStationData()
    }
  }

  //顶部搜索
  @Builder
  chargeSearchBar() {
    ChargeSearchBarView({
      name: this.name,
      maxLength: 30,
      inputController: this.inputController,
      searchButtonCallback: (value) => {
        if (!value) {
          toast($r('app.string.please_search_content'))
        } else {
          this.name = value.replace(/\s+/g, '')
          this.getStationData()
        }
      }
    })
  }

  //搜索结果
  @Builder
  chargeSearchList() {
    ChargeListView({ chargeList: this.chargeList })
  }

  build() {
    Column() {
      this.chargeSearchBar() //顶部搜索
      Stack({ alignContent: Alignment.Top }) {
        if (this.isLoading) {
          LoadingView()
            .margin({ bottom: 95 + ScreenUtils.getInstance().getStatusBarHeight() })
        } else {
          this.chargeSearchList() //搜索结果
        }
      }
      .width(StyleConst.FULL_WIDTH)
      .height(StyleConst.FULL_HEIGHT)
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .backgroundColor($r('app.color.color_f0f2f4'))
  }
}