<template>
  <div class="login-page">
    <div class="login-wrag">
      <div class="logo">
        <img class="logo-img" src="@/assets/images/Groundhog.jpg" alt="" />
      </div>
    </div>
    <img class="left" src="@/assets/images/left.png" />
    <div class="login-rect">
      <el-form class="login_form" ref="loginForms" :model="loginForm" :rules="rules">
        <h1 class="title">
          Hello
          <span>没有账号？<span class="register" @click="goRegister">去注册</span></span>
        </h1>
        <h2>欢迎来到土拨鼠充电平台</h2>
        <el-form-item label="选择租户">
          <el-select v-model="selectedTenantId" @click="teSelect" placeholder="请选择租户">
            <el-option v-for="item in useStore.teArr" :key="item.id" :label="item.name" :value="item.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="username">
          <el-input :prefix-icon="User" v-model="loginForm.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input type="password" :prefix-icon="Lock" show-password v-model="loginForm.password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="login_btn" type="primary" size="default" @click="login">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { User, Lock } from '@element-plus/icons-vue'
import { reactive, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElNotification } from 'element-plus'
//引入获取当前时间的函数
import { getTime } from '@/utils/time'
//引入用户相关的小仓库
import useUserStore from '@/store/module/use'
let useStore = useUserStore()
//收集账号与密码的数据
let loginForm = reactive({ username: '', password: '' })
//获取路由器
let $router = useRouter()
//路由对象
let $route = useRoute()
//获取el-form组件
let loginForms = ref()
//保存下拉框选择的租户id
let selectedTenantId = ref('')
//点击下拉框，通知仓库获取租户信息
let teSelect = () => {
  useStore.tenantSelect()
}
//登录按钮的回调函数
let login = async () => {
   // 保存选中的租户id到仓库
   useStore.setTenantId(selectedTenantId.value)
  await loginForms.value.validate()
    try {
      //保证登录成功 useStore.userLogin(loginForm)打印的是promise
      await useStore.userLogin(loginForm)
      //用户登录获得当前权限
      await useStore.setCurrentPermission()
      //保存用户名到仓库
      useStore.setUsername(loginForm.username)
      //编程式导航跳转到展示数据首页
      //判断登录的时候,路由路径当中是否有query参数，如果有就往query参数挑战，没有跳转到首页
      let redirect: any = $route.query.redirect
      $router.push({ path: redirect || '/' })
      //登录成功提示信息
      ElNotification({
        type: 'success',
        message: '欢迎回来',
        title: `HI,${getTime()}好`
      })
    } catch (error) {
      ElNotification({
        type: 'error',
        message: (error as Error).message
      })
    }
}

//定义自定义校验规则
let validatePass = (_rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请输入密码'))
  } else if (value.length < 5) {
    callback(new Error('密码长度至少五位'))
  } else {
    callback()
  }
}

//校验规则
let rules = {
  username: [{ required: true, trigger: 'blur', message: '请输入用户名' }],
  password: [{ required: true, trigger: 'blur', validator: validatePass }]
}
//注册按钮
let goRegister = () => {
  $router.push('/register')
}
</script>

<style scoped lang="scss">
.login-page {
  width: 100vw;
  height: 100vh;
  background: url('@/assets/images/rect.png');
  background-size: 100%;
  padding-top: 40px;
  .login-wrag {
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
  .left {
    max-width: 568px;
    left: -400px;
    top: 240px;
    position: absolute;
    opacity: 0.3;
  }
  .login-rect {
    border: 1.5px solid #fff;
    border-radius: 8px;
    background: hsla(0, 0%, 100%, 0.5);
    width: 400px;
    margin: 0 auto;
    padding: 20px;
    .title {
      color: rgba(0, 0, 0, 0.85);
      font-size: 25px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      font-weight: 600;
      span {
        color: rgba(0, 0, 0, 0.85);
        font-size: 15px;
        font-weight: 400;
        .register {
          color: #06f;
          cursor: pointer;
          font-weight: 900;
        }
      }
    }
    h2 {
      font-size: 15px;
      color: rgba(0, 0, 0, 0.85);
      margin: 10px 0px;
      font-weight: 600;
    }
    .login_btn {
      width: 100%;
    }
  }
}
</style>
