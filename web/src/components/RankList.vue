<template>
  <el-table :data="ranks">
    <el-table-column props="index" label="Rank"></el-table-column>
    <el-table-column prop="UserName" label="用户名"></el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "RankList",
  props: {
    ranks: {
      type: Array,
      default: () => []
    }
  }
};
</script>

<style scoped></style>