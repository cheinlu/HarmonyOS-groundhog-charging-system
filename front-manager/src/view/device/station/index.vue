<template>
  <el-card style="height: 70px">
    <el-form :inline="true" class="form">
      <el-form-item :label="$t('equip.nameLabel')">
        <el-input :placeholder="$t('equip.nameLabel')" v-model="name"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search" :disabled="name ? false : true">{{ $t('button.search') }}</el-button>
        <el-button type="primary" @click="reset">{{ $t('button.reset') }}</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card style="margin: 10px 0px">
    <el-button type="primary" size="default" @click="addStation" v-has="'ChargeStationAdd'">{{ $t('equip.addStation') }}</el-button>
    <!-- table展示充电站信息 -->
    <el-table style="margin: 10px 0px" border :data="stationArr">
      <el-table-column label="#" align="center" type="index"></el-table-column>
      <el-table-column label="ID" align="center" prop="id" width="80"></el-table-column>
      <el-table-column :label="$t('equip.nameLabel')" align="center" prop="name" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('equip.addressLabel')" align="center" prop="address" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('equip.viewMap')" align="center" prop="prop" width="100">
        <template #="{ row }">
          <el-button type="primary" icon="Position" size="small" @click="showMap(row)"></el-button>
        </template>
      </el-table-column>
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
      <el-table-column :label="$t('tabel.operate')" width="180px" align="center">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="updateStation(row)" v-has="'ChargeStationUpdate'">{{ $t('button.edit') }}</el-button>
          <el-popconfirm :title="`你确定要删除${row.name}?`" @confirm="deleteStation(row.id)" width="260px">
            <template #reference>
              <el-button type="primary" size="small" icon="Delete" v-has="'ChargeStationDel'">{{ $t('button.delete') }}</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页器-->
    <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[3, 5, 7, 10]" :background="true" layout="prev, pager, next, jumper,->,sizes,total" :total="total" @current-change="getHasStation" @size-change="handler" />
  </el-card>
  <!-- 抽屉结构:完成添加新的用户账号|更新已有的账号信息 -->
  <el-drawer v-model="drawer" size="50%">
    <!-- 头部标题:将来文字内容应该动态的 -->
    <template #header>
      <h4>{{ $t(chargeForm.id ? 'equip.editStation' : 'equip.addStation') }}</h4>
    </template>
    <!-- 身体部分 -->
    <template #default>
      <el-form :model="chargeForm" ref="chargeFormRef" :rules="chargeRules" label-width="77px">
        <el-form-item :label="$t('equip.nameLabel')" prop="name">
          <el-input :placeholder="$t('equip.nameLabel')" v-model="chargeForm.name"></el-input>
        </el-form-item>
        <el-form-item :label="$t('equip.addressLabel')" prop="address">
          <el-input :placeholder="$t('equip.addressLabel')" v-model="chargeForm.address"></el-input>
        </el-form-item>
        <el-form-item :label="$t('equip.imageUrl')" prop="imageUrl">
          <el-upload class="avatar-uploader" :show-file-list="false" :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
            <img v-if="chargeForm.imageUrl" :src="chargeForm.imageUrl" class="avatar" />
            <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item :label="$t('equip.coordinate')" prop="coordinate">
          <template #="{}">
            <baidu-map class="map" :center="center" :zoom="zoom" @ready="ready" @click="getClickInfo" :scroll-wheel-zoom="true">
              <bm-marker :position="selectedPoint"></bm-marker>
              <bm-circle v-show="showDefaultCircle" :center="center" :radius="3" :options="{ strokeColor: 'Red', strokeWeight: 6, strokeOpacity: 1, color: 'Red', fillColor: '#f03' }"></bm-circle>
            </baidu-map>
          </template>
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancel">{{ $t('pop.cancel') }}</el-button>
        <el-button type="primary" @click="save">{{ $t('pop.confirm') }}</el-button>
      </div>
    </template>
  </el-drawer>
  <!-- 查看地图的弹窗 -->
  <el-dialog v-model="dialogVisible" width="660px" :center="dialogCenter" @close="close">
    <baidu-map class="map" :center="center" :zoom="zoom" @ready="ready" @click="getClickInfo" :scroll-wheel-zoom="true">
      <bm-marker :position="selectedPoint"></bm-marker>
      <bm-circle v-show="showDefaultCircle" :center="center" :radius="3" :options="{ strokeColor: 'Red', strokeWeight: 6, strokeOpacity: 1, color: 'Red', fillColor: '#f03' }"></bm-circle>
    </baidu-map>
  </el-dialog>
</template>

<script setup lang="ts">
import BMap from 'vue-baidu-map-3x'
import { ref, onMounted, reactive, nextTick } from 'vue'
import { reqStation, reqRemoveStation, reqAddOrUpdateStation, reqUpload } from '@/api/device/station/index.ts'
import type { stationList, stations, stationAddOrUpdate } from '@/api/device/station/type'
import useLayOutSettingStore from '@/store/module/setting'
import { ElMessage } from 'element-plus'
import type { UploadProps } from 'element-plus'
import { formatDate } from '@/utils/time'
let settingStore = useLayOutSettingStore()
//设置地图的初始中心位置
let center = reactive({ lng: 113.909401, lat: 22.590799 })
//地图缩放级别
let zoom = ref(13)
//更新为所点击位置的经纬度坐标
let selectedPoint = ref({ lng: 0, lat: 0 })
//显示默认圆形覆盖物
let showDefaultCircle = ref<boolean>(true)
//在初始状态下没有任何圆形覆盖物
let defaultCircleOverlay = ref(null)
let chargeFormRef = ref()
//默认页码
let pageNo = ref<number>(1)
//一页展示几条数据
let pageSize = ref<number>(3)
//用户总个数
let total = ref<number>(0)
//抽屉默认关闭
let drawer = ref<boolean>(false)
//充电站名
let name = ref('')
//存储全部充电站的数组
let stationArr = ref<stations>([])
//点击查看地图按钮，对话框默认隐藏
let dialogVisible = ref<boolean>(false)
//实现在弹窗显示地图时定位在弹窗的中心
let dialogCenter = center
//站点信息的收集
let chargeForm = reactive<stationAddOrUpdate>({
  name: '',
  address: '',
  coordinate: '',
  imageUrl: '',
  id: 0
})
onMounted(() => {
  getHasStation()
})

//获取充电站列表数据
let getHasStation = async (pager = 1) => {
  pageNo.value = pager
  let res: stationList = await reqStation(pageNo.value, pageSize.value, name.value)
  if (res.code == 0) {
    stationArr.value = res.data.List
    pageNo.value = res.data.PageNo
    pageSize.value = res.data.PageSize
    total.value = res.data.TotalCount
  }
}
//分页器下拉菜单的自定义事件的回调
let handler = () => {
  getHasStation()
}

//搜索按钮的回调
let search = () => {
  getHasStation()
  name.value = ''
}

//重置按钮的回调
let reset = () => {
  settingStore.refsh = !settingStore.refsh
}
//添加充电站按钮
let addStation = () => {
  //打开抽屉
  drawer.value = true
  //添加按钮时，默认没有选中的经纬度
  selectedPoint.value.lng = 0
  selectedPoint.value.lat = 0
  //清空表单数据
  Object.assign(chargeForm, {
    name: '',
    address: '',
    coordinate: '',
    imageUrl: '',
    id: 0
  })
  //清空上一次表单校验错误提示
  nextTick(() => {
    chargeFormRef.value.clearValidate('name')
    chargeFormRef.value.clearValidate('address')
    chargeFormRef.value.clearValidate('coordinate')
    chargeFormRef.value.clearValidate('imageUrl')
  })
}
//修改充电站按钮
let updateStation = (row: any) => {
  //打开抽屉
  drawer.value = true
  //合并参数
  Object.assign(chargeForm, row)
  // 更新selectedPoint的值
  selectedPoint.value.lng = row.coordinate.split(',')[0]
  selectedPoint.value.lat = row.coordinate.split(',')[1]
  //更新经纬度后定位在弹窗中心
  dialogCenter.lng = row.coordinate.split(',')[0]
  dialogCenter.lat = row.coordinate.split(',')[1]

  // 调用ready方法重新渲染地图
  ready({ BMap, map: null }) // 传入null代表不对map进行操作，只是触发更新
}
//删除充电站按钮
let deleteStation = async (id: number) => {
  let res: any = await reqRemoveStation(id)
  if (res.code == 0) {
    ElMessage({ type: 'success', message: '删除站点成功' })
    getHasStation(stationArr.value.length > 1 ? pageNo.value : pageNo.value - 1)
  }
}
//点击添加|修改充电站抽屉的取消按钮
let cancel = () => {
  //重新刷新，清空表单数据
  settingStore.refsh = !settingStore.refsh
  //关闭抽屉
  drawer.value = false
}
//点击添加|修改充电站抽屉的确定按钮
let save = async () => {
  //表单校验合格再发请求
  await chargeFormRef.value.validate()
  let res: any = await reqAddOrUpdateStation(chargeForm)
  if (res.code == 0) {
    //抽屉关闭
    drawer.value = false
    //提示添加成功
    ElMessage({ type: 'success', message: chargeForm.id ? '修改充电站成功' : '添加充电站成功' })
    //获取数据
    getHasStation()
  } else {
    //抽屉关闭
    drawer.value = false
    ElMessage({ type: 'error', message: chargeForm.id ? '修改充电站失败' : '添加充电站失败' })
  }
}
//地图的回调
const ready = ({ BMap, map }: { BMap: any; map: any }) => {
  if (!BMap) {
    console.warn('BMap 对象还未准备就绪')
    return
  }
  //创建地图上的点
  let point = new BMap.Point(center.lng, center.lat)
  //将地图中心设置为这个点
  map.setCenter(point)
  //缩放度13
  map.setZoom(13)
  //判断 defaultCircleOverlay.value 是否为真，来添加或删除红色圆形覆盖物。
  if (defaultCircleOverlay.value) {
    //下面2行用来创建圆形覆盖物，并添加到地图上
    const circle = new BMap.Circle(point, 3, { strokeColor: 'Red', strokeWeight: 6, strokeOpacity: 1, color: 'Red', fillColor: '#f03' })
    map.addOverlay(circle)
    //下面2行用来标记默认圆形覆盖物的存在
    defaultCircleOverlay.value = circle
    showDefaultCircle.value = true
  }
}
//点击地图触发的事件
let getClickInfo = (e: any) => {
  //获取点击位置的经纬度
  selectedPoint.value.lng = e.point.lng
  selectedPoint.value.lat = e.point.lat
  //将获取到的经纬度转换成字符串形式，并用逗号进行拼接赋值给chargeForm的经纬度
  chargeForm.coordinate = `${e.point.lng},${e.point.lat}`
  //隐藏默认圆形覆盖物
  showDefaultCircle.value = false
}
//校验规则
let chargeRules = {
  name: [{ required: true, message: '请输入站名', trigger: 'blur' }],
  address: [{ required: true, message: '请输入地址', trigger: 'blur' }],
  coordinate: [{ required: true, message: '请确认位置', trigger: 'blur' }],
  imageUrl: [{ required: true, message: '请上传图片', trigger: 'blur' }]
}
//点击查看地图按钮，打开地图
let showMap = (row: any) => {
  chargeForm.coordinate = row.coordinate
  // 弹出地图对话框
  dialogVisible.value = true
  // 更新selectedPoint的值
  selectedPoint.value.lng = row.coordinate.split(',')[0]
  selectedPoint.value.lat = row.coordinate.split(',')[1]
  //更新经纬度后定位在弹窗中心
  dialogCenter.lng = row.coordinate.split(',')[0]
  dialogCenter.lat = row.coordinate.split(',')[1]
  // 调用ready方法重新渲染地图
  // ready({ BMap, map: null }) // 传入null代表不对map进行操作，只是触发更新
}
//展示地图弹窗关闭时，刷新页面
let close = () => {
  settingStore.refsh = !settingStore.refsh
}

//图片上传成功的钩子
const handleAvatarSuccess: UploadProps['onSuccess'] = () => {
  chargeFormRef.value.clearValidate('imageUrl')
}
//上传图片组件->上传图片之前触发的钩子函数
const beforeAvatarUpload: UploadProps['beforeUpload'] = async (rawFile: any) => {
  //请求上传文件的接口
  let res = await reqUpload(rawFile)
  //将接口的地址赋值给表单并呈现
  chargeForm.imageUrl = res.data.url
  //上传图片格式和大小要求  png|jpg  4M
  if (rawFile.type !== 'image/png' || rawFile.type == 'image/jpg') {
    ElMessage.error('上传文件格式务必PNG|JPG')
    return false
  } else if (rawFile.size / 1024 / 1024 > 4) {
    ElMessage.error('上传文件大小小于4M')
    return false
  }
  // 取消默认的上传请求
  return false
}
</script>
<script lang="ts">
export default {
  name: '充电站'
}
</script>
<style scoped>
.form {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
/* 地图容器必须设置宽和高属性 */
.map {
  width: 600px;
  height: 400px;
}
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>
