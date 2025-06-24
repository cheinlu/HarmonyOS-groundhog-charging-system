import { RoutePageConst, StyleConst, toast } from '@tbs/common'
import { router } from '@kit.ArkUI'
import { getRegister } from './viewModel/RegisterViewModel'

@Entry
@Component
struct RegisterPage {
  @State adminUsername: string = '' //用户名
  @State adminPassword: string = '' //密码
  @State validCode: string = '' //验证码
  @State name: string = '' //企业名称
  @State domain: string = '' //网站

  //注册
  async getRegisterInfo() {
    if (!this.adminUsername) {
      toast($r('app.string.please_enter_the_username'))
      return
    }
    if (!this.adminPassword) {
      toast($r('app.string.please_enter_the_password'))
      return
    }
    if (!this.validCode) {
      toast($r('app.string.please_enter_the_verification_code'))
      return
    }
    if (!this.name) {
      toast($r('app.string.please_enter_the_company_name'))
      return
    }
    if (!this.domain) {
      toast($r('app.string.please_enter_the_company_website'))
      return
    }

    let data = await getRegister(this.name, this.domain, this.adminUsername, this.adminPassword, this.validCode)
    if (data != null && Number(data) == 0) {
      toast($r('app.string.registered_successfully'))
      router.pushUrl({
        url: RoutePageConst.LoginPage
      })
    }
  }

  build() {
    Column({ space: 15 }) {
      Image($r('app.media.Groundhog'))
        .width($r('app.float.size_100'))
        .borderRadius($r('app.float.size_50'))
        .margin({ top: $r('app.float.size_120') })
      Text($r('app.string.welcome_register_groundhog'))
        .fontColor(Color.Gray)
        .fontSize($r('app.float.size_text_20'))
        .textAlign(TextAlign.Center)
      TextInput({
        placeholder: $r('app.string.please_enter_the_username')
      })
        .width(StyleConst.PERCENT_90)
        .onChange((val: string) => {
          this.adminUsername = val
        })
      Row() {
        TextInput({ placeholder: $r('app.string.verification_code') })
          .width(StyleConst.PERCENT_50)
          .onChange((val: string) => this.validCode = val)
        Button($r('app.string.get_code'))
          .width(StyleConst.PERCENT_40)
          .backgroundColor($r('app.color.color_0aa671'))
      }.justifyContent(FlexAlign.SpaceBetween)
      .width(StyleConst.PERCENT_90)

      TextInput({
        placeholder: $r('app.string.password_length')
      })
        .width(StyleConst.PERCENT_90)
        .onChange((val: string) => this.adminPassword = val).type(InputType.Password)
      TextInput({
        placeholder: $r('app.string.company_name')
      })
        .width(StyleConst.PERCENT_90)
        .onChange((val: string) => {
          this.name = val
        })
      TextInput({
        placeholder: $r('app.string.company_website')
      })
        .width(StyleConst.PERCENT_90)
        .onChange((val: string) => {
          this.domain = val
        })

      Button($r('app.string.register_immediately'))
        .width(StyleConst.PERCENT_90)
        .backgroundColor($r('app.color.color_0aa671'))
        .opacity(!this.adminUsername && !this.adminPassword && !this.validCode && !this.name &&
          !this.domain ? 1 : 0.5)
        .onClick(() => {
          this.getRegisterInfo()
        })
      Row() {
        Text($r('app.string.existing_account'))
        Text($r('app.string.login_in')).fontWeight(FontWeight.Bold).fontColor($r('app.color.color_0aa671'))
          .onClick(() => {
            router.back()
          })
      }

      LoadingProgress()
        .color($r('app.color.color_0aa671')).width($r('app.float.size_50')).height($r('app.float.size_50'))
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .backgroundImage($r('app.media.rect'))
  }
}