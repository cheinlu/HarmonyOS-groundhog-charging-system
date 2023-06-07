<template>
  <div>
    <!-- 面包屑 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>订单管理</el-breadcrumb-item>
      <el-breadcrumb-item>充电订单</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片 -->
    <el-card>
      <!-- 标签 -->
      <el-row :gutter="22">
        <el-col :span="3">
          <el-input placeholder="订单编号" clearable v-model.trim="chargeQueryInfo.orderCode" @clear="getChargeOrderList"> </el-input>
        </el-col>
        <el-col :span="3">
          <el-select placeholder="用户名" v-model="chargeQueryInfo.userId">
            <el-option :label="user.nickname" :value="user.id" v-for="user in userList" :key="user.id"></el-option>
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select placeholder="充电站名" v-model="chargeQueryInfo.stationId" @change="changeStation">
            <el-option :label="charge.name" :value="charge.id" v-for="charge in chargeList" :key="charge.id"></el-option>
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-select placeholder="充电桩id" v-model="chargeQueryInfo.pileId">
            <el-option :label="pile.code" :value="pile.id" v-for="pile in filteredPileList" :key="pile.id"></el-option>
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-date-picker v-model="chargeQueryInfo.beforeAt" type="datetime" placeholder="选择开始时间" @clear="getChargeOrderList"> </el-date-picker>
        </el-col>
        <el-col :span="4">
          <el-date-picker v-model="chargeQueryInfo.afterAt" type="datetime" placeholder="选择结束时间" @clear="getChargeOrderList"> </el-date-picker>
        </el-col>
        <el-col :span="2">
          <el-select placeholder="订单状态" v-model="chargeQueryInfo.state">
            <el-option label="充电中" value="0"></el-option>
            <el-option label="已完成" value="1"></el-option>
            <el-option label="未付帐" value="2"></el-option>
          </el-select>
        </el-col>
        <el-button icon="el-icon-search" type="primary" @click="getChargeOrderList"></el-button>
      </el-row>
      <!-- 列表 -->
      <el-table style="width: 100%" border :data="chargeOrderList">
        <el-table-column type="index" label="序号" width="80" align="center"> </el-table-column>
        <el-table-column prop="orderCode" label="订单编号" width="width"> </el-table-column>
        <el-table-column prop="nickname" label="名字" width="120"> </el-table-column>
        <el-table-column prop="stationName" label="站名" width="120"> </el-table-column>
        <el-table-column prop="pileCode" label="充电桩编号" width="100"> </el-table-column>
        <el-table-column prop="startAt" label="充电开始时间" width="width"> </el-table-column>
        <el-table-column prop="stopAt" label="充电结束时间" width="width"> </el-table-column>
        <el-table-column prop="state" label="订单状态" width="80">
          <template slot-scope="{ row }">
            <el-tag v-if="row.state === 0">充电中</el-tag>
            <el-tag v-else-if="row.state === 1" type="success">已完成</el-tag>
            <el-tag v-else type="danger">未付账</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="价格" width="80"> </el-table-column>
        <el-table-column prop="createAt" label="创建时间" width="width"> </el-table-column>
        <el-table-column prop="updateAt" label="修改时间" width="width"> </el-table-column>
        <el-table-column prop="prop" label="操作" width="80">
          <template slot-scope="{ row }">
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="editChargeOrder(row)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页器 -->
      <el-pagination
        style="margin-top: 10px; text-align: center"
        :current-page.sync="chargeQueryInfo.pageNo"
        :page-count="5"
        :page-sizes="[3, 5, 10]"
        :page-size="chargeQueryInfo.pageSize"
        layout=" prev, pager, next, jumper,->,sizes,total"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-card>
    <!-- 修改的对话框 -->
    <el-dialog title="修改充电订单" :visible.sync="dialogFormVisible" width="30%">
      <el-form ref="pileFormRef" label-width="70px" :model="chargeOrderForm">
        <el-form-item label="订单状态">
          <el-select placeholder="请选择订单状态" v-model="chargeOrderForm.state">
            <el-option label="充电中" value="0"></el-option>
            <el-option label="已完成" value="1"></el-option>
            <el-option label="未付账" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="价格" prop="code" style="width: 290px">
          <el-input v-model.trim="chargeOrderForm.price" type="number" placeholder="请输入价格"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="updateChargeOrder">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ChargeOrder',
  data() {
    return {
      chargeQueryInfo: {
        //订单编号
        orderCode: '',
        //用户id，使用下拉框
        userId: '',
        //充电站id，使用下拉框
        stationId: '',
        //充电桩编号，使用下拉框，下拉框与 充电站下拉框联动
        pileId: '',
        //充电开始时间（大于10点
        beforeAt: '',
        //充电开始时间（小于12点
        afterAt: '',
        //订单状态（0:充电中,1:已完成,2:未付账）
        state: null,
        //当前页
        pageNo: 1,
        //每页展示的条数
        pageSize: 10
      },
      //用于存储用户选择的充电站id
      selectedStationId: '',
      //总条数
      total: 0,
      //订单管理列表
      chargeOrderList: [],
      //获取用户列表的数据
      queryInfo: {
        //当前页
        pageNo: 1,
        //当前页展示的条数
        pageSize: 5
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
      //用户列表数据
      userList: [],
      //充电站列表数据
      chargeList: [],
      //充电桩列表
      pileList: [],
      //对话框表单的数据
      chargeOrderForm: {
        id: '',
        state: null,
        price: null
      },
      //对话框默认隐藏
      dialogFormVisible: false
    }
  },
  computed: {
    //对 pileList 进行过滤，只保留与 selectedStationId 相等的充电桩
    filteredPileList() {
      return this.pileList.filter((pile) => pile.stationId === this.selectedStationId)
    }
  },
  mounted() {
    this.getChargeOrderList()
    this.getUserList()
    this.getChargeList()
    this.getPileList()
  },
  methods: {
    //充电订单列表
    async getChargeOrderList() {
      let { data: res } = await this.$API.order.chargeOrderList(this.chargeQueryInfo)
      if (res.code == 0) {
        this.chargeOrderList = res.data.List
        this.chargeQueryInfo.pageNo = res.data.PageNo
        this.chargeQueryInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      }
    },
    //每条页数发生变化时触发
    handleSizeChange(pageSize) {
      this.chargeQueryInfo.pageSize = pageSize
      this.getChargeOrderList()
    },
    //当前页数
    handleCurrentChange(pageNo) {
      this.chargeQueryInfo.pageNo = pageNo
      this.getChargeOrderList()
    },
    //获取用户列表的数据
    async getUserList() {
      //发请求
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
    //充电站列表数据
    async getChargeList() {
      //发请求
      let { data: res } = await this.$API.charge.chargeList(this.queryInfo)
      if (res.code == 0) {
        this.chargeList = res.data.List
        this.queryInfo.pageNo = res.data.PageNo
        this.queryInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      } else {
        console.log('err')
      }
      console.log('充电站列表', this.chargeList)
    },
    //充电桩列表数据
    async getPileList() {
      //发请求
      let { data: res } = await this.$API.charge.pileList(this.pileInfo)
      if (res.code == 0) {
        this.pileList = res.data.List
        this.pileInfo.pageNo = res.data.PageNo
        this.pileInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      }
    },
    //修改充电订单按钮
    editChargeOrder(row) {
      //打开对话框
      this.dialogFormVisible = true
      //对表单里面的数据进行赋值
      this.chargeOrderForm.id = row.id
      this.chargeOrderForm.price = row.price
      this.chargeOrderForm.state = row.state
    },
    //修改订单的回调函数
    async updateChargeOrder() {
      let { data: res } = await this.$API.order.updateChargeOrder(this.chargeOrderForm)
      if (res.code == 0) {
        this.$message.success('修改充电订单成功')
        this.dialogFormVisible = false
        this.getChargeOrderList()
      }
    },
    //监听用户选择充电站时更新 selectedStationId 的值
    changeStation() {
      this.selectedStationId = this.chargeQueryInfo.stationId
    }
  }
}
</script>

<style></style>
