<template>
  <view class="charge-container">
    <!-- 设备编号 -->
    <view class="serial-list station">
      <!-- 编号名字 -->
      <view class="serial-title">充电站</view>
      <!-- 对应的编号 -->
      <view class="serial">{{ stationName }}</view>
    </view>
    <view class="serial-list">
      <!-- 编号名字 -->
      <view class="serial-title">设备编号</view>
      <!-- 对应的编号 -->
      <view class="serial">{{ code }}</view>
    </view>

    <!-- 充电桩基本信息 -->
    <view class="charge-list">
      <view class="charge-price">
        <view class="charge-title">收费标准</view>
        <view @click="showPopPrice = true">收费规则>></view>
      </view>
      <view class="charge-price">
        <view class="charge-title">充电时常</view>
        <view class="price">
          <uni-data-select style="font-size: 20rpx" v-model="chargingTime" :localdata="timeList" @change="onTimeAndPayMethodChange"> </uni-data-select>
        </view>
      </view>
    <!--  <view class="charge-price">
        <view class="charge-title">支付方式</view>
        <view class="price">
          <uni-data-select style="font-size: 10rpx" v-model="payMethod" :localdata="payMethodList" @change="onTimeAndPayMethodChange"> </uni-data-select>
        </view>
      </view> -->
       <view style="color: red;">PS：此功能不涉及支付，只做演示使用</view>
    </view>
    <!-- 开始充电 -->
    <view class="recharge" @click="goChargeOrder">充电</view>
   
    <!-- 收费标准弹窗 -->
    <pop v-if="showPopPrice" :list="chargePriceList" @confirm="confimCharge"></pop>
  </view>
</template>

<script setup>
import { onLoad } from '@dcloudio/uni-app'
import { requestStartCharge, requestChargePrice } from '@/utils/api/charge.js'
import { ref, reactive } from 'vue'
import { defineProps, onMounted } from 'vue'
const props = defineProps(['id', 'code', 'pricePerHour','stationName'])
let scannedId = ref(null) //扫码进入页面时的id值
let showPopPrice = ref(false)
let pageSize = ref(10)
let pageNo = ref(1)
let chargePriceList = ref([])
let chargingTime = ref(0)
let payMethod = ref(0)
let payMethodList = reactive([{ value: 0, text: '无' }])
let timeList = reactive([{ value: 0, text: '充满即止' }])
onLoad((options) => {
  //判断是否通过扫码进入页面并包含id
  if (options.pileId) {
    scannedId.value = options.pileId
  }
})
// 在mounted钩子函数中获取并使用参数的值
onMounted(() => {
  //获取充电价格的数据
  getChargePrice()
})

let onTimeAndPayMethodChange = (e) => {
  console.log('onTimeAndPayMethodChange time change:', e)
}
//获取充电价格
let getChargePrice = async () => {
  let { data: res } = await requestChargePrice(pageNo.value, pageSize.value)
  if (res.code == 0) {
    chargePriceList.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
  }
}
//接收子组件pop传递的事件，关闭弹窗
let confimCharge = () => {
  showPopPrice.value = false
}
//充电按钮,跳转去充电记录列表页面
let goChargeOrder = async () => {
  let id = scannedId.value || props.id
  if (!id) {
    uni.$showMsg('没有获取到有效的充电桩ID')
    return
  }
  let { data: res } = await requestStartCharge(id)
  if (res.code == 0) {
    uni.$showMsg('充电成功')
    uni.navigateTo({
      url: '/subpkg/chargeOrder/chargeOrder'
    })
  } else {
    uni.$showMsg(res.data.message)
    //跳转到充电桩列表
    uni.switchTab({
      url: '/pages/charge/charge'
    })
  }
}
</script>

<style lang="scss">
 page{
  overflow: hidden; 
}
.charge-container {
  height: 100vh;
  padding: 30rpx;
  background-color: #f4f4f4;
  
  .serial-list {
    background-color: #ffffff;
    padding: 20rpx;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 15rpx;
    font-size: 30rpx;
  }
  .station{
    margin-bottom: 15px;
  }
  .charge-list {
    background-color: #ffffff;
    padding: 20rpx;
    border-radius: 15rpx;
    font-size: 30rpx;
    margin: 30rpx 0;
    .charge-price {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin: 30rpx 0;
      .price {
        width: 200rpx;
      }
    }
  }
  .recharge {
    background-color: #0aa671;
    padding: 20rpx;
    border-radius: 15rpx;
    font-size: 30rpx;
    text-align: center;
    color: white;
  }
}
</style>
