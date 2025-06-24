import { abilityAccessCtrl, common, PermissionRequestResult, Permissions, Want } from "@kit.AbilityKit";
import { BusinessError } from "@kit.BasicServicesKit";
import { appBundleManager } from "./AppBundleManager";

class PermissionManager {
  atManager: abilityAccessCtrl.AtManager = abilityAccessCtrl.createAtManager();

  // 每次访问受目标权限保护的接口之前，都需要使用requestPermissionsFromUser()接口请求相应的权限。
  // 权限组：当应用请求权限时，同一个权限组的权限将会在一个弹窗内一起请求用户授权。用户点击同意，即同意了整个权限组。
  // 如定位权限组：['ohos.permission.APPROXIMATELY_LOCATION', 'ohos.permission.LOCATION']
  async checkPermissions(permissions: Array<Permissions>): Promise<boolean> {
    console.log(`check permission ${permissions}`);

    try {
      let grantStatusArray: abilityAccessCtrl.GrantStatus[] = [];
      for (let i = 0; i < permissions.length; i++) {
        grantStatusArray[i] = await this.checkPermissionInner(permissions[i]);
      }
      console.log(`checkPermissions: ${permissions}, grantStatus ${grantStatusArray} `);

      // 全同意才算同意
      const denied = grantStatusArray.indexOf(abilityAccessCtrl.GrantStatus.PERMISSION_DENIED) > -1;
      return !denied;
    } catch (e) {
      return Promise.reject(e);
    }
  }

  private async checkPermissionInner(permission: Permissions): Promise<abilityAccessCtrl.GrantStatus> {
    let grantStatus: abilityAccessCtrl.GrantStatus = abilityAccessCtrl.GrantStatus.PERMISSION_DENIED;

    // 校验应用是否被授予权限
    try {
      grantStatus = await this.atManager.checkAccessToken(appBundleManager.getTokenId(), permission);
    } catch (error) {
      const err: BusinessError = error as BusinessError;
      console.log(`Failed to check access token. Code is ${err.code}, message is ${err.message}`);
      return Promise.reject(`checkPermissionInner: checkAccessToken error: ${JSON.stringify(error)}`,);
    }

    return grantStatus;
  }

  async reqPermissionsFromUser(permissions: Array<Permissions>, context: Context): Promise<boolean> {
    let atManager: abilityAccessCtrl.AtManager = abilityAccessCtrl.createAtManager();
    // requestPermissionsFromUser会判断权限的授权状态来决定是否唤起弹窗
    try {
      let result: PermissionRequestResult = await atManager.requestPermissionsFromUser(context, permissions);
      let grantStatus: Array<number> = result.authResults;
      for (let i = 0; i < grantStatus.length; i++) {
        if (grantStatus[i] === 0) {
          // 用户授权，可以继续访问目标操作
          console.log(`reqPermissionsFromUser: ${permissions[i]} is granted`);
        } else {
          console.log(`reqPermissionsFromUser: ${permissions[i]} is denied`);
          return false;
        }
      }
      return true;
    } catch (e) {
      console.log(`reqPermissionsFromUser error: Code is ${e.code}, message is ${e.message}`);
      return Promise.reject(e);
    }
  }

  //被拒绝后，再次点击全权限申请使用该方法
  async requestPermissionAgain(permissions: Array<Permissions>, context: Context): Promise<boolean> {
    try {
      const result = await this.atManager.requestPermissionOnSetting(context, permissions)
      return result.every((status) => {
        return status === abilityAccessCtrl.GrantStatus.PERMISSION_GRANTED
      })
    } catch (e) {
      console.log(`reqPermissionsFromUser error: Code is ${e.code}, message is ${e.message}`);
      return Promise.reject(e);
    }
  }

  async requestPermission(permissions: Array<Permissions>, context: common.UIAbilityContext, success: () => void,
    failure: () => void) {
    let grantStatus = await this.checkPermissions(permissions);
    if (grantStatus) {
      success()
    } else {
      let granted =
        await this.reqPermissionsFromUser(permissions, context)
      if (granted) {
        success()
      } else {
        failure()
      }
    }
  }

  /*
 * 打开系统设置页面
 * */
  openSettings(context: common.UIAbilityContext) {
    this.openSystemSettings(context, "application_info_entry")
  }

  /*
  * 打开系统设置页面
  * uri：根据‘设置’应用配置的界面信息，选择不同的uri
  * */
  openSystemSettings(context: common.UIAbilityContext, uri?: string) {
    let wantInfo: Want = {}
    wantInfo = {
      bundleName: 'com.huawei.hmos.settings',
      abilityName: 'com.huawei.hmos.settings.MainAbility',
      uri: uri,
      parameters: { pushParams: context.abilityInfo.bundleName }
    }

    context.startAbility(wantInfo).then(() => {
      // 从设置页返回
      console.info(`after opening system settings`);
    }).catch((err: BusinessError) => {
      console.error(`Failed to Start settings ability. Code is ${err.code}, message is ${err.message}`);
    })
  }
}

let permissionManager = new PermissionManager();

export { permissionManager };