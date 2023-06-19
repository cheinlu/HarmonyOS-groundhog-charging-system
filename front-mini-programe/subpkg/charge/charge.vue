<template>
  <view class="charge-container">
    <!-- 设备编号 -->
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
        <view @click="showPopPrice=true">收费规则>></view>
      </view>
      <view class="charge-price">
        <view class="charge-title">充电时常</view>
        <view class="price">
          <uni-data-select style="font-size: 20rpx" v-model="chargingTime" :localdata="timeList" @change="onTimeAndPayMethodChange"> </uni-data-select>
        </view>
      </view>
      <view class="charge-price">
        <view class="charge-title">支付方式</view>
        <view class="price">
          <uni-data-select style="font-size: 10rpx" v-model="payMethod" :localdata="payMethodList" @change="onTimeAndPayMethodChange"> </uni-data-select>
        </view>
      </view>
    </view>
    <!-- 开始充电 -->
    <view class="recharge" @click="goChargeOrder">充电</view>
    <!-- popPrice组件 -->
  <popPrice v-if="showPopPrice" :list="chargePriceList" @confirm="confimCharge"></popPrice>
  </view>
</template>

<script>
import { requestStartCharge,requestChargePrice } from '@/utils/api/home.js'
export default {
  name: 'charge',
  data() {
    return {
      showPopPrice:false,
      pageNo:1,
      pageSize:10,
      chargePriceList:[],
      //充电桩id
      id: null,
      //充电桩编号
      code: '',
      //每个充电站当前时间价格
      pricePerHour: null,
      //支付方式
      payMethod: 0,
      payMethodList: [{ value: 0, text: '余额支付' }],
      //充电时长
      chargingTime: 0,
      timeList: [{ value: 0, text: '充满即止' }],
    }
  },
  onLoad(pile) {
    //获取当前时间价格
    this.pricePerHour = pile.pricePerHour
    // 获取 id 参数
    this.id = pile.id
    // 获取 code 参数
    this.code = pile.code
    //获取充电价格的数据
    this.getChargePrice()
  },
  methods: {
    //接收PopPrice组件传递的事件，关闭弹窗
    confimCharge(){
      this.showPopPrice = false
    },
    onTimeAndPayMethodChange(e) {
      console.log('onTimeAndPayMethodChange time change:', e)
    },
    //跳转至充电订单页面
    async goChargeOrder() {
      let { data: res } = await requestStartCharge(this.id)
      if (res.code == 0) {
        uni.$showMsg('充电成功')
        uni.navigateTo({
          url: '/subpkg/chargeOrder/chargeOrder'
        })
      } else {
        uni.$showMsg(res.data.message)
        //跳转到充电桩列表
        uni.switchTab({
          url: '/pages/home/home'
        })
      }
    },
    //获取充电价格
    async getChargePrice(){
     let that = this
     let {pageNo,pageSize} = that
     let {data:res} = await requestChargePrice(pageNo,pageSize)
      if(res.code==0){
       that.chargePriceList = res.data.List
       that.pageNo = res.data.PageNo
       that.pageSize = res.data.PageSize
     }
    }
  }
}
</script>

<style lang="scss">
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
