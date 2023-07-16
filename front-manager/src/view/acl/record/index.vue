<template>
  <el-card>
    <el-table style="margin: 10px 0px" border :data="payArr">
      <el-table-column label="#" align="center" type="index"></el-table-column>
      <el-table-column label="ID" align="center" prop="id" width="80"></el-table-column>
      <el-table-column label="订单编号" align="center" prop="payCode" show-overflow-tooltip></el-table-column>
      <el-table-column label="姓名" align="center" prop="username" width="80"></el-table-column>
      <el-table-column label="支付价格" align="center" prop="price" width="100"></el-table-column>
      <el-table-column label="充值状态" align="center" prop="state" width="100">
        <template #="{ row }">
          <el-tag v-if="row.state === 0">正在充值</el-tag>
          <el-tag v-else-if="row.state === 1" type="success">充值成功</el-tag>
          <el-tag v-else type="danger">取消充值</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="支付时间" align="center" show-overflow-tooltip> 
        <template #="{row}">
          <div>{{ formatDate(row.payAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" show-overflow-tooltip>
        <template #="{row}">
          <div>{{ formatDate(row.createAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="修改时间" align="center" show-overflow-tooltip>
        <template #="{row}">
          <div>{{ formatDate(row.updateAt) }}</div>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页器  -->
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3, 5, 7]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getHasPay" @size-change="handler" />
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted} from 'vue'
import { reqPayList } from '@/api/acl/user/index'
import type {pays,payRecord} from '@/api/acl/user/type'
import { formatDate } from '@/utils/time'
//默认页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(3)
//用户总个数
let total = ref<number>(0)
//充值列表
let payArr = ref<pays>([])

//组件挂载
onMounted(() => {
  getHasPay()
})
//获取充值列表数据
let getHasPay = async (pager = 1) => {
  pageNo.value = pager
  let res:payRecord = await reqPayList(pageNo.value, pageSize.value)
  if (res.code == 0) {
    payArr.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//分页器下拉菜单的自定义事件的回调
let handler = () => {
  getHasPay()
}

</script>

<style scoped></style>
