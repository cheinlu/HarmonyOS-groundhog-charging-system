import { geoLocationManager } from "@kit.LocationKit";
import { PermissionConst } from "./PermissionConst";
import { permissionManager } from "./PermissionManager";

class LocationManager {

  async requestPermissionAgain() {
    return await permissionManager.requestPermissionAgain(PermissionConst.locationPermissions, getContext(this))
  }
  // 检查是否有定位权限，外部可直接调用，异常时返回false
  async checkLocationPermissions(): Promise<boolean> {
    try {
      return await permissionManager.checkPermissions(PermissionConst.locationPermissions) &&
      geoLocationManager.isLocationEnabled();
    } catch (e) {
      console.log(`lucy== hasLocationPermissions, ${JSON.stringify(e)}`)
      return false;
    }
  }

  // 校验并申请权限，如果申请权限返回false，则抛出异常，使用者需要使用catch接收。
  async checkAndApplyPermissions(context: Context): Promise<boolean> {
    let isGranted = await this.checkLocationPermissions();
    if (isGranted) {
      return true;
    }
    try {
      let userGranted = await permissionManager.reqPermissionsFromUser(PermissionConst.locationPermissions, context);
      if (userGranted) {
        return true;
      } else {
        return false;
      }
    } catch (e) {
      return Promise.reject(e);
    }
  }
}

export const locationManager = new LocationManager()