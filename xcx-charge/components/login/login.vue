<template>
  <view class="login-container">
    <!-- 图标 -->
     <image :src="require('./static/Groundhog.jpg')" class="logo"></image>
    <!-- 登录按钮 -->
    <button type="primary" class="btn-login" @click="getUserInfo">微信用户快捷登录</button>
    <!-- 登录提示 -->
    <view class="tips-text">登录后尽享更多权益</view>
  </view>
</template>

<script>
import { requestLogin } from '@/utils/api/user.js'
import { mapMutations } from 'vuex'
export default {
  name: 'login',
  data() {
    return {}
  },
  methods: {
    ...mapMutations('m_user', ['updateUserInfo', 'updateToken']),
    //获取用户信息并登录
    getUserInfo() {
      let that = this
      uni.getUserProfile({
        desc: '获取用户信息',
        success(res) {
          that.updateUserInfo(res.userInfo)
          that.getAccessToken(res)
        }
      })
    },
    //获取token
    getAccessToken(userInfo) {
      let that = this
      uni.login({
        provider: 'weixin',
        success(res) {
          let code = res.code
          requestLogin({ code }).then((res) => {
            if (res.data.code == 0) {
              that.updateToken(res.data.data.token)
            } else if (res.data.code === 404) {
              //如果未注册，则跳转到注册页面
              that.gotoRegister()
            }
          })
        }
      })
    },
    //跳转到注册页面
    gotoRegister() {
      uni.showModal({
        title: '提示',
        content: '需要进行手机验证获得验证码',
        success: function (res) {
          if (res.confirm) {
            uni.navigateTo({
              url: '/subpkg/register/register'
            })
          } else if (res.cancel) {
            console.log('用户点击取消')
          }
        }
      })
    }
  }
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
    height: 40px;
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
    border-radius: 100px;
    margin: 25px 0;
  }
  .tips-text {
    font-size: 12px;
    color: gray;
  }
}
</style>
