<template>
  <div>
     <!-- 面包屑 -->
     <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>监控大盘</el-breadcrumb-item>
    </el-breadcrumb>
    <Card></Card>
    <Sale></Sale>
  </div>
</template>

<script>
import Card from './card'
import Sale from './sale'
export default {
  name: 'dashboard',
  components: {
    Card,
    Sale
  }
}
</script>

<style></style>
