<script setup lang="ts">
import eventBus from "@/plugin/event-bus/event-bus";
import { ref, onMounted } from 'vue';
import axios from 'axios';

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id' },
  { title: 'Name', dataIndex: 'name', key: 'name' },
  { title: 'Short Name', dataIndex: 'short_name', key: 'short_name' },
  { title: 'Description', dataIndex: 'desc', key: 'desc' },
  { title: 'Parent ID', dataIndex: 'parent_id', key: 'parent_id' },
  { title: 'Depth', dataIndex: 'depth', key: 'depth' },
];

const data = ref([]);
const pagination = ref({
  current: 1,
  pageSize: 10,
  total: data.value.length,
});
const rowSelection = ref({
  onChange: (selectedRowKeys, selectedRows) => {
    console.log('selectedRowKeys:', selectedRowKeys);
    console.log('selectedRows:', selectedRows);
  },
});
const stickyHeader = ref(100);

onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:3001/api/Category/list/10/1');
    data.value = response.data.data;
    console.log(response.data.data)
    //pagination.value.total = response.data.length;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
});

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
  <a-table
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :row-selection="rowSelection"
      :sticky-header="stickyHeader"
      row-key="id"
  />
</template>

<style scoped>

</style>