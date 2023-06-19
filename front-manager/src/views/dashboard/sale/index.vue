<template>
  <el-card class="box-card" style="margin: 10px 0px">
    <div slot="header" class="clearfix">
      <!--  @tab-click="handleClick" -->
      <!-- 头部左侧内容 -->
      <el-tabs class="tab" v-model="name">
        <el-tab-pane label="设备统计" name="device"></el-tab-pane>
        <el-tab-pane label="用户统计" name="user"></el-tab-pane>
        <el-tab-pane label="充值统计" name="recharge"></el-tab-pane>
      </el-tabs>
    </div>
    <div>
      <!-- 搜索区域 -->
      <el-row :gutter="20">
        <el-col :span="3">
          <el-select placeholder="请选择" value="">
            <el-option label="北京" value="bj">选项1</el-option>
            <el-option label="北京" value="bj">选项2</el-option>
          </el-select>
        </el-col>
        <el-col :span="6">
          <el-date-picker value="" class="date" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" value-format="yyyy-MM-dd" />
        </el-col>
        <el-col :span="4">
          <el-button type="primary" icon="el-icon-search">搜索</el-button>
        </el-col>
      </el-row>
      <!-- 图表区域 -->
      <el-row :gutter="10">
        <el-col :span="24">
          <!-- 容器 -->
          <div class="charts" ref="charts"></div>
        </el-col>
      </el-row>
    </div>
  </el-card>
</template>

<script>
import echarts from 'echarts'
export default {
  name: '',
  data() {
    return {
      name: 'device',
      mycharts: null
    }
  },
  mounted() {
    this.mycharts = echarts.init(this.$refs.charts)
    this.mycharts.setOption({
      title: {
        text: this.title
      },
      tooltip: {
        trigger: 'axis'
      },
      legend: {},
      toolbox: {
        show: true,
        feature: {
          dataZoom: {
            yAxisIndex: 'none'
          },
          dataView: { readOnly: false },
          magicType: { type: ['line', 'bar'] },
          restore: {},
          saveAsImage: {}
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
      },
      yAxis: {
        type: 'value',
        axisLabel: {
          formatter: '{value} °C'
        }
      },
      series: [
        {
          name: 'Highest',
          type: 'line',
          data: [10, 11, 13, 11, 12, 12, 9],
          markPoint: {
            data: [
              { type: 'max', name: 'Max' },
              { type: 'min', name: 'Min' }
            ]
          },
          markLine: {
            data: [{ type: 'average', name: 'Avg' }]
          }
        },
        {
          name: 'Lowest',
          type: 'line',
          data: [1, -2, 2, 5, 3, 2, 0],
          markPoint: {
            data: [{ name: '周最低', value: -2, xAxis: 1, yAxis: -1.5 }]
          },
          markLine: {
            data: [
              { type: 'average', name: 'Avg' },
              [
                {
                  symbol: 'none',
                  x: '90%',
                  yAxis: 'max'
                },
                {
                  symbol: 'circle',
                  label: {
                    position: 'start',
                    formatter: 'Max'
                  },
                  type: 'max',
                  name: '最高点'
                }
              ]
            ]
          }
        }
      ]
    })
  },
  computed: {
    // 计算属性-标题
    title() {
      // this.activeName == 'sale' ? '销售额' : '访问量'
      // return this.name == 'device'? '设备统计' : '用户统计'
      return this.name == 'device' ? '设备统计' : this.name == 'user' ? '用户统计' : '充值统计'
    }
  },
  watch: {
    title() {
      this.mycharts.setOption({
        title: {
          text: this.title
        }
      })
    }
  }
}
</script>

<style scoped>
.el-card__body {
  border-bottom: none;
}
.clearfix {
  position: relative;
  display: flex;
  justify-content: space-between;
}
.tab {
  width: 100%;
}
.right {
  position: absolute;
  right: 0px;
}
.date {
  /* width: 250px;
  margin: 0px 20px; */
}
.right span {
  margin: 0px 10px;
}
.charts {
  width: 100%;
  height: 300px;
}
ul {
  list-style: none;
  width: 100%;
  height: 300px;
  padding: 0px;
}
ul li {
  height: 8%;
  margin: 10px 0px;
}
.rindex {
  float: left;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: black;
  color: white;
  text-align: center;
}
.rvalue {
  float: right;
}
</style>
