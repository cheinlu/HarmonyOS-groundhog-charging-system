<template>
  <el-card style="height: 170px">
    <el-form :inline="true" class="form" label-width="90px" :model="changeOrder">
      <el-form-item label="订单编号:">
        <el-input placeholder="请输入订单编号" v-model="changeOrder.orderCode"></el-input>
      </el-form-item>
      <el-form-item label="用户名:">
        <el-select placeholder="请选择用户名" v-model="changeOrder.userId">
          <el-option :label="user.nickname" :value="user.id" v-for="user in userArr" :key="user.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="充电站名:">
        <el-select placeholder="请选择充电站名" v-model="changeOrder.stationId" @change="chargeStation">
          <el-option v-for="item in stationArr" :key="item.id" :label="item.name" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="充电桩编号:">
        <el-select placeholder="请输入充电桩编号" v-model="changeOrder.pileId" :disabled="changeOrder.stationId ? false : true">
          <el-option :label="pile.code" :value="pile.id" v-for="pile in filterPile" :key="pile.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="选择时间:">
        <el-date-picker v-model="changeOrder.beforeAt" type="datetime" placeholder="选择开始时间" />
        <span style="margin: 0 7px"></span>
        <el-date-picker v-model="changeOrder.afterAt" type="datetime" placeholder="选择结束时间" />
      </el-form-item>
      <el-form-item label="订单状态:">
        <el-select placeholder="请选择订单状态" v-model="changeOrder.state">
          <el-option label="充电中" value="0"></el-option>
          <el-option label="已完成" value="1"></el-option>
          <el-option label="未付帐" value="2"></el-option>
        </el-select>
      </el-form-item>
      <span style="margin: 0 10px"></span>
      <el-form-item>
        <el-button type="primary" @click="search" :disabled="isButtonDisabled">搜索</el-button>
        <el-button type="primary" @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card style="margin: 10px 0px">
    <!-- table展示用户信息 -->
    <el-table style="margin: 10px 0px" border :data="orderArr">
      <el-table-column label="#" align="center" type="index" width="35"></el-table-column>
      <el-table-column label="订单编号" align="center" prop="orderCode" width="120" show-overflow-tooltip></el-table-column>
      <el-table-column label="名字" align="center" prop="nickname" show-overflow-tooltip></el-table-column>
      <el-table-column label="站名" align="center" prop="stationName" show-overflow-tooltip></el-table-column>
      <el-table-column label="桩编号" align="center" prop="pileCode" show-overflow-tooltip> </el-table-column>
      <el-table-column label="充电开始时间" align="center" show-overflow-tooltip width="130">
        <template #="{ row }">
          <div>{{ formatDate(row.startAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="充电结束时间" align="center" show-overflow-tooltip width="110">
        <template #="{ row }">
          <div>{{ formatDate(row.stopAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="订单状态" align="center" prop="state" show-overflow-tooltip width="90">
        <template #="{ row }">
          <el-tag v-if="row.state === 0">充电中</el-tag>
          <el-tag v-else-if="row.state === 1" type="success">已完成</el-tag>
          <el-tag v-else type="danger">未付账</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="价格" align="center" prop="price" show-overflow-tooltip> </el-table-column>
      <el-table-column label="创建时间" align="center" show-overflow-tooltip width="130">
        <template #="{ row }">
          <div>{{ formatDate(row.createAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="修改时间" align="center" show-overflow-tooltip width="130">
        <template #="{ row }">
          <div>{{ formatDate(row.updateAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="100" fixed="right">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="updateOrder(row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页器  -->
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3, 5, 7]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getHasOrder" @size-change="handler" />
  </el-card>
  <!-- 编辑订单的弹窗提示 -->
  <el-dialog v-model="dialogVisible" title="修改充电订单" width="30%">
    <el-form ref="form" :rules="rules" :model="orderForm" label-width="90px">
      <el-form-item label="订单状态:" prop="state">
        <el-select v-model="orderForm.state">
          <el-option label="充电中" value="0"></el-option>
          <el-option label="已完成" value="1"></el-option>
          <el-option label="未付账" value="2"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="价格:" prop="price" style="width: 295px">
        <el-input placeholder="请输入价格" v-model="orderForm.price"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" size="default" @click="dialogVisible = false">取消</el-button>
      <el-button type="primary" size="default" @click="save">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, computed, nextTick } from 'vue'
import { reqOrder, reqUpdateOrder } from '@/api/opration/index'
import { reqUserInfo } from '@/api/acl/user/index'
import type { Records } from '@/api/acl/user/type'
import { reqStation } from '@/api/device/station/index.ts'
import type { stations } from '@/api/device/station/type'
import { reqPile } from '@/api/device/pile/index'
import type { piles } from '@/api/device/pile/type'
import type { charges, chargeList } from '@/api/opration/type'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/time'
import useLayOutSettingStore from '@/store/module/setting'
let settingStore = useLayOutSettingStore()
//默认页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(3)
//用户总个数
let total = ref<number>(0)
//充电订单数据
let orderArr = ref<charges>([])
//存储全部用户的数组
let userArr = ref<Records>([])
//存储全部充电站的数组
let stationArr = ref<stations>([])
//存储充电桩数据
let pileArr = ref<piles>([])
//用于存储用户选择的充电站id
let selectId = ref(0)
let dialogVisible = ref<boolean>(false)
let form = ref()
//充电订单参数收集
let changeOrder: { [key: string]: any } = reactive({
  orderCode: '',
  userId: '',
  stationId: '',
  pileId: '',
  beforeAt: '',
  afterAt: '',
  state: ''
})
//修改表单的数据收集
let orderForm = reactive({
  id: 0,
  state: null,
  price: null
})
//组件挂载
onMounted(() => {
  getHasOrder()
  getUserList()
  getHasStation()
  getHasPile()
})
//对 充电桩数据 进行过滤，只保留与 selectedStationId 相等的充电桩
let filterPile = computed(() => {
  return pileArr.value.filter((pile) => {
    pile.stationId === selectId.value
  })
})
//搜索按钮默认禁用，只要changeOrder有一个参数有值就不禁用了
const isButtonDisabled = computed(() => {
  for (let key in changeOrder) {
    if (changeOrder[key]) {
      return false //只要有一个属性有值，就不禁用按钮
    }
  }
  //所有属性都没有值，禁用按钮
  return true
})
//获取订单的数据
let getHasOrder = async (pager = 1) => {
  pageNo.value = pager
  let res: chargeList = await reqOrder(pageNo.value, pageSize.value, changeOrder.orderCode, changeOrder.userId, changeOrder.stationId, changeOrder.pileId, changeOrder.beforeAt, changeOrder.afterAt, changeOrder.state)
  if (res.code == 0) {
    orderArr.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//分页器下拉菜单的自定义事件的回调
let handler = () => {
  getHasOrder()
}
//获取用户信息
let getUserList = async () => {
  let res: any = await reqUserInfo(1, 100, '')
  if (res.code == 0) {
    userArr.value = res.data.List
  }
}
//获取充电站列表数据
let getHasStation = async () => {
  let res: any = await reqStation(1, 100, '')
  if (res.code == 0) {
    stationArr.value = res.data.List
  }
}
//获取充电桩列表数据
let getHasPile = async () => {
  let res: any = await reqPile(1, 100, 0, '')
  if (res.code == 0) {
    pileArr.value = res.data.List
  }
}
//当选择充电站时，获取选择值的id
let chargeStation = () => {
  selectId.value = changeOrder.stationId
  console.log('filterPile', filterPile)
}
//修改订单的按钮回调
let updateOrder = (row: any) => {
  dialogVisible.value = true
  //参数赋值
  orderForm.id = row.id
  orderForm.state = row.state
  orderForm.price = row.price
  //清空上一次表单校验错误结果
  nextTick(() => {
    form.value.clearValidate('state')
    form.value.clearValidate('price')
  })
}
//搜索按钮
let search = () => {
  getHasOrder()
  //清空数据
  Object.assign(changeOrder, {
    orderCode: '',
    userId: '',
    stationId: '',
    pileId: '',
    beforeAt: '',
    afterAt: '',
    state: ''
  })
}
//重置按钮
let reset = () => {
  settingStore.refsh = !settingStore.refsh
}
let validatePrice = (_rule: any, value: any, callback: any) => {
  if (value == null) {
    return callback(new Error('价格不能为空'))
  }
  setTimeout(() => {
    if (isNaN(value)) {
      callback(new Error('请输入数值'))
    } else {
      if (value < 0) {
        callback(new Error('价格必须大于0'))
      } else {
        callback()
      }
    }
  }, 100)
}
//校验规则
let rules = {
  state: [{ required: true, message: '请选择订单状态', trigger: 'change' }],
  price: [{ required: true, trigger: 'blur', validator: validatePrice }]
}
//点击确定按钮
let save = async () => {
  await form.value.validate()
  let res: any = await reqUpdateOrder(orderForm.id, orderForm.price, orderForm.state)
  if (res.code == 0) {
    //弹窗关闭
    dialogVisible.value = false
    //提示添加|修改成功
    ElMessage({ type: 'success', message: '更新充电订单成功' })
    //获取数据
    getHasOrder()
  } else {
    //弹窗关闭
    dialogVisible.value = false
    ElMessage({ type: 'error', message: '更新充电订单失败' })
  }
}
</script>
<script lang="ts">
export default {
  name: '订单'
}
</script>
<style scoped></style>
