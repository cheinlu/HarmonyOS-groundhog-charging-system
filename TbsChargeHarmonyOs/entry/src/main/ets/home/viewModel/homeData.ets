
export class ButtonInfoModel {
  index: number = 0;
  img: Resource = $r('app.media.home');
  selectImg: Resource = $r('app.media.home_active');
  title: Resource = $r('app.string.charge');
}

const buttonInfo: ButtonInfoModel[] = [
  {
    index: 0,
    img: $r("app.media.home"),
    selectImg: $r("app.media.home_active"),
    title: $r('app.string.charge')
  },
  {
    index: 1,
    img: $r("app.media.scan2"),
    selectImg: $r("app.media.scan2"),
    title: $r('app.string.scan')
  },
  {
    index: 2,
    img: $r("app.media.my"),
    selectImg: $r("app.media.my_active"),
    title: $r('app.string.mine')
  }
]

export { buttonInfo }