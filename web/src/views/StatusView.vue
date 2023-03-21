<template>
  <div style="width: 80vw">
    <h1>提交记录</h1>
    <el-table :data="submission" style="width: 100%">
      <el-table-column prop="submitTime" label="提交时间"></el-table-column>
      <el-table-column prop="submitID" label="提交ID"></el-table-column>
      <el-table-column prop="problemID" label="题目id"></el-table-column>
      <el-table-column prop="status" label="状态"></el-table-column>
      <el-table-column prop="timeUsed" label="用时"></el-table-column>
      <el-table-column prop="submitter" label="提交者"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'StatusView',
  props: {
    submission: {
      type: Array,
      default: () => []
    }
  }
}
</script>

<style scoped>
</style>
