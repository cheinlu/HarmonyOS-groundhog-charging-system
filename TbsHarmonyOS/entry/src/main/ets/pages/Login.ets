import router from '@ohos.router'
import Prompt from '@system.prompt'
import UserModel from '../model/UserModel'
import { Tenants } from '../viewModel/Tenant'
@Entry
@Component
struct Login {
  @State tenant:string = ''
  @State account:string = ''
  @State password:string = ''
  @State Tenants:Tenants = []

  aboutToAppear(){
    this.TenantInfo()
  }

  build() {
    Column({space:15}){
      Image($r('app.media.Groundhog')).width('100vp').borderRadius(50).margin({top:120})
      Text('欢迎登录土拨鼠充电平台').fontColor(Color.Gray).fontSize(20)
        Select(this.Tenants)
          .selected(0).value('请选择租户').font({size:16,weight:500}).fontColor('#182431').selectedOptionFont({ size: 16, weight: 400 })
          .optionFont({ size: 16, weight: 400 })
          .onSelect((index: number) => {
            this.tenant = (index + 1).toString()
          })
        TextInput({
          placeholder:'请输入用户名'
        }).width(300)
          .onChange((val:string) => {
          this.account = val
            AppStorage.SetOrCreate('account',val)
        })
        TextInput({
          placeholder:'请输入密码'
        }).width(300).onChange((val:string) => this.password = val).type(InputType.Password)
      Button('登录').width("90%").backgroundColor($r('app.color.charge_header')).onClick(()=>{
        this.LoginInfo()
      })
      Row(){
        Text('没有账户？')
        Text('去注册').fontWeight(FontWeight.Bold).fontColor($r('app.color.charge_header'))
          .onClick(()=>{
            router.pushUrl({
              url:'pages/Register'
            })
          })
      }
      LoadingProgress()
        .color($r('app.color.charge_header')).width(50).height(50)
    }
    .width('100%')
    .height('100%')
    .backgroundImage($r('app.media.rect'))
  }

  //登录按钮
  async LoginInfo(){
    let res:any = await UserModel.getToken(this.tenant,this.account,this.password)
    if(res.responseCode === 200){
      let loginRes = JSON.parse(res.result.toString())
      let token = loginRes.data.token
      AppStorage.SetOrCreate('TOKEN', token);
      Prompt.showToast({
        duration: 2000,
        message:'登录成功'
      })
      router.replaceUrl({
        url:'pages/Index'
      })
    }
  }

  //获取租户信息
  async TenantInfo(){
    let res:any = await UserModel.getTenant()
    if(res.responseCode === 200){
      let tenantRes = JSON.parse(res.result.toString())
      let trArr = tenantRes.data.Data
      this.Tenants = trArr.map(item => {
        return {value:item.name}
      })
    }
  }
}