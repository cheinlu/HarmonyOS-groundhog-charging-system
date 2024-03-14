import router from '@ohos.router'
import UserModel from '../model/UserModel'
@Entry
@Component
struct Register {
  @State name:string = '' //企业名称
  @State domain:string = '' //网站
  @State adminUsername:string = '' //用户名
  @State adminPassword:string = '' //密码
  @State validCode:string = '' //验证码

  build() {
    Column({space:15}){
      Image($r('app.media.Groundhog')).width('100vp').borderRadius(50).margin({top:120})
      Text('欢迎注册土拨鼠充电平台').fontColor(Color.Gray).fontSize(20)
      TextInput({
        placeholder:'请输入用户名'
      }).width(300)
        .onChange((val:string) => {
          this.adminUsername = val
        })
      Row(){
        TextInput({placeholder:'验证码:666666'}).width(150).onChange((val:string)=>this.validCode = val)
        Button('获取验证码').backgroundColor($r('app.color.charge_header'))
      }.justifyContent(FlexAlign.SpaceBetween).width(300)
      TextInput({
        placeholder:'密码:长度至少5位'
      }).width(300).onChange((val:string) => this.adminPassword = val).type(InputType.Password)
      TextInput({
        placeholder:'公司名称'
      }).width(300)
        .onChange((val:string) => {
          this.name = val
        })
      TextInput({
        placeholder:'公司网站,如www.example.com'
      }).width(300)
        .onChange((val:string) => {
          this.domain = val
        })

      Button('立即注册').width("90%").backgroundColor($r('app.color.charge_header'))
      .onClick(()=>{
        this.RegisterInfo()
      })
      Row(){
        Text('已有账户？')
        Text('去登录').fontWeight(FontWeight.Bold).fontColor($r('app.color.charge_header'))
          .onClick(()=>{
            router.pushUrl({
              url:'pages/Login'
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
  //注册
  async RegisterInfo(){
    let res:any = await UserModel.getRegister(this.name,this.domain,this.adminUsername,this.adminPassword,this.validCode)
    if(res.responseCode === 200){
      let registerRes = JSON.parse(res.result.toString())
      if(registerRes.code == 0){
        router.pushUrl({
          url:'pages/Login'
        })
      }
    }
  }
}