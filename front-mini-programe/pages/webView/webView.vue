<template>
	<view>
		<web-view :src="url"></web-view>
	</view>
</template>
 
<script>
	export default {
		data() {
			return {
				url:''
			};
		},
		onLoad(e) {
            //接收传递的访问地址
			this.url = e.url
		}
	}
</script>