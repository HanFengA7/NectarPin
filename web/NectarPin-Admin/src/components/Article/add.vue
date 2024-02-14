<script setup>
import eventBus from "@/plugin/event-bus/event-bus";
import {reactive, ref} from "vue";
import router from "@/router";

/*
接收父组件数据
*/
//[1]UserInfoData数据
const props = defineProps(['userInfo']);

/*
公共数据
 */
//分类数据
const CategoryForm = reactive([
  {
    cid: '1',
    label: '默认分类',
  }, {
    cid: '2',
    label: '胡言乱语',
  }, {
    cid: '3',
    label: '悠然心喜',
  }, {
    cid: '4',
    label: '悄悄代码',
  }])
//标签数据
const TagsRef = ref([]);
const TagsHandleChange = () => {
  ArticleForm.tags = TagsRef.value.toString()
};
//文章评论开关
const aif_commentRef = ref(true);
const aif_commentHandleChange = () => {
  if (aif_commentRef.value === true){
    ArticleForm.aif_comment = 1
  }else {
    ArticleForm.aif_comment = 0
  }
};
//文章隐藏开关
const aif_hideRef = ref(false);
const aif_hideHandleChange = () => {
  if (aif_hideRef.value === true){
    ArticleForm.aif_hide = 1
  }else {
    ArticleForm.aif_hide = 0
  }
};
//文章加密开关|密码
const aif_encryptRef = ref(false);
const aif_encryptHandleChange = () => {
  if (aif_encryptRef.value === true){
    ArticleForm.aif_encrypt = 1
  }else {
    ArticleForm.aif_encrypt = 0
  }
};

//文章数据
const ArticleForm = reactive({
  uid: props.userInfo["id"],
  cid: '1',
  title: '',
  tags: '',
  desc: '',
  content: '',
  img_url: '',
  aif_comment: 1,
  aif_hide: 0,
  aif_encrypt: 0,
  aif_encrypt_pwd: '',
})

/*
传数据给父组件
*/
//[1]设置侧边栏选择选项
let SelectedKeys = ref(["Article/add"]);
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

  <div class="ArticleAdd-body-PC">
    <a-form :model="ArticleForm">
      <div class="ArticleAdd-body-config-PC">
        <a-row>
          <a-col :span="12">
            <a-form-item field="title" label="文章标题" label-col-flex="100px">
              <a-input v-model="ArticleForm.title"/>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="cid" label="文章分类" label-col-flex="100px">
              <a-select v-model="ArticleForm.cid" placeholder="Please select ...">
                <a-option v-for="item of CategoryForm" :value="item.cid" :label="item.label"/>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="desc" label="文章简介" label-col-flex="100px">
              <a-textarea v-model="ArticleForm.desc" :max-length="200" :show-word-limit="true" auto-size/>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="tags" label="文章标签" label-col-flex="100px">
              <a-input-tag v-model="TagsRef" :size="'medium'" placeholder="填写一些文章的标签"
                           @press-enter="TagsHandleChange" allow-clear/>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="img_url" label="文章展图" label-col-flex="100px">
              <a-input v-model="ArticleForm.img_url" placeholder="填写一个图片外链"/>
            </a-form-item>
          </a-col>
          <a-col :span="4">
            <a-form-item field="aif_comment" label="评论开关" label-col-flex="100px">
              <a-switch
                  type="round"
                  v-model="aif_commentRef"
                  @change="aif_commentHandleChange"
                  />
            </a-form-item>
          </a-col>
          <a-col :span="4">
            <a-form-item field="aif_comment" label="文章隐藏" label-col-flex="100px">
              <a-switch
                  type="round"
                  v-model="aif_hideRef"
                  @change="aif_hideHandleChange"
              />
            </a-form-item>
          </a-col>
          <a-col :span="4">
            <a-form-item field="aif_comment" label="文章加密" label-col-flex="100px">
              <a-switch
                  type="round"
                  v-model="aif_encryptRef"
                  @change="aif_encryptHandleChange"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12" v-if="ArticleForm.aif_encrypt === 1">
            <a-form-item field="aif_encrypt_pwd" label="文章密码" label-col-flex="100px">
              <a-input-password v-model="ArticleForm.aif_encrypt_pwd"/>
            </a-form-item>
          </a-col>
        </a-row>
      </div>
      <v-md-editor
          left-toolbar="undo redo clear | h bold italic strikethrough quote | ul ol table hr | link image code emoji todo-list | save"
          v-model="ArticleForm.Content"
          height="700px"
          style="border-radius: 10px;"
      />
    </a-form>
  </div>


</template>

<style scoped>
.ArticleAdd-heard-PC {
  background: #ffffff;
  padding: 0 50px 5px 25px;
  height: 100px;
}

.ArticleAdd-body-PC {
  border-radius: 15px;
  background: #ffffff;
  margin: 25px;
}

.ArticleAdd-body-config-PC {
  padding: 20px;
}
</style>