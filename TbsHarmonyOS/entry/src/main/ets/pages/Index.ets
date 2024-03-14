import BreakpointType from '../common/bean/BreanpointType'
import BreakpointConstants from '../common/constants/BreakpointConstants'
import BreakpointSystem from '../common/utils/BreakpointSystem'
import ChargeIndex from '../view/charge/ChargeIndex'
import ScanIndex from '../view/scan/ScanIndex'
import UserIndex from '../view/user/UserIndex'
@Entry
@Component
struct Index {
  @State currentIndex:number = 0
  private breakpointSystem:BreakpointSystem = new BreakpointSystem()
  @StorageProp('currentBreakpoint') currentBreakpoint:string = BreakpointConstants.BREAKPOINT_SM

  aboutToAppear(){
    this.breakpointSystem.register()
  }

  aboutToDisappear(){
    this.breakpointSystem.unregister()
  }

  @Builder TabBarBuilder(title: ResourceStr, image: ResourceStr, index: number){
    Column({space:8}){
      Image(image)
        .width(22)
        .fillColor(this.selectColor(index))
      Text(title)
        .fontSize(14)
        .fontColor(this.selectColor(index))
    }
  }

  selectColor(index: number){
    return this.currentIndex === index ? $r('app.color.primary_color') : $r('app.color.gray')
  }

  build() {
    Tabs({ barPosition: BreakpointConstants.BAR_POSITION.getValue(this.currentBreakpoint) }){
      TabContent(){
        ChargeIndex()
      }
      .tabBar(this.TabBarBuilder($r('app.string.tab_record'), $r('app.media.ic_calendar'), 0))
      TabContent(){
        ScanIndex()
      }
      .tabBar(this.TabBarBuilder($r('app.string.tab_discover'), $r('app.media.discover'), 1))
      TabContent(){
        UserIndex()
      }
      .tabBar(this.TabBarBuilder($r('app.string.tab_user'), $r('app.media.ic_user_portrait'), 2))
    }
    .width('100%')
    .height('100%')
    .onChange(index => this.currentIndex = index)
    .vertical(new BreakpointType({
      sm: false,
      md: true,
      lg: true
    }).getValue(this.currentBreakpoint))
  }
}
