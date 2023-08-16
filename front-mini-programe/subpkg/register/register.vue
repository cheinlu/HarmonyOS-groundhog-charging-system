<template>
  <view class="container">
    <view class="form-group">
      <text class="label">手机号</text>
      <input type="number" class="input" @input="handleInput" v-model="phoneNumber" placeholder="请输入手机号"/>
    </view>
    <view class="form-group">
      <text class="label">验证码</text>
      <input type="number" class="input" maxlength="6" v-model="verificationCode" placeholder="请输入验证码" :disabled="codeInputDisabled" />
      <button class="btn-verification-code" @click="sendVerificationCode" :disabled="disableBtn">{{btnText}}</button>
    </view>
      <button v-if="submitDisabled" class="btn-submit disabled">注册</button>
      <button v-else class="btn-submit" @click="handleSubmit">注册</button>
    <p style="color: red;">备注：手机号可以随意填写，并不会真的发验证码，验证码只能输 666666</p>
  </view>
</template>

<script setup>
import {requestVerificationCode,requestRegister} from '@/utils/api/user.js'
import {ref,reactive} from 'vue'
import useUserStore from '@/store/user.js'
let useStore = useUserStore()
let phoneNumber = ref('')
let verificationCode = ref('')
let codeInputDisabled = ref(true)
let submitDisabled = ref(true)
let disableBtn = ref(false)
let btnText = ref('获取验证码')
let countdown = ref(5)
let userInfo = reactive({
  nickName:'',
  avatarUrl:'',
  gender:''
})
//效验手机号格式
let handleInput = ()=>{
  const reg = /^1[0-9]{10}$/
  if (reg.test(phoneNumber.value)) {
    disableBtn.value = false
  } else {
    disableBtn.value = true
  }
}
//获取验证码按钮
let sendVerificationCode = ()=>{
   if(phoneNumber.value == ''){
     uni.$showMsg('请输入手机号')
     disableBtn.value = false
   }else{
     // 输入手机号后，发送验证码
     btnText.value = `${countdown.value}s 后重新获取`
     disableBtn.value = true
     codeInputDisabled.value = false //使验证码输入框变为可输入状态
     submitDisabled.value = false
     const timer = setInterval(() => {
       if (countdown.value <= 1) {
         clearInterval(timer)
         btnText.value = '重新获取'
         countdown.value = 5
         disableBtn.value = false
         codeInputDisabled.value = false //使验证码输入框变为可输入状态
       } else {
         countdown.value--
         btnText.value = `${countdown.value}s 后重新获取`
       }
     }, 1000)
     //调用验证码的接口
     requestVerificationCode({
       params:{
         phone:phoneNumber.value
       }
     }).then(res=>{
       if(res.data.code==0){
         uni.$showMsg('验证码获取成功')
       }
     })
   }
}
let handleSubmit = ()=>{
  //如果手机号与验证码为空时，则返回false
  if(phoneNumber.value==''&&verificationCode.value!== '666666'){
    uni.$showMsg('请输入手机号和验证码')
    return false
  }else{
    //获取用户信息
    uni.getUserProfile({
      desc:'用户信息',
      success(res){
        useStore.updateUserInfo(res.userInfo)
        userInfo = res.userInfo
        uni.login({
          provider:'weixin',
          success(res) {
            let regisCode = res.code
            //整理参数
            let params = {
              code:regisCode,
              phone:phoneNumber.value,
              verifyCode:verificationCode.value,
              nickname:userInfo.nickName,
              avatarUrl:userInfo.avatarUrl,
              gender:userInfo.gender
            }
            //发起注册请求
            requestRegister(params).then(res=>{
              if(res.data.code==0){
                //存储token
                useStore.updateToken(res.data.data.token)
                 uni.$showMsg('注册成功')
                //跳转至首页
                uni.switchTab({
                  url:'/pages/my/my'
                })
              }else{
                 uni.$showMsg(res.data.message)
              }
            }).catch(err => {
              console.error(err)
              uni.$showMsg('服务器出错，请稍后再试')
              })
          }
        })
      }
    })
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
  font-size: 30rpx;
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
