<template>
  <div class="login-page">
    <div class="login-wrag">
      <div class="logo">
        <img class="logo-img" src="@/assets/images/Groundhog.jpg" alt="" />
      </div>
    </div>
    <img class="left" src="@/assets/images/left.png" />
    <div class="login-rect">
      <el-form
        ref="loginForms"
        :model="loginForm"
        :rules="rules"
        label-width="65px"
      >
        <h1 class="title">
          Hello
          <span>
            {{ $t('login.account') }}？
            <span class="register" @click="goRegister">
              {{ $t('login.register') }}
            </span>
          </span>
        </h1>
        <h2>{{ $t('login.welcome') }}</h2>
        <el-form-item :label="$t('login.tenement')">
          <el-select
            v-model="selectedTenantId"
            @click="teSelect"
            @change="getSelectedTenantName"
            :placeholder="$t('login.tenements')"
          >
            <el-option
              v-for="item in useStore.teArr"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="username" :label="$t('login.accounts')">
          <el-input
            :prefix-icon="User"
            v-model="loginForm.username"
            :placeholder="$t('login.usernamePlaceholder')"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password" :label="$t('login.password')">
          <el-input
            type="password"
            :prefix-icon="Lock"
            show-password
            v-model="loginForm.password"
            :placeholder="$t('login.passwordPlaceholder')"
          ></el-input>
        </el-form-item>
        <el-form-item :label="$t('login.code')">
          <el-input
            :placeholder="$t('login.verificationCode')"
            v-model="sidentifyMode"
          ></el-input>
        </el-form-item>
        <div
          @click="refreshCode"
          style="text-align: center; margin-bottom: 10px"
        >
          <SIdentify :identifyCode="identifyCode"></SIdentify>
        </div>
        <el-button
          class="login_btn"
          type="primary"
          size="default"
          @click="login"
        >
          {{ $t('login.text') }}
        </el-button>
      </el-form>
    </div>
    <el-dropdown class="language-rect">
      <span class="el-dropdown-link">
        <svg-icon name="language" width="30px" height="20px">1</svg-icon>
        <span>{{ $t('login.language') }}</span>
        <el-icon class="el-icon--right">
          <arrow-down />
        </el-icon>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item @click="changeZh">
            {{ $t('login.chinese') }}
          </el-dropdown-item>
          <el-dropdown-item @click="changeEn">
            {{ $t('login.english') }}
          </el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
</template>

<script setup lang="ts">
import { User, Lock } from '@element-plus/icons-vue'
import { reactive, ref, getCurrentInstance, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElNotification } from 'element-plus'
//引入获取当前时间的函数
import { getTime } from '@/utils/time'
//引入用户相关的小仓库
import useUserStore from '@/store/module/use'
let useStore = useUserStore()
let sidentifyMode = ref('') //输入框验证码
let identifyCode = ref('')
let identifyCodes = ref('1234567890abcdefjhijklinopqrsduvwxyz')
//收集账号与密码的数据
let loginForm = reactive({ username: 'admin', password: 'admin' })
//获取路由器
let $router = useRouter()
//路由对象
let $route = useRoute()
//获取el-form组件
let loginForms = ref()
//保存下拉框选择的租户id
let selectedTenantId = ref('')
//存储选中租户的名字
let selectedTenantName = ref('')
//点击下拉框，通知仓库获取租户信息
let teSelect = () => {
  useStore.tenantSelect()
}
//根据选中的id获取租户的名字
let getSelectedTenantName = () => {
  let selectedTenant = useStore.teArr.find(
    (item: any) => item.id == selectedTenantId.value
  )
  if (selectedTenant) {
    selectedTenantName.value = selectedTenant.name
  }
}
//登录按钮的回调函数
let login = async () => {
  // 保存选中的租户id到仓库
  useStore.setTenantId(selectedTenantId.value)
  //保存选中的租户名字到仓库
  useStore.setTenantName(selectedTenantName.value)
  await loginForms.value.validate()
  //验证验证码不为空
  if (!sidentifyMode.value) {
    ElNotification({ type: 'error', message: '验证码不能为空！' })
    return
  }
  //验证验证码是否正确
  if (sidentifyMode.value != identifyCode.value) {
    ElNotification({ type: 'error', message: '验证码错误' })
    return
  }
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
      title: `HI,${getTime()}`
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
//getCurrentInstance函数获取当前组件实例,并通过.appContext.config.globalProperties访问全局属性
//?号表示前面如果返回null，后面不执行
let $this = getCurrentInstance()?.appContext.config.globalProperties as any

//切换中文
const changeZh = () => {
  $this.$i18n.locale = 'zh'
  localStorage.setItem('lang', 'zh')
}
//切换英文
const changeEn = () => {
  $this.$i18n.locale = 'en'
  localStorage.setItem('lang', 'en')
}
//组件挂载
onMounted(() => {
  identifyCode.value = ''
  makeCode(identifyCodes.value, 4)
})
// 生成随机数
const randomNum = (min: any, max: any) => {
  max = max + 1
  return Math.floor(Math.random() * (max - min) + min)
}
// 更新验证码
const refreshCode = () => {
  identifyCode.value = ''
  makeCode(identifyCodes.value, 4)
}
// 随机生成验证码字符串

const makeCode = (o: any, l: any) => {
  for (let i = 0; i < l; i++) {
    identifyCode.value += o[randomNum(0, o.length)]
  }
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
  .language-rect {
    position: absolute;
    width: 150px;
    color: #202d40;
    cursor: pointer;
    outline: none;
    left: 50%;
    margin-left: -40px;
    margin-top: 50px;
    text-align: center;
    .el-dropdown-link {
      display: flex;
      align-items: center;
    }
  }
}
</style>
