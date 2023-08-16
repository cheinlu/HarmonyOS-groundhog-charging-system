<template>
	<!--查看收费标准的弹窗-->
	<view class="pop-dlg-rect">
		<view class="dlg-rect">
			<view class="dlg-text">收费标准</view>
      <scroll-view scroll-y="true">
        <view class="row">
          <view class="cell header">收费时段</view>
          <view class="cell header">收费(元/度)</view>
        </view>
        <view class="row" v-for="item in list" :key="item.id">
          <view class="cell">{{item.startHour}}-{{item.endHour}}</view>
          <view class="cell">{{item.price}}</view>
        </view>
      </scroll-view>
			<view class="dlg-btns-rect">
				<view class="btn-rect" @click="confim">
					<view class="btn-text">我知道了</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
let $emit = defineEmits(['confirm'])
let props = defineProps(['list'])

let confim = ()=>{
  $emit('confirm')
}
</script>

<style lang="scss" scoped>
.pop-dlg-rect {
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background: rgba(0, 0, 0, 0.4);
	display: flex;
	align-items: center;
	justify-content: center;
	z-index: 1000;
	.dlg-rect {
		width: 500rpx;
		background: #FBFBFB;
		border-radius: 17rpx;
		padding: 25rpx 56rpx 52rpx;
		.dlg-text {
			font-size: 38rpx;
			font-family: Source Han Sans CN-Regular, Source Han Sans CN;
			font-weight: 400;
			color: #3D3D3D;
			line-height: 44px;
      text-align: center;
		}
		.dlg-btns-rect {
			margin-top: 21rpx;
			width: 100%;
			display: flex;
			flex-wrap: nowrap;
			justify-content: space-around;
			.btn-rect {
				width: 200rpx;
				height: 70rpx;
				background: #0aa671;
				border-radius: 10rpx;
				display: flex;
				align-items: center;
				justify-content: center;
				.btn-text {
					font-size: 25rpx;
					font-family: Source Han Sans CN-Regular, Source Han Sans CN;
					font-weight: 400;
					color: #FBFBFB;
				}
			}
		}
	}
}
scroll-view{
  height:300rpx;
  white-space: nowrap;
  // border-right: 2rpx solid #ccc;
}
.row {
  display: flex;
  align-items: center;
  width: 495rpx;
  height: 60rpx;
  border: 2rpx solid #ccc;
}

.cell {
  flex: 1;
  padding: 0 10rpx;
  text-align: center;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.header {
  font-weight: bold;
  height: 100%;
  line-height: 60rpx;
  background-color: #eee;
}

</style>
