import { Permissions } from '@kit.AbilityKit';

export class PermissionConst {
  //获取位置权限
  static readonly locationPermissions: Array<Permissions> = [
    'ohos.permission.APPROXIMATELY_LOCATION',
    'ohos.permission.LOCATION',
  ];
  //获取相机权限
  static readonly cameraPermission: Permissions = 'ohos.permission.CAMERA';
}