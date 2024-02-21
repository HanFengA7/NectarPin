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
let SelectedKeys = ref(["Category"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);


/*
[HeardCardOnBack] 返回函数
*/
const HeardCardOnBack = () => {
  router.push({name: 'Category'})
}
</script>

<template>
分类
</template>

<style scoped>

</style>