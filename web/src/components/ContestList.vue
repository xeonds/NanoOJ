<template>
  <el-table :data="contests">
    <el-table-column prop="id" label="测试标题"></el-table-column>
    <el-table-column prop="title" label="开始时间"></el-table-column>
    <el-table-column prop="difficulty" label="测试时长"></el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "ContestList",
   props: {
    contests: {
      type: Array,
      default: () => []
    }
   }
};
</script>

<style scoped>
</style>