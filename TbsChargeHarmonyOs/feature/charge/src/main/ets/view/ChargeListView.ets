import { BottomPop, StyleConst, ThirdApp, thirdAppManager, toast, locationManager, ImagePreviewView } from "@tbs/common"
import { AddressData, ChargeType, mapListItem, PileType } from "../model/ChargeModel"
import { getPiles } from "../viewModel/ChargeViewModel"
import { ChargeItemDialogView } from "./ChargeItemDialogView"
import { router } from "@kit.ArkUI"
import '../page/ChargingPage'
import { OpenLinkOptions, common } from "@kit.AbilityKit"
import { BusinessError } from "@kit.BasicServicesKit"
import { geoLocationManager } from "@kit.LocationKit"
import { MapDialogCall } from "./MapDialogCall"

@Component
export struct ChargeListView {
  @Prop chargeList: ChargeType[]
  @State Piles: PileType[] = []
  @State destinationLongitude: number = 0
  @State destinationLatitude: number = 0
  @State destinationName: string = ''
  @State mapList: mapListItem[] = []
  locationData: AddressData = {}
  private selectMapDialogController?: CustomDialogController
  @State imageUrl: ResourceStr = ''

  async getPileData(id: number) {
    let data = await getPiles(id)
    if (data) {
      this.Piles = data
    }
  }

  aboutToDisappear(): void {
    this.chargeItemController = null
    this.selectMapDialogController?.close()
    this.selectMapDialogController = undefined
  }

  chargeItemController: CustomDialogController | null = new CustomDialogController({
    builder: BottomPop({
      title: $r('app.string.selection_of_charging_pile'),
      bgColor: $r('app.color.color_f0f2f4'),
      titleBgColor: $r('app.color.color_f0f2f4'),
      contentBgColor: $r('app.color.color_f0f2f4'),
      keyboardAvoidMode: KeyboardAvoidMode.NONE,
      childBuilderParam: () => {
        this.chargeItemDialog()
      }
    }),
    customStyle: true,
    alignment: DialogAlignment.Bottom,
    maskColor: $r('app.color.color_80000000'),
  })

  @Builder
  chargeItemDialog() {
    ChargeItemDialogView({
      Piles: this.Piles,
      clickPileItemCallBack: (item: PileType) => {
        this.clickPileItem(item)
      }
    })
  }

  //单个桩点击处理
  clickPileItem(item: PileType) {
    if (item.state === 0) {
      this.chargeItemController?.close()
      router.pushNamedRoute({
        name: 'ChargingPage',
        params: {
          "pileId":item.id,
          "pileCode":item.code,
          "stationName":item.stationName
        }
      })
    } else if (item.state === 1) {
      toast('此桩正在充电中～')
    } else {
      toast('桩故障，无法充电哦～')
    }
  }

  @Builder
  chargeItem(item: ChargeType) {
    Column({ space: 12 }) {
      Text(item.name).fontWeight(FontWeight.Bold)
        .onClick(() => {
          this.chargeItemController?.open()
          if (item.id) {
            this.getPileData(item.id)
          }
        })
      Row({ space: 8 }) {
        Text(`￥ ${item.pricePerHour}`).fontColor($r('app.color.color_0aa671'))
        Text('/度').fontColor($r('app.color.color_adadb3'))
        Blank()
        Text(item.tenantName).fontSize($r('app.float.size_text_18')).fontColor($r('app.color.color_0aa671'))
        Blank()
        Image($r('app.media.fast')).width($r('app.float.size_20'))
        Text(`闲${item.freePileTotal}`).fontWeight(FontWeight.Bold)
        Text(`/${item.pileTotal}`)
      }
      .width(StyleConst.FULL_WIDTH)
      .onClick(() => {
        this.chargeItemController?.open()
        if (item.id) {
          this.getPileData(item.id)
        }
      })

      Row({ space: 8 }) {
        Row({ space: 8 }) {
          if (item.distance) {
            Text(`${item.distance.toFixed(2)}km`)
          }
          Image($r('app.media.location')).width($r('app.float.size_20'))
        }
        .onClick(() => {
          this.handleCheckInstallation(item)
        })

        Image($r('app.media.imgUrl')).width($r('app.float.size_20'))
          .onClick(() => {
            if (item.imageUrl) {
              this.imageUrl = item.imageUrl
              this.previewCustomDialog.open()
            } else {
              toast('没配图片哦，无法预览～')
            }
            console.log('lucy== img', item.id, item.imageUrl)
          })
      }

    }
    .backgroundColor($r('app.color.color_fff'))
    .padding({
      left: $r('app.float.size_10'),
      right: $r('app.float.size_10'),
      top: $r('app.float.size_15'),
      bottom: $r('app.float.size_15')
    })
    .borderRadius($r('app.float.size_20'))
    .alignItems(HorizontalAlign.Start)
  }

  //有数据页面
  @Builder
  haveDataCardView() {
    //列表
    List({ space: 10 }) {
      ForEach(this.chargeList, (item: ChargeType) => {
        ListItem() {
          this.chargeItem(item)
        }

      })
    }
    .scrollBar(BarState.Off)
    .edgeEffect(EdgeEffect.None)
    .width(StyleConst.FULL_WIDTH)
    .padding($r('app.float.size_10'))
    .layoutWeight(1)
  }

  //无数据页面
  @Builder
  noDataCardView() {
    Column() {
      Image($r('app.media.empty_charge1'))
        .width($r('app.float.size_100'))
        .height($r('app.float.size_100'))
      Text($r('app.string.no_data'))
        .fontSize($r('app.float.size_text_20'))
        .fontColor(Color.Black)
        .fontWeight(FontWeight.Bold)
        .margin({ top: $r('app.float.size_20'), bottom: $r('app.float.size_24') })
    }
    .margin({ top: $r('app.float.size_200') })
  }

  //点击导航图标事件
  handleCheckInstallation(item: ChargeType) {
    this.isAppInstalled(item)
    if (this.mapList.length === 0) {
      toast($r('app.string.visit_store_install_maps_first'))
    } else {
      this.permissionRequest()
    }
  }

  isAppInstalled(model: ChargeType) {
    const mapList = [ThirdApp.Amapuri, ThirdApp.Baidumap, ThirdApp.Maps].map((key: ThirdApp) => {
      let name: Length = ''
      let link: string = ""
      switch (key) {
        case ThirdApp.Amapuri:
          link =
            `amapuri://route/plan?dlat=${model.coordinate?.split(',')[0]}&dname=${model.name}&slat=${this.locationData.latitude}&dlon=${model.coordinate?.split(',')[1]}&slon=${this.locationData.longitude}&sname=我的位置&t=0&sourceApplication=com.huawei.hmos.maps.app`
          name = $r('app.string.visit_store_map')
          break
        case ThirdApp.Baidumap:
          name = $r('app.string.visit_store_baidu_map')
          link =
            `baidumap://map/direction?destination=latlng:${model.coordinate?.split(',')[0]},${model.coordinate?.split(',')[1]}|name:${model.name}&mode=driving`
          break
        case ThirdApp.Maps:
          name = $r('app.string.visit_store_petal_map')
          link =
            `https://www.petalmaps.com/routes/?daddr=${model.coordinate?.split(',')[0]},${model.coordinate?.split(',')[1]}(${model.name})&coordinateType=0&type=drive&utm_source=com.huawei.hmos.maps.app`
          break
        default:
          name = ""
      }
      return {
        name,
        link,
        isAppInstall: thirdAppManager.isAppInstall(key)
      } as mapListItem
    })

    console.log('lucy== map', JSON.stringify(mapList))
    this.mapList = mapList.filter(ele => ele.isAppInstall)
  }

  async permissionRequest() {
    try {
      let isGranted = await locationManager.checkAndApplyPermissions(getContext(this))
      if (!isGranted) {
        isGranted = await locationManager.requestPermissionAgain()
      }

      if (isGranted) {
        this.selectMapDialogController = new CustomDialogController({
          builder: MapDialogCall({
            title: '请选择地图',
            mapList: this.mapList,
            onClickCallBack: (item) => {
              this.handleStartNavigation(item)
            }
          }),
          customStyle: true,
          autoCancel: true,
          alignment: DialogAlignment.Bottom,
          keyboardAvoidMode: KeyboardAvoidMode.NONE
        })
        this.selectMapDialogController.open()
        const request: geoLocationManager.SingleLocationRequest = {
          locatingPriority: geoLocationManager.LocatingPriority.PRIORITY_LOCATING_SPEED,
          locatingTimeoutMs: 10000
        };
        await geoLocationManager.getCurrentLocation(request).then((location) => {
          // 获取到的经纬度
          this.locationData.latitude = location.latitude;
          this.locationData.longitude = location.longitude;
        }).catch((err: object) => {
          console.error("lucy== 获取位置失败:", JSON.stringify(err));
        });

      }
    } catch (e) {
      console.log('lucy== Address request location Permission fail' + e.message)
    }
  }

  handleStartNavigation(item: mapListItem) {
    this.selectMapDialogController?.close()
    this.jumpToMapApp(item.link)
  }

  jumpToMapApp(link: string) {
    try {
      const openLinkOptions: OpenLinkOptions = {
        appLinkingOnly: false
      };
      const context = getContext(this) as common.UIAbilityContext
      context.openLink(link, openLinkOptions)
        .then(() => {
          console.log('open link success.');
        }).catch((err: BusinessError) => {
        console.log(`open link failed. Code is ${err.code}, message is ${err.message}`);
      });
    } catch (paramError) {
      console.log(`Failed to start link. Code is ${paramError.code}, message is ${paramError.message}`);
    }
  }

  //点击图片预览事件
  private previewCustomDialog: CustomDialogController = new CustomDialogController({
    builder: ImagePreviewView({
      imageUrl: this.imageUrl,
    }),
    customStyle: true, //自由设置样式
    alignment: DialogAlignment.Center,
    autoCancel: false,
  })

  build() {
    if (this.chargeList.length > 0) {
      this.haveDataCardView()
    } else {
      this.noDataCardView()
    }
  }
}