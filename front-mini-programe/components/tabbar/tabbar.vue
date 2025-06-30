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
    iconPath: '/static/tab_icons/scan-active.png',
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
	console.log('lucy== token',useStore.token);
  if (useStore.activeTab !== index) {
    //如果点击的是扫描按钮
	// pileId=2&code=02&stationName=宝安新安充电站 草料二维码，选中文本输入左侧内容即可，具体参数根据实际场景替换
    if (index === 1) {
		if(useStore.token){
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
		}else{
			uni.showModal({
				title:'未登录',
				content: '请先登录后再进行扫码操作',
				confirmText: '去登录',
				 success: (res) => {
					 if (res.confirm) {
						 useStore.setActive(2)
						 uni.switchTab({
						 	url:'/pages/my/my'
						 })
					 }
				 }
			})
		}
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
