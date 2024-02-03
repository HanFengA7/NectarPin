<script setup>
//接收父组件的UserInfoData数据
const props = defineProps(['userInfo']);
//传数据给父组件
const emit = defineEmits();
emit('child-event', '这是子组件传递给父组件的数据');
</script>

<template>
  <div class="personalCenter-heard-PC">

  </div>
  <div class="personalCenter-heardBox-PC">

  </div>
</template>

<style scoped>
.personalCenter-heard-PC {
  background: #ffffff;
  padding: 15px 50px 15px 60px;
  height: 100px;
}
.personalCenter-heardBox-PC{
  background: #ffffff;
  height: 500px;
}
</style>