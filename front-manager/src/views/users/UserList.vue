<template>
  <div>
    <!-- 面包屑 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>运营管理</el-breadcrumb-item>
      <el-breadcrumb-item>用户中心页面</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片 -->
    <el-card>
      <!-- input标签 -->
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入搜索内容" clearable v-model.trim="queryInfo.username" @clear="getUserList">
            <el-button slot="append" icon="el-icon-search" @click="getUserList"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addShow">添加用户</el-button>
        </el-col>
      </el-row>
      <!-- 列表 -->
      <el-table style="width: 100%" border :data="userList">
        <el-table-column type="index" label="序号" width="180" align="center"> </el-table-column>
        <el-table-column prop="username" label="姓名" width="180"> </el-table-column>
        <el-table-column prop="nickname" label="昵称" width="180"> </el-table-column>
        <el-table-column prop="balance" label="余额" width="180"> </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="width"> </el-table-column>
        <el-table-column prop="prop" label="操作" width="width">
          <template slot-scope="{ row }">
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="editShow(row)"></el-button>
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="handleDelet(row)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页器 -->
      <el-pagination
        style="margin-top: 10px; text-align: center"
        :current-page.sync="queryInfo.pageNo"
        :page-count="5"
        :page-sizes="[3, 5, 10]"
        :page-size="queryInfo.pageSize"
        layout=" prev, pager, next, jumper,->,sizes,total"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-card>
    <!-- 添加|修改用户的对话框 -->
    <el-dialog :title="addForm.id ? '修改用户' : '添加用户'" :visible.sync="dialogFormVisible" width="50%">
      <el-form :model="addForm" ref="addFormRef" :rules="rules" label-width="70px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model.trim="addForm.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="昵称" prop="nickname">
          <el-input v-model.trim="addForm.nickname" placeholder="请输入昵称"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model.trim="addForm.password" placeholder="请输入密码"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="addOrUpdateUser">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'UserList',
  data() {
    return {
      //获取用户列表的数据
      queryInfo: {
        //名字
        username: '',
        //当前页
        pageNo: 1,
        //当前页展示的条数
        pageSize: 10
      },
      //总条数
      total: 0,
      //用户列表数据
      userList: [],
      //添加用户的表单数据
      addForm: {
        username: '',
        nickname: '',
        password: ''
      },
      //校验规则
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 11 个字符', trigger: 'blur' }
        ],
        nickname: [
          { required: true, message: '请输入昵称', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 11 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 15, message: '长度在 6 到 15 个字符', trigger: 'blur' }
        ]
      },
      // 控制添加用户对话框显示与隐藏
      dialogFormVisible: false
    }
  },
  mounted() {
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
    //每条页数发生变化时触发
    handleSizeChange(pageSize) {
      this.queryInfo.pageSize = pageSize
      this.getUserList()
    },
    //当前页数
    handleCurrentChange(pageNo) {
      this.queryInfo.pageNo = pageNo
      this.getUserList()
    },
    //点击添加用户的按钮，展现dialog
    addShow() {
      //对话框展现
      this.dialogFormVisible = true
      //清除表单数据
      this.addForm = { username: '', nickname: '', createTime: '' }
    },
    //添加|修改用户
    addOrUpdateUser() {
      //表单验证通过再发请求
      this.$refs.addFormRef.validate(async (valid) => {
        if (valid) {
          let { data: res } = await this.$API.user.arrOrUpdateUserList(this.addForm)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: this.addForm.id ? '修改用户成功' : '添加用户成功'
            })
            //隐藏对话框
            this.dialogFormVisible = false
            //重新调列表数据
            this.getUserList()
          } else {
            this.$message.error('添加用户失败')
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    //删除用户
    handleDelet(row) {
      this.$confirm(`你确定删除${row.username}吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          //当点击确定按钮时，向服务器发请求
          let { data: res } = await this.$API.user.deletUserList(row.id)
          if (res.code == 0) {
            this.$message.success('删除成功！')
            //⭐1.利用Math.ceil算出新的分页 total为什么要-1,因为删完后total还没更新
            let totalPage = Math.ceil((this.total - 1) / this.queryInfo.pageSize)
            //⭐2.计算出当前页码
            this.queryInfo.pageNo = this.queryInfo.pageNo > totalPage ? totalPage : this.queryInfo.pageNo
            //⭐3.如果页码为1的话，则无效
            this.queryInfo.pageNo = this.queryInfo.pageNo < 1 ? 1 : this.queryInfo.pageNo
            //获取列表数据
            this.getUserList()
          }
        })
        .catch(() => {
          this.$message.info('已取消删除')
        })
    },
    //点击修改用户的按钮，展现dialog
    editShow(row) {
      //打开对话框
      this.dialogFormVisible = true
      //将已有的用户信息赋值给修改用户的表单
      this.addForm = { ...row }
    }
  }
}
</script>

<style scoped></style>
