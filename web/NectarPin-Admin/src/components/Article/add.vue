<script setup lang="ts">
import eventBus from "@/plugin/event-bus/event-bus";
import {ref} from "vue";

/*
传数据给父组件
[1]设置侧边栏选择选项
*/
let SelectedKeys: any = ref(["Article/add"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);
</script>

<template>
<h1>撰写文章</h1>
</template>

<style scoped>

</style>