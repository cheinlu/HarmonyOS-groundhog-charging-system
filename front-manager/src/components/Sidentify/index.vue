<template>
  <div class="s-canvas">
    <canvas id="s-canvas" :width="props.contentWidth" :height="props.contentHeight"></canvas>
  </div>
</template>

<script setup lang="ts">
import { onMounted, watch } from 'vue'

const props = defineProps({
  identifyCode: {
    type: String,
    default: '1234'
  },
  fontSizeMin: {
    type: Number,
    default: 25
  },
  fontSizeMax: {
    type: Number,
    default: 35
  },
  backgroundColorMin: {
    type: Number,
    default: 255
  },
  backgroundColorMax: {
    type: Number,
    default: 255
  },
  colorMin: {
    type: Number,
    default: 0
  },
  colorMax: {
    type: Number,
    default: 160
  },
  lineColorMin: {
    type: Number,
    default: 40
  },
  lineColorMax: {
    type: Number,
    default: 180
  },
  dotColorMin: {
    type: Number,
    default: 0
  },
  dotColorMax: {
    type: Number,
    default: 255
  },
  contentWidth: {
    type: Number,
    default: 112
  },
  contentHeight: {
    type: Number,
    default: 40
  }
})
// 生成一个随机数
const randomNum = (min: number, max: number) => {
  return Math.floor(Math.random() * (max - min) + min)
}
// 生成一个随机的颜色
const randomColor = (min: number, max: number) => {
  let r = randomNum(min, max)
  let g = randomNum(min, max)
  let b = randomNum(min, max)
  return 'rgb(' + r + ',' + g + ',' + b + ')'
}
//绘制干扰线
const drawLine = (ctx: { strokeStyle: string; beginPath: () => void; moveTo: (arg0: number, arg1: number) => void; lineTo: (arg0: number, arg1: number) => void; stroke: () => void }) => {
  // 绘制干扰线
  for (let i = 0; i < 5; i++) {
    ctx.strokeStyle = randomColor(props.lineColorMin, props.lineColorMax)
    ctx.beginPath()
    ctx.moveTo(randomNum(0, props.contentWidth), randomNum(0, props.contentHeight))
    ctx.lineTo(randomNum(0, props.contentWidth), randomNum(0, props.contentHeight))
    ctx.stroke()
  }
}
const drawText = (ctx: { fillStyle: string; font: string; translate: (arg0: number, arg1: number) => void; rotate: (arg0: number) => void; fillText: (arg0: any, arg1: number, arg2: number) => void }, txt: string, i: number) => {
  ctx.fillStyle = randomColor(props.colorMin, props.colorMax)
  ctx.font = randomNum(props.fontSizeMin, props.fontSizeMax) + 'px SimHei'
  let x = (i + 1) * (props.contentWidth / (props.identifyCode.length + 1))
  let y = randomNum(props.fontSizeMax, props.contentHeight - 5)
  let deg = randomNum(-45, 45)
  // 修改坐标原点和旋转角度
  ctx.translate(x, y)
  ctx.rotate((deg * Math.PI) / 180)
  ctx.fillText(txt, 0, 0)
  // 恢复坐标原点和旋转角度
  ctx.rotate((-deg * Math.PI) / 180)
  ctx.translate(-x, -y)
}
//绘制干扰点
const drawDot = (ctx: { fillStyle: string; beginPath: () => void; arc: (arg0: number, arg1: number, arg2: number, arg3: number, arg4: number) => void; fill: () => void }) => {
  // 绘制干扰点
  for (let i = 0; i < 80; i++) {
    ctx.fillStyle = randomColor(0, 255)
    ctx.beginPath()
    ctx.arc(randomNum(0, props.contentWidth), randomNum(0, props.contentHeight), 1, 0, 2 * Math.PI)
    ctx.fill()
  }
}
//绘制背景
const drawPic = () => {
  let canvas = document.getElementById('s-canvas') as HTMLCanvasElement;
  let ctx:any = canvas.getContext('2d')
  ctx.textBaseline = 'bottom'
  // 绘制背景
  ctx.fillStyle = randomColor(props.backgroundColorMin, props.backgroundColorMax)
  ctx.fillRect(0, 0, props.contentWidth, props.contentHeight)
  // 绘制文字
  for (let i = 0; i < props.identifyCode.length; i++) {
    drawText(ctx, props.identifyCode[i], i)
  }
  drawLine(ctx)
  drawDot(ctx)
  
}
//监听
watch(
  () => props.identifyCode,
  () => {
    drawPic()
  }
)
//组件挂载
onMounted(() => {
  drawPic()
})
</script>

<style scoped>
.s-canvas {
  cursor: pointer;
}
</style>
