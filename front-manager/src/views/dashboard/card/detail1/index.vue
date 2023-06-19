<template>
  <div class="detail1-head">
    <span class="detail1-title">{{ title }}</span>
    <span>{{ count }}</span>
  </div>
</template>

<script>
export default {
  name: 'detail1',
  props: ['title', 'count']
}
</script>

<style lang="less" scoped>
.detail1-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 300px;
  margin-bottom: 10px;
  .detail1-title {
    // margin-right: 80px;
  }
}
</style>
