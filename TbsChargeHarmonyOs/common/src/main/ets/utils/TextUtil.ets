import measure from '@ohos.measure'
export class TextUtil{
  static getTextWidth(textContent: string,
    fontSize: number | string | Resource = $r('app.float.size_text_14'), maxLines: number = 1,
    weight: FontWeight | number = FontWeight.Regular): number {

    if (!textContent || textContent.length === 0) {
      return 0
    }

    const width: number = measure.measureText({
      textContent: textContent,
      fontSize: fontSize,
      fontWeight: weight,
      maxLines: maxLines
    })
    return px2vp(width)
  }
}