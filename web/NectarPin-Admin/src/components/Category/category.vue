<script setup lang="ts">
import eventBus from "@/plugin/event-bus/event-bus";
import {reactive, ref} from "vue";
import router from "@/router";
import {GetCategoryList} from "@/api/Category/get";
import {Message, Notification} from "@arco-design/web-vue";
import {CreateCategory} from "@/api/Category/post";


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
//分类数据定义
const CategoryForm = reactive([{
  id: 1,
  name: '',
  short_name: '',
  desc: '',
  parent_id: 0,
  depth: 1,
}]);
//分类树数据定义[Max：三级]
const CategoryTreeData = ref([{
  title: '',
  key: '',
  children: [
    {
      title: '',
      key: '',
      children: [
        {
          title: '',
          key: ''
        },
      ],
    },
  ],
}])
//分类数据获取
const GetCategoryListDataFunc = () => {
  GetCategoryList(200, 1).then((res: any) => {
    if (res.data.code == 200) {
      for (let i = 0; i < res.data.data.length; i++) {
        CategoryForm[i] = res.data.data[i]
        console.log(CategoryForm)
      }
      let TreeDataF1 = res.data.data.filter((item: any) => item.depth === 1)
      let TreeDataF2 = res.data.data.filter((item: any) => item.depth === 2)
      let TreeDataF3 = res.data.data.filter((item: any) => item.depth === 3)

      //[depth]:1
      for (let j = 0; j < TreeDataF1.length; j++) {
        CategoryTreeData.value[j] = TreeDataF1[j]
        CategoryTreeData.value[j].title = TreeDataF1[j].name
        CategoryTreeData.value[j].key = TreeDataF1[j].id
        CategoryTreeData.value[j].children = []
        //[depth]:2
        if (TreeDataF2.length != 0) {
          for (let k = 0; k < TreeDataF2.length; k++) {
            if (CategoryTreeData.value[j].key === TreeDataF2[k].parent_id) {
              CategoryTreeData.value[j].children[k] = TreeDataF2[k]
              CategoryTreeData.value[j].children[k].title = TreeDataF2[k].name
              CategoryTreeData.value[j].children[k].key = TreeDataF2[k].id
              CategoryTreeData.value[j].children[k].children = []
              //[depth]:3
              if (TreeDataF3.length != 0) {
                for (let i = 0; i < TreeDataF3.length; i++) {
                  if (CategoryTreeData.value[j].children[k].key === TreeDataF3[i].parent_id) {
                    CategoryTreeData.value[j].children[k].children[i] = TreeDataF3[i]
                    CategoryTreeData.value[j].children[k].children[i].title = TreeDataF3[i].name
                    CategoryTreeData.value[j].children[k].children[i].key = TreeDataF3[i].id
                  }
                }
              }
            }
          }
        }
      }
    } else {
      Message.error('获取分类列表数据异常')
    }
  })
}
GetCategoryListDataFunc()

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
const addCategoryHandleBeforeOk = (done: any) => {
  CreateCategory(addCategoryForm).then((res: any) => {
    console.log(res.data.code)
    if (res.data.code == 200) {
      Notification.success('添加分类成功！')
      //刷新分类列表数据
      GetCategoryListDataFunc()
      addCategoryFormRef.value['resetFields'](['name','short_name','desc','parent_id','depth'])
    } else {
      Notification.error('添加分类失败，请检查数据是否合法!')
      addCategoryFormRef.value['resetFields'](['name','short_name','desc','parent_id','depth'])
    }
  })
  window.setTimeout(() => {
    done()
  }, 2000)
};
//[添加分类][取消事件]
const addCategoryHandleCancel = () => {
  //刷新分类列表数据
  GetCategoryListDataFunc()
  addCategoryVisible.value = false;
  addCategoryFormRef.value['resetFields'](['name','short_name','desc','parent_id','depth'])
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
            <a-input v-model="addCategoryForm.name"/>
          </a-form-item>
          <a-form-item field="name" label="分类缩略名">
            <a-input v-model="addCategoryForm.short_name"/>
          </a-form-item>
          <a-form-item field="parent_id" label="父级分类">
            <a-select v-model="addCategoryForm.parent_id">
              <a-option :value="0">不选择</a-option>
              <a-option v-for="item of CategoryForm" :value="item.id" :label="item.name"/>
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
        {{CategoryTreeData}}
        <a-tree
            :data="CategoryTreeData"
            :show-line="true"
        />

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