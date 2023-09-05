<template>
  <!--弹框信息确认-->
  <u-popup :show="show" mode="bottom" closeable @close="confimPile">
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
          <uni-tr v-for="pile in list" :key="pile.id">
            <uni-td align="center">{{ pile.code }}</uni-td>
            <uni-td align="center">
              <template v-slot>
                <button size="mini" type="default" v-if="pile.state == 0">空闲中</button>
                <button size="mini" type="primary" v-else-if="pile.state == 1">充电中</button>
                <button size="mini" type="warn" v-else>故障中</button>
              </template>
            </uni-td>
            <uni-td align="center">
              <view @click="goCharging(pile)">
                <button v-if="pile.state == 0" class="uni-button" size="mini" type="primary">充电</button>
                <button v-else class="uni-button" size="mini" type="primary" disabled>充电</button>
              </view>
            </uni-td>
          </uni-tr>
        </uni-table>
      </view>
    </scroll-view>
  </u-popup>
</template>

<script setup>
import {defineProps,ref} from 'vue'
let $emit = defineEmits(['confirm'])
let props = defineProps(['list', 'pricePerHour','show'])
let isLoadingPile = ref(false)
let confimPile = () => {
  $emit('confirm')
}
//去充电页面
let goCharging = (pile) => {
  // 构建跳转URL并拼接参数
  let url = '/subpkg/charging/charging?id=' + pile.id + '&code=' + pile.code + '&pricePerHour=' + props.pricePerHour+'&stationName='+pile.stationName
  uni.navigateTo({
    url
  })
}
</script>
