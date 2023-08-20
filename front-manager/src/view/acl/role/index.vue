<template>
  <el-card style="height: 70px">
    <el-form :inline="true" class="form">
      <el-form-item :label="$t('system.roleName')">
        <el-input :placeholder="$t('system.roleName')" v-model="keyword"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search" :disabled="keyword ? false : true">{{$t('button.search')}}</el-button>
        <el-button type="primary" @click="reset">{{$t('button.reset')}}</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card style="margin: 10px 0">
    <el-button type="primary" icon="Plus" @click="addRole" v-has="'RoleAdd'">{{$t('system.addRole')}}</el-button>
    <el-table border style="margin: 10px 0px" :data="allRole">
      <el-table-column label="#" type="index" align="center"></el-table-column>
      <el-table-column prop="id" label="id" align="center"></el-table-column>
      <el-table-column prop="name" :label="$t('system.roleName')" align="center" show-overflow-tooltip></el-table-column>
      <el-table-column prop="bindUsers" :label="$t('system.member')" align="center" show-overflow-tooltip></el-table-column>
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
      <el-table-column :label="$t('tabel.operate')" align="center" width="290px">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="User" @click="setPermisstion(row.id)">{{$t('button.privileges')}}</el-button>
          <el-button type="primary" size="small" icon="Edit" @click="updateRole(row)" v-has="'RoleUpdate'">{{$t('button.edit')}}</el-button>
          <el-popconfirm :title="`你确定要删除${row.name}?`" width="260px" @confirm="removeRole(row.id)" >
            <template #reference>
              <el-button type="primary" size="small" icon="Delete" v-has="'RoleDel'">{{$t('button.delete')}}</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3,5,7,10]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getHasRole" @size-change="sizeChange" />
  </el-card>
  <!-- 添加职位与更新已有职位的结构:对话框 -->
  <el-dialog v-model="dialogVisite" :title="$t(RoleParams.id ? 'system.editRole' : 'system.addRole')">
    <el-form :model="RoleParams" ref="form" :rules="rules">
      <el-form-item :label="$t('system.roleName')" prop="name">
        <el-input :placeholder="$t('system.roleName')" v-model="RoleParams.name"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" size="default" @click="dialogVisite = false">{{$t('pop.cancel')}}</el-button>
      <el-button type="primary" size="default" @click="save">{{$t('pop.confirm')}}</el-button>
    </template>
  </el-dialog>
 
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, nextTick } from 'vue'
import { formatDate } from '@/utils/time.ts'
import { reqAllRoleList,reqAddOrUpdateRole ,reqRemoveRole} from '@/api/acl/role/index.ts'
import type {roles,roleList,roleAddOrUp} from '@/api/acl/role/type'
import useLayOutSettingStore from '@/store/module/setting.ts'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

// import $bus from '@/bus/index.ts'
let $router = useRouter()
//当前页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(10)
//职位总个数
let total = ref<number>(0)
//搜索职位关键字
let keyword = ref<string>('')
//存储全部已有的职位
let allRole = ref<roles>([])
//layou组件相关配置小仓库
let settingStore = useLayOutSettingStore()
//控制对话框的显示与隐藏
let dialogVisite = ref<boolean>(false)
const form = ref<any>()

//收集新增岗位数据
let RoleParams = reactive<roleAddOrUp>({
  name: ''
})
//组件挂载完毕
onMounted(() => {
  //获取职位请求
  getHasRole()
})
//获取全部用户信息的方法|分页器当前页码发生变化的回调
let getHasRole = async (pager = 1) => {
  //修改当前页码
  pageNo.value = pager
  let res: roleList = await reqAllRoleList(pageNo.value, pageSize.value, keyword.value)
  if (res.code == 0) {
    allRole.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//下拉菜单的回调
let sizeChange = () => {
  getHasRole()
}
//搜索按钮
let search = () => {
  getHasRole()
  keyword.value = ''
}
//重置按钮
let reset = () => {
  settingStore.refsh = !settingStore.refsh
}
//添加职位
let addRole = () => {
  //弹窗打开
  dialogVisite.value = true
  //清空数据
  Object.assign(RoleParams, {
    name: '',
    id: 0
  })
  //清空上一次表单校验错误结果
  nextTick(() => {
    form.value.clearValidate('name')
  })
}
//编辑按钮
let updateRole = (row: any) => {
  //显示出对话框
  dialogVisite.value = true
  //存储已有的职位----带有ID的
  Object.assign(RoleParams, row)
  //清空上一次表单校验错误结果
  nextTick(() => {
    form.value.clearValidate('name')
  })
}
//自定义校验规则的回调
let validateName = (_rule: any, value: any, callback: any) => {
  if (value.trim().length >= 2) {
    callback()
  } else {
    callback(new Error('职位名称至少两位'))
  }
}
//职位相关的校验规则
let rules = {
  name: [
    {
      required: true,
      trigger: 'blur',
      validator: validateName
    }
  ]
}
//弹窗的确定按钮
let save = async () => {
  //表单校验结果,结果通过在发请求、结果没有通过不应该在发生请求
  await form.value.validate()
  //添加职位|更新职位的请求
  let res: any = await reqAddOrUpdateRole(RoleParams)
  if (res.code == 0) {
    //提示文字
    ElMessage({ type: 'success', message: RoleParams.id ? '更新成功' : '添加成功' })
    //对话框显示
    dialogVisite.value = false
    //再次获取全部的已有的职位
    getHasRole(RoleParams.id ? pageNo.value : 1)
  }
}

//删除角色
let removeRole = async (id: number) => {
  let res: any = await reqRemoveRole(id)
  if (res.code == 0) {
    //提示信息
    ElMessage({ type: 'success', message: '删除成功' })
    getHasRole(allRole.value.length > 1 ? pageNo.value : pageNo.value - 1)
  }
}
//分配权限，跳转至权限管理页面
let setPermisstion = (id:number)=>{
  $router.push({
    path: '/acl/authority',
    query: {
      roleId: id
    }
  })
}
</script>

<style scoped>
.form {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 50px;
}
</style>
