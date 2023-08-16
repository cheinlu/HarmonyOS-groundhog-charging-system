<template>
  <view class="login-container">
    <!-- 图标 -->
     <image src="./images/Groundhog.jpg" class="logo"></image>
    <!-- 登录按钮 -->
    <button type="primary" class="btn-login" @click="getAccessToken">微信用户快捷登录</button>
    <!-- 登录提示 -->
    <view class="tips-text">登录后尽享更多权益</view>
  </view>
</template>

<script setup>
import {ref} from 'vue'
import useUserStore from '@/store/user.js'
let useStore = useUserStore()
//保存code
let code = ref('')
let getAccessToken = ()=>{
  uni.login({
     provider: 'weixin',
     success(res) {
      let lcode = res.code
       useStore.userLogin(lcode)
     }
  })
}
</script>

<style lang="scss">
.login-container {
  height: 750rpx;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #f8f8f8;
  position: relative;
  overflow: hidden;

  &::after {
    content: ' ';
    display: block;
    width: 100%;
    height: 80rpx;
    background-color: white;
    position: absolute;
    bottom: 0;
    left: 0;
    border-radius: 100%;
    transform: translateY(50%);
  }
  .logo {
    width: 200rpx;
    height: 200rpx;
  }
  .btn-login {
    background-color: #0aa671;
    width: 90%;
    border-radius: 200rpx;
    margin: 50rpx 0;
  }
  .tips-text {
    font-size: 24rpx;
    color: gray;
  }
}
</style>
