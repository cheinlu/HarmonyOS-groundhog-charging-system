<template>
  <!-- 顶部左侧静态 -->
  <el-icon style="margin-right: 10px" @click="changeIcon">
    <component :is="LayOutSettingStore.fold ? 'Fold' : 'Expand'"></component>
  </el-icon>
  <!-- 左侧面包屑 -->
  <el-breadcrumb separator-icon="ArrowRight">
    <el-breadcrumb-item v-for="(item, index) in $route.matched" :key="index" v-show="item.meta.title" :to="item.path">
      <el-icon>
        <component :is="item.meta.icon"></component>
      </el-icon>
      <span>{{ item.meta.title }}</span>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts">
//引入路由
import { useRoute } from 'vue-router'
import useLayOutSettingStore from '../../../store/module/setting'
let LayOutSettingStore = useLayOutSettingStore()
let $route = useRoute()
//点击图标的方法
const changeIcon = () => {
  //图标进行切换
  LayOutSettingStore.fold = !LayOutSettingStore.fold
}
</script>

<style scoped>
.el-breadcrumb__inner a, .el-breadcrumb__inner.is-link{
  display: flex;
}
</style>
