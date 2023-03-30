<template>
  <el-table :data="problems">
    <el-table-column label="问题ID">
      <template v-slot="{ row }">
        <router-link :to="'/problem/' + row.id">{{ row.id }}</router-link>
      </template>
    </el-table-column>
    <el-table-column prop="title" label="问题标题"></el-table-column>
    <el-table-column prop="difficulty" label="题目难度"></el-table-column>
    <el-table-column prop="passRate" label="通过率"></el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "ProblemList",
  props: {
    problems: {
      type: Array,
      default: () => []
    }
  }
};
</script>

<style scoped></style>
