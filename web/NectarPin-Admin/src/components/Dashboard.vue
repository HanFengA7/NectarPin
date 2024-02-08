<script lang="ts" setup>
import {computed, ref, reactive} from 'vue'
import axios from "axios";
import eventBus from "@/plugin/event-bus/event-bus.js";


const data = [
  ['2024-1-8', 1],
  ['2024-1-9', 5],
  ['2024-1-10', 1],
  ['2024-1-11', 9],
  ['2024-1-16', 3],
  ['2024-1-24', 2],
];
const option = computed(() => {
  return {
    color: {
      0: "",
      1: "",
      2: "",
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
      target: {
        inRange: {
          color: {
            1: "#9be9a8",
            2: "#40c463",
            3: "#30a14e",
            4: "#216e39",
          }
        }
      },
      controller: {
        inRange: {
          color: {
            1: "#9be9a8",
            2: "#40c463",
            3: "#30a14e",
            4: "#216e39",
          }
        }
      },
    },
    calendar: {
      top: 75,
      left: 70,
      right: 10,
      cellSize: ["auto", 14],
      range: "2024",
      itemStyle: {
        borderWidth: 0.5
      },
      yearLabel: {
        show: true
      },
      dayLabel: {
        show: false,
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




//头部卡片信息
const heardCardInfo = reactive({
  hitokoto:"",
  date:"",
  solarDay:"",
  sign:"",
  lastIpProvince:"",
})
axios.get('https://international.v1.hitokoto.cn/?c=d')
    .then(res =>{
      heardCardInfo.hitokoto = res.data["hitokoto"] + "--《" +res.data["from"]+"》"
    })

//接收父组件的UserInfoData数据
const props = defineProps(['userInfo']);

//传数据给父组件
//eventBus.emit('child-data-selectedKeys', ["Dashboard"]);
/*设置侧边栏选择选项*/
let SelectedKeys: any = ref(["Dashboard"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);
</script>

<template>
  <div class="dashboard-heard-PC">
    <h2>🤗&nbsp;欢迎回来&nbsp;,&nbsp;{{ props.userInfo["nickname"] }}&nbsp;,&nbsp;快来写一篇吧!</h2>

    <h5>
      {{ heardCardInfo.hitokoto }}
    </h5>

  </div>
  <div class="dashboard-body-PC">
    <a-row :gutter="24">
      <a-col :span="24">
        <div class="dashboard-card2-PC">
          <a-statistic title="文章总数" :value="15" :precision="0" :value-from="0" :start="true" animation>
            <template #prefix>
              <icon-file style="padding-right: 5px"/>
            </template>
            <template #suffix>篇</template>
          </a-statistic>
          <a-statistic title="分类总数" :value="5" :precision="0" :value-from="0" :start="true" animation
                       style="padding-left: 150px">
            <template #prefix>
              <icon-folder style="padding-right: 5px"/>
            </template>
            <template #suffix>个</template>
          </a-statistic>
          <a-statistic title="评论总数" :value="66" :precision="0" :value-from="0" :start="true" animation
                       style="padding-left: 150px">
            <template #prefix>
              <icon-message style="padding-right: 5px"/>
            </template>
            <template #suffix>条</template>
          </a-statistic>
          <a-statistic title="获赞总数" :value="999" :precision="0" :value-from="0" :start="true" animation
                       style="padding-left: 150px">
            <template #prefix>
              <icon-thumb-up style="padding-right: 5px"/>
            </template>
            <template #suffix>个</template>
          </a-statistic>
          <a-statistic title="今日访客量" :value="2" :precision="0" :value-from="0" :start="true" animation
                       style="padding-left: 150px">
            <template #prefix>
              <icon-user style="padding-right: 5px"/>
            </template>
            <template #suffix>位</template>
          </a-statistic>

        </div>
      </a-col>
    </a-row>
    <br/>
    <a-row :gutter="24">
      <a-col :span="13">
        <div class="dashboard-chart1-PC">
          <e-charts class="chart" :option="option" style="width: 700px;height: 100%"/>
        </div>
      </a-col>
      <a-col :span="11">
        <div class="dashboard-card1-PC">
          <a-page-header
              :style="{ background: 'var(--color-bg-2)' }"
              title="快捷入口"
              subtitle="便捷导航，一键访问"
              :show-back="false"
          >
            <br/>
            <a-row class="dashboard-card1-box-PC">
              <a-col :lg="{span: 2}">
                <a-tooltip content="写作">
                  <div class="dashboard-card1-box-btn-bg-PC">
                    <div class="dashboard-card1-box-btn-PC">
                      <icon-edit/>
                    </div>
                  </div>
                </a-tooltip>
              </a-col>

              <a-col :lg="{span: 2, offset: 2}">
                <a-tooltip content="文章管理">
                  <div class="dashboard-card1-box-btn-bg-PC">
                    <div class="dashboard-card1-box-btn-PC">
                      <icon-sort/>
                    </div>
                  </div>
                </a-tooltip>
              </a-col>

              <a-col :lg="{span: 2, offset: 2}">
                <a-tooltip content="博客首页">
                  <div class="dashboard-card1-box-btn-bg-PC">
                    <div class="dashboard-card1-box-btn-PC">
                      <icon-home/>
                    </div>
                  </div>
                </a-tooltip>
              </a-col>

              <a-col :lg="{span: 2, offset: 2}">
                <a-tooltip content="Todo">
                  <div class="dashboard-card1-box-btn-bg-PC">
                    <div class="dashboard-card1-box-btn-PC">
                      <icon-mind-mapping/>
                    </div>
                  </div>
                </a-tooltip>
              </a-col>

              <a-col :lg="{span: 2, offset: 2}">
                <a-tooltip content="基本设置">
                  <div class="dashboard-card1-box-btn-bg-PC">
                    <div class="dashboard-card1-box-btn-PC">
                      <icon-settings/>
                    </div>
                  </div>
                </a-tooltip>
              </a-col>

              <a-col :lg="{span: 2, offset: 2}">
                <a-tooltip content="NectarPin Github">
                  <div class="dashboard-card1-box-btn-bg-PC">
                    <div class="dashboard-card1-box-btn-PC">
                      <icon-github/>
                    </div>
                  </div>
                </a-tooltip>
              </a-col>
            </a-row>
          </a-page-header>
        </div>
      </a-col>
    </a-row>
    <br/>
    <!--第三层-->
    <a-row :gutter="24">
      <a-col :xs="8">
        <div class="dashboard-card3-PC" style="margin-left: 0">
          <a-page-header
              :style="{ background: 'var(--color-bg-2)' }"
              title="热门文章"
              subtitle="热文尽览，洞察世界。"
              :show-back="false"
          >
          </a-page-header>
        </div>
      </a-col>
      <a-col :xs="8">
        <div class="dashboard-card3-PC">
          <a-page-header
              :style="{ background: 'var(--color-bg-2)' }"
              title="最新评论"
              subtitle="见解留声，新评即现。"
              :show-back="false"
          >
          </a-page-header>
        </div>
      </a-col>
      <a-col :xs="8">
        <div class="dashboard-card3-PC" style="margin-right: 0">
          <a-page-header
              :style="{ background: 'var(--color-bg-2)' }"
              title="Todo看板"
              subtitle="任务有序，生活更简单。"
              :show-back="false"
          >
          </a-page-header>
        </div>
      </a-col>

    </a-row>

  </div>
</template>

<style scoped>
.dashboard-heard-PC {
  background: #ffffff;
  padding: 15px 50px 15px 60px;
  height: 100px;
}

.dashboard-heard-PC h2 {
  font-weight: 300;
}

.dashboard-heard-PC h5 {
  font-weight: 200;
}

.dashboard-body-PC {
  margin: 30px 45px;
}

.dashboard-chart1-PC {
  padding: 5px;
  background: #ffffff;
  border-radius: 15px;
  height: 200px
}

.dashboard-card1-PC {
  padding: 5px;
  background: #ffffff;
  border-radius: 15px;
  height: 200px
}

.dashboard-card2-PC {
  padding: 20px;
  background: #ffffff;
  border-radius: 15px;
  height: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.dashboard-card3-PC {
  margin-left: 20px;
  margin-right: 20px;
  padding: 10px;
  background: #ffffff;
  border-radius: 15px;
  height: 300px;
}

.dashboard-card1-box-btn-bg-PC {
  width: 40px;
  height: 40px;
  padding: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f6f6f7;
  border-radius: 5px;
  transition: 3s ease-in-out;
}

.dashboard-card1-box-btn-bg-PC:hover {
  border-radius: 20px;
  box-shadow: 0 0 20px 13px rgb(246 246 247 / 25%);
}

.dashboard-card1-box-btn-bg-PC:active {
  border-radius: 20px;
  box-shadow: 0 0 20px 13px rgb(246 246 247 / 25%);
  transition: 1s ease-in-out;
}
</style>