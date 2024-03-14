import UserHeader from './UserHeader'
import UserPage from './UserPage'
@Component
export default struct UserIndex {
  build() {
    Column(){
      //1.头部
      UserHeader()
      //2.主体
      UserPage()
    }
    .width('100%')
    .height('100%')
    .backgroundColor('#EFEFEF')
  }
}