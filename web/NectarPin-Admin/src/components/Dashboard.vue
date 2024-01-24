<script setup>
import { ref,computed } from 'vue'



const data = [
  ['2024-1-8',1],
  ['2024-1-9',5],
  ['2024-1-10',1],
  ['2024-1-11',9],
  ['2024-1-16',3],
  ['2024-1-24',2],
];
const option = computed(()=>{
  return {
    color:{
      0:"",
      1:"",
      2:"",
    },
    tooltip: {},
    visualMap: {
      min: 0,
      max: 10,
      type: "piecewise",
      orient: "horizontal",
      left: "center",
      top: 20,
      textStyle: {
        color: "#000"
      },
      target:{
        inRange:{
          color:{
            1:"#9be9a8",
            2:"#40c463",
            3:"#30a14e",
            4:"#216e39",
          }
        }
      },
      controller:{
        inRange:{
          color:{
            1:"#9be9a8",
            2:"#40c463",
            3:"#30a14e",
            4:"#216e39",
          }
        }
      },
    },
    calendar: {
      top: 80,
      left: 30,
      right: 30,
      cellSize: ["auto", 14],
      range: "2024",
      itemStyle: {
        borderWidth: 0.5
      },
      yearLabel: {
        show: true
      },
      dayLabel: {
        show:false,
        nameMap: 'ZH'
      },
      monthLabel: {
        nameMap: 'ZH'
      }
    },
    series: {
      type: "heatmap",
      coordinateSystem: "calendar",
      data: data
    }
  }
})
</script>

<template>
  <div class="dashboard-heard-PC">
    <h2>🤗&nbsp;欢迎回来&nbsp;,&nbsp;火山喵&nbsp;,&nbsp;你有2天没写博文了&nbsp;,&nbsp;快来写一篇吧!</h2>
    <h5>今天大部分地区多云。 最低气温6°，天气寒冷，哥哥注意添加衣物哟！</h5>
  </div>
  <div class="dashboard-body-PC">
    <div class="dashboard-chart1-PC">
    <e-charts class="chart" :option="option" style="width: 700px;height: 100%"/>
    </div>
  </div>
</template>

<style scoped>
.dashboard-heard-PC{
  background: #ffffff;
  padding: 15px 50px 15px 60px;
  height: 100px;
}
.dashboard-heard-PC h2{
  font-weight: 300;
}
.dashboard-heard-PC h5{
  font-weight: 200;
}
.dashboard-body-PC{
  margin: 30px 45px;
}
.dashboard-chart1-PC{
  padding: 5px;
  background: #ffffff;
  border-radius: 15px;
  width: 700px;
  height: 200px
}
</style>