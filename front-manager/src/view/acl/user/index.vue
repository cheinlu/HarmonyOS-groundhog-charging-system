<template>
  <el-card style="height: 70px">
    <el-form :inline="true" class="form">
      <el-form-item label="用户名">
        <el-input placeholder="请输入用户名" v-model="username"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search" :disabled="username ? false : true">搜索</el-button>
        <el-button type="primary" @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card style="margin: 10px 0px">
    <el-button type="primary" size="default" @click="addUser" v-has="'UserAdd'">添加用户</el-button>
    <!-- table展示用户信息 -->
    <el-table style="margin: 10px 0px" border :data="userArr">
      <el-table-column label="#" align="center" type="index"></el-table-column>
      <el-table-column label="ID" align="center" prop="id"></el-table-column>
      <el-table-column label="用户名字" align="center" prop="username" show-overflow-tooltip></el-table-column>
      <el-table-column label="用户名称" align="center" prop="nickname" show-overflow-tooltip></el-table-column>
      <el-table-column label="余额" align="center" prop="balance" show-overflow-tooltip>
        <template #="{ row }">
          <div class="balance">
            <span>{{ row.balance }}</span>
            <el-button type="primary" @click="addPay(row)" size="small">充值</el-button>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" show-overflow-tooltip width="250">
        <template #="{ row }">
          <div>{{ formatDate(row.createTime) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200px" align="center">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="updateUser(row)" v-has="'UserUpdate'">编辑</el-button>
          <el-popconfirm :title="`你确定要删除${row.username}?`" width="260px" @confirm="deleteUser(row.id)" >
            <template #reference>
              <el-button type="primary" size="small" icon="Delete" v-has="'UserDel'">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页器 -->
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3, 5, 7, 10]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getUserList" @size-change="handler" />
  </el-card>
  <!-- 抽屉结构:完成添加新的用户账号|更新已有的账号信息 -->
  <el-drawer v-model="drawer">
    <!-- 头部标题:将来文字内容应该动态的 -->
    <template #header>
      <h4>{{ userParams.id ? '更新用户' : '添加用户' }}</h4>
    </template>
    <!-- 身体部分 -->
    <template #default>
      <el-form ref="formRef" :model="userParams" :rules="rules">
        <el-form-item label="用户姓名" prop="username">
          <el-input placeholder="请您输入用户姓名" v-model="userParams.username"></el-input>
        </el-form-item>
        <el-form-item label="用户昵称" prop="nickname">
          <el-input placeholder="请您输入用户昵称" v-model="userParams.nickname"></el-input>
        </el-form-item>
        <el-form-item label="用户密码" prop="password" v-if="!userParams.id">
          <el-input placeholder="请您输入用户密码" v-model="userParams.password"></el-input>
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="drawer = false">取消</el-button>
        <el-button type="primary" @click="save">确定</el-button>
      </div>
    </template>
  </el-drawer>
  <!-- 充值按钮的弹窗 -->
  <el-dialog v-model="dialogVisible" title="充值" width="30%">
    <el-form ref="form" label-width="90px" :rules="payRules" :model="payForm">
      <el-form-item label="用户名:" prop="userId">
        <span>{{ userParams.username }}</span>
      </el-form-item>
      <el-form-item label="充值金额:" prop="price">
        <el-input placeholder="请输入价格" v-model="payForm.price"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" size="default" @click="dialogVisible = false">取消</el-button>
      <el-button type="primary" size="default" @click="savePay">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, nextTick } from 'vue'
import { reqUserInfo, reqAddOrUpdateUser, reqRemoveUser, reqUserPay } from '@/api/acl/user/index'
import type {userAddOrUp, User, UserResponseData, Records, userPay } from '@/api/acl/user/type'
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
//定义响应式数据:收集用户输入进来的关键字
let username = ref<string>('')
//存储全部用户的数组
let userArr = ref<Records>([])
//抽屉默认关闭
let drawer = ref<boolean>(false)
let dialogVisible = ref<boolean>(false)
//用户信息的收集
let userParams = reactive<userAddOrUp>({
  username: '',
  nickname: '',
  password: ''
})
//收集充值弹窗数据
let payForm = reactive<userPay>({
  userId: '',
  price: 0
})
//获取充值表单
let form = ref()
//获取表单
const formRef = ref<any>()
//组件挂载
onMounted(() => {
  getUserList()
})
//获取用户信息
let getUserList = async (pager = 1) => {
  pageNo.value = pager
  let res: UserResponseData = await reqUserInfo(pageNo.value, pageSize.value, username.value)
  if (res.code == 0) {
    userArr.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//分页器下拉菜单的自定义事件的回调
let handler = () => {
  getUserList()
}
//搜索按钮的回调
let search = () => {
  getUserList()
  username.value = ''
}
//重置按钮的回调
let reset = () => {
  settingStore.refsh = !settingStore.refsh
}

//添加用户的按钮回调
let addUser = () => {
  //抽屉打开
  drawer.value = true
  //清空表单数据
  Object.assign(userParams, {
    username: '',
    nickname: '',
    password: '',
    id: 0
  })
  //清空上一次表单校验错误提示
  nextTick(() => {
    formRef.value.clearValidate('username')
    formRef.value.clearValidate('nickname')
    formRef.value.clearValidate('password')
  })
}
//更新用户的按钮回调
let updateUser = (row: User) => {
  //打开抽屉
  drawer.value = true
  //存储已有的用户数据
  Object.assign(userParams, row)
  //清空上一次表单校验错误结果
  nextTick(() => {
    formRef.value.clearValidate('username')
    formRef.value.clearValidate('nickname')
    formRef.value.clearValidate('password')
  })
}
//删除用户的按钮回调
let deleteUser = async (id: number) => {
  let res: any = await reqRemoveUser(id)
  if (res.code == 0) {
    ElMessage({ type: 'success', message: '删除用户成功' })
    getUserList(userArr.value.length > 1 ? pageNo.value : pageNo.value - 1)
  }
}

//点击添加|修改用户抽屉的确定按钮
let save = async () => {
  //校验合格再发请求
  await formRef.value.validate()
  let res: any = await reqAddOrUpdateUser(userParams)
  if (res.code == 0) {
    //抽屉关闭
    drawer.value = false
    //提示添加|修改成功
    ElMessage({ type: 'success', message: userParams.id ? '更新成功' : '添加成功' })
    //获取数据
    getUserList()
  } else {
    drawer.value = false
    ElMessage({ type: 'error', message: userParams.id ? '更新失败' : '添加失败' })
  }
}
//密码校验
const validatorPassword = (_rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请输入密码'))
  } else if (value.length < 5) {
    callback(new Error('密码长度至少五位'))
  } else {
    callback()
  }
}
//添加|修改用户校验规则
let rules = {
  username: [{ required: true, trigger: 'blur', message: '请输入用户名' }],
  nickname: [{ required: true, trigger: 'blur', message: '请输入昵称' }],
  password: [{ required: true, trigger: 'blur', validator: validatorPassword }]
}
//价格校验
let validatePrice = (_rule: any, value: any, callback: any) => {
  if (value == null) {
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
//充值校验规则
let payRules = {
  price: [{ required: true, trigger: 'blur', validator: validatePrice }]
}
//充值弹窗的确定按钮
//弹窗确定按钮
let savePay = async () => {
  await form.value.validate()
  let res: any = await reqUserPay(payForm)
  if (res.code == 0) {
    //弹窗关闭
    dialogVisible.value = false
    //提示充值成功
    ElMessage({ type: 'success', message: '充值成功' })
    //获取数据
    getUserList()
  } else {
    //弹窗关闭
    dialogVisible.value = false
    ElMessage({ type: 'error', message: res.message })
  }
}
//充值按钮的回调
let addPay = (row: any) => {
  dialogVisible.value = true
  //存储已有表单数据
  payForm.userId = row.id
  payForm.price = row.balance
  userParams.username = row.username
  //清空上一次表单校验错误结果
  nextTick(() => {
    form.value.clearValidate('userId')
    form.value.clearValidate('price')
  })
}
</script>

<style scoped>
.form {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.balance {
  width: 130px;
  padding: 0 15px;
  display: flex;
  justify-content: space-between;
}
</style>
