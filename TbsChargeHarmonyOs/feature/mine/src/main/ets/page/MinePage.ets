import { StyleConst } from "@tbs/common";
import { MineContentView } from "../view/MineContentView";
import { MineHeaderView } from "../view/MineHeaderView";

@Component
export struct MinePage {
  build() {
    Column() {
      //1.头部
      MineHeaderView()
      //2.主体
      MineContentView()
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .backgroundColor($r('app.color.color_efefef'))
  }
}