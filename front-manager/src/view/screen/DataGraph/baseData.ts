import * as echarts from "echarts";
export const option1 = {
  tooltip: {
    trigger: "item",
    axisPointer: {
      // 坐标轴指示器，坐标轴触发有效
      type: "shadow", // 默认为直线，可选为：'line' | 'shadow'
    },
  },
  grid: {
    left: "1%",
    right: "18%",
    bottom: "4%",
    top: "20%",
    containLabel: true,
  },
  legend: {
    data: ["销售成功", "订单失败", "交易终止", "交易成功率"],
    left: "0%",
    top: "2%",
    textStyle: {
      color: "#666666",
    },
    itemWidth: 15,
    itemHeight: 10,
    itemGap: 25,
  },
  xAxis: {
    type: "category",
    data: [ "2016", "2017", "2018", "2019", "2020", "2021", "2022","2023"],
    axisLine: {
      lineStyle: {
        color: "#cdd5e2",
      },
    },
    axisLabel: {
      textStyle: {
        color: "#666666",
      },
    },
  },

  yAxis: [
    {
      type: "value",
      axisLine: {
        show: false,
        lineStyle: {
          color: "#cdd5e2",
        },
      },
      splitLine: {
        show: false,
      },
      axisLabel: {
        textStyle: {
          color: "#666666",
        },
      },
    },
    {
      type: "value",
      name: "百分比",
      nameTextStyle: {
        color: "#666666",
      },
      position: "right",
      axisLine: {
        lineStyle: {
          color: "#cdd5e2",
        },
      },
      splitLine: {
        show: false,
      },
      axisLabel: {
        show: true,
        formatter: "{value} %", //右侧Y轴文字显示
        textStyle: {
          color: "#666666",
        },
      },
    },
  ],
  series: [
    {
      name: "销售成功",
      type: "bar",
      barWidth: "12px",
      itemStyle: {
        normal: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: "#29acff",
            },
            {
              offset: 1,
              color: "#4bdfff",
            },
          ]),
          barBorderRadius: 6,
        },
      },
      data: [400, 400, 300, 300, 300, 400, 400, 400, 300],
    },
    {
      name: "订单失败",
      type: "bar",
      barWidth: "12px",
      itemStyle: {
        normal: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: "#3d93f2",
            },
            {
              offset: 1,
              color: "#5dc1fd",
            },
          ]),
          barBorderRadius: 6,
        },
      },
      data: [400, 500, 500, 500, 500, 400, 400, 500, 500],
    },
    {
      name: "交易终止",
      type: "bar",
      barWidth: "12px",
      itemStyle: {
        normal: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: "#01c871",
            },
            {
              offset: 1,
              color: "#55f49c",
            },
          ]),
          barBorderRadius: 6,
        },
      },
      data: [400, 600, 700, 700, 1000, 400, 400, 600, 700],
    },
    {
      name: "交易成功率",
      type: "line",
      yAxisIndex: 1, //使用的 y 轴的 index，在单个图表实例中存在多个 y轴的时候有用
      smooth: false, //平滑曲线显示

      symbol: "circle", //标记的图形为实心圆
      symbolSize: 8, //标记的大小
      itemStyle: {
        normal: {
          color: "#ffa43a",
          borderColor: "rgba(255, 234, 0, 0.5)", //圆点透明 边框
          borderWidth: 5,
        },
      },
      lineStyle: {
        color: "#ffa43a",
      },

      data: [3, 2.5, 3.8, 3.5, 4.9, 2.8, 5, 4],
    },
  ],
};

export const option2 = {
  tooltip: {
    trigger: "axis",
  },
  grid: {
    left: "2%",
    right: "15%",
    top: "10%",
    bottom: "10%", //也可设置left和right设置距离来控制图表的大小
  },
  xAxis: [
    {
      type: "category",
      data: [
        "00:00",
        "00:02",
        "00:04",
        "00:06",
        "00:08",
        "00:10",
        "00:12",
        "00:14",
        "00:16",
        "00:18",
        "00:20",
        "00:22",
        "00:24",
      ],
      axisLine: {
        lineStyle: {
          color: "#ddd",
        },
      },
      axisTick: {
        show: false,
      },
      axisLabel: {
        interval: 0,
        textStyle: {
          color: "#333",
        },
        margin: 15,
      },
    },
  ],
  yAxis: [
    {
      type: "value",
      axisTick: {
        show: false,
      },
      axisLine: {
        show: false,
      },
      axisLabel: {
        textStyle: {
          color: "#333",
        },
      },
      splitLine: {
        lineStyle: {
          color: "#dadada",
        },
      },
    },
  ],
  series: [
    {
      name: "访问量统计",
      type: "line",
      data: [13, 10, 8, 12, 15, 14, 10, 13, 10, 13, 12, 15, 17],
      symbolSize: 6,
      symbol: "circle",
      lineStyle: {
        color: "#3CC69B",
      },
      itemStyle: {
        normal: {
          color: "#90e4ca",
          borderColor: "#3CC69B",
        },
      },
      areaStyle: {
        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          {
            offset: 0,
            color: "#49d2a7e3",
          },
          {
            offset: 0.6,
            color: "#9be8d0ed",
          },
          {
            offset: 1,
            color: "#ddece8ed",
          },
        ]),
      },
    },
  ],
};
