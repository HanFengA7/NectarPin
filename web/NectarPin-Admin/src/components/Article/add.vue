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
let SelectedKeys: any = ref(["Article/add"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);

/*
[HeardCardOnBack] 返回函数
*/
const HeardCardOnBack = () => {
  router.push({name: 'Dashboard'})
}
</script>

<template>
  <div class="ArticleAdd-heard-PC">
    <a-page-header
        title="撰写文章"
        subtitle="WriteArticle"
        :show-back="true"
        @back=HeardCardOnBack
    >
      <template #breadcrumb>
        <a-breadcrumb>
          <a-breadcrumb-item>后台</a-breadcrumb-item>
          <a-breadcrumb-item>撰写文章</a-breadcrumb-item>
        </a-breadcrumb>
      </template>

      <template #extra>
        <a-space>
          <a-button @click="">重置文章</a-button>
        </a-space>
      </template>
    </a-page-header>
  </div>

  <br/>

</template>

<style scoped>
.ArticleAdd-heard-PC{
  background: #ffffff;
  padding: 0 50px 5px 25px;
  height: 100px;
}
</style>