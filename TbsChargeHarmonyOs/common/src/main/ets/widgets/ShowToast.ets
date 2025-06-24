import { promptAction } from '@kit.ArkUI'

//轻提示工具类

export const toast = (hintContent: string | ResourceStr, alignment: Alignment = Alignment.Center): void => {
  promptAction.showToast({ message: hintContent, alignment: alignment })
}
