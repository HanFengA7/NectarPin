<script setup lang="ts">
import eventBus from "@/plugin/event-bus/event-bus";
import {reactive, ref} from "vue";
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
分类数据
 */
const CategoryForm = reactive([{
  id: 1,
  name: '默认分类',
  short_name: 'mysql',
  desc: 'mysql',
  parent_id: 0,
  depth: 1,
},{
  id: 2,
  name: '默认分类1',
  short_name: 'mysql',
  desc: 'mysql',
  parent_id: 0,
  depth: 1,
},{
  id: 3,
  name: '默认分类2',
  short_name: 'mysql',
  desc: 'mysql',
  parent_id: 0,
  depth: 1,
}]);


/*
添加分类
 */
//[添加分类][表单数据]
const addCategoryFormRef = ref()
const addCategoryForm = reactive({
  name: '',
  short_name: '',
  desc: '',
  parent_id: 0,
  depth: 1,
});
//[添加分类][模态框显示]
const addCategoryVisible = ref(false);
//[添加分类][点击添加分类事件]
const addCategoryHandleClick = () => {
  addCategoryVisible.value = true;
};
//[添加分类][添加事件]
const addCategoryHandleBeforeOk = (done:any) => {
  console.log(addCategoryForm)
  window.setTimeout(() => {
    done()
    // prevent close
    // done(false)
  }, 3000)
};
//[添加分类][取消事件]
const addCategoryHandleCancel = () => {
  addCategoryVisible.value = false;
  addCategoryFormRef.value['resetFields'](['name'])
}


/*
[HeardCardOnBack] 返回函数
*/
const HeardCardOnBack = () => {
  router.push({name: 'Dashboard'})
}
</script>

<template>
  <a-row>

    <!--电脑端 [Start]-->
    <a-col :xs="{span: 0}" :sm="{span: 0}" :md="{span: 0}" :lg="{span: 24}" :xl="{span: 24}" :xxl="{span: 24}">

      <!--[0] 模态框 [Start]-->

      <!--[0-1] 添加分类-->
      <a-modal
          v-model:visible="addCategoryVisible"
          title="添加分类"
          @cancel="addCategoryHandleCancel"
          @before-ok="addCategoryHandleBeforeOk">
        <a-form ref="addCategoryFormRef" :model="addCategoryForm">
          <a-form-item field="name" label="分类名称">
            <a-input v-model="addCategoryForm.name" />
          </a-form-item>
          <a-form-item field="name" label="分类缩略名">
            <a-input v-model="addCategoryForm.short_name" />
          </a-form-item>
          <a-form-item field="parent_id" label="父级分类">
            <a-select v-model="addCategoryForm.parent_id">
              <a-option :value="0">不选择</a-option>
              <a-option v-for="item of CategoryForm" :value="item.id" :label="item.name" />
            </a-select>
          </a-form-item>
          <a-form-item field="depth" label="分类层级">
            <a-select v-model="addCategoryForm.depth">
              <a-option :value="1">一级分类</a-option>
              <a-option :value="2">二级分类</a-option>
              <a-option :value="3">三级分类</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="desc" label="分类描述">
            <a-textarea v-model="addCategoryForm.desc" :max-length="30" allow-clear show-word-limit auto-size/>
          </a-form-item>
        </a-form>
      </a-modal>

      <!--[0] 模态框 [End]-->

      <!--[1] 页头卡片 [Start]-->
      <div class="Category-heard-PC">
        <a-page-header
            title="分类管理"
            subtitle="CategoryManagement"
            :show-back="true"
            @back=HeardCardOnBack
        >
          <template #breadcrumb>
            <a-breadcrumb>
              <a-breadcrumb-item>后台</a-breadcrumb-item>
              <a-breadcrumb-item>文章管理</a-breadcrumb-item>
            </a-breadcrumb>
          </template>

          <template #extra>
            <a-space>
              <a-button @click="addCategoryHandleClick">添加分类</a-button>
              <a-button @click="">批量删除</a-button>
            </a-space>
          </template>
        </a-page-header>
      </div>
      <!--[1] 页头卡片 [End]-->

      <!--[2] 文章列表 [Start]-->
      <div class="Category-Table-PC">

      </div>
      <!--[2] 文章列表 [End]-->
    </a-col>
    <!--电脑端 [End]-->

    <!--手机端 [Start]-->
    <a-col :xs="{span: 24}" :sm="{span: 24}" :md="{span: 24}" :lg="{span: 0}" :xl="{span: 0}" :xxl="{span: 0}">
      <!--[1] 页头卡片 [Start]-->
      <div class="Article-heard-PE">
        <a-page-header
            title="文章管理"
            subtitle="ArticleManagement"
            :show-back="true"
            @back=HeardCardOnBack
        >
          <template #breadcrumb>
            <a-breadcrumb>
              <a-breadcrumb-item>后台</a-breadcrumb-item>
              <a-breadcrumb-item>文章管理</a-breadcrumb-item>
            </a-breadcrumb>
          </template>
        </a-page-header>
      </div>
      <!--[1] 页头卡片 [End]-->

      <!--[2] 文章列表 [Start]-->
      <div class="Article-Table-PC">

      </div>
      <!--[2] 文章列表 [End]-->
    </a-col>
    <!--手机端 [End]-->

  </a-row>
</template>

<style scoped>
/*
[1] 页头卡片 CSS
*/
.Category-heard-PC {
  background: #ffffff;
  padding: 0 50px 5px 25px;
  height: 100px;
}

.Category-heard-PE {
  background: #ffffff;
  padding: 0 50px 5px 15px;
  height: 100px;
}

/*
[2] 文章表格 CSS
*/
.Category-Table-PC {
  margin: 30px 35px;
  padding: 10px;
  background: #ffffff;
  border-radius: 15px;
}
</style>