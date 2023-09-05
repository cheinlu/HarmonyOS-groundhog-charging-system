<template>
  <view>
    <login v-if="!useStore.token"></login>
    <info v-else></info>
    <tabbar></tabbar>
  </view>
</template>

<script setup>
import { onShow } from '@dcloudio/uni-app'
import useUserStore from '@/store/user.js'
let useStore = useUserStore()
onShow(() => {
  uni.hideTabBar({
    animation: false
  })
  useStore.activeTab = 2
})


</script>

<style lang="scss"></style>
