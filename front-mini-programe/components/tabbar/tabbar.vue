<template>
  <view>
    <u-tabbar :active="useStore.activeTab" :fixed="true" :placeholder="true" :safeAreaInsetBottom="true">
      <u-tabbar-item v-for="(item, index) in tabbarItems" :key="index" :icon="getTabbarIcon(item, index)" :text="item.text" @click="handleTabbarItemClick(item, index)"></u-tabbar-item>
    </u-tabbar>
  </view>
</template>

<script setup>
//引入pinia仓库
import useUserStore from '@/store/user.js'
let useStore = useUserStore()

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
    iconPath: '/static/tab_icons/scan.png',
    selectedIconPath: '/static/tab_icons/scan-active.png'
  },
  {
    pagePath: '/pages/my/my',
    text: '我的',
    iconPath: '/static/tab_icons/my.png',
    selectedIconPath: '/static/tab_icons/my-active.png'
  }
]
//点击tabbar按钮
const handleTabbarItemClick = (item, index) => {
  if (useStore.activeTab !== index) {
    //如果点击的是扫描按钮
    if (index === 1) {
      uni.scanCode({
        success(res) {
          const scanUrl = res.result
          const regExp = /pileId=(\d+)&code=(\d+)&stationName=([^&]+)/;
          const matchResult = scanUrl.match(regExp)
          let pileId = ''
          let code = ''
          let stationName = ''
          if (matchResult) {
            pileId = matchResult[1]
            code = matchResult[2]
            stationName = matchResult[3]
          }
          uni.reLaunch({
            url: '/subpkg/charging/charging?pileId=' + pileId + '&code=' + code+'&stationName='+stationName
          })
        }
      })
    } else {
      useStore.setActive(index)
      const path = item.pagePath
      uni.switchTab({
        url: path
      })
    }
  }
}
//图标的切换
const getTabbarIcon = (item, index) => {
  return useStore.activeTab === index ? item.selectedIconPath : item.iconPath
}
</script>
