<template>
  <view class="my-user-info">
    <view class="top-box">
      <button class="avatar" open-type="chooseAvatar" @chooseavatar="onChooseAvatar">
        <image :src="avatarUrl" class="avatar-img"></image>
      </button>
      <input class="nickname" type="nickname" :value="nickname" @blur="bindblur" placeholder="请输入昵称" />
    </view>
    <!-- 内容区域 -->
    <view class=".panel-list">
      <view class="panel">
        <view class="panel-top">
          <view class="panel-item">
            <text>{{ balance.toFixed(2) }}元</text>
            <text>充电余额></text>
          </view>
        </view>
      </view>
      <view class="panel">
        <view class="panel-body">
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
      <view class="panel">
        <view class="panel-body2">
          <view v-for="(item, index) in panelData" :key="index" class="panel-item2" @click="developing">
            <image :src="item.image" class="icon2"></image>
            <text>{{ item.text }}</text>
          </view>
        </view>
      </view>
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
import BASE_URL from '@/env.js'
import { onMounted, ref } from 'vue'
import { requestUserInfo, requestNickname } from '@/utils/api/user.js'
import useUserStore from '@/store/user.js'
import agreement from './images/agreement.png'
import certification from './images/certification.png'
import feat from './images/feat.png'
import message from './images/message.png'
import person from './images/person.png'
import rule from './images/rule.png'
let useStore = useUserStore()
let token = useUserStore().token
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
  uni.$showMsg('开发中...')
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

//更改头像
const onChooseAvatar = async e => {
  const tempFilePath = e.detail.avatarUrl //上传的图片地址
  const maxSizeInBytes = 1024 * 1024 // 限制大小为1MB
  uni.getFileInfo({
    filePath: tempFilePath,
    success: res => {
      const fileSize = res.size
      if (fileSize > maxSizeInBytes) {
        //如果上传的图片大小超过1MB，进行提示
        uni.$showMsg('请上传小于1MB的图片')
        return
      }
      //图片大小符合，替换图片
      avatarUrl.value = tempFilePath
      //将更改的图片上传到服务器
      uni.uploadFile({
        url: BASE_URL + '/wx-api/upload/avatar',
        filePath: avatarUrl.value,
        name: 'file',
        header: {
          Authorization: 'Bearer ' + token // 将 token 添加到请求的 header 中
        },
        success(res) {
          let fileRes = JSON.parse(res.data)
          uni.$showMsg(fileRes.message)
        }
      })
    }
  })
}

//获取微信昵称
const bindblur = async e => {
  const newNickname = (nickname.value = e.detail.value)
  //将更改的昵称上传给接口
  let res = await requestNickname({ newNickname })
  uni.$showMsg(res.data.message)
}
</script>

<style lang="scss">
.my-user-info {
  height: 100vh;
  background-color: #f4f4f4;
  .top-box {
    height: 240rpx;
    background-color: #0aa671;
    display: flex;
    flex-direction: column;
    align-items: center;
    .avatar {
      width: 100rpx;
      height: 100rpx;
      border-radius: 50rpx;
      border: 2rpx solid white;
      box-shadow: 0 1rpx 5rpx black;
      padding: 0;
      .avatar-img {
        width: 100%;
        height: 100%;
      }
    }
    .nickname {
      color: white;
      margin-top: 20rpx;
      text-align: center;
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
