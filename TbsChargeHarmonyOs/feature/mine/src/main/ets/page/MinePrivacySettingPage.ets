import { permissionManager, StyleConst, TbsTitleBar, locationManager, PermissionConst } from '@tbs/common'
import { MinePrivacyTitleItemView } from '../view/MinePrivacyTitleItemView'
import { common } from '@kit.AbilityKit'

@Entry({ routeName: "MinePrivacySettingPage" })
@Component
struct MinePrivacySettingPage {
  @State isOpenLocation: boolean = false
  @State isOpenCamera: boolean = false

  onPageShow(): void {
    //每次进来检查定位是否打开，精准定位打开才算开启，反之未开启
    locationManager.checkLocationPermissions().then((isOpen) => {
      this.isOpenLocation = isOpen
    })
    //每次进来检查相机是否打开
    permissionManager.checkPermissions([PermissionConst.cameraPermission]).then((isOpen) => {
      this.isOpenCamera = isOpen
    })
  }

  @Builder
  buildPageContent() {
    Column() {
      //位置定位
      MinePrivacyTitleItemView({
        title: $r('app.string.privacy_settings_address'),
        subTitleLeft: $r('app.string.privacy_settings_address_description'),
        rightTitle: this.isOpenLocation ? $r('app.string.enabled') :
        $r('app.string.setup'),
        showBottomLine: true,
        eventClick: () => {
          permissionManager.openSettings(getContext(this) as common.UIAbilityContext)
        }
      })
        .margin({ top: $r('app.float.size_12') })

      //相机
      MinePrivacyTitleItemView({
        title: $r('app.string.privacy_settings_camera'),
        subTitleLeft: $r('app.string.privacy_settings_camera_description'),
        rightTitle: this.isOpenCamera ? $r('app.string.enabled') :
        $r('app.string.setup'),
        eventClick: () => {
          permissionManager.openSettings(getContext() as common.UIAbilityContext)
        }
      })
    }
    .height(StyleConst.FULL_HEIGHT)
  }

  build() {
    Stack({ alignContent: Alignment.TopStart }) {
      Column() {
        TbsTitleBar({
          title: $r('app.string.setting'),
          showBackIcon: true
        })
        Row() {
          this.buildPageContent()
        }.alignItems(VerticalAlign.Top)
        .width(StyleConst.FULL_WIDTH)
        .layoutWeight(1)
      }
      .justifyContent(FlexAlign.Start)
      .width(StyleConst.FULL_WIDTH)
      .height(StyleConst.FULL_HEIGHT)
      .backgroundColor($r('app.color.color_f0f2f4'))
    }
  }
}