<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="4">
        <el-card>
          <Detail title="注册用户" count="1331人" :icon="headSvg"> </Detail>
        </el-card>
      </el-col>
      <el-col :span="4">
        <el-card>
          <Detail title="总订单" count="2314元" :icon="moneySvg"></Detail>
        </el-card>
      </el-col>
      <el-col :span="4">
        <el-card>
          <Detail title="总销售额" count="2314元" :icon="moneySvg"></Detail>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <div class="flex-card">
            <Detail1 title="今日订单" count="234元"></Detail1>
            <Detail1 title="今日销售额" count="234元"></Detail1>
            <Detail1 title="今日访问量" count="234元"></Detail1>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Detail from './detail'
import Detail1 from './detail1'
export default {
  name: 'card',
  data() {
    return {
      headSvg:
        ' <svg t="1686227950225" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2006" width="20" height="20"><path d="M631.119772 619.422812l0.344028-16.685369c79.298505-55.388544 117.829666-153.092558 117.829666-297.756425 0-139.503444-87.211154-222.758273-233.423148-222.758273l-7.912649 0c-146.03998 0-233.423148 83.254829-233.423148 222.758273 0 143.975811 38.875189 241.679825 118.86175 297.412397l0 16.857383c-133.138922 17.889467-266.621871 70.869814-266.621871 151.88846 0 113.357299 129.526625 170.810012 385.139593 170.810012s385.139593-57.452713 385.139593-170.810012C897.053586 690.292626 763.914665 637.31228 631.119772 619.422812z" fill="#575B66" p-id="2007"/></svg>',
      moneySvg:
        '<svg t="1686301963042" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4634" width="20" height="20"><path d="M487.3 588h79.4c14.6 0 26.4-11.8 26.4-26.4s-11.8-26.4-26.4-26.4h-68l66.3-66.3c10.3-10.3 10.3-27.1 0-37.4-10.3-10.3-27.1-10.3-37.4 0l-62 61.9-62-61.9c-10.3-10.3-27.1-10.3-37.4 0-10.3 10.3-10.3 27.1 0 37.4l66.3 66.3H355c-14.6 0-26.4 11.8-26.4 26.4S340.4 588 355 588h79.4v52.9H355c-9.4 0-18.2 5-22.9 13.2-4.7 8.2-4.7 18.2 0 26.4s13.4 13.2 22.9 13.2h79.4v83.1c0 7 2.8 13.7 7.7 18.7 5 5 11.7 7.7 18.7 7.7h0.1c7 0 13.7-2.8 18.7-7.7 5-5 7.7-11.7 7.7-18.7v-83.1h79.4c9.4 0 18.2-5 22.9-13.2 4.7-8.2 4.7-18.2 0-26.4s-13.4-13.2-22.9-13.2h-79.4V588zM162.8 387.1L427 166.9c19.6-16.3 48.1-16.3 67.7 0L759 387.1c12.1 10 19 24.9 19 40.6v450.9c0 29.2-23.7 52.9-52.9 52.9H196.6c-29.2 0-52.9-23.7-52.9-52.9V427.7c0-15.7 7-30.5 19.1-40.6z m352.9-258.9l17.1-14.2c19.6-16.3 48.1-16.3 67.6 0l264.3 220.2c12.1 10 19 24.9 19 40.6v450.9c0 29.2-23.7 52.9-52.9 52.9V401.3c0-15.7-7-30.6-19-40.6L547.5 140.5c-8.9-7.5-20.1-11.8-31.8-12.3z m0 0" p-id="4635"></path></svg>'
    }
  },
  components: {
    Detail,
    Detail1
  }
}
</script>

<style lang="less" scoped>
.wel {
  font-size: 30px;
}
.flex-card {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}
</style>
