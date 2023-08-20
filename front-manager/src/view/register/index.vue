<template>
  <div class="register-page">
    <div class="register-wrag">
      <div class="logo">
        <img class="logo-img" src="@/assets/images/Groundhog.jpg" alt="" />
      </div>
    </div>
    <img class="right" src="@/assets/images/right.png" />
    <div class="register-rect">
      <el-form label-width="80px" ref="formRef" :model="registerParams" :rules="rules">
        <h4>{{$t('register.welcome')}}</h4>
        <el-form-item :label="$t('register.username')" prop="adminUsername">
          <el-input :prefix-icon="User" :placeholder="$t('register.username')" v-model="registerParams.adminUsername"></el-input>
        </el-form-item>
        <el-row>
          <el-col :span="13">
            <el-form-item :label="$t('register.validCode')" prop="validCode">
              <el-input :placeholder="$t('register.validCode')" v-model="registerParams.validCode"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item prop="empty" style="margin-bottom: 0">
              <el-button type="primary" size="small">{{$t('register.getCode')}}</el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item :label="$t('register.password')" prop="adminPassword">
          <el-input type="password" :placeholder="$t('register.password')" :prefix-icon="Lock" show-password v-model="registerParams.adminPassword"></el-input>
        </el-form-item>
        <h3>{{$t('register.companyInformation')}}</h3>
        <el-form-item :label="$t('register.companyName')" prop="name">
          <el-input :prefix-icon="House" :placeholder="$t('register.companyName')" v-model="registerParams.name"></el-input>
        </el-form-item>
        <el-form-item :label="$t('register.website')" prop="domain">
          <el-input :prefix-icon="Link" :placeholder="$t('register.website')" v-model="registerParams.domain"></el-input>
        </el-form-item>
      </el-form>
      <el-button class="register_btn" type="primary" size="default" @click="register">{{$t('register.registerButton')}}</el-button>
      <div class="login">{{$t('register.existingAccount')}}? <span @click="goLogin">{{$t('register.login')}}</span></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { User, Lock, House, Link } from '@element-plus/icons-vue'
import { reactive,ref } from 'vue'
import { reqRegister } from '@/api/acl/tenant/index'
import type {registerFormData} from '@/api/acl/tenant/type'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
//获取路由器
let $router = useRouter()
let formRef = ref()
//收集注册信息
let registerParams = reactive<registerFormData>({
  //企业名称
  name: '',
  //网站
  domain: "",
  //用户名
  adminUsername: '',
  //密码
  adminPassword: '',
  //验证码
  validCode: ''
})
//获取注册的数据
let register = async () => {
  //校验通过再发请求
  await formRef.value.validate()
  let res: any = await reqRegister(registerParams)
  if (res.code == 0) {
    $router.push('/login')
    ElMessage({ type: 'success', message: '注册成功' })
  } else {
    ElMessage({ type: 'error', message: res.message })
  }
}
//去登录
let goLogin = () => {
  $router.push('/login')
}
//网站
let validateWebsite = (_rule: any, value: any, callback: any) => {
  const pattern = /^([A-Za-z0-9-]+\.){1,2}[A-Za-z]{2,}$/
  if (!value || pattern.test(value)) {
    callback()
  } else {
    callback(new Error('请输入有效的公司网站,如www.example.com'))
  }
}
//验证码
let validateCode = (_rule: any, value: any, callback: any) => {
  if (value == '') {
    callback(new Error('请输入验证码'))
  } else if (value !== '666666') {
    callback(new Error('验证码为666666'))
  } else {
    callback()
  }
}
//密码
let validatePass = (_rule:any, value:any, callback:any) => {
  if (value === '') {
    callback(new Error('请输入密码'));
  } else if (value.length < 5) {
    callback(new Error('密码长度至少五位'));
  } else{
    callback();
  }
};
//校验规则
let rules = {
  name: [{required: true, message: '请输入公司名称', trigger: 'blur'}],
  domain: [
    { required: true, message: '请输入公司网站', trigger: 'blur' },
    { validator: validateWebsite, trigger: 'blur' }
  ],
  adminUsername: [{required: true,message:'请输入用户名', trigger: 'blur'}],
  adminPassword: [ { required: true, trigger: 'blur',validator: validatePass },],
  validCode: [{ required: true, trigger: 'blur', validator: validateCode }]
}
</script>

<style scoped lang="scss">
.register-page {
  width: 100vw;
  height: 100vh;
  background: url('@/assets/images/rect.png') no-repeat;
  background-size: 100%;
  padding-top: 40px;
  .register-wrag {
    width: 100%;
    min-width: 320px;
    margin-bottom: 30px;
    .logo {
      text-align: center;
      .logo-img {
        max-width: 145px;
        height: 50px;
        cursor: pointer;
      }
    }
  }
  .right{
    max-width: 320px;
    // right: -100px;
    top: -60px;
    position: absolute;
    opacity: .4;
  }
  .register-rect {
    border: 1.5px solid #fff;
    border-radius: 8px;
    background: hsla(0, 0%, 100%, 0.5);
    width: 400px;
    margin: 0 auto;
    padding: 20px;
    h4 {
      height: 33px;
      margin: 0 0 24px;
      font-size: 20px;
      font-weight: 600;
      line-height: normal;
    }
    h3 {
      margin: 0 0 15px;
      color: rgba(0, 0, 0, 0.85);
      font-weight: 600;
    }
    .register_btn {
      width: 100%;
    }
    .login {
      margin: 24px 0 0;
      color: #8592a5;
      font-size: 14px;
      text-align: center;
      span {
        color: #06f;
        cursor: pointer;
        font-weight: 900;
      }
    }
  }
}
</style>
