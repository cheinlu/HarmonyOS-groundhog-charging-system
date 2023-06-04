<template>
   <view>
     <login v-if="!token"></login>
     <userInfo v-else></userInfo>
   </view>
</template>

<script>
  import {mapState} from 'vuex'
  export default {
    name:"my",
    data() {
      return {
        
      };
    },
    computed:{
      ...mapState('m_user',['token'])
    }
  }
</script>

<style lang="scss">
</style>