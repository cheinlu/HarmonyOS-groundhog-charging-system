import { AbilityConstant, ConfigurationConstant, UIAbility, Want } from '@kit.AbilityKit';
import { hilog } from '@kit.PerformanceAnalysisKit';
import { window } from '@kit.ArkUI';
import {
  commonStorage,
  GlobalContext,
  initDataPreference,
  RoutePageConst,
  ScreenUtils,
  StorageConst,
  WindowModel
} from '@tbs/common';
import { BusinessError } from '@kit.BasicServicesKit';

const DOMAIN = 0x0000;

export default class EntryAbility extends UIAbility {
  onCreate(want: Want, launchParam: AbilityConstant.LaunchParam): void {
    this.context.getApplicationContext().setColorMode(ConfigurationConstant.ColorMode.COLOR_MODE_NOT_SET);
    hilog.info(DOMAIN, 'testTag', '%{public}s', 'Ability onCreate');
    // 绑定上下文信息
    GlobalContext.getContext().setObject(GlobalContext.CONTEXT, this.context);

    initDataPreference()
  }

  onDestroy(): void {
    hilog.info(DOMAIN, 'testTag', '%{public}s', 'Ability onDestroy');
  }

  onWindowStageCreate(windowStage: window.WindowStage): void {
    // Main window is created, set main page for this ability
    hilog.info(DOMAIN, 'testTag', '%{public}s', 'Ability onWindowStageCreate');
    this.setMainWindowImmersive(windowStage)
    //有登录信息->跳转至App充电页，无登录信息->登录页
    let destPageUrl = RoutePageConst.AppHomePage
    if (!commonStorage.getStringValueSync(StorageConst.TOKEN, "")) {
      destPageUrl = RoutePageConst.LoginPage
    }
    windowStage.loadContent(destPageUrl, (err) => {
      windowStage.getMainWindow((err: BusinessError, data)=>{
        let windowInfo: window.Window = data;
        if (!windowInfo.getUIContext()) {
          return
        }
        ScreenUtils.getInstance().setWindow(windowInfo)
      })
      if (err.code) {
        hilog.error(DOMAIN, 'testTag', 'Failed to load the content. Cause: %{public}s', JSON.stringify(err));
        return;
      }
      hilog.info(DOMAIN, 'testTag', 'Succeeded in loading the content.');
    });
  }

  onWindowStageDestroy(): void {
    // Main window is destroyed, release UI related resources
    hilog.info(DOMAIN, 'testTag', '%{public}s', 'Ability onWindowStageDestroy');
  }

  onForeground(): void {
    // Ability has brought to foreground
    hilog.info(DOMAIN, 'testTag', '%{public}s', 'Ability onForeground');
  }

  onBackground(): void {
    // Ability has back to background
    hilog.info(DOMAIN, 'testTag', '%{public}s', 'Ability onBackground');
  }

  /**
   * 设置沉浸式状态栏
   * 默认情况：状态栏、导航栏透明，状态栏字体颜色黑色
   * @param windowStage
   */
  setMainWindowImmersive(windowStage: window.WindowStage) {
    WindowModel.getInstance().setWindowStage(windowStage)
    WindowModel.getInstance().setMainWindowImmersive(true)
    let type = window.AvoidAreaType.TYPE_SYSTEM;
    try {
      let mainWindow = windowStage.getMainWindowSync()
      let avoidArea = mainWindow.getWindowAvoidArea(type)
      GlobalContext.getContext().setObject(GlobalContext.STATUS_BAR_HEIGHT, avoidArea.topRect.height); //状态栏高度

      //获取导航条避让遮挡的区域
      let t = window.AvoidAreaType.TYPE_NAVIGATION_INDICATOR
      let avoidAreaIndicator = mainWindow.getWindowAvoidArea(t);
      let bottomRectHeight = avoidAreaIndicator.bottomRect.height; // 获取到导航条区域的高度
      GlobalContext.getContext().setObject(GlobalContext.NAVIGATION_HEIGHT, bottomRectHeight);
    } catch (e) {
      console.error('lucy++ EntryAbility getMainWindowSync error:' + e.message)
    }
  }
}