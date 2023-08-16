<template>
  <view>
    <tabbar></tabbar>
  </view>
</template>

<script setup>
import {onShow} from '@dcloudio/uni-app'
onShow(()=>{
  uni.hideTabBar({
    animation:false
  })
})
</script>

<style lang="scss" scoped>

</style>