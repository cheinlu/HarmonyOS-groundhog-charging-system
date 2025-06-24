import display from '@ohos.display';
import { GlobalContext } from '.';
import window from '@ohos.window';
import { StorageConst } from '../manager';


export class ScreenUtils {
  private static instance: ScreenUtils;

  setWindow(windowInfo: window.Window) {
    // 获取窗口属性
    const properties = windowInfo.getWindowProperties();
    this.setScreen(properties.windowRect.width, properties.windowRect.height)
    windowInfo.on("windowSizeChange", (size: Size) => {
      this.setScreen(size.width, size.height)
    })
  }

  private constructor() {
    this.defaultDisplay = display.getDefaultDisplaySync()
    this.defaultDisplay.getCutoutInfo((e, data) => {
      this.cutoutInfo = data
    })
    let callback: Callback<display.FoldDisplayMode> = (data: display.FoldDisplayMode) => {
      this.setFoldDisplayMode(data)
    };
    display.on("foldDisplayModeChange", callback)
  }

  private setScreen(width: number, height: number) {
    AppStorage.setOrCreate(StorageConst.SCREEN_WIDTH_VP, px2vp(width));
    AppStorage.setOrCreate(StorageConst.SCREEN_HEIGHT_VP, px2vp(height));
  }

  private setFoldDisplayMode(data: display.FoldDisplayMode) {
    AppStorage.setOrCreate(StorageConst.FOLD_DISPLAY_MODE, data);
  }

  public static getInstance(): ScreenUtils {
    if (!ScreenUtils.instance) {
      ScreenUtils.instance = new ScreenUtils();
    }
    return ScreenUtils.instance;
  }

  //屏幕数据
  private defaultDisplay: display.Display | null = null
  //挖孔屏状态栏区域
  private cutoutInfo: display.CutoutInfo | null = null

  getDisplay(): display.Display | null {
    return this.defaultDisplay
  }

  getScreenWidthPx(): number {
    return this.defaultDisplay?.width ?? -1
  }

  getScreenHeightPx(): number {
    return this.defaultDisplay?.height ?? -1
  }

  getScreenContentHeightPX(): number {
    return vp2px(this.getScreenHeightVp() - this.getStatusBarHeight())
  }

  getScreenContentHeightVp(): number {
    return px2vp(this.getScreenHeightPx())
  }

  getScreenHeightVp(): number {
    return px2vp(this.getScreenHeightPx())
  }

  getScreenWidthVp(): number {
    return px2vp(this.getScreenWidthPx())
  }

  /**
   * 获取刘海屏状态栏最大的区域
   * @returns
   */
  getStatusBarHeight(): number {
    const statusBar = GlobalContext.getContext().getObject(GlobalContext.STATUS_BAR_HEIGHT) as number;
    return px2vp(statusBar)
  }

  // 底部安全距离
  getSafeBottomHeight() {
    const navigation = GlobalContext.getContext().getObject(GlobalContext.NAVIGATION_HEIGHT) as number;
    return px2vp(navigation)
  }
}