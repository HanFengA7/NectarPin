<script setup lang="ts">
import eventBus from "@/plugin/event-bus/event-bus";
import {computed, reactive, ref} from 'vue';
import router from "@/router";
import {GetCategoryList} from "@/api/Category/get";
import {Message} from "@arco-design/web-vue";
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
[HeardCardOnBack] 返回函数
*/
const HeardCardOnBack = () => {
  router.push({name: 'Category'})
}


/*
[模块]添加分类
*/
//[添加分类][数据]
const formRef = ref()
const form = reactive({
  name: '',
  short_name: '',
  desc: '',
  parent_id: 0,
  depth: 1,
})
const CategoryFormListAdd = <any>ref([])
//[添加分类][获取数据]
const GetCategoryListData = async () => {
  await GetCategoryList(100000, 1).then((res: any) => {
    if (res.data.code === 200) {
      CategoryFormListAdd.value = res.data.data.filter((item: any) => item.depth === 1 || item.depth === 2)
    } else {
      Message.error('获取分类数据异常！')
    }
  })
}
GetCategoryListData()
//[添加分类]获取反转后的CategoryFormListAdd数组
const reversedCategoryFormListAdd = computed(() => CategoryFormListAdd.value.slice().reverse())
//[添加分类][添加数据][提交事件]
const handleSubmit = (data: any) => {
  if (data.error === undefined) {
    CreateCategory(data.values).then((res: any) => {
      if (res.data.code == 200) {
        Message.success(res.data.msg)
        formRef.value.resetFields()
      } else {
        if (res.data.msg["name"] != undefined) {
          Message.error(res.data.msg["name"])
        }
        if (res.data.msg["short_name"] != undefined) {
          Message.error(res.data.msg["short_name"])
        }
        if (res.data.msg["exist_category"] != undefined) {
          Message.error(res.data.msg["exist_category"])
        }
      }
    })
  }
};
//[点击父级分类事件]
const addCategoryParentChange = (id: any) => {
  //id -> depth
  if (id === 0) {
    form.depth = 1
  } else {
    form.depth = CategoryFormListAdd.value.find((item: any) => item.id === id).depth + 1
  }
}
const addCategoryParentVisibleChange = (bool:boolean) => {
  if (bool){
    GetCategoryListData()
  }
}
</script>

<template>
  <a-row>

    <!--电脑端 [Start]-->
    <a-col :xs="{span: 0}" :sm="{span: 0}" :md="{span: 0}" :lg="{span: 24}" :xl="{span: 24}" :xxl="{span: 24}">

      <a-form ref="formRef" :model="form" @submit="handleSubmit">
        <!--[1] 页头卡片 [Start]-->
        <div class="Category-heard-PC">
          <a-page-header
              title="添加分类"
              subtitle="AddCategory"
              :show-back="true"
              @back=HeardCardOnBack
          >
            <template #breadcrumb>
              <a-breadcrumb>
                <a-breadcrumb-item>后台</a-breadcrumb-item>
                <a-breadcrumb-item>分类管理</a-breadcrumb-item>
                <a-breadcrumb-item>添加分类</a-breadcrumb-item>
              </a-breadcrumb>
            </template>

            <template #extra>
              <a-space>
                <a-button html-type="submit">添加分类</a-button>
                <a-button @click="formRef.resetFields()">重置表单</a-button>
              </a-space>
            </template>
          </a-page-header>
        </div>
        <!--[1] 页头卡片 [End]-->

        <!--[2] 添加分类表单 [Start]-->
        <div class="Category-Body-PC">
          <a-form-item field="name" tooltip="填写分类名称" label="分类名称">
            <a-input v-model="form.name"/>
          </a-form-item>
          <a-form-item field="short_name" tooltip="填写分类缩略名,建议使用英文字母" label="分类缩略名">
            <a-input v-model="form.short_name"/>
          </a-form-item>

          <a-form-item field="parent_id" tooltip="请选择父级分类,若不选择则是父级" label="父级分类">
            <a-select
                v-model="form.parent_id"
                @change="addCategoryParentChange"
                @popup-visible-change="addCategoryParentVisibleChange"
            >
              <a-option :value="0">不选择</a-option>
              <a-option v-for="item in reversedCategoryFormListAdd" :value="item.id" :label="item.name"/>
            </a-select>
          </a-form-item>

          <a-form-item v-if="form.parent_id > 0 && form.depth === 1" field="depth"
                       tooltip="请选择分类层级"
                       label="分类层级">
            <a-select v-model="form.depth">
              <a-option :value="1" disabled>一级分类</a-option>
              <a-option :value="2">二级分类</a-option>
              <a-option :value="3">三级分类</a-option>
            </a-select>
          </a-form-item>
          <a-form-item v-if="form.parent_id > 0 && form.depth === 2" field="depth"
                       tooltip="当前是二级分类"
                       label="分类层级">
            <a-select v-model="form.depth">
              <a-option :value="1" disabled>一级分类</a-option>
              <a-option :value="2">二级分类</a-option>
              <a-option :value="3" disabled>三级分类</a-option>
            </a-select>
          </a-form-item>
          <a-form-item v-if="form.parent_id > 0 && form.depth === 3" field="depth"
                       tooltip="当前是三级分类"
                       label="分类层级">
            <a-select v-model="form.depth">
              <a-option :value="1" disabled>一级分类</a-option>
              <a-option :value="2" disabled>二级分类</a-option>
              <a-option :value="3">三级分类</a-option>
            </a-select>
          </a-form-item>
          <a-form-item v-if="form.parent_id === 0" field="depth"
                       tooltip="当前是一级分类"
                       label="分类层级">
            <a-select v-model="form.depth">
              <a-option :value="1">一级分类</a-option>
              <a-option :value="2" disabled>二级分类</a-option>
              <a-option :value="3" disabled>三级分类</a-option>
            </a-select>
          </a-form-item>

          <a-form-item field="desc" tooltip="填写分类描述" label="分类描述">
            <a-textarea v-model="form.desc" :max-length="30" allow-clear show-word-limit/>
          </a-form-item>
        </div>
        <!--[2] 添加分类表单 [End]-->
      </a-form>
    </a-col>
    <!--电脑端 [End]-->

    <!--手机端 [Start]-->
    <a-col :xs="{span: 24}" :sm="{span: 24}" :md="{span: 24}" :lg="{span: 0}" :xl="{span: 0}" :xxl="{span: 0}">
      <!--[1] 页头卡片 [Start]-->
      <div class="Category-heard-PE">
        <a-page-header
            title="分类管理"
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

      <!--[2] 添加分类表单 [Start]-->
      <div class="Article-Table-PC">

      </div>
      <!--[2] 添加分类表单 [End]-->
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
[2] 主体表单 CSS
*/
.Category-Body-PC {
  margin: 30px 25%;
  padding: 15px;
  background: #ffffff;
  border-radius: 15px;
}
</style>