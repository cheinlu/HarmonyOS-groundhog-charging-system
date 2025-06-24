import { BottomLoadingBuilder, formatDate, StyleConst, TbsTitleBar, toast } from "@tbs/common"
import { OrderType } from "../model/ChargeModel"
import { getChargeOrders, getStopCharge } from "../viewModel/ChargeViewModel"

@Entry({ routeName: 'ChargeOrderPage' })
@Component
export struct ChargeOrderPage {
  @State Orders: OrderType[] = []
  @State loading: boolean = false // 加载状态
  @State pageNum: number = 1 //当前页，起始值为1
  @State pageSize: number = 5
  @State haveNextPage: boolean = false //是否还有下一页数据

  onPageShow(): void {
    this.resetPage()
    this.getChargeOrdersData(false)
  }

  // 重置分页状态
  resetPage() {
    this.pageNum = 1
    this.haveNextPage = true
    this.Orders = []
  }

  // 获取充电列表数据
  async getChargeOrdersData(isLoadMore: boolean = false) {
    // 如果没有下一页，不再请求
    if (isLoadMore && !this.haveNextPage) {
      return
    }

    this.loading = true
    try {
      // 计算当前请求的页码
      const currentPage: number = isLoadMore ? this.pageNum + 1 : 1

      let data = await getChargeOrders(currentPage, this.pageSize)

      if (data && data.List) {
        // 更新分页状态
        this.pageNum = currentPage
        this.haveNextPage = data.List.length >= this.pageSize

        // 更新订单列表
        if (isLoadMore) {
          this.Orders = [...this.Orders, ...data.List]
        } else {
          this.Orders = data.List
        }
      }
    } catch (error) {
      toast($r('app.string.load_order_error'))
      console.error('获取订单失败:', error)
    } finally {
      this.loading = false
    }
  }

  //结束充电
  async getStopChargeInfo(id: number) {
    let data = await getStopCharge(id)
    if (data != null && Number(data) === 0) {
      this.Orders = this.Orders.map(item => {
        if (item.id == id) {
          return {
            id: item.id,
            pileCode: item.pileCode,
            orderCode: item.orderCode,
            price: item.price,
            startAt: item.startAt,
            state: 1,
            stationName: item.stationName,
            stopAt: item.stopAt
          }
        }
        return item
      })
      toast($r('app.string.stop_charge_success'))
    } else {
      toast($r('app.string.stop_charge_fail'))
    }
  }

  @Builder
  OrderCard() {
    List({ space: 10 }) {
      ForEach(this.Orders, (item: OrderType) => {
        ListItem() {
          RelativeContainer() {
            Text(item.orderCode)
              .fontColor(Color.White)
              .height($r('app.float.size_40'))
              .padding($r('app.float.size_10'))
              .borderRadius({ bottomRight: $r('app.float.size_10'), bottomLeft: $r('app.float.size_10') })
              .backgroundColor($r('app.color.color_0aa671'))
              .id('code_text')
              .alignRules({
                right: { anchor: "__container__", align: HorizontalAlign.End },
              })

            Column({ space: 8 }) {
              if (item.state === 0) {
                Text($r('app.string.charging_in_progress')).fontWeight(FontWeight.Bold).fontColor(Color.Red)
              } else {
                Text($r('app.string.completed')).fontWeight(FontWeight.Bold).fontColor($r('app.color.color_0aa671'))
              }
              Row({ space: 10 }) {
                Text($r('app.string.pile_code'))
                Text(item.pileCode)
              }

              Row({ space: 10 }) {
                Text($r('app.string.charge_name'))
                Text(item.stationName)
              }

              Row({ space: 10 }) {
                Text($r('app.string.price'))
                Text(`${item.price}`)
              }

              Row({ space: 10 }) {
                Text($r('app.string.order_status'))
                Text(item.state == 0 ? '充电中' : item.state == 1 ? '已完成' : '未付账')
              }

              Row({ space: 10 }) {
                Text($r('app.string.start_time'))
                Text(formatDate(item.startAt ?? ''))
              }

              Row({ space: 10 }) {
                Text($r('app.string.end_time'))
                Text(item.stopAt && item.state !== 0 ? formatDate(item.stopAt) : '')
              }

              if (item.state === 0) {
                Button($r('app.string.stop_charge')).width($r('app.float.size_90')).backgroundColor(Color.Red)
                  .onClick(() => {
                    if (item.id) {
                      this.getStopChargeInfo(item.id)
                    }
                  })
              } else {
                Text('')
              }
            }
            .alignItems(HorizontalAlign.Start)
            .id('content')
            .alignRules({
              top: { anchor: "code_text", align: VerticalAlign.Bottom }
            })
          }
          .width(StyleConst.FULL_WIDTH)
          .height('auto')
          .padding({ left: $r('app.float.size_10'), bottom: $r('app.float.size_10') })
          .borderRadius($r('app.float.size_10'))
          .backgroundColor('#FFF')
        }
        .width(StyleConst.FULL_WIDTH)
      })
      ListItem() {
        BottomLoadingBuilder({
          haveNextPage: this.haveNextPage,
          loading: this.loading,
          dataLength: this.Orders.length
        })
      }
    }
    .width(StyleConst.FULL_WIDTH)
    .padding({ left: $r('app.float.size_10'), right: $r('app.float.size_10') })
    .layoutWeight(1)
    .scrollBar(BarState.Off)
    .edgeEffect(EdgeEffect.None)
    .onReachEnd(() => {
      if (this.haveNextPage && !this.loading) {
        this.getChargeOrdersData(true)
      }
    })
  }

  build() {
    Column({ space: 10 }) {
      //1.标题
      TbsTitleBar({ title: $r('app.string.order'), showBackIcon: true, titleBarBgColor: $r('app.color.color_0aa671') })
      //2.订单
      this.OrderCard()
    }
    .width(StyleConst.FULL_WIDTH)
    .height(StyleConst.FULL_HEIGHT)
    .backgroundColor($r('app.color.color_efefef'))
  }
}