<script setup lang="ts">
import eventBus from "@/plugin/event-bus/event-bus";
import {ref} from "vue";
import router from "@/router";

/*
接收父组件数据
*/
//[1]UserInfoData数据
const props = defineProps(['userInfo']);

/*
传数据给父组件
*/
//[1]设置侧边栏选择选项
let SelectedKeys = ref(["Article"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);

/*
[HeardCardOnBack] 返回函数
*/
const HeardCardOnBack = () => {
  router.push({name: 'Dashboard'})
}
</script>

<template>
  <div>
    <a-row>
      <a-col :xs="{span: 5, offset: 1}" :lg="{span: 6, offset: 2}">
        Col
      </a-col>
      <a-col :xs="{span: 11, offset: 1}" :lg="{span: 6, offset: 2}">
        Col
      </a-col>
      <a-col :xs="{span: 5, offset: 1}" :lg="{span: 6, offset: 2}">
        Col
      </a-col>
    </a-row>
  </div>
</template>

<style scoped>

</style>