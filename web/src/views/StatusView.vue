<template>
  <div>
    <h1>提交记录</h1>
    <el-table :data="submissions">
      <el-table-column prop="CreatedAt" label="提交时间"></el-table-column>
      <el-table-column prop="ID" label="提交ID"></el-table-column>
      <el-table-column prop="ProblemID" label="题目id"></el-table-column>
      <el-table-column label="状态">
        <template #default="{ row }">
          <el-tag class="ml-2" type="success">Accepted</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="Time" label="用时"></el-table-column>
      <el-table-column prop="UserID" label="提交者"></el-table-column>
    </el-table>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  created: function () {
    this.fetchSubmissions();
  },
  computed: {
    ...mapGetters({ submissions: "getSubmissions" }),
  },
  methods: {
    ...mapActions({ fetchSubmissions: "fetchSubmissions", getUserById: "fetchUserInfo" }),
    Username: function (id) {
      return this.getUserById(id).Username;
    }
  },
};
</script>

<style scoped></style>
