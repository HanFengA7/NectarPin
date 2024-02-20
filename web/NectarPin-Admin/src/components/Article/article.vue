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
let SelectedKeys = ref(["Article"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);

/*
文章列表
 */
//公共数据
const rowSelection = reactive({
  type: 'checkbox',
  showCheckedAll: true,
});
const selectedKeys = ref([]);
const pagination = {pageSize: 5}
const scroll = {
  x: 1000,
  y: 200
}
const columns = [
  {
    title: '标题',
    dataIndex: 'title',
    fixed: 'left',
    width: 10
  },
  {
    title: '分类',
    dataIndex: 'category',
  },
  {
    title: '标签',
    dataIndex: 'tags',
  },
  {
    title: '创作时间',
    dataIndex: 'date',
  },
  {
    title: '更多操作',
    slotName: 'optional',
    fixed: 'right',
    width: 50
  },
]
const data = reactive([{
  key: '1',
  title: '每个人都是自己命运的工匠',
  tags: '1,2,3',
  category: '励志鸡汤',
  date: '2023年9月2日 13:19',
}, {
  key: '2',
  title: '每个人都是自己命运的工匠2',
  tags: '',
  category: '励志鸡汤',
  date: '2023年9月2日 13:19',
}, {
  key: '3',
  title: '每个人都是自己命运的工匠3',
  tags: '',
  category: '励志鸡汤',
  date: '2023年9月2日 13:19',
}])
//文章列表数据


//选择器
const GetSelectedKey = (key: any) => {
  selectedKeys.value = key
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
      <!--[1] 页头卡片 [Start]-->
      <div class="Article-heard-PC">
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

          <template #extra>
            <a-space>
              <a-button @click="router.push({name:'Article/add'})">发布文章</a-button>
              <a-button @click="console.log(selectedKeys)">批量删除</a-button>
            </a-space>
          </template>
        </a-page-header>
      </div>
      <!--[1] 页头卡片 [End]-->

      <!--[2] 文章列表 [Start]-->
      <div class="Article-Table-PC">
        <a-table
            row-key="key"
            :columns="columns"
            :data="data"
            :row-selection="rowSelection"
            :pagination="pagination"
            :sticky-header="100"
            @select="GetSelectedKey"
        >
          <template #columns>
            <a-table-column title="标题" data-index="title"></a-table-column>
            <a-table-column title="分类" data-index="category"></a-table-column>
            <a-table-column title="标签">
              <template #cell="{ record }">
                <a-tag
                    v-for="tag in record.tags.split(',')"
                    :key="tag"
                    :color="'gray'"
                    bordered
                    style="margin-right: 5px;"
                >
                  {{ tag }}
                </a-tag>
              </template>
            </a-table-column>
            <a-table-column title="时间" data-index="date"></a-table-column>
            <a-table-column title="更多操作">
              <template #cell="{ record }">
                <a-button style="right: 10px">编辑</a-button>
                <a-button>删除</a-button>
              </template>
            </a-table-column>
          </template>
        </a-table>
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
        <a-table
            row-key="name"
            :columns="columns"
            :data="data"
            v-model:selectedKeys="selectedKeys"
            :row-selection="rowSelection"
            :pagination="pagination"
            :sticky-header="100"
            :size="'mini'"
            :scroll="scroll"
        >
          <template #columns>
            <a-table-column title="标题" data-index="title"></a-table-column>
            <a-table-column title="分类" data-index="category"></a-table-column>
            <a-table-column title="时间" data-index="date"></a-table-column>
            <a-table-column title="">
              <template #cell="{ record }">
                <a-button style="right: 10px">编辑</a-button>
                <a-button>删除</a-button>
              </template>
            </a-table-column>
          </template>
        </a-table>
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
.Article-heard-PC {
  background: #ffffff;
  padding: 0 50px 5px 25px;
  height: 100px;
}

.Article-heard-PE {
  background: #ffffff;
  padding: 0 50px 5px 15px;
  height: 100px;
}

/*
[2] 文章表格 CSS
*/
.Article-Table-PC {
  margin: 30px 35px;
  padding: 10px;
  background: #ffffff;
  border-radius: 15px;
}
</style>