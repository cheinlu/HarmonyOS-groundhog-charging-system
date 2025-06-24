import { StyleConst, toast } from '@tbs/common';
import { TenantItem, TenantValueItem } from './model/LoginModel';
import { router } from '@kit.ArkUI';
import { getTenant, getToken } from './viewModel/LoginViewModel';

@Entry
@Component
struct LoginPage {
  @State tenant: string = ''
  @State username: string = ''
  @State password: string = ''
  @State Tenants: TenantValueItem[] = []

  aboutToAppear() {
    this.getTenantData()
  }

  //获取租户信息
  async getTenantData() {
    let data = await getTenant()
    if (data) {
      this.Tenants = data.map((item: TenantItem): TenantValueItem => ({
        value: item.name || ''
      })) || [];
    }
  }

  //去登录
  async login() {
    if (!this.tenant) {
      toast($r('app.string.please_select_the_tenant'))
      return
    }
    if (!this.username) {
      toast($r('app.string.please_enter_the_username'))
      return
    }
    if (!this.password) {
      toast($r('app.string.please_enter_the_password'))
      return
    }
    let data = await getToken(this.tenant, this.username, this.password)
    if (data) {
      router.replaceUrl({
        url: 'home/HomePage'
      })
    }
  }

  build() {
    Column({ space: 15 }) {
      Image($r('app.media.Groundhog'))
        .width($r('app.float.size_100'))
        .borderRadius($r('app.float.size_50'))
        .margin({ top: $r('app.float.size_120') })
      Text($r('app.string.welcome_login_groundhog'))
        .fontColor(Color.Gray)
        .fontSize($r('app.float.size_text_20'))
        .textAlign(TextAlign.Center)
      Select(this.Tenants)
        .selected(0)
        .value($r('app.string.please_select_the_tenant'))
        .font({ size: $r('app.float.size_16'), weight: FontWeight.Bold })
        .fontColor($r('app.color.color_182431'))
        .selectedOptionFont({ size: $r('app.float.size_16'), weight: FontWeight.Regular })
        .optionFont({ size: $r('app.float.size_16'), weight: FontWeight.Regular })
        .width(StyleConst.PERCENT_60)
        .onSelect((index: number) => {
          this.tenant = (index + 1).toString()
        })
      TextInput({
        placeholder: $r('app.string.please_enter_the_username')
      })
        .width(StyleConst.PERCENT_90)
        .onChange((val: string) => {
          this.username = val
        })
      TextInput({
        placeholder: $r('app.string.please_enter_the_password')
      })
        .width(StyleConst.PERCENT_90)
        .onChange((val) => {
          this.password = val
        }).type(InputType.Password)
      Button($r('app.string.login'))
        .width(StyleConst.PERCENT_90)
        .backgroundColor($r('app.color.color_0aa671'))
        .opacity(!this.tenant && !this.username && !this.password ? 1 : 0.5)
        .onClick(() => {
          this.login()
        })
      Row() {
        Text($r('app.string.no_account'))
        Text($r('app.string.register')).fontWeight(FontWeight.Bold).fontColor($r('app.color.color_0aa671'))
          .onClick(() => {
            router.pushUrl({ url: 'register/RegisterPage' })
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