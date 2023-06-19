<template>
  <div>
    <!-- 面包屑 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>设备信息</el-breadcrumb-item>
      <el-breadcrumb-item>充电站</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片 -->
    <el-card>
      <!-- input标签 -->
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入搜索内容" v-model.trim="chargeInfo.name" clearable @clear="getChargeList">
            <el-button slot="append" icon="el-icon-search" @click="getChargeList"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addShow">添加充电站</el-button>
        </el-col>
      </el-row>
      <!-- 列表 -->
      <el-table style="width: 100%" border :data="chargeList">
        <el-table-column type="index" label="序号" width="80" align="center"> </el-table-column>
        <el-table-column prop="name" label="站名" width="180"> </el-table-column>
        <el-table-column prop="address" label="地址" width="width"> </el-table-column>
        <el-table-column prop="coordinate" label="坐标" width="width"> </el-table-column>
        <el-table-column prop="prop" label="查看地图" width="90" align="center">
          <template slot-scope="{ row }">
            <el-button type="success" icon="el-icon-position" size="mini" @click="showMap(row)"></el-button>
          </template>
        </el-table-column>
        <el-table-column prop="createAt" label="创建时间" width="width"> </el-table-column>
        <el-table-column prop="updateAt" label="修改时间" width="width"> </el-table-column>
        <el-table-column prop="prop" label="操作" width="160">
          <template slot-scope="{ row }">
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="chargeUpdate(row)"></el-button>
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="chargeDelet(row)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页器 -->
      <!--  -->
      <el-pagination
        style="margin-top: 10px; text-align: center"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page.sync="chargeInfo.pageNo"
        :page-count="5"
        :page-sizes="[3, 5, 10]"
        :page-size="chargeInfo.pageSize"
        layout="prev, pager, next, jumper,->,sizes,total"
        :total="total"
      >
      </el-pagination>
    </el-card>
    <!-- 添加用户的对话框 -->
    <el-dialog :title="chargeForm.id ? '修改充电站' : '添加充电站'" :visible.sync="dialogFormVisible" width="50%">
      <el-form :model="chargeForm" ref="chargeFormRef" :rules="chargeRules" label-width="70px">
        <el-form-item label="站名" prop="name">
          <el-input v-model.trim="chargeForm.name" placeholder="请输入站名"></el-input>
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model.trim="chargeForm.address" placeholder="请输入地址"></el-input>
        </el-form-item>
        <el-form-item label="坐标" prop="coordinate">
          <template>
            <baidu-map :center="center" :zoom="zoom" @ready="handler" style="height: 400px; width: 700px" @click="getClickInfo" :scroll-wheel-zoom="true"> </baidu-map>
          </template>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="addOrUpdateCharge">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 显示地图的对话框 -->
    <el-dialog :visible.sync="mapDialogVisible" width="750px">
      <baidu-map :center="center" :zoom="zoom" @ready="handler" style="height: 400px; width: 700px" @click="getClickInfo" :scroll-wheel-zoom="true"> </baidu-map>
    </el-dialog>
  </div>
</template>

<script>
import BaiduMap from 'vue-baidu-map/components/map/Map.vue'
export default {
  name: 'charge',
  components: {
    BaiduMap
  },
  data() {
    return {
      map: null,
      BMap: null,
      //默认坐标
      center: { lng: 113.909401, lat: 22.590799 },
      zoom: 13,
      //获取充电站列表的数据
      chargeInfo: {
        name: '',
        address: '',
        //当前页
        pageNo: 1,
        //当前页展示的条数
        pageSize: 10
      },
      //总条数
      total: 0,
      //充电站列表数据
      chargeList: [],
      //添加|修改的对话框
      dialogFormVisible: false,
      //展示地图的对话框
      mapDialogVisible: false,
      //表单展现的数据
      chargeForm: {
        name: '',
        address: '',
        coordinate: ''
      },
      //验证规则
      chargeRules: {
        name: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        address: [{ required: true, message: '请输入地址', trigger: 'blur' }]
      }
    }
  },
  mounted() {
    this.getChargeList()
  },
  methods: {
    //充电站列表数据
    async getChargeList() {
      let { data: res } = await this.$API.charge.chargeList(this.chargeInfo)
      console.log('res', res)
      if (res.code == 0) {
        this.chargeList = res.data.List
        this.chargeInfo.pageNo = res.data.PageNo
        this.chargeInfo.pageSize = res.data.PageSize
        this.total = res.data.TotalCount
      } else {
        console.log('err')
      }
    },
    //每条页数发生变化时触发
    handleSizeChange(pageSize) {
      this.chargeInfo.pageSize = pageSize
      this.getChargeList()
    },
    //当前页数
    handleCurrentChange(pageNo) {
      this.chargeInfo.pageNo = pageNo
      this.getChargeList()
    },
    //删除充电站
    chargeDelet(row) {
      this.$confirm(`你确定删除${row.name}吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          let { data: res } = await this.$API.charge.deletChargeList(row.id)
          this.$message.success('删除成功！')
          //⭐1.利用Math.ceil算出新的分页 total为什么要-1,因为删完后total还没更新
          let totalPage = Math.ceil((this.total - 1) / this.chargeInfo.pageSize)
          //⭐2.计算出当前页码
          this.chargeInfo.pageNo = this.chargeInfo.pageNo > totalPage ? totalPage : this.chargeInfo.pageNo
          //⭐3.如果页码为1的话，则无效
          this.chargeInfo.pageNo = this.chargeInfo.pageNo < 1 ? 1 : this.chargeInfo.pageNo
          //获取列表数据
          this.getChargeList()
          console.log('充电站删除', res)
        })
        .catch(() => {
          this.$message.info('已取消删除')
        })
    },
    //点击添加用户的按钮，展现dialog
    addShow() {
      //对话框展现
      this.dialogFormVisible = true
      //清除表单数据
      this.chargeForm = { name: '', address: '', coordinate: '' }
    },
    //添加|修改的确认按钮
    addOrUpdateCharge() {
      this.$refs.chargeFormRef.validate(async (valid) => {
        if (valid) {
          console.log('valid', valid)
          let { data: res } = await this.$API.charge.addOrUpdateChargeList(this.chargeForm)
          console.log('this.chargeForm', this.chargeForm)
          console.log('添加', res)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: this.chargeForm.id ? '修改充电站成功' : '添加充电站成功'
            })
            //隐藏对话框
            this.dialogFormVisible = false
            //如果修改停留在当前页，添加停留在第一页
            this.getChargeList(this.chargeForm.id ? this.chargeInfo.pageNo : 1)
          }
        }
      })
    },
    //修改某一个充电站
    chargeUpdate(row) {
      //对话框展示
      this.dialogFormVisible = true
      this.chargeForm = { ...row }
    },
    //地图的回调
    handler({ BMap, map }) {
      console.log('1 mapmap ', map)
      this.map = map
      this.BMap = BMap
      console.log('2 this.map--', this.map)
      if (!BMap) {
        console.warn('BMap 对象还未准备就绪')
        return
      }

      if (this.chargeForm.coordinate) {
        console.log('@', this.chargeForm.coordinate)
        // 构建 BMap.Point 对象
        // const point = new BMap.Point(this.chargeForm.coordinate.split(',')[0], this.chargeForm.coordinate.split(',')[1])
        // // 根据 BMap 实例的 setCenter 方法将地图中心设置为 point 坐标，并调整地图缩放比例
        // this.center = { lng: point.lng, lat: point.lat }
        // this.zoom = 16
        let point = new BMap.Point(this.chargeForm.coordinate.split(',')[0], this.chargeForm.coordinate.split(',')[1])
        map.centerAndZoom(point, 13)
        var marker = new BMap.Marker(point) // 创建标注
        map.addOverlay(marker) // 将标注添加到地图中
        var circle = new BMap.Circle(point, 6, { strokeColor: 'Red', strokeWeight: 6, strokeOpacity: 1, Color: 'Red', fillColor: '#f03' })
        map.addOverlay(circle)
      } else {
        var point = new BMap.Point(113.909401, 22.590799)
        map.centerAndZoom(point, 13)
      }
    },
    getClickInfo(e) {
      this.center.lng = e.point.lng
      this.center.lat = e.point.lat
      let strLng = String(e.point.lng)
      let strLat = String(e.point.lat)
      let newCenter = strLng.concat(',', strLat)
      this.chargeForm.coordinate = newCenter
    },
    showMap(row) {
      this.chargeForm.coordinate = row.coordinate
      // 弹出地图对话框
      this.mapDialogVisible = true

      this.$nextTick(() => {
        console.log('3 this.map++', this.map)
        if (this.map) {
          //清除标注
          this.map.clearOverlays()
          let point = new this.BMap.Point(this.chargeForm.coordinate.split(',')[0], this.chargeForm.coordinate.split(',')[1])
          this.map.centerAndZoom(point, 13)
          var marker = new this.BMap.Marker(point) // 创建标注
          this.map.addOverlay(marker) // 将标注添加到地图中
          var circle = new this.BMap.Circle(point, 6, { strokeColor: 'Red', strokeWeight: 6, strokeOpacity: 1, Color: 'Red', fillColor: '#f03' })
          this.map.addOverlay(circle)
        }
      })
    }
  }
}
</script>

<style>
.charge-container {
  text-align: center;
}
</style>
