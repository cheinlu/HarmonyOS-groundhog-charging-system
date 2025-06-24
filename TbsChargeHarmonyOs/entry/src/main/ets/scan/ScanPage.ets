import { permissionManager, PermissionConst, StyleConst, RoutePageConst, ScreenUtils, toast } from '@tbs/common';
import { CustomScanView } from './view/CustomScanView';
import { customScan, scanBarcode, scanCore } from '@kit.ScanKit';
import { BusinessError } from '@kit.BasicServicesKit';
import { common } from '@kit.AbilityKit';
import { router } from '@kit.ArkUI';
import { ChargingQrData } from './ScanPageModel';
import '@tbs/charge/src/main/ets/page/ChargingPage'
//扫码--已实现单个二维码处理逻辑
@Entry
@Component
struct ScanPage {
  @State surfaceId: string = ''
  @State scanResults: scanBarcode.ScanResult[] = []
  @State returnIndex: number = 0
  @State isFlashlightOn: boolean = false // 手电筒开启灯光状态，默认关闭
  @State isShowFlashlight: boolean = false // 是否显示手电筒（环境暗时需自动显示），默认不显示
  @State userGrant: boolean = false // 用户权限记录
  isScanned: boolean = false // 是否已扫出结果
  // 设置预览流宽度，默认单位：vp
  cameraWidth: number = ScreenUtils.getInstance().getScreenWidthVp()
  // 设置预览流高度，默认单位：vp
  cameraHeight: number = ScreenUtils.getInstance().getScreenHeightVp()
  xComponentController: XComponentController = new XComponentController();
  options: scanBarcode.ScanOptions = {
    scanTypes: [scanCore.ScanType.ALL], //设置扫码类型，默认扫码ALL（全部码类型
    enableMultiMode: true,
    enableAlbum: true
  };
  viewControl: customScan.ViewControl = {
    width: this.cameraWidth,
    height: this.cameraHeight,
    surfaceId: this.surfaceId
  }

  async onPageShow() {
    let param = router.getParams() as Record<string, Object>;
    if (param && param.returnIndex) {
      this.returnIndex = param.returnIndex as number
    }
    this.isScanned = false
    let context = getContext() as common.UIAbilityContext
    await permissionManager.requestPermission([PermissionConst.cameraPermission], context, () => {
      this.initCustomScan()
    }, () => {
      permissionManager.requestPermissionAgain([PermissionConst.cameraPermission], context).then(() => {
      }).catch((err: BusinessError) => {
        console.log('lucy== 失败了吗', JSON.stringify(err))
      })
    })
  }

  //释放资源
  onPageHide(): void {
    this.isShowFlashlight = false
    this.releaseScan()
  }

  handleBack() {
    router.back({
      url: RoutePageConst.AppHomePage,
      params: { restoreIndex: this.returnIndex }
    })
  }

  // 初始化自定义扫码器
  initCustomScan() {
    this.userGrant = true
    try {
      customScan.init(this.options)
    } catch (error) {
      console.error('lucy++ 扫码初始化失败', error.code)
    }
    this.startCustomScan()
  }

  //开启自定义扫码
  startCustomScan() {
    this.viewControl = { width: this.cameraWidth, height: this.cameraHeight, surfaceId: this.surfaceId };
    customScan.start(this.viewControl).then((scanResult: Array<scanBarcode.ScanResult>) => {
      this.handlerScanResult(scanResult)
      this.onLightingFlash()
    }).catch((error: BusinessError) => {
      console.error("lucy== 自定义扫码 开始失败:", error.code, error.message, error.name);
    });
  }

  //扫码结果处理
  handlerScanResult(result: scanBarcode.ScanResult[]) {
    if (this.isScanned) {
      return
    }
    if (result.length > 0) {
      this.stopScan()
      this.isScanned = true
    }
    this.scanResults = result
    if (result.length === 1) {
      this.handlerQRCodeLogic(result[0])
    }
  }

  //停止扫描(关闭相机闪光灯)
  stopScan() {
    if (this.isFlashlightOn) {
      try {
        //关闭相机闪光灯
        customScan.closeFlashLight()
        this.offLightingFlash()
        this.isShowFlashlight = false
        this.isFlashlightOn = false
      } catch (e) {
        console.error('lucy== 关闭手电筒失败', e.code)
      }
      try {
        //暂停扫码相机流
        customScan.stop().then(() => {
          console.info('lucy== 扫码已停止')
        }).catch((error: BusinessError) => {
          console.error('lucy== 扫码失败', error.code)
        })
      } catch (e) {
        console.error('扫码 停止失败:', e.code)
      }
    }
  }

  //释放扫码资源
  async releaseScan() {
    await this.stopScan()
    try {
      await customScan.release().then(() => {
        console.info('lucy== 自定义扫码 已释放')
      }).catch((e: BusinessError) => {
        console.error('lucy== 自定义扫码 释放失败:', e.code)
      })
    } catch (e) {
      console.error('lucy== 自定义扫码 释放失败:', e.code)
    }
  }

  handlerQRCodeLogic(result: scanBarcode.ScanResult) {
    let scanResult = result.originalValue
    let parsedResult = JSON.parse(scanResult) as ChargingQrData
    toast($r('app.string.scan_success'))
    console.log('lucy== scanResult', scanResult, JSON.stringify(scanResult))
    if (scanResult.startsWith('http://') || scanResult.startsWith('https://')) {
      router.replaceUrl({ url: "webview/WebViewPage", params: { url: scanResult } })
    } else if (parsedResult.type === 'charging') {
      router.replaceNamedRoute({
        name: 'ChargingPage',
        params: {
          pileId: parsedResult.pileId,
          pileCode: parsedResult.pileCode,
          stationName: parsedResult.stationName
        }
      });
    }
  }

  //订阅闪光灯状态监听事件
  onLightingFlash() {
    try {
      customScan.on('lightingFlash', (e, isLightingFlash) => {
        if (e) {
          console.error(`lucy== 监听弱光环境失败, code：${e.code}, message：${e.message}`)
          return
        }
        if (!this.isFlashlightOn) {
          console.log('lucy== isLightingFlash', isLightingFlash)
          this.isShowFlashlight = isLightingFlash
        }
      });
    } catch (e) {
      console.error(`lucy== 监听弱光环境失败, code：${e.code}, message：${e.message}`)
    }
  }

  //注销手电筒监听状态
  offLightingFlash() {
    try {
      customScan.off('lightingFlash')
    } catch (e) {
      console.error('lucy== 自定义扫码 手电筒释放失败:', e.code);
    }
  }

  build() {
    Stack() {
      if (this.userGrant) {
        Column() {
          XComponent({ type: XComponentType.SURFACE, controller: this.xComponentController })
            .onLoad(() => {
              this.surfaceId = this.xComponentController.getXComponentSurfaceId()
              // 开启扫码
              this.startCustomScan()
            })
        }.width(StyleConst.FULL_WIDTH).height(StyleConst.FULL_HEIGHT)
      }

      CustomScanView({
        isFlashlightOn: this.isFlashlightOn,
        isShowFlashlight: this.isShowFlashlight,
        isScanned: this.isScanned,
        onBack: () => {
          console.log('lucy== 返回')
          this.handleBack()
        },
        handlerLamplightCompletion: (isOpen) => {
          if (isOpen) {
            customScan.openFlashLight()
          } else {
            customScan.closeFlashLight()
            this.isShowFlashlight = false
          }
        },
      })
        .width(StyleConst.FULL_WIDTH)
        .height(StyleConst.FULL_HEIGHT)
        .visibility(this.scanResults.length > 1 ? Visibility.Hidden : Visibility.Visible)
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
  }
}