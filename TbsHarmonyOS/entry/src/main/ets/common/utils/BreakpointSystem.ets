import mediaQuery from '@ohos.mediaquery';
import BreakpointConstants from '../constants/BreakpointConstants'

export default class BreakpointSystem{
  //定义3个屏幕监听器
  private smListener: mediaQuery.MediaQueryListener = mediaQuery.matchMediaSync(BreakpointConstants.RANGE_SM)
  private mdListener: mediaQuery.MediaQueryListener = mediaQuery.matchMediaSync(BreakpointConstants.RANGE_MD)
  private lgListener: mediaQuery.MediaQueryListener = mediaQuery.matchMediaSync(BreakpointConstants.RANGE_LG)

  //定义回调函数
  smListenerCallback(result: mediaQuery.MediaQueryResult){
    if(result.matches){
      this.updateCurrentBreakpoint(BreakpointConstants.BREAKPOINT_SM)
    }
  }

  mdListenerCallback(result: mediaQuery.MediaQueryResult){
    if(result.matches){
      this.updateCurrentBreakpoint(BreakpointConstants.BREAKPOINT_MD)
    }
  }

  lgListenerCallback(result: mediaQuery.MediaQueryResult){
    if(result.matches){
      this.updateCurrentBreakpoint(BreakpointConstants.BREAKPOINT_LG)
    }
  }
  //定义存储函数，以上3个函数中调用
  updateCurrentBreakpoint(breakpoint: string){
    AppStorage.SetOrCreate(BreakpointConstants.CURRENT_BREAKPOINT, breakpoint)
  }
  //监听
  register(){
    this.smListener.on('change', this.smListenerCallback.bind(this))
    this.mdListener.on('change', this.mdListenerCallback.bind(this))
    this.lgListener.on('change', this.lgListenerCallback.bind(this))
  }
  //取消监听
  unregister(){
    this.smListener.off('change', this.smListenerCallback.bind(this))
    this.mdListener.off('change', this.mdListenerCallback.bind(this))
    this.lgListener.off('change', this.lgListenerCallback.bind(this))
  }
}

/*
 * 1、定义3个屏幕监听器
 * 2、分别监听不同的屏幕获取宽度范围
 * （先定义3个屏幕的函数，定义函数时使用AppStorage进行存储）+（再进行绑定，利用on监听，off取消监听）
 * */