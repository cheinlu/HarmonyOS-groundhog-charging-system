<template>
  <view class="login-container">
    <!-- 图标 -->
    <image src="./images/Groundhog.jpg" class="logo"></image>
    <!-- 登录按钮 -->
    <button type="primary" class="btn-login" @click="getLogin">微信用户快捷登录</button>
    <!-- 登录提示 -->
    <view class="tips-text">登录后尽享更多权益</view>
  </view>
</template>

<script setup>
//引入接口
import { requestLogin, requestRegister } from '@/utils/api/user.js'

//引入本地存储
import { SET_INFO,SET_TOKEN } from '@/utils/token'
import useUserStore from '@/store/user.js'
let useStore = useUserStore()
//登录按钮
const getLogin = () => {
  uni.login({
    provider: 'weixin',
    success(res) {
      const code = { code: res.code }
      requestLogin(code).then((lres) => {
        if (lres.data.code == 0) {
          useStore.token = lres.data.data.token
          SET_TOKEN(lres.data.data.token)
        } else {
          uni.showModal({
            title: '温馨提示',
            content: '微信授权登录后才能完整使用小程序功能',
            confirmText: '同意',
            cancelText: '拒绝',
            success: function (res) {
              if (res.confirm) {
                getRegister()
              } else if (res.cancel) {
                console.log('点击了拒绝按钮，不做操作');
              }
            }
          })
        }
      })
    }
  })
}
//注册逻辑
const getRegister = () => {
  uni.getUserProfile({
    desc: '用于完善会员信息',
    lang: 'zh_CN',
    success(res) {
      useStore.updateUserInfo(res.userInfo)
      SET_INFO(res.userInfo)
      let userInfo = res.userInfo
      uni.login({
        provider: 'weixin',
        success(res) {
          let regisCode = res.code
          //整理参数
          let params = {
            code: regisCode,
            nickname: userInfo.nickName,
            avatarUrl: userInfo.avatarUrl,
            gender: userInfo.gender
          }
          //发起注册请求
          requestRegister(params)
            .then((res) => {
              if (res.data.code == 0) {
                //存储token
                useStore.updateToken(res.data.data.token)
                getLogin()
              } else {
                // uni.$showMsg(res.data.message)
              }
            })
            .catch((err) => {
              console.error(err)
              uni.$showMsg('服务器出错，请稍后再试')
            })
        }
      })
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
