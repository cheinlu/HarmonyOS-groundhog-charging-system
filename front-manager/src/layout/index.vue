<template>
  <div class="layout_container">
    <!-- 左侧菜单 -->
    <div class="layout_slider" :class="{ fold: LayOutSettingStore.fold ? true : false }">
     <Logo></Logo>
      <!-- 滚动菜单 -->
      <el-scrollbar class="scrollbar">
        <el-menu :collapse="LayOutSettingStore.fold ? true : false"
          :default-active="$route.path"
          background-color="#001529"
          text-color="white"
          :class="{ fold: LayOutSettingStore.fold ? true : false }">
          <!--根据路由动态生成菜单-->
          <Menu :menuList="userStore.menuRoutes"></Menu>
        </el-menu>
      </el-scrollbar>
    </div>
    <!-- 顶部导航 -->
    <div class="layout_tabbar" :class="{ fold: LayOutSettingStore.fold ? true : false }">
      <!-- 顶部Tabbar组件 -->
      <Tabbar></Tabbar>
    </div>
    <!-- 内容展示区域 -->
    <div class="layout_main" :class="{ fold: LayOutSettingStore.fold ? true : false }">
      <Main></Main>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineAsyncComponent} from 'vue'
//引入路由
import { useRoute } from 'vue-router'
//使用异步组件的形式引入子组件，会在需要时延迟加载组件的代码，提高加载速度--Logo左侧菜单logo--Menu左侧菜单--Main右侧内容--tabbar顶部
let Logo = defineAsyncComponent(()=>import('./logo/index.vue'))
let Menu = defineAsyncComponent(()=>import('./menu/index.vue'))
let Tabbar = defineAsyncComponent(()=>import('./tabbar/index.vue'))
let Main = defineAsyncComponent(()=>import('./main/index.vue'))

//获取用户相关的小仓库
import useUserStore from '@/store/module/use'
//引入layout组件相关的小仓库
import useLayOutSettingStore from '@/store/module/setting'

let LayOutSettingStore = useLayOutSettingStore()

let userStore = useUserStore()
//获取路由
let $route = useRoute()
</script>

<style scoped lang="scss">
.layout_container {
  width: 100%;
  height: 100vh;

  .layout_slider {
    color: white;
    width: $base-menu-width;
    height: 100vh;
    background: $base-menu-background;
    transition: all 0.3s;

    .scrollbar {
      width: 100%;
      height: calc(100vh - $base-menu-logo-height);

      .el-menu {
        border-right: none;
      }
    }
    &.fold {
      width: $base-menu-min-width;
    }
  }

  .layout_tabbar {
    position: fixed;
    width: calc(100% - $base-menu-width);
    height: $base-tabbar-height;
    top: 0px;
    left: $base-menu-width;
    transition: all 0.3s;
    &.fold {
      width: calc(100vw - $base-menu-min-width);
      left: $base-menu-min-width;
    }
  }

  .layout_main {
    position: absolute;
    width: calc(100% - $base-menu-width);
    height: calc(100vh - $base-tabbar-height);
    left: $base-menu-width;
    top: $base-tabbar-height;
    padding: 20px;
    overflow: auto;
    transition: all 0.3s;

    &.fold {
      width: calc(100vw - $base-menu-min-width);
      left: $base-menu-min-width;
    }
  }
}
</style>

