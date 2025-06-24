import { StyleConst, TbsTitleBar } from "@tbs/common";
import { router } from "@kit.ArkUI";
import { webview } from "@kit.ArkWeb";

@Entry
@Component
export struct WebViewPage {
  @State url: string = '';
  @State title: string = ''
  controller: webview.WebviewController = new webview.WebviewController();

  aboutToAppear(): void {
    const params = router.getParams() as Record<string, Object>; // 获取传递过来的参数对象
    if (params && params.url) {
      this.url = params.url as string;
    }
  }

  @Builder
  WebViewBuilder() {
    Web({
      src: this.url,
      controller: this.controller,
      renderMode: RenderMode.ASYNC_RENDER
    })
      .onTitleReceive(event => {
        if (event) {
          this.title = event.title
        }
      })
  }

  build() {
    Column() {
      TbsTitleBar({
        title: this.title,
        showBackIcon: true
      })
      this.WebViewBuilder()
    }.width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
  }
}