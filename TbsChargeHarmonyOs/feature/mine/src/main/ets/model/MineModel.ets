export class UserInfoData {
  id?: number
  username?: string
  nickname?: string
  balance?: number
  avatarUrl?: string
}

export class panelData {
  type?: number
  name?: ResourceStr
  image?: ResourceStr
}

export const panelList: Array<panelData> = [
  { type: 1, name: $r('app.string.setting'), image: $r('app.media.setting') },
  { type: 2, name: $r('app.string.agreement'), image: $r('app.media.agreement') },
  { type: 3, name: $r('app.string.certification'), image: $r('app.media.certification') },
  { type: 4, name: $r('app.string.feat'), image: $r('app.media.feat') },
  { type: 5, name: $r('app.string.message'), image: $r('app.media.message') },
  { type: 6, name: $r('app.string.person'), image: $r('app.media.person') },
  { type: 7, name: $r('app.string.rule'), image: $r('app.media.rule') }
]

export const orderOrServiceList: Array<panelData> = [
  { type: 1, name: $r('app.string.order'), image: $r('app.media.order') },
  { type: 2, name: $r('app.string.service'), image: $r('app.media.service') }
]