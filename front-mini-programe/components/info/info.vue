<template>
  <view class="my-user-info">
    <view class="top-box">
      <image :src="avatarUrl" class="avatar-img"></image>
      <view class="nickname">{{ nickname }}</view>
    </view>
    <!-- 面板区域 -->
    <view class=".panel-list">
      <!-- 第一个面板 -->
      <view class="panel">
        <!-- panel 的主体区域 -->
        <view class="panel-top">
          <!-- panel 的 item 项 -->
          <view class="panel-item">
            <text>{{ balance.toFixed(2) }}元</text>
            <text>充电余额></text>
          </view>
        </view>
      </view>
      <!-- 第二个面板 -->
      <view class="panel">
        <!-- 主体 -->
        <view class="panel-body">
          <!-- 面板item项 -->
          <navigator url="/subpkg/chargeOrder/chargeOrder" class="panel-item">
            <image src="./images/order.png" class="icon"></image>
            <text>充电订单</text>
          </navigator>
          <view class="panel-item" @click="developing">
            <image src="./images/service.png" class="icon"></image>
            <text>客服中心</text>
          </view>
        </view>
      </view>
      <!-- 第三个面板 -->
      <view class="panel">
        <view class="panel-body2">
          <!-- 面板item项 -->
          <view v-for="(item, index) in panelData" :key="index" class="panel-item2" @click="developing">
            <image :src="item.image" class="icon2"></image>
            <text>{{ item.text }}</text>
          </view>
        </view>
      </view>
      <!-- 第四个面板 -->
      <view class="panel" @click="logout">
        <view class="panel-bottom">
          <text>退出登录</text>
          <uni-icons type="right" size="15"></uni-icons>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { requestUserInfo } from '@/utils/api/user.js'
import useUserStore from '@/store/user.js'
let useStore = useUserStore()
import agreement from './images/agreement.png'
import certification from './images/certification.png'
import feat from './images/feat.png'
import message from './images/message.png'
import person from './images/person.png'
import rule from './images/rule.png'
//以下两个数组为图表文字展示内容
const panelData = [
  { text: '用户协议', image: agreement },
  { text: '平台资质', image: certification },
  { text: '我要合作', image: feat },
  { text: '消息通知', image: message },
  { text: '个人信息', image: person },
  { text: '规则中心', image: rule }
]

let balance = ref(0)
let nickname = ref('')
let avatarUrl = ref('')
//未开发功能，点击提示正在开发中
let developing = () => {
  uni.$showMsg('正在开发中......')
}
//退出登录
let logout = () => {
  uni.showModal({
    title: '提示',
    content: '确认退出登录吗',
    success: function (res) {
      if (res.confirm) {
        useStore.userLogout()
      }
    }
  })
}
//组件挂载
onMounted(() => {
  getUserInfo()
})
//获取用户信息
let getUserInfo = async () => {
  let { data: res } = await requestUserInfo()
  if (res.code == 0) {
    balance.value = res.data.balance
    nickname.value = res.data.nickname
    avatarUrl.value = res.data.avatarUrl
  }
}
</script>

<style lang="scss">
page{
  overflow: hidden;
}
.my-user-info {
  height: 100vh;
  background-color: #f4f4f4;
  .top-box {
    height: 200rpx;
    background-color: #0aa671;
    display: flex;
    padding-left: 40rpx;
    .avatar-img {
      width: 100rpx;
      height: 100rpx;
      border-radius: 50rpx;
      border: 2rpx solid white;
      box-shadow: 0 1rpx 5rpx black;
      margin-right: 20rpx;
    }
    .nickname {
      color: white;
      margin-top: 20rpx;
      font-size: 30rpx;
      font-weight: bold;
    }
  }
}
.panel-list {
  padding: 0 20rpx;
  position: relative;
  top: -40rpx;

  .panel {
    background-color: white;
    border-radius: 15rpx;
    margin-bottom: 20rpx;
    .panel-top {
      display: flex;
      padding-left: 30rpx;
      .panel-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        font-size: 25rpx;
        padding: 20rpx 0;
      }
    }
    .panel-body {
      display: flex;
      justify-content: space-around;

      .panel-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
        font-size: 25rpx;
        padding: 20rpx 0;

        .icon {
          width: 70rpx;
          height: 70rpx;
        }
      }
    }
    .panel-body2 {
      display: flex;
      justify-content: flex-start;
      flex-wrap: wrap;
      .panel-item2 {
        display: flex;
        flex-direction: column;
        align-items: center;
        font-size: 25rpx;
        padding: 20rpx 0;
        flex-basis: calc(25%);
        .icon2 {
          width: 50rpx;
          height: 50rpx;
        }
      }
    }
    .panel-bottom {
      display: flex;
      justify-content: space-between;
      padding: 35rpx;
      font-size: 25rpx;
    }
  }
}
</style>
