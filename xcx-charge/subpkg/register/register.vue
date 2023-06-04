<template>
  <view class="container">
    <view class="form-group">
      <text class="label">手机号</text>
      <input type="number" class="input" @input="handleInput" v-model="phoneNumber" placeholder="请输入手机号"/>
    </view>
    <view class="form-group">
      <text class="label">验证码</text>
      <input type="number" class="input" maxlength="6" v-model="verificationCode" placeholder="请输入验证码"/>
      <button class="btn-verification-code" @click="sendVerificationCode" :disabled="disableBtn">{{btnText}}</button>
    </view>
    <button class="btn-submit" @click="handleSubmit">注册</button>
    <p style="color: red;">备注：手机号可以随意填写，并不会真的发验证码，验证码只能输 666666</p>
  </view>
</template>

<script>
  import {requestVerificationCode,requestRegister} from '@/utils/api/user.js'
export default {
  name: 'register',
  data() {
    return {
      phoneNumber: '',
      verificationCode: '',
      btnText: '获取验证码',
      disableBtn: false,
      submitDisabled: true,
      countdown: 5,
      //获取用户信息
      userInfo:{
        nickname:'',
        avatarUrl:'',
        gender:''
      },
    }
  },
  methods: {
    //效验手机号格式
    handleInput() {
      // 校验手机号格式，如果正确则启用发送验证码按钮
      const reg = /^1[0-9]{10}$/
      if (reg.test(this.phoneNumber)) {
        this.disableBtn = false
      } else {
        this.disableBtn = true
      }
    },
    //获取验证码按钮
    sendVerificationCode() {
      //如果手机号为空时，获取验证码按钮点击无效，并提示输入手机号
      if(this.phoneNumber == ''){
        uni.showToast({
          icon:'none',
          title:'请输入手机号'
        })
        this.disableBtn = false
      }else{
        // 输入手机号后，发送验证码
        this.btnText = `${this.countdown}s 后重新获取`
        this.disableBtn = true
        const timer = setInterval(() => {
          if (this.countdown <= 1) {
            clearInterval(timer)
            this.btnText = '重新获取'
            this.countdown = 5
            this.disableBtn = false
          } else {
            this.countdown--
            this.btnText = `${this.countdown}s 后重新获取`
          }
        }, 1000)
        //调用验证码的接口
        requestVerificationCode({
          params:{
            phone:this.phoneNumber
          }
        }).then(res=>{
          if(res.data.code==0){
            uni.showToast({
              icon:'success',
              title:'验证码'
            })
          }
        })
      }
    },
    //点击注册按钮
    handleSubmit() {
      let that = this
      //如果手机号与验证码为空时，则返回false
      if(that.phoneNumber==''&&that.verificationCode==''){
        return false
      }else{
        //获取用户信息
      uni.getUserProfile({
        desc:'用户信息',
        success(res) {
          that.userInfo = res.userInfo,
          uni.login({
            provider:'weixin',
            success(res) {
              let regisCode = res.code
              //整理参数
              let params = {
                code:regisCode,
                phone:that.phoneNumber,
                verifyCode:that.verificationCode,
                nickname:that.userInfo.nickname,
                avatarUrl:that.userInfo.avatarUrl,
                gender:that.userInfo.gender
              }
              //发起注册请求
              requestRegister(params).then(res=>{
                console.log('res',res);
                if(res.data.code==0){
                  uni.showToast({
                     icon:'none',
                     title:'注册成功'
                   });
                   //存储用户信息
                   uni.setStorageSync('userInfo',that.userInfo)
                  //跳转至首页
                  uni.switchTab({
                    url:'/pages/my/my'
                  })
                }
              })
            }
          })
        }
      })
    }
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  padding: 40rpx;
}

.form-group {
  display: flex;
  align-items: center;
  margin-bottom: 40rpx;
}

.label {
  width: 120rpx;
  font-size: 25rpx;
}

.input {
  flex: 1;
  height: 80rpx;
  padding: 0 40rpx;
  font-size: 25rpx;
  border: 2rpx solid #ccc;
}

.btn-verification-code {
  width: 180rpx;
  height: 80rpx;
  line-height: 80rpx;
  margin-left: 40rpx;
  font-size: 25rpx;
  color: #fff;
  background-color: #0aa671;
  border: none;
  outline: none;
}

.btn-submit {
  display: block;
  width: 100%;
  height: 100rpx;
  line-height: 100rpx;
  font-size: 25rpx;
  color: #fff;
  background-color: #0aa671;
  border: none;
  outline: none;
  cursor: pointer;
}

.btn-submit.disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>
