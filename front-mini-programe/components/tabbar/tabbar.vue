<template>
  <view>
    <u-tabbar :active="active" :fixed="true" :placeholder="true" :safeAreaInsetBottom="true">
      <u-tabbar-item v-for="(item, index) in tabbarItems" :key="index" :icon="getTabbarIcon(item, index)" :text="item.text" @click="handleTabbarItemClick(item, index)"></u-tabbar-item>
    </u-tabbar>
  </view>
</template>

<script setup>
import { ref } from 'vue'
//当前匹配项，默认设置为2，即我的页面
let active = ref(2)

const tabbarItems = [
  {
    pagePath: '/pages/charge/charge',
    text: '充电',
    iconPath: '/static/tab_icons/home.png',
    selectedIconPath: '/static/tab_icons/home-active.png'
  },
  {
    pagePath: '/pages/scan/scan',
    text: '扫码',
    iconPath: '/static/tab_icons/scan-active.png'
  },
  {
    pagePath: '/pages/my/my',
    text: '我的',
    iconPath: '/static/tab_icons/my.png',
    selectedIconPath: '/static/tab_icons/my-active.png'
  }
]
//点击tabbar
const handleTabbarItemClick = (item, index) => {
  if (active.value !== index) {
    if (index === 1) {
      uni.scanCode({
        success(res) {
          const scanUrl = res.result
          const regExp = /pileId=(\d+)&code=(\d+)/
          const matchResult = scanUrl.match(regExp)
          let pileId = ''
          let code = ''
          if (matchResult) {
            pileId = matchResult[1]
            code = matchResult[2]
          }
          uni.reLaunch({
            url: '/subpkg/charging/charging?pileId=' + pileId + '&code=' + code
          })
        }
      })
    } else {
      active.value = index // 更新选中项的索引
      const path = item.pagePath
      uni.switchTab({
        url: path
      })
    }
  }
}
//图标的切换
const getTabbarIcon = (item, index) => {
  return active.value === index ? item.selectedIconPath : item.iconPath
}
</script>
