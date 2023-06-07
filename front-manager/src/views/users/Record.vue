<template>
  <div>
    <!-- 面包屑 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>运营管理</el-breadcrumb-item>
      <el-breadcrumb-item>充值页面</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片 -->
    <el-card>
      <!-- 按钮 -->
      <el-button type="warning" style="margin-bottom: 15px" @click="showDialog">充值</el-button>
      <!-- 表格 -->
      <el-table style="width: 100%" border :data="recordList">
        <el-table-column type="index" label="序号" width="100" align="center"> </el-table-column>
        <el-table-column prop="payCode" label="订单编号" width="width"> </el-table-column>
        <el-table-column prop="username" label="姓名" width="100"> </el-table-column>
        <el-table-column prop="price" label="支付价格" width="120"> </el-table-column>
        <el-table-column prop="state" label="充值状态" width="120">
          <template slot-scope="{ row }">
            <el-tag v-if="row.state === 0">正在充值</el-tag>
            <el-tag v-else-if="row.state === 1" type="success">充值成功</el-tag>
            <el-tag v-else type="danger">取消充值</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="payAt" label="支付时间" width="width"> </el-table-column>
        <el-table-column prop="createAt" label="创建时间" width="width"> </el-table-column>
        <el-table-column prop="updateAt" label="修改时间" width="width"> </el-table-column>
      </el-table>
      <!-- 分页器 -->
      <el-pagination
        style="margin-top: 10px; text-align: center"
        :current-page.sync="recordInfo.pageNo"
        :page-count="5"
        :page-sizes="[3, 5, 10]"
        :page-size="recordInfo.pageSize"
        layout=" prev, pager, next, jumper,->,sizes,total"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-card>
    <!-- 充值的对话框 -->
    <el-dialog title="充值" :visible.sync="dialogFormVisible" width="30%">
      <el-form ref="addFormRef" label-width="80px" :model="payForm" :rules="payRules">
        <el-form-item label="用户名" prop="userId">
          <el-select placeholder="用户名" v-model="payForm.userId">
            <el-option :label="name.username" :value="name.id" v-for="name in userList" :key="name.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="充值金额" prop="price">
          <el-input v-model.number="payForm.price" placeholder="请输入充值金额"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="recordPay">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Record',
  data() {
    return {
      //获取用户列表的数据
      queryInfo: {
        username: '',
        //当前页
        pageNo: 1,
        //当前页展示的条数
        pageSize: 100
      },
      //获取用户充值列表
      recordInfo: {
        pageNo: 1,
        pageSize: 10
      },
      //用户列表数据
      userList: [],
      //充值列表
      recordList: [],
      //总条数
      total: 0,
      //充值的对话框
      dialogFormVisible: false,
      //充值展现的表单
      payForm: {
        userId: null,
        price: null
      },
      //验证规则
      payRules: {
        userId: [{ required: true, message: '请选择用户', trigger: 'change' }],
        price: [{ required: true, message: '请输入充值金额', trigger: 'blur' }]
      }
    }
  },
  mounted() {
    this.getRecordList()
    this.getUserList()
  },
  methods: {
    //获取用户列表的数据
    async getUserList() {
      let { data: res } = await this.$API.user.userList(this.queryInfo)
      if (res.code == 0) {
        this.userList = res.data.List
        this.queryInfo.pageNo = res.data.PageNo
        this.queryInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      } else {
        localStorage.clear()
        this.$router.push('/login')
      }
    },
    //获取用户充值的列表
    async getRecordList() {
      let { data: res } = await this.$API.user.recordList(this.recordInfo)
      if (res.code == 0) {
        this.recordList = res.data.List
        this.total = res.data.TotalCount
        this.recordInfo.pageNo = res.data.PageNo
        this.recordInfo.pageSize = res.data.PageSize
      }
    },
    //每条页数发生变化时触发
    handleSizeChange(pageSize) {
      this.recordInfo.pageSize = pageSize
      this.getRecordList()
    },
    //当前页数
    handleCurrentChange(pageNo) {
      this.recordInfo.pageNo = pageNo
      this.getRecordList()
    },
    //点击充值按钮，展现对话框
    showDialog() {
      //展现对话框
      this.dialogFormVisible = true
      //清除表单上一次遗留的数据
      this.payForm = { userId: null, price: null }
    },
    //充值对话框的确认按钮
    recordPay() {
      //表单验证通过，再发请求
      this.$refs.addFormRef.validate(async (valid) => {
        if (valid) {
          let { data: res } = await this.$API.user.recordPay(this.payForm)
          if (res.code == 0) {
            this.$message.success('充值成功')
            this.dialogFormVisible = false
            this.getRecordList()
          } else {
            this.$message.error('充值失败')
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style></style>
