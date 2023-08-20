<template>
  <el-card style="height: 70px">
    <el-form :inline="true" class="form">
      <el-form-item :label="$t('equip.nameLabel')">
        <el-select :placeholder="$t('equip.nameLabel')" v-model="pileData.stationId">
          <el-option v-for="item in stationArr" :key="item.id" :label="item.name" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search" :disabled="pileData.stationId ? false : true">{{ $t('button.search') }}</el-button>
        <el-button type="primary" @click="reset">{{ $t('button.reset') }}</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card style="margin: 10px 0px">
    <el-button type="primary" size="default" @click="addPile" v-has="'ChargePileAdd'">{{ $t('equip.addPile') }}</el-button>
    <!-- table展示用户信息 -->
    <el-table style="margin: 10px 0px" border :data="pileArr">
      <el-table-column label="#" align="center" type="index"></el-table-column>
      <el-table-column label="ID" align="center" prop="id" width="50"></el-table-column>
      <el-table-column :label="$t('equip.serialNumber')" align="center" prop="code" width="90" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('equip.nameLabel')" align="center" width="150" prop="stationName" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('equip.state')" align="center" prop="state" show-overflow-tooltip width="150">
        <template #="{ row }">
          <el-tag v-if="row.state === 0">空闲中</el-tag>
          <el-tag v-else-if="row.state === 1" type="success">充电中</el-tag>
          <el-tag v-else type="danger">故障中</el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="$t('tabel.createAt')" align="center" show-overflow-tooltip>
        <template #="{ row }">
          <div>{{ formatDate(row.createAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('tabel.updateAt')" align="center" show-overflow-tooltip>
        <template #="{ row }">
          <div>{{ formatDate(row.updateAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('tabel.operate')" align="center">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="updatePile(row)" v-has="'ChargePileUpdate'">{{ $t('button.edit') }}</el-button>
          <el-popconfirm :title="`你确定要删除编号为${row.code}的${row.stationName}?`" width="300px" @confirm="deleteUser(row.id)">
            <template #reference>
              <el-button type="primary" size="small" icon="Delete" v-has="'ChargePileDel'">{{ $t('button.delete') }}</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页器 -->
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3, 5, 7, 10]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getHasPile" @size-change="handler" />
  </el-card>
  <!-- 添加|修改充电桩弹框 -->
  <el-dialog v-model="dialogVisible" :title="$t(pileForm.id?'equip.editPile':'equip.addPile')" width="30%">
    <el-form label-width="120px" :rules="rules" :model="pileForm" ref="pileFormRef">
      <el-form-item :label="$t('equip.nameLabel')" prop="stationId">
        <el-select v-model="pileForm.stationId" :placeholder="$t('equip.nameLabel')">
          <el-option v-for="item in stationArr" :key="item.id" :label="item.name" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('equip.serialNumber')" style="width: 320px" prop="code">
        <el-input :placeholder="$t('equip.serialNumber')" v-model="pileForm.code"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">{{ $t('pop.cancel') }}</el-button>
        <el-button type="primary" @click="save"> {{ $t('pop.confirm') }} </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, nextTick } from 'vue'
import { reqPile, reqAddOrUpdatePile, reqRemovePile } from '@/api/device/pile/index'
import type { pileList, piles, pileAddOrUpdate } from '@/api/device/pile/type'
import { reqStation } from '@/api/device/station/index.ts'
import type { stationList, stations } from '@/api/device/station/type'
import useLayOutSettingStore from '@/store/module/setting'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/time'
let settingStore = useLayOutSettingStore()
//默认页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(3)
//用户总个数
let total = ref<number>(0)
//存储充电桩数据
let pileArr = ref<piles>([])
//存储全部充电站的数组
let stationArr = ref<stations>([])
//弹窗默认隐藏
let dialogVisible = ref(false)
//收集站点id和code用来搜索
let pileData = reactive({
  stationId: '',
  code: ''
})
//收集添加|编辑充电桩数据
let pileForm = reactive<pileAddOrUpdate>({
  stationId: 0,
  code: ''
})
//获取表单
let pileFormRef = ref()
//组件挂载
onMounted(() => {
  getHasPile()
  getHasStation()
})
//获取充电桩列表数据
let getHasPile = async (pager = 1) => {
  pageNo.value = pager
  let res: pileList = await reqPile(pageNo.value, pageSize.value, pileData.stationId, pileData.code)
  if (res.code == 0) {
    pileArr.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//获取充电站列表数据
let getHasStation = async () => {
  let res: stationList = await reqStation(1, 100, '')
  if (res.code == 0) {
    stationArr.value = res.data.List
  }
}
//分页器下拉菜单的自定义事件的回调
let handler = () => {
  getHasPile()
}
//搜索充电站
let search = () => {
  getHasPile()
  //清空数据
  pileData.stationId = ''
}
//重置按钮的回调
let reset = () => {
  settingStore.refsh = !settingStore.refsh
}
//添加充电桩按钮回调
let addPile = () => {
  dialogVisible.value = true
  //清空表单数据
  Object.assign(pileForm, {
    stationId: null,
    code: ''
  })
  //清空上一次表单校验错误结果
  nextTick(() => {
    pileFormRef.value.clearValidate('stationId')
    pileFormRef.value.clearValidate('code')
  })
}
//编辑充电桩按钮回调
let updatePile = (row: any) => {
  Object.assign(pileForm, row)
  dialogVisible.value = true
  //清空上一次表单校验错误结果
  nextTick(() => {
    pileFormRef.value.clearValidate('stationId')
    pileFormRef.value.clearValidate('code')
  })
}
//删除充电桩按钮回调
let deleteUser = async (id: number) => {
  let res: any = await reqRemovePile(id)
  if (res.code == 0) {
    ElMessage({ type: 'success', message: '删除充电桩成功' })
    getHasPile(pileArr.value.length > 1 ? pageNo.value : pageNo.value - 1)
  }
}
//验证规则
let rules = {
  stationId: [{ required: true, message: '请选择充电站', trigger: 'change' }],
  code: [{ required: true, message: '请输入编号', trigger: 'blur' }]
}
//弹窗的确定按钮
let save = async () => {
  await pileFormRef.value.validate()
  let res: any = await reqAddOrUpdatePile(pileForm)
  if (res.code == 0) {
    //弹窗关闭
    dialogVisible.value = false
    //提示添加|修改成功
    ElMessage({ type: 'success', message: pileForm.id ? '更新充电桩成功' : '添加充电桩成功' })
    //获取数据
    getHasPile()
  } else {
    //弹窗关闭
    dialogVisible.value = false
    ElMessage({ type: 'error', message: pileForm.id ? '更新充电桩失败' : '添加充电桩失败' })
  }
}
</script>

<style scoped>
.form {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
