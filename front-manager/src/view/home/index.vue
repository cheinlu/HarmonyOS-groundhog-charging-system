<template>
  <el-card>
    <div class="box">
      <img src="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif" alt="" class="avatar">
      <div class="bottom">
        <h3 class="title">{{ getTime() }}~ {{ userStore.username }}</h3>
        <p class="subtitle">{{$t('homes.welcomeMessage')}}--{{ userStore.tenantName }}</p>
      </div>
    </div>
  </el-card>
  <div class="bottoms">
    <img src="@/assets/images/charge.jpg" alt="">
  </div>
</template>

<script setup lang="ts">
import { getTime } from '@/utils/time';
//引入用户相关的仓库,获取当前用户的头像、昵称
import useUserStore from '@/store/module/use';
//获取存储用户信息的仓库对象
let userStore = useUserStore();

</script>

<style scoped lang="scss">
.box {
  display: flex;

  .avatar {
    width: 100px;
    height: 100px;
    border-radius: 50%;
  }

  .bottom {
    margin-left: 20px;

    .title {
      font-size: 30px;
      font-weight: 900;
      margin-bottom: 30px;
    }

    .subtitle {
      font-style: italic;
      color: #64D9D6;
      font-weight: 900;
      font-size: 20px;
    }
  }
}
.bottoms{
  margin-top:30px;
  margin-left: 25%;
  width: 550px;
  height: 550px;
  img{
    width: 100%;
    height: 100%;
    border-radius: 20px;
  }
}
</style>