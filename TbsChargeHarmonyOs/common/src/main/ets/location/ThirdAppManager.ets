import { BusinessError } from "@kit.BasicServicesKit"
import { bundleManager } from "@kit.AbilityKit"

export enum ThirdApp {
  Alipay = 'alipays://',
  Amapuri = "amapuri://", // 高德地图的 scheme
  Baidumap = "baidumap://map/direction", // 百度地图的 scheme
  Maps = "maps://", // 华为地图的 scheme
  Https = "https://" // App link scheme
}

class ThirdAppManager{
  /** 第三方应用是否安装 */
  isAppInstall(thirdApp: ThirdApp): boolean {
    try {
      return bundleManager.canOpenLink(thirdApp)
    } catch (e) {
      let code = (e as BusinessError)?.code
      let msg = (e as BusinessError)?.message ?? ''
      if (code !== undefined) {
        if (code === 17700056) {
          msg += ` Please include "${thirdApp.replace('://',
            '')}" inside the "querySchemes" element of module.json5 in your app module.`
        }
        console.log(`isAppInstall get error${msg}`)
      } else {
        console.log(`isAppInstall get error${e}`)
      }
      return false
    }
  }
}

export const thirdAppManager = new ThirdAppManager()