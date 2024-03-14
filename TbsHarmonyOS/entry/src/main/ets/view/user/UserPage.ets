import router from '@ohos.router'
import ChargeModel from '../../model/ChargeModel'
import UserModel from '../../model/UserModel'

class panelData{
  name: string
  image: ResourceStr

  constructor(name: string,image: ResourceStr) {
    this.name = name
    this.image = image
  }
}
@Component
export default struct UserPage {

  private items:Array<panelData> = [
    new panelData('用户协议',$r('app.media.agreement')),
    new panelData('平台资质',$r('app.media.certification')),
    new panelData('我要合作',$r('app.media.feat')),
    new panelData('消息通知',$r('app.media.message')),
    new panelData('个人信息',$r('app.media.person')),
    new panelData('规则中心',$r('app.media.rule')),
  ]

  @State balance:number = 0
  //局部公共样式函数
  @Styles userPageSty(){
    .width('95%')
    .padding(10)
    .borderRadius(10)
    .backgroundColor('#FFF')
  }

  aboutToAppear(){
    this.UserInfo()
  }
  build() {
    Column({space:10}){
      //1.充电余额
      Column({space:6}){
          Text(`${this.balance.toFixed(2)}元`)
          Row(){
            Text('充电余额')
            Text('>')
          }
      }
      .margin({top:-20})
      .height(60)
      .alignItems(HorizontalAlign.Start)
      .userPageSty()
      //2.充电订单
      Row({space:10}){
        Column(){
          Image($r('app.media.order')).width(45)
          Text('充电订单')
        }.onClick(() => {
          router.pushUrl(
            {
              url:'pages/ChargeOrder',
            },
            err => {
              if(err){
                console.log(`路由失败，errCode:${err.code},errMsg:${err.message}`)
              }
            }
          )
        })
        Column(){
          Image($r('app.media.service')).width(45)
          Text('客服中心')
        }
      }
      .height(100)
      .justifyContent(FlexAlign.SpaceAround)
      .userPageSty()
      //3.用户协议
      Flex({wrap:FlexWrap.Wrap,alignContent:FlexAlign.SpaceAround}){
        ForEach(this.items,(item:panelData)=>{
          Column({space:4}){
            Image(item.image).width(30)
            Text(item.name)
          }.width('25%')
        })
      }
      .height(150)
      .width('95%')
      .padding(10)
      .borderRadius(10)
      .backgroundColor('#FFF')
      //4.退出登录
      Row(){
        Text('退出登录')
        Image($r('app.media.icon_arrow_right')).width(20).fillColor(Color.Black)
      }
      .height(60)
      .justifyContent(FlexAlign.SpaceBetween)
      .userPageSty()
      .onClick(()=>{
       AppStorage.Delete('TOKEN')
        router.replaceUrl({
          url:'pages/Login'
        })
      })
    }
  }
  async UserInfo(){
    let res:any = await UserModel.getUserInfo()
    if(res.responseCode === 200){
      let userRes = JSON.parse(res.result.toString())
      this.balance = userRes.data.balance
    }
  }
}