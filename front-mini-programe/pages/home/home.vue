<template>
  <view class="home-container">
    <!-- 搜索区域 -->
    <view class="search-box">
      <uni-search-bar @input="handleInput" :radius="100" cancelButton="none"></uni-search-bar>
    </view>
    <scroll-view scroll-y="true" :style="{ height: windowHeight + 'px' }" @scrolltolower="loadNextPage" :scroll-top="scrollTop"
@scrolltoupper="refreshData"
ref="scrollView">
      <view class="scroll-item" v-for="station in stationList" :key="station.id">
        <view class="pile-rect" @click="handleClickPile(station)">
          <view class="pile-title">{{ station.name }}</view>
          <view class="pile-box">
            <view class="price">
              <view class="pri"> ￥ {{ station.pricePerHour }} </view>
              <view class="limit"> /度 </view>
            </view>
            <view class="resi">
              <image class="fast" :src="require('./static/fast.png')" mode=""></image>
              <view class="free"> 闲{{ station.freePileTotal }} </view>
              <view class="total"> /{{ station.pileTotal }} </view>
            </view>
          </view>
        </view>
        <view class="pile-distance">
          <view class="distance">{{ station.distance.toFixed(2) }}km</view>
          <view class="navi" @click="showLocation(station)">
            <image class="location" :src="require('./static/location.png')" mode=""></image>
          </view>
        </view>
      </view>
    </scroll-view>

    <!-- 弹窗 -->
    <u-popup :show="isShowPopup" @close="closePopup" closeable>
      <scroll-view scroll-y="true">
        <view>
          <uni-table ref="table" :loading="isLoadingPile" border stripe emptyText="暂无更多数据">
            <!-- 表头行 -->
            <uni-tr>
              <uni-th align="center" width="80">编号</uni-th>
              <uni-th align="center" width="100">状态</uni-th>
              <uni-th align="center" width="100">充电</uni-th>
            </uni-tr>
            <!-- 表格数据行 -->
            <uni-tr v-for="pile in pileList" :key="pile.id">
              <uni-td align="center">{{ pile.code }}</uni-td>
              <uni-td align="center">
                <template v-slot>
                  <button class="mini-btn" size="mini" type="default" v-if="pile.state == 0">空闲中</button>
                  <button class="mini-btn" size="mini" type="primary" v-else-if="pile.state == 1">充电中</button>
                  <button class="mini-btn" size="mini" type="warn" v-else>故障中</button>
                </template>
              </uni-td>
              <uni-td align="center">
                <navigator :url="'../../subpkg/charge/charge?id=' + pile.id + '&code=' + pile.code + '&pricePerHour=' + pricePerHour">
                  <button v-if="pile.state == 0" class="uni-button" size="mini" type="primary">充电</button>
                  <button v-else class="uni-button" size="mini" type="primary" disabled>充电</button>
                </navigator>
              </uni-td>
            </uni-tr>
          </uni-table>
        </view>
      </scroll-view>
    </u-popup>
  </view>
</template>

<script>
import { requestStations, requestPiles } from '@/utils/api/home.js'
export default {
  name: 'home',
  data() {
    return {
      scrollTop:0,
      //延时器的timerId
      timer: null,
      //输入的关键字
      kw: '',
      // 是否显示弹窗
      isShowPopup: false,
      // 窗口高度
      windowHeight: 0,
      // 当前站点ID
      currentStationId: null,
      // 当前站点的价格
      pricePerHour: null,
      //站点信息
      stationInfo: {
        name: '',
        address: '',
        coordinate: ''
      },
      //站点列表
      stationList: [],
      // 每页数据量
      pageSize: 10,
      // 当前页
      pageNo: 1,
      // 数据总量
      total: 0,
      //充电桩列表数据
      pileList: [],
      // 是否正在加载充电桩信息
      isLoadingPile: false,
    }
  },
  onLoad() {
    let sysInfo = uni.getSystemInfoSync()
    this.windowHeight = sysInfo.windowHeight - 50
    //获取充电站信息
    this.getStationList()
    //获取充电站信息
    this.getPileList()
  },
  methods: {
    //搜索站点
    handleInput(keyword) {
      if (this.kw !== keyword) {
        this.kw = keyword
        //清除timer对应的定时器
        clearTimeout(this.timer)
        //重启一个延时器，并把timerid赋值给this.timer
        this.timer = setTimeout(() => {
          //根据关键词，查询搜索建议列表
          this.getStationList({ params: { name: this.kw } })
        }, 600)
      }
    },
    // 封装获取充电站列表数据的方法
    getStationList() {
      let that = this;
      // 获取当前位置
      uni.getLocation({
        type: 'wgs84',
        timeout: 10000, // 设置10秒超时
        success: function (res) {
          that.stationInfo.coordinate = res.longitude + ',' + res.latitude;
          that.requestStations().then();
        },
        fail: function (err) {
          that.requestStations().then();
          uni.$showMsg('获取当前位置失败');
        }
      });
    },
    
    // 封装请求获取充电站数据的方法
    requestStations() {
      let that = this;
      return new Promise((resolve, reject) => {
        let params = {
          name: that.kw,
          address: that.stationInfo.address,
          pageNo: that.pageNo,
          pageSize: that.pageSize,
          coordinate: that.stationInfo.coordinate
        };
    
        requestStations(params).then((res) => {
          if (res.data.code == 0) {
            that.stationList = that.stationList.concat(res.data.data.List);
            that.pageNo = res.data.data.PageNo;
            that.pageSize = res.data.data.PageSize;
            that.total = res.data.data.TotalCount;
            resolve(res.data);
          } else {
            reject(res.data);
          }
        });
      });
    },

    //点击每个充电站的回调
    handleClickPile(station) {
      let that = this
      //将每个充电站的id赋值给currentStationId
      that.currentStationId = station.id
      //当前时间价格
      that.pricePerHour = station.pricePerHour
      that.getPileList()
      //打开弹窗
      that.isShowPopup = true
    },
    //显示位置
    showLocation(station) {
      //整理参数
      const str = station.coordinate
      //获得经纬度
      const [lng, lat] = str.split(',').map(parseFloat)
      //查看位置
      uni.openLocation({
        latitude: lat, //纬度 - 目的地/坐标点
        longitude: lng //经度 - 目的地/坐标点
      })
    },
    //弹窗关闭定义
    closePopup() {
      this.isShowPopup = false
    },
    //充电桩列表数据
    async getPileList() {
      try {
        let { data: res } = await requestPiles(this.currentStationId)
        if (res.code == 0) {
          this.pileList = res.data.List
          this.pageSize = res.data.PageSize
          this.pageNo = res.data.PageNo
          this.total = res.data.TotalCount
        }
      } catch (err) {
        console.log(err)
      }
    },
    //滚动到底部，加载下一页充电站数据
    loadNextPage(){
      // 利用Math.ceil算出新的分页
      let totalPage = Math.ceil((this.total) / this.pageSize)
      if(this.pageNo>=totalPage){
        uni.$showMsg('已经到最底部了')
        return false
      }else{
        this.pageNo++
        this.getStationList()
      }
    },
    refreshData(){
      this.scrollTop = 0
      this.stationList = []
      this.getStationList()
    }
  }
}
</script>

<style lang="scss">
.home-container {
  background-color: #f7f7f7;
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
}
.uni-table-td {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
