<template>
  <div>
    <router-view />
  </div>
</template>

<script setup lang="ts">
import { onBeforeMount, onMounted } from 'vue';
import { useRouter } from 'vue-router';

//获取用户相关的小仓库
import useUserStore from '@/store/module/use'
let userStore = useUserStore()

const router = useRouter();

onBeforeMount(()=>{
  userStore.setCurrentPermission()
})

onMounted(() => {
  // 刷新页面时，手动触发路由跳转
  router.replace(router.currentRoute.value.fullPath);
});
</script>

<style scoped>

</style>
