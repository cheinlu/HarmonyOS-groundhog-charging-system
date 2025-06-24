import { window } from "@kit.ArkUI";
import { BusinessError } from "@kit.BasicServicesKit";

export class WindowModel {
  private constructor() {
  }

  /**
   * 默认样式
   */
  private readonly DEFAULT: window.SystemBarProperties = {
    statusBarColor: '#00000000',
    navigationBarColor: '#00000000',
    // 以下两个属性从API Version 8开始支持
    statusBarContentColor: '#000000',
    navigationBarContentColor: '#00000000'
  };
  /**
   * WindowModel 单例
   */
  private static instance?: WindowModel;

  /**
   * 获取WindowModel单例实例
   * @returns {WindowModel} WindowModel
   */
  static getInstance(): WindowModel {
    if (!WindowModel.instance) {
      WindowModel.instance = new WindowModel();
    }

    return WindowModel.instance;
  }

  /**
   * 缓存的当前WindowStage实例
   */
  private windowStage?: window.WindowStage;

  /**
   * 缓存windowStage
   * @param windowStage 当前WindowStage实例
   * @returns {void}
   */
  setWindowStage(windowStage: window.WindowStage): void {
    this.windowStage = windowStage;
  }

  getWindowStage(): window.WindowStage | undefined {
    return this.windowStage
  }

  /**
   * 当前主窗口是否开启沉浸模式
   * @param enable 是否开启
   * @returns {void}
   */
  setMainWindowImmersive(enable: boolean, properties?: window.SystemBarProperties): void {
    if (this.windowStage === undefined) {
      console.error('windowStage is undefined.');
      return;
    }
    // 1.获取应用主窗口。
    let windowClass: window.Window | null = null;
    this.windowStage.getMainWindow((err: BusinessError, data) => {
      let errCode: number = err.code;
      if (errCode) {
        console.error('Failed to obtain the main window. Cause: ' + JSON.stringify(err));
        return;
      }
      windowClass = data;
      console.info('Succeeded in obtaining the main window data.');

      // 2.实现沉浸式效果。方式二：设置窗口为全屏布局，配合设置导航栏、状态栏的透明度、背景/文字颜色及高亮图标等属性，与主窗口显示保持协调一致。
      windowClass.setWindowLayoutFullScreen(enable, (err: BusinessError) => {
        let errCode: number = err.code;
        if (errCode) {
          console.error('Failed to set the window layout to full-screen mode. Cause:' + JSON.stringify(err));
          return;
        }
        console.info('Succeeded in setting the window layout to full-screen mode.');
      });
      const p = properties ? properties : this.DEFAULT
      windowClass.setWindowSystemBarProperties(p, (err: BusinessError) => {
        let errCode: number = err.code;
        if (errCode) {
          console.error('Failed to set the system bar properties. Cause: ' + JSON.stringify(err));
          return;
        }
        console.info('Succeeded in setting the system bar properties.');
      });
    })
  }

  /**
   * 软键盘是否显示/隐藏
   * @param callback
   */
  getKeyBoardVisible(callback: (visible: boolean) => void) {
    window.getLastWindow(getContext(this)).then(currentWindow => {
      // 监听软键盘的隐藏和显示
      currentWindow.on('avoidAreaChange', data => {
        if (data.type == window.AvoidAreaType.TYPE_KEYBOARD) {
          const keyBoardHeight = px2vp(data.area.bottomRect.height);
          if (keyBoardHeight > 0) {
            callback(true)
          } else {
            callback(false)
          }
        }
      })
    })
  }
}