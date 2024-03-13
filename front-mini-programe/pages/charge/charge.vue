<template>
  <view class="home-container">
    <!-- 加载中 -->
    <view v-if="loading" class="loading">
      <image src="./images/loding.gif" class="loading-icon"></image>
      <text class="loading-text">加载中...</text>
    </view>
    <!-- 搜索区域 -->
    <view class="search-box">
      <uni-search-bar @input="handleInput" :radius="100" cancelButton="none"></uni-search-bar>
    </view>
    <scroll-view scroll-y="true" :style="{ height: windowHeight + 'px' }" @scrolltolower="loadNextPage" :scroll-top="scrollTop" @scrolltoupper="refreshData" ref="scrollView">
      <view class="scroll-item" v-for="station in stationList" :key="station.id">
        <view class="pile-rect" @click="handleClickPile(station)">
          <view class="pile-title">{{ station.name }}</view>
          <view class="pile-box">
            <view class="price">
              <view class="pri"> ￥ {{ useStore.token?station.pricePerHour:'***' }} </view>
              <view class="limit"> /度 </view>
            </view>
            <view class="pile-tanan">{{ station.tenantName }}</view>
            <view class="resi">
              <image class="fast" src="./images/fast.png" mode=""></image>
              <view class="free"> 闲{{ station.freePileTotal }} </view>
              <view class="total"> /{{ station.pileTotal }}</view>
            </view>
          </view>
        </view>
        <view class="pile-distance">
          <view class="distance">{{ station.distance.toFixed(2) }}km</view>
          <view class="navi" @click="showLocation(station)">
            <image class="location" src="./images/location.png" mode=""></image>
          </view>
          <view class="navi" @click="showImg(station)">
            <image class="location" src="./images/imgUrl.png" mode=""></image>
          </view>
        </view>
      </view>
    </scroll-view>

    <!-- 弹窗 -->
    <popPile :show="show" @confirm="confimCharge" :list="pileList" :pricePerHour="pricePerHour"></popPile>
    <!-- 导航栏 -->
    <tabbar></tabbar>
  </view>
</template>

<script setup>
import { onShow } from '@dcloudio/uni-app'
import { ref, onMounted, reactive } from 'vue'
import { requestStations, requestPiles } from '@/utils/api/charge.js'
import useUserStore from '@/store/user.js'
let useStore = useUserStore()
let loading = ref(false)
let show = ref(false)
let pricePerHour = ref(null)
let kw = ref('')
let timer = ref(null)
let windowHeight = ref(0)
let scrollTop = ref(0)
let isShowPopup = ref(false)
let stationInfo = reactive({
  name: '',
  address: '',
  coordinate: ''
})
let stationList = ref([])
let pageSize = ref(10)
let pageNo = ref(1)
let total = ref(0)
let pileList = ref([])

onShow(() => {
  uni.hideTabBar({
    animation: false
  })
  useStore.activeTab = 0
})
//组件挂载
onMounted(() => {
  let sysInfo = uni.getSystemInfoSync()
  windowHeight.value = sysInfo.windowHeight - 50
  //获取充电站信息
  getStationList()
})
//搜索站点
let handleInput = (keyword) => {
  if (kw.value !== keyword) {
    kw.value = keyword
    //清除timer对应的定时器
    clearTimeout(timer.value)
    //重启一个延时器，并把timerid赋值给timer
    timer.value = setTimeout(() => {
      // 对关键词进行解码
      let decodedKeyword = decodeURIComponent(keyword)
      //根据关键词，查询搜索建议列表
      getStationList({ params: { name: decodedKeyword } })
    }, 600)
  }
}
//获取充电站数据
let getStationList = () => {
  // 显示加载中状态
  loading.value = true
  //请求获取地理位置权限
  uni.authorize({
    scope: 'scope.userLocation',
    success() {
      stationList.value = []
      //授权成功，获取位置
      uni.getLocation({
        type: 'wgs84',
        timeout: 10000, // 设置10秒超时
        success: function (res) {
          //成功获取经纬度，拼接后给到接口里的位置字段
          stationInfo.coordinate = res.longitude + ',' + res.latitude
          getHasStations().then()
        },
        fail: function () {
          getHasStations().then()
          uni.$showMsg('请授权地理位置权限')
          // 再次调用getStationList进行再次授权
          
        }
      })
    }
  })
}
//封装请求获取充电站数据的方法
let getHasStations = () => {
  return new Promise((resolve, reject) => {
    //整理参数
    let params = {
      name: kw.value,
      address: stationInfo.address,
      pageNo: pageNo.value,
      pageSize: pageSize.value,
      coordinate: stationInfo.coordinate
    }
    //调取充电站的接口
    requestStations(params)
      .then((res) => {
        if (res.data.code == 0) {
          stationList.value = stationList.value.concat(res.data.data.List)
          pageNo.value = res.data.data.PageNo
          pageSize.value = res.data.data.PageSize
          total.value = res.data.data.TotalCount
          resolve(res.data)
        } else {
          reject(res.data)
        }
      })
      .finally(() => {
        // 取消加载中状态
        loading.value = false
      })
  })
}
//点击每个充电站的回调
let handleClickPile = async (station) => {
  //打开弹窗
  show.value = !show.value
  pricePerHour.value = station.pricePerHour
  let { data: res } = await requestPiles(station.id)
  if (res.code == 0) {
    pileList.value = res.data.List
    pageSize.value = res.data.PageSize
    pageNo.value = res.data.PageNo
    total.value = res.data.TotalCount
  } else {
    uni.$showMsg('res.message')
  }
}
//显示位置
let showLocation = (station) => {
  //整理参数
  const str = station.coordinate
  //获得经纬度
  const [lng, lat] = str.split(',').map(parseFloat)
  console.log('lng',lng,'lat',lat);
  //查看位置
  uni.openLocation({
    latitude: lat, //纬度 - 目的地/坐标点
    longitude: lng //经度 - 目的地/坐标点
  })
}
//显示充电站预览图
let showImg = (station) => {
  if (station.imageUrl == '') {
    uni.$showMsg('图片未上传')
  } else {
    uni.previewImage({
      urls: [station.imageUrl],
      current: station.imageUrl
    })
  }
}
//滚动到底部，加载下一页充电站数据
let loadNextPage = () => {
  // 利用Math.ceil算出新的分页
  let totalPage = Math.ceil(total.value / pageSize.value)
  if (pageNo.value >= totalPage) {
    uni.$showMsg('已经到最底部了')
    return false
  } else {
    pageNo.value++
    getStationList()
  }
  isShowPopup.value = false
}
let refreshData = () => {
  scrollTop.value = 0
  stationList.value = []
  getStationList()
}
//popPile通知charge关闭弹窗
let confimCharge = () => {
  show.value = false
}
</script>

<style lang="scss">
.home-container {
  background-color: #f7f7f7;
  width: 100vw;
  height: 100vh;
  .search-box {
    position: sticky;
    top: 0;
    z-index: 999;
    background-color: #0aa671;
    .mini-btn {
      height: 36px;
      line-height: 36px;
      border-radius: 15px;
      margin: 25rpx;
    }
  }
  .scroll-item {
    background-color: #ffffff;
    padding: 25rpx;
    margin: 20rpx;
    font-size: 20rpx;
    margin-bottom: 20rpx;
    border-radius: 15rpx;
    .pile-tanan {
      font-size: 30rpx;
      text-align: center;
      color: #0aa671;
    }
    .pile-title {
      font-size: 25rpx;
      font-weight: bold;
    }
    .pile-box {
      margin: 10rpx 0;
      display: flex;
      justify-content: space-between;
      .price {
        display: flex;
        align-items: center;
        .pri {
          font-size: 30rpx;
          color: #0aa671;
          margin: 0 10rpx;
        }
        .limit {
          color: #adadb3;
          font-size: 20rpx;
        }
      }
      .resi {
        display: flex;
        font-size: 25rpx;
        .fast {
          width: 30rpx;
          height: 30rpx;
        }
        .free {
          font-weight: bold;
          margin: 0 10rpx;
        }
      }
    }
    .pile-distance {
      display: flex;
      align-items: center;
      .distance {
        font-size: 25rpx;
      }
      .navi {
        margin-left: 20rpx;

        .location {
          width: 40rpx;
          height: 40rpx;
        }
      }
    }
  }
  // loding区域
  .loading {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    display: flex;
    flex-direction: column;
    align-items: center;
    .loading-icon {
      width: 100px;
      height: 100px;
      margin-bottom: 10px;
    }
    .loading-text {
      font-size: 14px;
      color: #999;
    }
  }
}
</style>
