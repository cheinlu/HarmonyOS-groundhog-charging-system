import { componentUtils, inspector } from "@kit.ArkUI"

export class ComponentUtils {
  /**
   * 同步方式获取组件渲染后的实际高度
   * @param componentId 组件ID
   * @param callback 注册的回调函数
   */
  static getComponentHeightSync(componentId: string, callback: (width: number) => void): void {
    // 监听组件
    const listener = inspector.createComponentObserver(componentId)
    // 注册回调
    listener.on('layout', () => {
      const componentInfo = componentUtils.getRectangleById(componentId)
      const height = px2vp(componentInfo.size.height)
      callback?.(height)
      // 取消注册回调
      listener.off('layout')
    })
  }
}