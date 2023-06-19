<template>
  <div>
    <div class="card-header">
      <span>{{ title }}</span>
    </div>
    <div class="card-content">
      <span>{{ count }}</span>
      <span v-html="icon"></span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'detail',
  props: ['title', 'count', 'icon']
}
</script>

<style lang="less" scoped>
.cart-header {
  font-size: 30px;
  padding: 10px 0;
}
.card-content {
  margin-top: 15px;
  display: flex;
  justify-content: space-between;
}
</style>
