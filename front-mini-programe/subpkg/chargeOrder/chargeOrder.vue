<template>
  <view>
    <!-- 无充电记录 -->
    <view class="record-rect" v-if="chargeOrderList.length === 0">
    			<view class="record">
    				<image src="/static/empty.png" class="empty-img" mode=""></image>
    				<view class="anonymous">暂无订单</view>
    				<view class="anonymous-order">您还没有订单记录</view>
    			</view>
    		</view>
    <!-- 有充电记录 -->
    <view v-else >
    <scroll-view scroll-y="true" :enable-back-to-top="true" :refresher-enabled="true" @refresh="onRefresh" @scrolltolower="loadMore">
      <view class="list-rect-item" v-for="item in chargeOrderList" :key="item.id">
        <view class="list-item">
          <view class="list-item-title">{{ item.orderCode }}</view>
          <view class="list-item-info">
            <view class="item-top" v-if="item.state === 0" style="color: #e64340">充电中</view>
            <view class="item-top" v-else>已完成</view>
            <view class="item-bottom">充电桩编号：{{ item.pileCode }}</view>
            <view class="item-bottom">充电站名：{{ item.stationName }}</view>
            <view class="item-bottom">价格：{{ item.price }}</view>
            <view class="item-bottom"
              >订单状态：
              {{ item.state === 0 ? '充电中' : item.state === 1 ? '已完成' : '未付账' }}
            </view>
            <view class="item-bottom">开始时间：{{ formatDate(item.startAt) }}</view>
            <view class="item-bottom">结束时间：{{ item.state === 0 ? '' : formatDate(item.stopAt) }}</view>
            <button class="mini-btn" type="warn" size="mini" @click="stopCharge(item)" v-if="item.state === 0">停止充电</button>
            <view v-else></view>
          </view>
        </view>
      </view>
    </scroll-view>
    </view>
  </view>
</template>

<script>
import { requestChargeOrders, requestStopCharge } from '@/utils/api/home.js'
export default {
  name: 'chargeOrder',
  data() {
    return {
      // 充电订单列表
      chargeOrderList: [],
      // 列表总数
      total: 0,
      chargeQury: {
        pageNo: 1,
        pageSize: 10
      }
    }
  },
  onLoad() {
    this.getChargeOrderList()
  },
  methods: {
    //简化时间格式
    formatDate(dateStr) {
      const date = new Date(dateStr)
      const year = date.getFullYear()
      const month = date.getMonth() + 1
      const day = date.getDate()
      const hour = date.getHours()
      const minute = date.getMinutes()
      const formattedDate = `${year}/${month}/${day} ${hour}:${minute}`
      return formattedDate
    },
    //获取充电订单列表
    async getChargeOrderList() {
      let { data: res } = await requestChargeOrders()
      if (res.code == 0) {
        this.chargeOrderList = res.data.List
        this.chargeQury.pageNo = res.data.PageNo
        this.chargeQury.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      }
    },
    //下拉刷新
    async onRefresh() {
      const pageNo = this.chargeQury.pageNo + 1
      const pageSize = this.chargeQury.pageSize
      const { data: res } = await requestChargeOrders({ pageNo, pageSize })
      if (res.code === 0 && res.data.List.length > 0) {
        // 将新获取的数据插入到原有数据的最前面
        this.chargeOrderList = res.data.List.concat(this.chargeOrderList)
      } else {
        console.log('已经全部加载完成')
      }
    },
    //上拉加载更多
    async loadMore() {
      console.log('上拉加载更多')
      if (this.chargeOrderList.length >= this.total) {
        console.log('已经全部加载完成')
        return
      }
      const pageNo = this.chargeQury.pageNo + 1
      const pageSize = this.chargeQury.pageSize
      const { data: res } = await requestChargeOrders({ pageNo, pageSize })
      if (res.code === 0 && res.data.List.length > 0) {
        // 将新获取的数据合并到原有数据列表中
        this.chargeOrderList = [...this.chargeOrderList, ...res.data.List]
        this.chargeQury.pageNo = res.data.PageNo
      } else {
        console.log('已经全部加载完成')
      }
    },
    //停止充电
    async stopCharge(item) {
      console.log('停止充电', item)
      let { data: res } = await requestStopCharge(item.id)
      console.log('res', res)
      if (res.code == 0) {
        item.state = 1
        uni.showToast({
          icon: 'none',
          title: '已停止充电'
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.record-rect{
  height: 100vh;
  background: #f4f6f4;
		.record{
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
       color: #3D3D3D;
			.empty-img{
        margin-top: 300rpx;
				width: 300rpx;
				height: 300rpx;
			}
			.anonymous{
				font-size: 30rpx;
				font-family: Source Han Sans CN-Regular, Source Han Sans CN;
				font-weight: 700;
				margin: 33rpx;
			}
      .anonymous-order{
        font-size: 20rpx;
      }
		}
	}
.list-rect-item {
  padding: 30rpx;
  background: #f4f6f4;
  .list-item {
    padding: 30rpx;
    background: #fbfbfb;
    box-shadow: 0px 0px 10rpx -4rpx rgba(0, 0, 0, 0.3);
    border-radius: 25rpx;
    position: relative;
    .list-item-title {
      width: 260rpx;
      height: 69rpx;
      font-size: 23rpx;
      line-height: 69rpx;
      text-align: center;
      color: #fbfbfb;
      background: #0aa671;
      border-radius: 0rpx 25rpx 0rpx 25rpx;
      position: absolute;
      top: 0;
      right: 0;
    }
    .list-item-info {
      font-weight: 400;
      width: 100%;
      font-size: 29rpx;
      font-family: Source Han Sans CN-Regular, Source Han Sans CN;
      .item-top {
        color: #0aa671;
        font-weight: bold;
        margin-bottom: 15rpx;
      }
      .item-bottom {
        margin-bottom: 15rpx;
      }
    }
  }
}
</style>
