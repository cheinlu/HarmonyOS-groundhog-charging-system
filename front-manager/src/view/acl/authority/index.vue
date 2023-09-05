<template>
  <el-card>
    <el-table border :data="authArr">
      <el-table-column label="#" type="index" align="center" width="80px"></el-table-column>
      <el-table-column prop="model" :label="$t('system.Category')" align="center" width="150px"></el-table-column>
      <el-table-column prop="permissionList" :label="$t('system.functionalPermissions')" align="center" show-overflow-tooltip>
        <template #="{ row }">
          <el-checkbox-group v-model="row.checkedCities" @change="handleCheckedCitiesChange(row)">
            <el-checkbox v-for="(item, index) in row.permissionList" :key="index" :label="item">{{ item.name }}</el-checkbox>
          </el-checkbox-group>
        </template>
      </el-table-column>
      <el-table-column :label="$t('system.selectAll')" align="center" width="200px">
        <template #="{ row }">
          <el-checkbox v-model="row.checkAll" :indeterminate="row.isIndeterminate" @change="row.permissionList && handleCheckAllChange(row)">{{$t('system.selectAll')}}</el-checkbox>
        </template>
      </el-table-column>
    </el-table>
    <div style="display: flex; justify-content: flex-end; margin-top: 15px">
      <el-button type="primary" icon="Check" @click="bindPermi"  :disabled="!isButtonEnabled" v-has="'PermissionViewBind'">{{$t('button.bindPermissions')}}</el-button>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted ,computed} from 'vue'
import { reqAllPermiList,reqBindPermi } from '@/api/acl/authority/index.ts'
import { useRoute,useRouter} from 'vue-router'
import { ElMessage } from 'element-plus'
let $route = useRoute()
let $router = useRouter()

// 列表数据
let authArr = ref<any[]>([])
//组件挂载
onMounted(() => {
  getHasAuth()
})
//获取权限列表数据
let getHasAuth = async () => {
  let res: any = await reqAllPermiList()
  if (res.code == 0) {
    authArr.value = res.data.data.map((item: any) => ({
      ...item,
      permissionList: item.permissionList || [] // 确保 permissionList 是一个数组，如为空则置为空数组
    }))
  }
}
//全选按钮操作
let handleCheckAllChange = (row: any) => {
  //indeterminate  设置不确定状态，仅负责样式控制
  //如果选中的小按钮与全选按钮相等，显示全部数组，否则空数组
  row.checkedCities = row.checkAll ? row.permissionList : []
  //true(全选)|false(没有全选)
  row.isIndeterminate = false
}
//小按钮--单选操作
let handleCheckedCitiesChange = (row: any) => {
  if (row.permissionList) {
    // 添加判断，确保 permissionList 不为 undefined
    //checkedCount是复选框的勾选数据-长度
    const checkedCount = row.checkedCities.length
    // 勾选上的项目个数与全部的职位个数相等，复选框全勾上
    row.checkAll = checkedCount === row.permissionList.length
    //不确定的样式
    row.isIndeterminate = checkedCount !== 0 && checkedCount !== row.permissionList.length
  }
  
}
// 绑定权限
let bindPermi = async () => {
  let id =  $route.query.roleId
  const selected: any = []
  for (const row of authArr.value) {
    if (row.checkedCities && row.checkedCities.length > 0) {
      selected.push(...row.checkedCities)
    }
  }
  const selectedKeyValues = selected.map((item: any) => item.key)
  let res:any = await reqBindPermi((id as any), selectedKeyValues)
  if(res.code==0){
    //提示文字
    ElMessage({ type: 'success', message: '绑定权限成功' })
    //跳转到角色管理页面
    $router.push('/acl/role')
  }
}
// 判断是否启用"绑定权限"按钮
const isButtonEnabled = computed(() => {
  for (const row of authArr.value) {
    if (row.checkedCities && row.checkedCities.length > 0) {
      return true
    }
  }
  return false
})
</script>
<script lang="ts">
export default {
  name: '权限管理'
}
</script>
<style scoped></style>
