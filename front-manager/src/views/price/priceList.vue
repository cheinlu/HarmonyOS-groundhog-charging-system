<template>
  <div>
    <!-- 面包屑 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>价格规则管理</el-breadcrumb-item>
      <el-breadcrumb-item>区间充电价格</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片 -->
    <el-card>
      <el-button type="primary" style="margin-bottom: 15px" @click="showDialog">添加充电价格</el-button>
      <!-- 列表 -->
      <el-table style="width: 100%" border :data="priceList">
        <el-table-column type="index" label="序号" width="80" align="center"> </el-table-column>
        <el-table-column prop="startHour" label="开始时间" width="120"> </el-table-column>
        <el-table-column prop="endHour" label="结束时间" width="120"> </el-table-column>
        <el-table-column prop="price" label="价格" width="120"> </el-table-column>
        <el-table-column prop="createAt" label="创建时间" width="width"> </el-table-column>
        <el-table-column prop="updateAt" label="修改时间" width="width"> </el-table-column>
        <el-table-column prop="prop" label="操作" width="width">
          <template slot-scope="{ row }">
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="editPrice(row)"></el-button>
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="handleDelet(row)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页器 -->
      <!-- -->
      <el-pagination
        style="margin-top: 10px; text-align: center"
        :current-page="priceInfo.pageNo"
        :page-count="5"
        :page-sizes="[3, 5, 10]"
        :page-size="priceInfo.pageSize"
        layout="prev, pager, next, jumper,->,sizes,total"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-card>
    <!-- 添加价格的对话框 -->
    <el-dialog title="添加价格" :visible.sync="dialogFormVisible" width="50%">
      <el-form ref="priceRef" :rules="priceRules" label-width="70px" :model="priceForm">
        <el-form-item label="开始时间" prop="startHour">
          <el-input v-model.number="priceForm.startHour" placeholder="请输入充值金额"></el-input>
        </el-form-item>
        <el-form-item label="结束时间" prop="endHour">
          <el-input v-model.number="priceForm.endHour" placeholder="请输入充值金额"></el-input>
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input v-model.number="priceForm.price" placeholder="请输入充值金额"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="addAndUpdatePrice">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'priceList',
  data() {
    //自定义验证规则
    var validateStart = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('开始时间不能为空'))
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('请输入数字值,且不能与列表重复'))
        } else {
          if (value > 24 || value < 0) {
            callback(new Error('必须在0-24之间'))
          } else {
            callback()
          }
        }
      }, 1000)
    }
    var validateEndHour = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('结束时间不能为空'))
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('请输入数字值,且不能与列表重复'))
        } else {
          if (value > 24 || value < 0) {
            callback(new Error('必须在0-24之间'))
          } else {
            callback()
          }
        }
      }, 1000)
    }
    var validatePrice = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('价格不能为空'))
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('请输入价格'))
        } else {
          if (value < 0) {
            callback(new Error('价格必须大于0'))
          } else {
            callback()
          }
        }
      }, 1000)
    }
    return {
      priceInfo: {
        //当前页
        pageNo: 1,
        //当前页展示的条数
        pageSize: 5
      },
      //总条数
      total: 0,
      //价格区间数据
      priceList: [],
      //对话框默认隐藏
      dialogFormVisible: false,
      //展示的表单数据
      priceForm: {
        startHour: null,
        endHour: null,
        price: null
      },
      //验证规则
      priceRules: {
        startHour: [{ validator: validateStart, trigger: 'blur' }],
        endHour: [{ validator: validateEndHour, trigger: 'blur' }],
        price: [{ validator: validatePrice, trigger: 'blur' }]
      }
    }
  },
  mounted() {
    this.getPriceList()
  },
  methods: {
    //价格区间列表数据
    async getPriceList() {
      //发请求
      let { data: res } = await this.$API.price.priceList(this.priceInfo)
      if (res.code == 0) {
        this.priceList = res.data.List
        this.priceInfo.pageNo = res.data.PageNo
        this.priceInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      }
    },
    //每条页数发生变化时触发
    handleSizeChange(pageSize) {
      this.priceInfo.pageSize = pageSize
      this.getPriceList()
    },
    //当前页数
    handleCurrentChange(pageNo) {
      this.priceInfo.pageNo = pageNo
      this.getPriceList()
    },
    //删除区间价格
    handleDelet(row) {
      this.$confirm(`你确定删除这个区间价格吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          //当点击确定按钮时，向服务器发请求
          let { data: res } = await this.$API.price.priceDelete(row.id)
          if (res.code == 0) {
            this.$message.success('删除成功！')
            //获取列表数据
            this.getPriceList()
          }
        })
        .catch(() => {
          this.$message.info('已取消删除')
        })
    },
    //点击添加|修改展现对话框
    showDialog() {
      //展现对话框
      this.dialogFormVisible = true
      //清除表单数据
      this.priceForm = { startHour: null, endHour: null, price: null }
    },
    //添加|修改的确定按钮
    addAndUpdatePrice() {
      //表单验证通过再发请求
      this.$refs.priceRef.validate(async (valid) => {
        if (valid) {
          let { data: res } = await this.$API.price.addOrUpdatePrice(this.priceForm)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: this.priceForm.id ? '修改充电价格成功' : '添加充电价格成功'
            })
            //隐藏对话框
            this.dialogFormVisible = false
            //如果修改停留在当前页，添加停留在第一页
            this.getPriceList()
          }
        }
      })
    },
    //修改的按钮
    editPrice(row) {
      //打开对话框
      this.dialogFormVisible = true
      //参数赋值
      this.priceForm = { ...row }
    }
  }
}
</script>

<style></style>
