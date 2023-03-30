<template>
  <el-table :data="problems">
    <el-table-column label="问题ID">
      <template v-slot="{ row }">
        <router-link :to="'/problem/' + row.id">{{ row.id }}</router-link>
      </template>
    </el-table-column>
    <el-table-column prop="title" label="问题标题"></el-table-column>
    <el-table-column label="题目难度">
      <template v-slot="{ row }">
        <el-rate v-model="row.difficulty" :colors="colors" disabled />
      </template>
    </el-table-column>
    <el-table-column prop="passRate" label="通过率"></el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "ProblemList",
  data: function () {
    return {
      colors: { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' },
    }
  },
  props: {
    problems: {
      type: Array,
      default: () => []
    }
  }
};
</script>
