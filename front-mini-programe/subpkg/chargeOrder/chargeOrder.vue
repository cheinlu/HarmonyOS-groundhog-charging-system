<template>
  <view class="order-content">
    <!-- 无充电记录 -->
    <view class="record-rect" v-if="chargeOrderList.length === 0">
      <view class="record">
        <image src="./images/empty.png" class="empty-img"></image>
        <view class="anonymous">暂无订单</view>
        <view class="anonymous-order">您还没有订单记录</view>
      </view>
    </view>
    <!-- 有充电记录 -->
    <view v-else>
      <scroll-view scroll-y="true" @scrolltolower="loadMore" style="height: 100vh">
        <view class="list-rect-item" v-for="item in chargeOrderList" :key="item.id">
          <view class="list-item">
            <view class="list-item-title">{{ item.orderCode }}</view>
            <view class="list-item-info">
              <view class="item-top" v-if="item.state === 0" style="color: #e64340">充电中</view>
              <view class="item-top" v-else>已完成</view>
              <view class="item-bottom">充电桩编号：{{ item.pileCode }}</view>
              <view class="item-bottom">充电站名：{{ item.stationName }}</view>
              <view class="item-bottom">价格：{{ item.price }}</view>
              <view class="item-bottom">订单状态： {{ item.state === 0 ? '充电中' : item.state === 1 ? '已完成' : '未付账' }} </view>
              <view class="item-bottom">开始时间：{{ formatDate(item.startAt) }}</view>
              <view class="item-bottom" v-if="item.stopAt !== undefined && item.stopAt !== null && item.state !== 0">结束时间：{{ item.state === 0 ? '' : formatDate(item.stopAt) }} </view>
              <button class="mini-btn" type="warn" size="mini" @click="stopCharge(item)" v-if="item.state === 0">停止充电</button>
              <view></view>
            </view>
          </view>
        </view>
      </scroll-view>
    </view>
  </view>
</template>

<script setup>
import { requestChargeOrders, requestStopCharge } from '@/utils/api/charge.js'
import { ref, onMounted } from 'vue'
import { formatDate } from '@/utils/time.js'
let pageSize = ref(10)
let pageNo = ref(1)
let total = ref(0)
let chargeOrderList = ref([])
//组件挂载
onMounted(() => {
  getChargeOrderList()
})
//获取充电订单列表
let getChargeOrderList = async () => {
  let { data: res } = await requestChargeOrders(pageNo.value, pageSize.value)
  if (res.code == 0) {
    chargeOrderList.value = chargeOrderList.value.concat(res.data.List)
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//停止充电
let stopCharge = async (item) => {
  let { data: res } = await requestStopCharge(item.id)
  if (res.code == 0) {
    item.state = 1
    uni.$showMsg('已停止充电')
  } else {
    uni.$showMsg(res.message)
  }
}
//上拉加载更多
let loadMore = () => {
  // 利用Math.ceil算出新的分页
  let totalPage = Math.ceil(total.value / pageSize.value)
  if (pageNo.value >= totalPage) {
    uni.$showMsg('已经到最底部了')
    return false
  } else {
    pageNo.value++
    getChargeOrderList()
  }
}
</script>

<style lang="scss" scoped>
.order-content {
  height: 100vh;
  background: #f4f6f4;
  .record-rect {
    height: 100vh;
    background: #f4f6f4;
    .record {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      color: #3d3d3d;
      .empty-img {
        margin-top: 300rpx;
        width: 300rpx;
        height: 300rpx;
      }
      .anonymous {
        font-size: 30rpx;
        font-family: Source Han Sans CN-Regular, Source Han Sans CN;
        font-weight: 700;
        margin: 33rpx;
      }
      .anonymous-order {
        font-size: 25rpx;
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
}
</style>
