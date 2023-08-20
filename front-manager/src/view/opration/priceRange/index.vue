<template>
  <el-card style="margin: 10px 0px">
    <el-button type="primary" size="default" @click="addPrice" v-has="'PriceAdd'">{{$t('opration.addPrice')}}</el-button>
    <!-- table展示用户信息 -->
    <el-table style="margin: 10px 0px" border :data="priceArr">
      <el-table-column label="#" align="center" type="index"></el-table-column>
      <el-table-column label="ID" align="center" prop="id"></el-table-column>
      <el-table-column :label="$t('opration.startTime')" align="center" prop="startHour" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('opration.endTime')" align="center" prop="endHour" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('opration.price')" align="center" prop="price" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('tabel.createAt')" align="center" show-overflow-tooltip width="250">
        <template #="{ row }">
          <div>{{ formatDate(row.createAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('tabel.updateAt')" align="center" show-overflow-tooltip width="250">
        <template #="{ row }">
          <div>{{ formatDate(row.updateAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('tabel.operate')" width="200px" align="center">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="updatePrice(row)" v-has="'PriceUpdate'">{{$t('button.edit')}}</el-button>
          <el-popconfirm :title="`你确定删除id为${row.id}的区间价格吗？`" width="250px" @confirm="deletePrice(row.id)" >
            <template #reference>
              <el-button type="primary" size="small" icon="Delete" v-has="'PriceDel'">{{$t('button.delete')}}</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页器  -->
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3, 5, 7, 10]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getHasPrice" @size-change="handler" />
  </el-card>
  <!-- 添加|修改充电价格弹窗 -->
  <el-dialog v-model="dialogVisible" :title="$t(priceParams.id?'opration.editPrice':'opration.addPrice')">
    <el-form :model="priceParams" ref="form" :rules="rules" label-width="95px">
      <el-form-item :label="$t('opration.startTime')" prop="startHour" >
        <el-input :placeholder="$t('opration.startTime')" v-model="priceParams.startHour"></el-input>
      </el-form-item>
      <el-form-item :label="$t('opration.endTime')" prop="endHour" >
        <el-input :placeholder="$t('opration.endTime')" v-model="priceParams.endHour"></el-input>
      </el-form-item>
      <el-form-item :label="$t('opration.price')" prop="price" >
        <el-input :placeholder="$t('opration.price')" v-model="priceParams.price"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" size="default" @click="dialogVisible = false">{{$t('pop.cancel')}}</el-button>
      <el-button type="primary" size="default" @click="save">{{$t('pop.confirm')}}</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, nextTick } from 'vue'
import { reqChargePrice, reqAddOrUpdatePrice, reqRemovePrice } from '@/api/opration/index'
import type { changePrice, price, priceList } from '@/api/opration/type'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/time'
//默认页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(3)
//用户总个数
let total = ref<number>(0)
//区间价格列表
let priceArr = ref<changePrice>([])
let form = ref()
//弹窗默认关闭
let dialogVisible = ref<boolean>(false)
//收集充电价格数据
let priceParams = reactive<price>({
  startHour: '',
  endHour: '',
  price: ''
})
//组件挂载
onMounted(() => {
  getHasPrice()
})
//获取充电价格数据
let getHasPrice = async (pager = 1) => {
  pageNo.value = pager
  let res: priceList = await reqChargePrice(pageNo.value, pageSize.value)
  if (res.code == 0) {
    priceArr.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//分页器下拉菜单的自定义事件的回调
let handler = () => {
  getHasPrice()
}
//添加充电价格按钮
let addPrice = () => {
  dialogVisible.value = true
  //清空表单数据
  Object.assign(priceParams, {
    startHour: '',
    endHour: '',
    price: '',
    id: 0
  })
  //清空上一次表单校验错误结果
  nextTick(() => {
    form.value.clearValidate('startHour')
    form.value.clearValidate('endHour')
    form.value.clearValidate('price')
  })
}
//编辑充电价格按钮
let updatePrice = (row: any) => {
  dialogVisible.value = true
  Object.assign(priceParams, row)
  //清空上一次表单校验错误结果
  nextTick(() => {
    form.value.clearValidate('startHour')
    form.value.clearValidate('endHour')
    form.value.clearValidate('price')
  })
}

//删除充电价格按钮
let deletePrice = async (id: number) => {
  let res: any = await reqRemovePrice(id)
  if (res.code == 0) {
    ElMessage({ type: 'success', message: '删除此区间价格成功' })
    getHasPrice(priceArr.value.length > 1 ? pageNo.value : pageNo.value - 1)
  }
}
//开始时间校验
let validatorStart = (_rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请选择开始时间'))
  } else if (!Number.isInteger(parseInt(value))) {
    callback(new Error('请输入整数类型的数字'))
  } else if (value < 0 || value > 24) {
    callback(new Error('请输入0-24之间的数字'))
  } else {
    callback()
  }
}
//结束时间校验
let validatorEnd = (_rule: any, value: any, callback: any) => {
  if (value == '') {
    callback(new Error('请选择结束时间'))
  } else if (!Number.isInteger(parseInt(value))) {
    callback(new Error('结束时间为0-24之间的数字'))
  } else if (value && value <= priceParams.startHour) {
    callback(new Error('结束时间必须晚于开始时间'))
  } else {
    callback()
  }
}
//价格校验
let validatorPrice = (_rule: any, value: any, callback: any) => {
  if (!value) {
    return callback(new Error('价格不能为空'))
  }
  setTimeout(() => {
    if (isNaN(value)) {
      callback(new Error('请输入数值'))
    } else {
      if (value <= 0) {
        callback(new Error('价格必须大于0'))
      } else {
        callback()
      }
    }
  }, 100)
}
//校验规则
let rules = {
  startHour: [{ required: true, trigger: 'blur', validator: validatorStart }],
  endHour: [{ required: true, trigger: 'blur', validator: validatorEnd }],
  price: [{ required: true, trigger: 'blur', validator: validatorPrice }]
}
//对话框确定按钮
let save = async () => {
  //校验合格再发请求
  await form.value.validate()
  let res: any = await reqAddOrUpdatePrice(priceParams)
  if (res.code == 0) {
    dialogVisible.value = false
    //提示添加|修改成功
    ElMessage({ type: 'success', message: priceParams.id ? '更新成功' : '添加成功' })
    getHasPrice()
  } else {
    dialogVisible.value = false
    ElMessage({ type: 'error', message: priceParams.id ? '更新失败' : '添加失败' })
  }
}
</script>

<style scoped></style>
