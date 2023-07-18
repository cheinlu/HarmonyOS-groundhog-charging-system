<template>
  <el-button size="small" icon="Refresh" circle @click="updateRefsh" />
  <el-button size="small" icon="FullScreen" circle @click="fullScreen"/>

  <el-popover placement="bottom" title="主题设置" :width="300" trigger="hover">
    <el-form>
      <el-form-item label="主题颜色">
        <el-color-picker @change="setColor" v-model="color" show-alpha :predefine="predefineColors" />
      </el-form-item>
      <el-form-item label="暗黑模式">
        <el-switch @change="changeDark" v-model="dark" class="mt-2" style="margin-left: 24px" inline-prompt active-icon="MoonNight" inactive-icon="Sunny" />
      </el-form-item>
    </el-form>
    <template #reference>
      <el-button size="small" icon="Setting" circle />
    </template>
  </el-popover>
  <img src="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif" style="width: 24px; height: 24px; margin: 0px 10px; border-radius: 50%" />
  <!-- 下拉菜单 -->
  <el-dropdown>
    <span class="el-dropdown-link">
      <span>{{ userStore.username }}</span>
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter,useRoute } from 'vue-router'
//获取骨架的小仓库
import useLayOutSettingStore from '@/store/module/setting.ts'
//获取用户名的仓库
import useUserStore from '@/store/module/use.ts'
let userStore = useUserStore()
let LayOutSettingStore = useLayOutSettingStore()
//获取路由信息
let $router = useRouter()
//获取路由对象
let $route = useRoute()
//刷新按钮点击回调
let updateRefsh = () => { 
  LayOutSettingStore.refsh = !LayOutSettingStore.refsh
}
//全屏按钮点击的回调
let fullScreen = () => {
  //DOM对象的一个属性:可以用来判断当前是不是全屏模式[全屏:true,不是全屏:false]
  let full = document.fullscreenElement
  //切换为全屏模式
  if (!full) {
    //文档根节点的方法requestFullscreen,实现全屏模式
    document.documentElement.requestFullscreen()
  } else {
    //变为不是全屏模式->退出全屏模式
    document.exitFullscreen()
  }
}
//退出登录的回调
let logout = ()=>{
  userStore.userLogout()
  //跳转到登录页面
  $router.push({path:'/login',query:{redirect:$route.path}})
}
//收集开关的数据
let dark = ref<boolean>(false)
//取色
const color = ref('rgba(255, 69, 0, 0.68)')
const predefineColors = ref(['#ff4500', '#ff8c00', '#ffd700', '#90ee90', '#00ced1', '#1e90ff', '#c71585', 'rgba(255, 69, 0, 0.68)', 'rgb(255, 120, 0)', 'hsv(51, 100, 98)', 'hsva(120, 40, 94, 0.5)', 'hsl(181, 100%, 37%)', 'hsla(209, 100%, 56%, 0.73)', '#c7158577'])
//暗黑模式的change事件
let changeDark = () => {
  //获取HTML根节点
  let html = document.documentElement
  //判断HTML标签是否有类名dark
  dark.value ? (html.className = 'dark') : (html.className = '')
}
//更改主题颜色
let setColor = ()=>{
  //获取HTML根节点
  let html = document.documentElement
  //设置css变量
  html.style.setProperty('--el-color-primary', color.value)
}

</script>

<style scoped></style>
