<template>
  <el-container>
    <!-- 头部区域 -->
    <el-header>
      <h2>土拔鼠充电管理平台</h2>
      <el-button type="info" @click="out">退出登录</el-button>
    </el-header>
    <!-- 主体区域 -->
    <el-container>
      <el-aside :width="isCollapse ? '64px' : '200px'">
        <div class="toggle-button" @click="toggleCollapse">|||</div>
        <el-menu class="el-menu-vertical-demo" background-color="#333744" text-color="#fff" active-text-color="#409EFF" unique-opened :collapse="isCollapse" :collapse-transition="false" router>
          <el-submenu index="1">
            <template slot="title">
              <i class="el-icon-s-home"></i>
              <span>系统首页</span>
            </template>
            <el-menu-item :index="'/home'">
              <template slot="title">
                <i class="el-icon-s-operation"></i>
                <span>监控大盘</span>
              </template>
            </el-menu-item>
            <el-menu-item :index="'/home/test'">
              <template slot="title">
                <i class="el-icon-s-operation"></i>
                <span>测试</span>
              </template>
            </el-menu-item>
          </el-submenu>
          <el-submenu index="2">
            <template slot="title">
              <i class="el-icon-s-cooperation"></i>
              <span>设备管理</span>
            </template>
            <el-menu-item :index="'/home/charge'">
              <template slot="title">
                <i class="el-icon-attract"></i>
                <span>充电站管理</span>
              </template>
            </el-menu-item>
            <el-menu-item :index="'/home/pile'">
              <template slot="title">
                <i class="el-icon-coordinate"></i>
                <span>充电桩管理</span>
              </template>
            </el-menu-item>
          </el-submenu>
          <el-submenu index="3">
            <template slot="title">
              <i class="el-icon-s-order"></i>
              <span>订单管理</span>
            </template>
            <el-menu-item :index="'/home/chargeOrder'">
              <template slot="title">
                <i class="el-icon-shopping-bag-1"></i>
                <span>充电订单</span>
              </template>
            </el-menu-item>
          </el-submenu>
          <el-submenu index="4">
            <template slot="title">
              <i class="el-icon-money"></i>
              <span>价格规则管理</span>
            </template>
            <el-menu-item :index="'/home/priceList'">
              <template slot="title">
                <i class="el-icon-price-tag"></i>
                <span>区间充电价格</span>
              </template>
            </el-menu-item>
          </el-submenu>
          <el-submenu index="5">
            <template slot="title">
              <i class="el-icon-s-cooperation"></i>
              <span>运营管理</span>
            </template>
            <el-menu-item :index="'/home/userList'">
              <template slot="title">
                <i class="el-icon-user"></i>
                <span>用户中心页面</span>
              </template>
            </el-menu-item>
            <el-menu-item :index="'/home/record'">
              <template slot="title">
                <i class="el-icon-s-finance"></i>
                <span>充值页面</span>
              </template>
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      // 是否折叠
      isCollapse: false
    }
  },
  methods: {
    //退出登录
    out() {
      // 清除token
      localStorage.clear()
      // 路由跳转到login
      this.$router.push('/login')
    },
    //折叠的回调
    toggleCollapse() {
      this.isCollapse = !this.isCollapse
    }
  }
}
</script>

<style lang="less" scoped>
.el-container {
  height: 100%;
}
.el-header {
  background-color: #373d41;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 0;
  color: #fff;
  > h2 {
    margin-left: 15px;
  }
}
.el-aside {
  background-color: #333744;
}
.el-main {
  background-color: #eaedf1;
}
.toggle-button {
  background-color: #4a5064;
  text-align: center;
  color: #fff;
  font-size: 12px;
  line-height: 24px;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>
