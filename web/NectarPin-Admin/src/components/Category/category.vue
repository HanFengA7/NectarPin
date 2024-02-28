<script setup lang="ts">
import eventBus from "@/plugin/event-bus/event-bus";
import {computed, onMounted, reactive, ref} from "vue";
import router from "@/router";
import {GetCategory, GetCategoryList} from "@/api/Category/get";
import {Message, Notification} from "@arco-design/web-vue";
import {DeleteCategory} from "@/api/Category/detele";
import {editCategory} from "@/api/Category/put";


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
const CategoryForm = <any>ref([]);
//分类数据获取
const GetCategoryFormData = async () => {
  const response = await GetCategoryList(CategoryTablePagination.value.pageSize, CategoryTablePagination.value.pageNum);
  CategoryForm.value = response.data.data.map((item: any) => ({
    ...item,
    parent_name: '父级',
  }));
  CategoryTablePagination.value.total = response.data.total
  for (let i = 0; i < CategoryForm.value.length; i++) {
    for (let j = 0; j < CategoryForm.value.length; j++) {
      if (CategoryForm.value[j].id === CategoryForm.value[i].parent_id) {
        CategoryForm.value[i].parent_name = CategoryForm.value[j].name
      }
    }
  }
}
onMounted(async () => {
  try {
    await GetCategoryFormData()
  } catch (error) {
    console.log(error)
  }
})

/*
分类表格
*/
const CategoryTableRowSelection = reactive({
  type: 'checkbox',
  showCheckedAll: true,
});
const CategoryTableSelectedKeys = <any>ref([]);
const CategoryTablePagination = ref({
  pageSize: 15,
  pageNum: 1,
  total: 0,
})
const CategoryTableScroll = {
  x: 1000,
  y: 200
}
const CategoryTableColumns = [
  {
    title: '分类ID',
    dataIndex: 'id',
    sortable: {
      sortDirections: ['ascend', 'descend'],
      defaultSortOrder: 'descend'
    },
    fixed: 'left',
    width: 100,
  },
  {
    title: '分类名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '分类缩略名',
    dataIndex: 'short_name',
    key: 'short_name',
  },
  {
    title: '父级分类',
    dataIndex: 'parent_name',
    key: 'parent_name',
  },
  {
    title: '分类层级',
    dataIndex: 'depth',
    key: 'depth',
  },
  {
    title: '分类描述',
    dataIndex: 'desc',
    key: 'desc',
  },
  {
    title: '更多操作',
    slotName: 'cell',
    fixed: 'right',
    width: 200
  },
]
//选择器
const CategoryTableGetSelectedKey = (key: any) => {
  CategoryTableSelectedKeys.value = key
}
const CategoryTableGetSelectedKeyALL = () => {
  let temp = <any>ref([])
  for (let i = 0; i < CategoryForm.value.length; i++) {
    temp.value[i] = CategoryForm.value[i]["id"];
  }
  CategoryTableSelectedKeys.value = temp.value
}
//分页
const CategoryTablePageChange = (page: number) => {
  CategoryTablePagination.value.pageNum = page
  GetCategoryFormData()
}

/*
[模块]编辑分类
 */
//[编辑分类][数据]
const CategoryFormListEdit = <any>ref([])
const editCategoryFormRef = ref()
const editCategoryForm = reactive({
  id: 0,
  name: '',
  short_name: '',
  desc: '',
  parent_id: 0,
  depth: 1,
})
//[编辑分类][模态框显示]
const editCategoryVisible = ref(false);
//[编辑分类][点击编辑分类事件]
const editCategoryHandleClick = async (id: number) => {
  await GetCategory(id).then((res: any) => {
    if (res.data.code === 200) {
      editCategoryForm.id = res.data.data[0].id
      editCategoryForm.name = res.data.data[0].name
      editCategoryForm.short_name = res.data.data[0].short_name
      editCategoryForm.desc = res.data.data[0].desc
      editCategoryForm.parent_id = res.data.data[0].parent_id
      editCategoryForm.depth = res.data.data[0].depth
    } else {
      Message.error('[01]获取分类数据异常！')
    }
  })
  await GetCategoryList(100000, 1).then((res: any) => {
    if (res.data.code === 200) {
      CategoryFormListEdit.value = res.data.data.filter((item: any) => item.depth === 1 || item.depth === 2)
    } else {
      Message.error('[02]获取分类数据异常！')
    }
  })
  editCategoryVisible.value = true
}
//[编辑分类][取消事件]
const editCategoryHandleCancel = () => {
  editCategoryVisible.value = false;
}
//[编辑分类][提交事件]
const editCategoryHandleBeforeOk = async (done: any) => {
  await editCategory(editCategoryForm.id, editCategoryForm).then((res: any) => {
    if (res.data.code === 200) {
      Notification.success(res.data.msg)
      GetCategoryFormData()
    } else {
      if (res.data.msg["name"] !== undefined) {
        Notification.error(res.data.msg["name"])
        return
      }

      if (res.data.msg["short_name"] !== undefined) {
        Notification.error(res.data.msg["short_name"])
        return
      }

      if (res.data.msg["short_name"] !== undefined && res.data.msg["name"]) {
        Notification.error(res.data.msg["name"])
        Notification.error(res.data.msg["short_name"])
        return
      }

      Notification.error(res.data.msg)

    }
  })


  window.setTimeout(() => {
    done()
  }, 3000)
}
//[编辑分类][点击父级分类事件]
const addCategoryParentChange = (id: any) => {
  //id -> depth
  if (id === 0) {
    editCategoryForm.depth = 1
  } else {
    editCategoryForm.depth = CategoryFormListEdit.value.find((item: any) => item.id === id).depth + 1
  }
}
//[编辑分类][获取反转后的CategoryFormListAdd数组]
const reversedCategoryFormListEdit = computed(() => CategoryFormListEdit.value.slice().reverse())


/*
[模块]删除分类
 */
//删除分类[单个]
const DeleteCategoryFunc = async (id: number) => {
  await DeleteCategory(id).then((res: any) => {
    if (res.data.code === 200) {
      Notification.success(res.data.msg)
    } else {
      Notification.error(res.data.msg)
    }
    GetCategoryFormData()
  })
}
//删除分类[批量删除]
const DeleteCategoryListFunc = async () => {
  for (let i = 0; i < CategoryTableSelectedKeys.value.length; i++) {
    await DeleteCategory(CategoryTableSelectedKeys.value[i]).then((res: any) => {
      if (res.data.code === 200) {
        Notification.success(res.data.msg)
      } else {
        Notification.error(res.data.msg)
      }
    })
  }
  await GetCategoryFormData()
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

      <!--[0-1] 编辑分类-->
      <a-modal
          v-model:visible="editCategoryVisible"
          title="编辑分类"
          @cancel="editCategoryHandleCancel"
          @before-ok="editCategoryHandleBeforeOk"
      >
        <a-form
            ref="editCategoryFormRef"
            :model="editCategoryForm"
        >
          <a-form-item field="name" label="分类名称">
            <a-input v-model="editCategoryForm.name"/>
          </a-form-item>
          <a-form-item field="short_name" label="分类缩略名">
            <a-input v-model="editCategoryForm.short_name"/>
          </a-form-item>
          <a-form-item field="parent_id" label="父级分类">
            <a-select
                v-model="editCategoryForm.parent_id"
                @change="addCategoryParentChange"
            >
              <a-option :value="0">不选择</a-option>
              <a-option v-for="item in reversedCategoryFormListEdit" :value="item.id" :label="item.name"/>
            </a-select>
          </a-form-item>

          <a-form-item v-if="editCategoryForm.parent_id === 0" field="depth"
                       label="分类层级">
            <a-select v-model="editCategoryForm.depth">
              <a-option :value="1">一级分类</a-option>
              <a-option :value="2" disabled>二级分类</a-option>
              <a-option :value="3" disabled>三级分类</a-option>
            </a-select>
          </a-form-item>
          <a-form-item v-if="editCategoryForm.parent_id > 0 && editCategoryForm.depth === 2" field="depth"
                       label="分类层级">
            <a-select v-model="editCategoryForm.depth">
              <a-option :value="1" disabled>一级分类</a-option>
              <a-option :value="2">二级分类</a-option>
              <a-option :value="3" disabled>三级分类</a-option>
            </a-select>
          </a-form-item>
          <a-form-item v-if="editCategoryForm.parent_id > 0 && editCategoryForm.depth === 3" field="depth"
                       label="分类层级">
            <a-select v-model="editCategoryForm.depth">
              <a-option :value="1" disabled>一级分类</a-option>
              <a-option :value="2" disabled>二级分类</a-option>
              <a-option :value="3">三级分类</a-option>
            </a-select>
          </a-form-item>

          <a-form-item field="desc" label="分类描述">
            <a-textarea v-model="editCategoryForm.desc" :max-length="30" allow-clear show-word-limit auto-size/>
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
              <a-breadcrumb-item>分类管理</a-breadcrumb-item>
            </a-breadcrumb>
          </template>

          <template #extra>
            <a-space>
              <a-button @click="router.push({name: 'Category/add'})">添加分类</a-button>
              <a-button @click="DeleteCategoryListFunc">批量删除</a-button>
            </a-space>
          </template>
        </a-page-header>
      </div>
      <!--[1] 页头卡片 [End]-->

      <!--[2] 文章列表 [Start]-->
      <div class="Category-Table-PC">
        <a-table
            row-key="id"
            :columns="CategoryTableColumns"
            :data="CategoryForm"
            :row-selection="CategoryTableRowSelection"
            :pagination="CategoryTablePagination"
            :sticky-header="100"
            :filter-icon-align-left="false"
            @select-all="CategoryTableGetSelectedKeyALL"
            @select="CategoryTableGetSelectedKey"
            @page-change="CategoryTablePageChange"
        >
          <template #cell="{ record }">
            <a-button style="right: 10px" @click="editCategoryHandleClick(record.id)">编辑</a-button>
            <a-button @click="DeleteCategoryFunc(record.id)">删除</a-button>
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
              <a-breadcrumb-item>分类管理</a-breadcrumb-item>
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