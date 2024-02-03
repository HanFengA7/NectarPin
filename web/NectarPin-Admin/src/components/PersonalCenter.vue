<script lang="ts" setup>
import eventBus from '@/plugin/event-bus/event-bus';
import {ref} from "vue";
//接收父组件的UserInfoData数据
const props = defineProps(['userInfo']);
//传数据给父组件
/*设置侧边栏选择选项*/
let SelectedKeys: any = ref(["PersonalCenter"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);
</script>

<template>
  <div class="personalCenter-heard-PC">

  </div>
  <div class="personalCenter-heardBox-PC">

  </div>
</template>

<style scoped>
.personalCenter-heard-PC {
  background: #ffffff;
  padding: 15px 50px 15px 60px;
  height: 100px;
}
.personalCenter-heardBox-PC{
  background: #ffffff;
  height: 500px;
}
</style>