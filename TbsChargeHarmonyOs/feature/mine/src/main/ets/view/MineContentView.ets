import { commonStorage, RoutePageConst, StorageConst, StyleConst, ConfirmDialog, toast } from "@tbs/common"
import UserCenterConst from "@tbs/common/src/main/ets/manager/UserCenterConst"
import { orderOrServiceList, panelData, panelList, UserInfoData } from "../model/MineModel"
import { router } from "@kit.ArkUI"
import '@tbs/charge/src/main/ets/page/ChargeOrderPage'
import '../page/MinePrivacySettingPage'

@Component
export struct MineContentView {
  userInfoData: UserInfoData | null = UserCenterConst.getUserInfoData()

  //局部公共样式函数
  @Styles
  userPageSty(){
    .width(StyleConst.PERCENT_95)
    .padding($r('app.float.size_10'))
    .borderRadius($r('app.float.size_10'))
    .backgroundColor($r('app.color.color_fff'))
  }

  // 退出登录-二次弹框
  logoutDialogController: CustomDialogController = new CustomDialogController({
    builder: this.buildLogoutPopView,
    customStyle: true,
    alignment: DialogAlignment.Center,
    autoCancel: false,
    maskColor: $r('app.color.color_80000000')
  })

  // 退出登录-二次弹框
  @Builder
  buildLogoutPopView() {
    ConfirmDialog({
      title: $r('app.string.friendly_reminder'),
      content: $r('app.string.logout_tip_detail'),
      leftBtn: $r('app.string.button_cancel'),
      rightBtn: $r('app.string.button_sure'),
      onTapLeft: () => {
        // 取消
        this.logoutDialogController.close()
      },
      onTapRight: () => {
        // 确定
        this.logoutDialogController.close()
        commonStorage.deleteValueSync(StorageConst.TOKEN)
        commonStorage.deleteValueSync(StorageConst.USER_INFO)
        router.replaceUrl({ url: RoutePageConst.LoginPage })
      }
    })
  }

  build() {
    Column({ space: 10 }) {
      //1.充电余额
      Column({ space: 6 }) {
        if (this.userInfoData?.balance) {
          Text(`${this.userInfoData?.balance.toFixed(2)}元`)
        }
        Row() {
          Text($r('app.string.charge_balance'))
          Text('>')
        }
      }
      .margin({ top: -20 })
      .alignItems(HorizontalAlign.Start)
      .userPageSty()

      //2.充电订单
      Row({ space: 10 }) {
        ForEach(orderOrServiceList, (item: panelData) => {
          Column() {
            Image(item.image)
              .objectFit(ImageFit.Contain)
              .size({ width: $r('app.float.size_35'), height: $r('app.float.size_35') })
            Text(item.name)
              .fontSize($r('app.float.size_text_12'))
              .textAlign(TextAlign.Center)
              .margin({ top: $r('app.float.size_6') })
          }.onClick(() => {
            if (item.type == 1) {
              router.pushNamedRoute({
                name: 'ChargeOrderPage'
              })
            } else {
              toast('功能开发中，敬请期待！')
            }
          })
        })
      }

      .justifyContent(FlexAlign.SpaceAround)
      .userPageSty()

      //3.用户协议
      Flex({ wrap: FlexWrap.Wrap, alignContent: FlexAlign.SpaceAround }) {
        ForEach(panelList, (item: panelData) => {
          Column({ space: 4 }) {
            Image(item.image)
              .objectFit(ImageFit.Contain)
              .size({ width: $r('app.float.size_35'), height: $r('app.float.size_35') })
            Text(item.name)
              .fontSize($r('app.float.size_text_12'))
              .textAlign(TextAlign.Center)
              .margin({ top: $r('app.float.size_6') })
          }.width(StyleConst.PERCENT_25).margin({ bottom: $r('app.float.size_10') })
          .onClick(() => {
            if (item.type == 1) {
              router.pushNamedRoute({
                name: "MinePrivacySettingPage"
              })
            } else {
              toast('功能开发中，敬请期待！')
            }
          })
        })
      }
      .width(StyleConst.PERCENT_95)
      .padding({ top: $r('app.float.size_10'), left: $r('app.float.size_10'), right: $r('app.float.size_10') })
      .borderRadius($r('app.float.size_10'))
      .backgroundColor($r('app.color.color_fff'))

      //4.退出登录
      Row() {
        Text($r('app.string.log_out')).margin({ left: $r('app.float.size_10') })
        Image($r('app.media.icon_arrow_right')).width($r('app.float.size_20')).fillColor(Color.Black)
      }
      .justifyContent(FlexAlign.SpaceBetween)
      .userPageSty()
      .padding({ top: $r('app.float.size_20'), bottom: $r('app.float.size_20') })
      .onClick(() => {
        this.logoutDialogController.open()
      })
    }
  }
}
