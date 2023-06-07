<template>
  <div>
    <!-- 面包屑 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>设备信息</el-breadcrumb-item>
      <el-breadcrumb-item>充电桩</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片 -->
    <el-card>
      <!-- input标签 -->
      <el-row :gutter="20">
        <el-col :span="5">
          <el-select v-model="pileInfo.stationId" filterable placeholder="请输入搜索充电站">
            <el-option v-for="item in chargeList" :key="item.id" :label="item.name" :value="item.id"></el-option>
          </el-select>
          <el-button icon="el-icon-search" @click="getPileList"></el-button>
        </el-col>
        <el-col :span="2">
          <el-button type="primary" @click="addShow">添加充电桩</el-button>
        </el-col>
      </el-row>
      <!-- 列表 -->
      <el-table style="width: 100%" border :data="pileList">
        <el-table-column type="index" label="序号" width="80" align="center"> </el-table-column>
        <el-table-column prop="stationName" label="充电站名" width="width"> </el-table-column>
        <el-table-column prop="code" label="编号" width="120"> </el-table-column>
        <el-table-column prop="state" label="状态" width="120">
          <template slot-scope="{ row }">
            <el-tag v-if="row.state === 0">空闲中</el-tag>
            <el-tag v-else-if="row.state === 1" type="success">充电中</el-tag>
            <el-tag v-else type="danger">故障中</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createAt" label="创建时间" width="width"> </el-table-column>
        <el-table-column prop="updateAt" label="修改时间" width="width"> </el-table-column>
        <el-table-column prop="prop" label="操作" width="width">
          <template slot-scope="{ row }">
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="editPile(row)"></el-button>
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="pileDelete(row)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页器 -->
      <!-- -->
      <el-pagination
        style="margin-top: 10px; text-align: center"
        :current-page="pileInfo.pageNo"
        :page-count="5"
        :page-sizes="[3, 5, 10]"
        :page-size="pileInfo.pageSize"
        layout="prev, pager, next, jumper,->,sizes,total"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-card>
    <!-- 添加用户|修改的对话框 -->
    <el-dialog :title="pileform.id ? '修改充电桩' : '添加充电桩'" :visible.sync="dialogFormVisible" width="30%">
      <el-form ref="pileFormRef" label-width="80px" :model="pileform" :rules="pileRules">
        <el-form-item label="充电站名" prop="stationId">
          <el-select v-model="pileform.stationId" placeholder="请选择充电站">
            <el-option v-for="item in chargeList" :key="item.id" :label="item.name" :value="item.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="编号" prop="code" style="width: 290px">
          <el-input v-model.trim="pileform.code" placeholder="请输入编号"></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="state" v-show="pileform.id">
          <el-select placeholder="请选择状态" v-model="pileform.state">
            <el-option label="空闲中" value="0"></el-option>
            <el-option label="充电中" value="1"></el-option>
            <el-option label="故障" value="2"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="addOrUpdatePile">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Pile',
  data() {
    return {
      //获取充电站列表的数据
      chargeInfo: {
        name: '',
        address: '',
        //当前页
        pageNo: 1,
        //当前页展示的条数
        pageSize: 100
      },
      //获取充电桩列表的数据
      pileInfo: {
        stationId: '',
        //编号
        code: '',
        //当前页
        pageNo: 1,
        //当前页展示的条数
        pageSize: 5
      },
      //总条数
      total: 0,
      //充电站列表
      chargeList: [],
      //充电桩列表
      pileList: [],
      //对话框默认隐藏
      dialogFormVisible: false,
      //展示表单数据
      pileform: {
        stationId: null,
        code: '',
        state: null
      },
      //验证规则
      pileRules: {
        stationId: [{ required: true, message: '请选择充电站', trigger: 'change' }],
        code: [{ required: true, message: '请输入编号', trigger: 'blur' }]
      }
    }
  },
  mounted() {
    this.getPileList()
    this.getChargeList()
  },
  methods: {
    //充电站列表数据
    async getChargeList() {
      //发请求
      let { data: res } = await this.$API.charge.chargeList(this.chargeInfo)
      if (res.code == 0) {
        this.chargeList = res.data.List
        this.chargeInfo.pageNo = res.data.PageNo
        this.chargeInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      } else {
        console.log('err')
      }
    },
    //充电桩列表数据
    async getPileList() {
      let { data: res } = await this.$API.charge.pileList(this.pileInfo)
      if (res.code == 0) {
        this.pileList = res.data.List
        this.pileInfo.pageNo = res.data.PageNo
        this.pileInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      }
    },
    //每条页数发生变化时触发
    handleSizeChange(pageSize) {
      this.pileInfo.pageSize = pageSize
      this.getPileList()
    },
    //当前页数
    handleCurrentChange(pageNo) {
      this.pileInfo.pageNo = pageNo
      this.getPileList()
    },
    //删除充电桩
    async pileDelete(row) {
      this.$confirm(`你确定删除-"${row.stationName}"-吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          let { data: res } = await this.$API.charge.deletePileList(row.id)
          this.$message.success('删除成功！')
          //⭐1.利用Math.ceil算出新的分页 total为什么要-1,因为删完后total还没更新
          let totalPage = Math.ceil((this.total - 1) / this.pileInfo.pageSize)
          //⭐2.计算出当前页码
          this.pileInfo.pageNo = this.pileInfo.pageNo > totalPage ? totalPage : this.pileInfo.pageNo
          //⭐3.如果页码为1的话，则无效
          this.pileInfo.pageNo = this.pileInfo.pageNo < 1 ? 1 : this.pileInfo.pageNo
          //获取列表数据
          this.getPileList()
          console.log('充电站删除', res)
        })
        .catch(() => {
          this.$message.info('已取消删除')
        })
    },
    //展现添加的对话框
    addShow() {
      this.dialogFormVisible = true
      //清除表单数据
      this.pileform = { stationId: null, code: '', state: null }
    },
    //添加|修改充电桩
    addOrUpdatePile() {
      //表单验证通过再发请求
      this.$refs.pileFormRef.validate(async (valid) => {
        if (valid) {
          let { data: res } = await this.$API.charge.addOrUpdatePile(this.pileform)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: this.pileform.id ? '修改充电桩成功' : '添加充电桩成功'
            })
            //隐藏对话框
            this.dialogFormVisible = false
            //发请求，获取列表数据
            this.getPileList()
          }
        }
      })
    },
    //点击修改按钮的回调
    editPile(row) {
      this.dialogFormVisible = true
      this.pileform = { ...row }
    }
  }
}
</script>

<style></style>
