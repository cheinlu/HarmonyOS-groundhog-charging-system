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
          <navigator url="../../subpkg/chargeOrder/chargeOrder" class="panel-item">
            <image :src="require('./static/order.png')" class="icon"></image>
            <text>充电订单</text>
          </navigator>
          <view class="panel-item" @click="developing">
            <image :src="require('./static/service.png')" class="icon"></image>
            <text>客服中心</text>
          </view>
        </view>
      </view>
      <!-- 第三个面板 -->
      <view class="panel" >
        <view class="panel-body2">
          <!-- 面板item项 -->
          <view class="panel-item2" @click="developing">
            <image :src="require('./static/person.png')" class="icon2"></image>
            <text>个人信息</text>
          </view>
          <!-- 面板item项 -->
          <view class="panel-item2" @click="developing">
            <image :src="require('./static/feat.png')" class="icon2"></image>
            <text>我要合作</text>
          </view>
          <!-- 面板item项 -->
          <view class="panel-item2" @click="developing">
            <image :src="require('./static/rule.png')" class="icon2"></image>
            <text>规则中心</text>
          </view>
          <!-- 面板item项 -->
          <view class="panel-item2" @click="developing">
            <image :src="require('./static/certification.png')" class="icon2"></image>
            <text>平台资质</text>
          </view>
          <!-- 面板item项 -->
          <view class="panel-item2" @click="developing">
            <image :src="require('./static/agreement.png')" class="icon2"></image>
            <text>用户协议</text>
          </view>
          <!-- 面板item项 -->
          <view class="panel-item2" @click="developing">
            <image :src="require('./static/message.png')" class="icon2"></image>
            <text>消息通知</text>
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

<script>
import { mapState, mapMutations } from 'vuex'
import { requestUserInfo } from '@/utils/api/user.js'
export default {
  name: 'userInfo',
  data() {
    return {
      //余额
      balance: null||0,
      nickname:'',
      avatarUrl:''
    }
  },
  computed: {
    ...mapState('m_user', ['userInfo'])
  },
  mounted() {
    this.getUserInfo()
  },
  methods: {
    ...mapMutations('m_user', ['updateUserInfo', 'updateToken']),
    //退出登录
    logout() {
      let that = this
      uni.showModal({
        title: '提示',
        content: '确认退出登录吗',
        success: function (res) {
          if (res.confirm) {
            //清除token
            that.updateToken('')
            //清除用户信息
            that.updateUserInfo({})
          }
        }
      })
    },
    //获取用户信息
    async getUserInfo() {
      try {
        let that = this
        let { data: res } = await requestUserInfo()
        console.log('res',res);
        if (res.code == 0) {
          that.balance = res.data.balance
          that.nickname = res.data.nickname
          that.avatarUrl = res.data.avatarUrl
        }else{
         uni.$showMsg(res.message)
         uni.navigateTo({
           url: '/subpkg/register/register'
         })
        }
      } catch (error) {
        console.log(error)
      }
    },
    developing(){
      //未开发的功能提示正在开发中
      uni.$showMsg('正在开发中......')
    }
  }
}
</script>

<style lang="scss">
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
