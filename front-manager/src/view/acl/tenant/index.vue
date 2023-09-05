<template>
  <el-card>
    <el-button type="primary" size="default" @click="addTenant" v-has="'TenantAdd'">{{$t('system.addTenant')}}</el-button>
    <!-- table展示用户信息 -->
    <el-table style="margin: 10px 0px" border :data="teList">
      <el-table-column label="#" align="center" type="index"></el-table-column>
      <el-table-column :label="$t('system.tenant')" align="center" prop="contactName" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('system.mailbox')" align="center" prop="contactEmail" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('system.phone')" align="center" prop="contactPhone" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('system.state')" align="center" prop="status" show-overflow-tooltip>
        <template #="{ row }">
          <el-tag v-if="row.status === 'active'">激活</el-tag>
          <el-tag v-else-if="row.status === 'inactive'">未激活</el-tag>
          <el-tag v-else>激活</el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="$t('system.companyName')" align="center" prop="name" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('system.website')" align="center" prop="domain" show-overflow-tooltip> </el-table-column>
      <el-table-column :label="$t('tabel.createAt')" align="center" show-overflow-tooltip width="160">
        <template #="{ row }">
          <div>{{ formatDate(row.createAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('tabel.updateAt')" align="center" show-overflow-tooltip width="160">
        <template #="{ row }">
          <div>{{ formatDate(row.updateAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('tabel.operate')" width="190px" align="center">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="updateTenant(row)" v-has="'TenantUpdate'">{{$t('button.edit')}}</el-button>
          <el-popconfirm :title="`你确定要删除${row.contactName}?`" width="260px" @confirm="deleteTenant(row.id)" v-has="'TenantDel'">
            <template #reference>
              <el-button type="primary" size="small" icon="Delete" v-has="'RoleDel'">{{$t('button.delete')}}</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页器 -->
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3, 5, 7, 10]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getHasTenant" @size-change="handler" />
  </el-card>
  <!-- 抽屉结构:完成添加新的用户账号|更新已有的账号信息 -->
  <el-drawer v-model="drawer">
    <!-- 头部标题:将来文字内容应该动态的 -->
    <template #header>
      <h4>{{ $t(tenantParams.id ? 'system.editTenant' : 'system.addTenant') }}</h4>
    </template>
    <!-- 身体部分 -->
    <template #default>
      <el-form ref="formRef" :model="tenantParams" :rules="rules" label-width="90px">
        <el-form-item :label="$t('system.tenant')" prop="contactName">
          <el-input :placeholder="$t('system.tenant')" v-model="tenantParams.contactName"></el-input>
        </el-form-item>
        <el-form-item :label="$t('system.mailbox')" prop="contactEmail">
          <el-input :placeholder="$t('system.mailbox')" v-model="tenantParams.contactEmail"></el-input>
        </el-form-item>
        <el-form-item :label="$t('system.phone')" prop="contactPhone">
          <el-input :placeholder="$t('system.phone')" v-model="tenantParams.contactPhone"></el-input>
        </el-form-item>
        <el-form-item :label="$t('system.state')" prop="status">
          <el-select v-model="tenantParams.status" :placeholder="$t('system.state')">
            <el-option label="激活" value="active"></el-option>
            <el-option label="未激活" value="inactive"></el-option>
            <el-option label="暂停" value="suspended"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('system.companyName')" prop="name">
          <el-input :placeholder="$t('system.companyName')" v-model="tenantParams.name"></el-input>
        </el-form-item>
        <el-form-item :label="$t('system.website')" prop="domain">
          <el-input :placeholder="$t('system.website')" v-model="tenantParams.domain"></el-input>
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="drawer = false">{{$t('pop.cancel')}}</el-button>
        <el-button type="primary" @click="save">{{$t('pop.confirm')}}</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup lang="ts">
import { formatDate } from '@/utils/time.ts'
import { ElMessage } from 'element-plus'
import { reqTeList, reqAddOrUpdateTe, reqDeleteTe } from '@/api/acl/tenant/index'
import {tenantList,teList,tenant} from '@/api/acl/tenant/type'
import { ref, onMounted, reactive, nextTick } from 'vue'

//默认页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(3)
//用户总个数
let total = ref<number>(0)
//租户数据
let teList = ref<teList>([])
//抽屉默认关闭
let drawer = ref<boolean>(false)
//获取表单
const formRef = ref<any>()
//租户信息的收集
let tenantParams = reactive<tenant>({
  name: '',
  domain: '',
  contactName: '',
  contactEmail: '',
  contactPhone: '',
  status: ''
})
//组件挂载
onMounted(() => {
  getHasTenant()
})
//获取租户数据
let getHasTenant = async (pager = 1) => {
  pageNo.value = pager
  let res: tenantList = await reqTeList()
  if (res.code == 0) {
    teList.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//分页器下拉菜单的自定义事件的回调
let handler = () => {
  getHasTenant()
}
//添加租户
let addTenant = () => {
  drawer.value = true
  //清空表单数据
  Object.assign(tenantParams, {
    name: '',
    domain: '',
    contactName: '',
    contactEmail: '',
    contactPhone: '',
    status: '',
    id: 0
  })
  //清空上一次表单校验错误结果
  nextTick(() => {
    formRef.value.clearValidate('name')
    formRef.value.clearValidate('domain')
    formRef.value.clearValidate('contactName')
    formRef.value.clearValidate('contactEmail')
    formRef.value.clearValidate('contactPhone')
    formRef.value.clearValidate('status')
  })
}
//修改租户
let updateTenant = (row: any) => {
  //打开抽屉
  drawer.value = true
  //存储已有的用户数据
  Object.assign(tenantParams, row)
   //清空上一次表单校验错误结果
   nextTick(() => {
    formRef.value.clearValidate('name')
    formRef.value.clearValidate('domain')
    formRef.value.clearValidate('contactName')
    formRef.value.clearValidate('contactEmail')
    formRef.value.clearValidate('contactPhone')
    formRef.value.clearValidate('status')
  })
}
//删除租户
let deleteTenant = async (id: number) => {
  let res: any = await reqDeleteTe(id)
  if (res.code == 0) {
    ElMessage({ type: 'success', message: '删除租户成功' })
    getHasTenant(teList.value.length > 1 ? pageNo.value : pageNo.value - 1)
  }
}
//电话
let validatePhone = (_rule:any, value:any, callback:any) => {
  const pattern = /^1[3456789]\d{9}$/
  if (!value || pattern.test(value)) {
    callback()
  } else {
    callback(new Error('请输入有效的电话号码(11位号码)'))
  }
}
//邮箱
let validateEmail = (_rule:any, value:any, callback:any) => {
  const pattern = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
  if (!value || pattern.test(value)) {
    callback()
  } else {
    callback(new Error('请输入有效的邮箱'))
  }
}
//网站
let validateWebsite = (_rule:any, value:any, callback:any) => {
  const pattern = /^([A-Za-z0-9-]+\.){1,2}[A-Za-z]{2,}$/
  if (!value || pattern.test(value)) {
    callback()
  } else {
    callback(new Error('请输入有效的公司网站,如www.example.com'))
  }
}
//校验规则
let rules = {
  name: [{ required: true, message: '请输入公司名称', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }],
  contactName: [{ required: true, message: '请输入租户名字', trigger: 'blur' }],
  domain: [
    { required: true, message: '请输入公司网站', trigger: 'blur' },
    // 自定义校验规则
    { validator: validateWebsite, trigger: 'blur' }
  ],
  contactEmail: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入有效的邮箱', trigger: ['blur', 'change'] }, // 自定义校验规则
    { validator: validateEmail, trigger: 'blur' }
  ],
  contactPhone: [
    { required: true, message: '请输入电话号码', trigger: 'blur' },
    // 自定义校验规则
    { validator: validatePhone, trigger: 'blur' }
  ]
}
//抽屉的确认按钮
let save = async () => {
  await formRef.value.validate()
  let res: any = await reqAddOrUpdateTe(tenantParams)
  if (res.code == 0) {
    //抽屉关闭
    drawer.value = false
    //提示添加|修改成功
    ElMessage({ type: 'success', message: tenantParams.id ? '更新成功' : '添加成功' })
    //获取数据
    getHasTenant()
  } else {
    drawer.value = false
    ElMessage({ type: 'error', message: tenantParams.id ? '更新失败' : '添加失败' })
  }
}
</script>
<script lang="ts">
export default {
  name: '租户管理'
}
</script>
<style scoped></style>
