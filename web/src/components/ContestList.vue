<template>
  <el-table :data="contests" @row-click="(row) => { this.$router.push(`/contests/${row.id}`) }">
    <el-table-column prop="id" label="ID"></el-table-column>
    <el-table-column prop="title" label="标题"></el-table-column>
    <el-table-column prop="start_time" label="开始时间"></el-table-column>
    <el-table-column prop="end_time" label="结束时间"></el-table-column>
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

<style scoped></style>