<template>
  <el-table :data="ranks">
    <el-table-column prop="index" label="Rank"></el-table-column>
    <el-table-column prop="username" label="用户名"></el-table-column>
    <el-table-column prop="profile" label="简介"></el-table-column>
  </el-table>
</template>

<script lang="ts" setup>
const props = defineProps({
  ranks: {
    type: Array,
    default: () => []
  }
});
</script>

<style scoped></style>