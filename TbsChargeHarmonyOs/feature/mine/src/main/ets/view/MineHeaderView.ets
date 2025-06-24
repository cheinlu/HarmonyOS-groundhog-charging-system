import { ScreenUtils, StyleConst } from "@tbs/common"
import { UserInfoData } from "../model/MineModel"
import UserCenterConst from '@tbs/common/src/main/ets/manager/UserCenterConst'

@Component
export struct MineHeaderView {
  userInfoData: UserInfoData | null = UserCenterConst.getUserInfoData()

  build() {
    Column() {
      Blank().height(ScreenUtils.getInstance().getStatusBarHeight()).backgroundColor($r('app.color.color_0aa671'))
      Image($r('app.media.Groundhog')).width($r('app.float.size_60')).borderRadius($r('app.float.size_50'))
      Text(this.userInfoData?.username).fontColor(Color.White)
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.PERCENT_20)
    .backgroundColor($r('app.color.color_0aa671'))
  }
}