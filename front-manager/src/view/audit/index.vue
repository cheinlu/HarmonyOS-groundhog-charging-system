<template>
  <el-card style="height: 70px">
    <el-form :inline="true" class="form">
      <el-form-item :label="$t('audit.changedUser')">
        <el-input :placeholder="$t('audit.changedUser')" v-model="auditParams.ChangedUser"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search" :disabled="auditParams.ChangedUser ? false : true">{{ $t('button.search') }}</el-button>
        <el-button type="primary" @click="reset">{{ $t('button.reset') }}</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card style="margin: 10px 0px">
    <el-table :data="auditArr" border>
      <el-table-column label="#" align="center" type="index"></el-table-column>
      <el-table-column :label="$t('audit.tableName')" align="center" prop="tableName" show-overflow-tooltip width="110px"></el-table-column>
      <el-table-column :label="$t('audit.changedUser')" align="center" prop="changedUser" show-overflow-tooltip width="120px"></el-table-column>
      <el-table-column :label="$t('audit.operationType')" align="center" prop="operationType" show-overflow-tooltip width="130px"></el-table-column>
      <el-table-column :label="$t('audit.oldValue')" align="center" prop="oldValue" show-overflow-tooltip width="150px"></el-table-column>
      <el-table-column :label="$t('audit.newValue')" align="center" prop="newValue" show-overflow-tooltip width="150px"></el-table-column>
      <el-table-column :label="$t('audit.operateSql')" align="center" prop="operateSql" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('audit.changedDate')" align="center" prop="changedDate" show-overflow-tooltip></el-table-column>
    </el-table>
  </el-card>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { reqAuditList } from '@/api/audit/index.ts'
import useLayOutSettingStore from '@/store/module/setting'
let settingStore = useLayOutSettingStore()
//默认页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(3)
//用户总个数
let total = ref<number>(0)
//列表的数组数据
let auditArr = ref([])
//获取列表需携带的参数
let auditParams = reactive({
  operationTypeL: '',
  ChangedUser: ''
})
//组件挂载
onMounted(() => {
  getHasAudit()
})
//获取审计列表数据
let getHasAudit = async () => {
  let res: any = await reqAuditList(auditParams)
  if (res.code == 0) {
    auditArr.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//搜索按钮的回调
let search = () => {
  getHasAudit()
  auditParams.ChangedUser = ''
}
//重置按钮的回调
let reset = () => {
  settingStore.refsh = !settingStore.refsh
}
</script>

<style scoped>
.form {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
