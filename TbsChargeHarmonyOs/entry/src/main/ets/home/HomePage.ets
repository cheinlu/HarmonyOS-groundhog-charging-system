import { RoutePageConst, StyleConst } from '@tbs/common/Index'
import { ChargePage } from '@tbs/charge'
import { MinePage } from '@tbs/mine'
import { buttonInfo, ButtonInfoModel } from './viewModel/homeData'
import { router } from '@kit.ArkUI'

@Entry
@Component
struct HomePage {
  controller: TabsController = new TabsController()
  @State currentIndex: number = 0
  // 用于标记是否正在跳转
  @State isNavigating: boolean = false

  build() {
    Flex({ direction: FlexDirection.Column, justifyContent: FlexAlign.End, alignItems: ItemAlign.Center }) {
      Tabs({ barPosition: BarPosition.End, index: 0, controller: this.controller }) {
        ForEach(buttonInfo, (item: ButtonInfoModel, index?: number) => {
          TabContent() {
            if (index == 0) {
              ChargePage()
            } else if (index == 2) {
              MinePage()
            }
          }.tabBar(this.BottomNavigation(item))
        })
      }
      .width(StyleConst.FULL_WIDTH)
      .height(StyleConst.FULL_HEIGHT)
      .vertical(false)
      .scrollable(false)
      .animationDuration(0)
      .divider({ strokeWidth: 1, color: $r('app.color.color_f0f2f4') })
      .barHeight($r('app.float.size_50'))
      .margin({ bottom: $r('app.float.size_20') })
      .onChange((index) => {
        // 如果点击的是第二个Tab（index=1）则跳转
        if (index === 1) {
          this.isNavigating = true
          const prevIndex = this.currentIndex
          router.pushUrl({
            url: RoutePageConst.ScanPage, // 替换为您的扫码页面路径
            params: { returnIndex: prevIndex }
          }, (err) => {
            if (err) {
              console.error('跳转失败:', err)
            }
            // 跳转完成后重置状态
            this.isNavigating = false
            this.currentIndex = this.currentIndex // 保持当前选中状态不变
          })
        } else {
          // 正常切换Tab
          this.currentIndex = index
        }
      })
    }
  }

  @Builder
  BottomNavigation(item: ButtonInfoModel) {
    Column({ space: 6 }) {
      Image(this.currentIndex === item.index ? item.selectImg : item.img)
        .objectFit(ImageFit.Contain)
        .width(item.index == 1 ? $r('app.float.size_28') : $r('app.float.size_24'))
        .height(item.index == 1 ? $r('app.float.size_28') : $r('app.float.size_24'))
      Text(item.title)
        .fontColor(this.currentIndex === item.index ? $r('app.color.color_0165b8') : Color.Black)
        .opacity(this.currentIndex === item.index ? 1 : 0.6)
        .fontWeight(FontWeight.Medium)
        .textAlign(TextAlign.Center)
        .fontSize($r('app.float.size_text_12'))
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .alignItems(HorizontalAlign.Center)
    .justifyContent(FlexAlign.Center)
    .onClick(() => {
      if (item.index === 1) {
        // 保存当前选中状态
        const prevIndex = this.currentIndex

        // 跳转到扫码页面
        router.pushUrl({
          url: RoutePageConst.ScanPage,
          params: { returnIndex: prevIndex } // 传递返回后应显示的Tab索引
        }, (err) => {
          if (err) {
            console.error('跳转失败:', err)
            // 如果跳转失败，恢复之前的状态
            this.currentIndex = prevIndex
          }
        })
      } else {
        // 正常切换Tab
        this.currentIndex = item.index
        this.controller.changeIndex(item.index)
      }
    })
  }
}
