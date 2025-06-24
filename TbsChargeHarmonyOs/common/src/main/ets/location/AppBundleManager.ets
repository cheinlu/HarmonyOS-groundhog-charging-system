import { bundleManager } from "@kit.AbilityKit";

class AppBundleManager{
  private bundleInfo?: bundleManager.BundleInfo

  getBundleInfo(): bundleManager.BundleInfo {
    if (this.bundleInfo == null || this.bundleInfo == undefined) {
      return bundleManager.getBundleInfoForSelfSync(bundleManager.BundleFlag.GET_BUNDLE_INFO_WITH_APPLICATION)
    }
    return this.bundleInfo
  }

  getTokenId(): number {
    let bundleInfo = this.getBundleInfo();
    return bundleInfo.appInfo.accessTokenId;
  }
}

export const appBundleManager = new AppBundleManager()