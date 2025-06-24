import { StyleConst } from "@tbs/common"
import { mapListItem } from "../model/ChargeModel"

interface TextStyle {
  height: Length
}

@Extend(Text)
//文本样式
function textStyles(item: TextStyle) {
  .fontColor($r('app.color.color_0165b8'))
  .textAlign(TextAlign.Center)
  .fontSize($r('app.float.size_18'))
  .height(item.height)
}

@CustomDialog
@Component
export struct MapDialogCall {
  mapList: mapListItem[] = []
  controller: CustomDialogController
  title: ResourceStr = '';
  onClickCallBack?: (ele: mapListItem) => void

  build() {
    Column({ space: 12 }) {
      Column() {
        Text(this.title)
          .fontColor(Color.Black)
          .fontSize($r('app.float.size_13'))
          .textAlign(TextAlign.Center)
          .width(StyleConst.FULL_WIDTH)
          .height($r('app.float.size_42'))

        ForEach(
          this.mapList,
          (ele: mapListItem) => {
            Text(ele.name)
              .textStyles({ height: $r('app.float.size_50')})
              .onClick(() => {
                this.onClickCallBack && this.onClickCallBack(ele)
              })
          }
        )
      }
      .borderRadius($r('app.float.size_8'))
      .justifyContent(FlexAlign.Center)
      .width(StyleConst.FULL_WIDTH)
      .backgroundColor(Color.White)
      .padding({top: $r('app.float.size_12'),
        bottom: 0})

      Row() {
        Text('取消')
          .textStyles({ height: $r('app.float.size_50')})
          .fontColor($r('app.color.color_0165b8'))
          .width(StyleConst.FULL_WIDTH)
      }
      .borderRadius($r('app.float.size_8'))
      .backgroundColor(Color.White)
      .height($r('app.float.size_50'))
      .margin({ bottom: $r('app.float.size_16') })
      .onClick(()=>{
        this.controller.close()
      })
    }
    .padding({
      left: $r('app.float.size_12'),
      right: $r('app.float.size_12'),
      bottom: $r('app.float.size_31'),
    })
    .expandSafeArea([SafeAreaType.SYSTEM], [SafeAreaEdge.BOTTOM])
    .width(StyleConst.FULL_WIDTH)
    .backgroundColor(Color.Transparent)
  }
}
