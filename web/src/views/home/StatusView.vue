<template>
  <div>
    <h1>提交记录</h1>
    <el-table :data="submissions.slice((currentPage - 1) * pageSize, currentPage * pageSize)">
      <el-table-column label="提交时间">
        <template #default="{ row }">
          <span type="info">{{ formatDate(row.CreatedAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="id" label="提交ID"></el-table-column>
      <el-table-column prop="problem_id" label="题目id"></el-table-column>
      <el-table-column label="状态">
        <template #default="{ row }">
          <el-tag class="ml-2" :type="statusTag(row.status)">{{ decodeStatus(row.status) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="信息">
        <template #default="{ row }">
          <el-button type="default" @click="showInfo(row)">查看</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="time" label="用时"></el-table-column>
      <el-table-column prop="user_id" label="提交者"></el-table-column>
    </el-table>
    <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :small="small" :disabled="disabled"
      :background="background" layout="prev, pager, next, jumper" :total="submissions.length" />
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import utils from '../../utils';

export default {
  created: function () {
    this.fetchSubmissions();
  },
  data: function () {
    return {
      currentPage: 1,
      pageSize: 10,
      small: false,
      disabled: false,
      background: false,
    };
  },
  computed: {
    ...mapGetters({ submissions: "getSubmissions" }),
  },
  methods: {
    ...mapActions({ fetchSubmissions: "fetchSubmissions", getUserById: "fetchUserInfo" }),
    Username: function (id) {
      return this.getUserById(id).Username;
    },
    decodeStatus: function (status) {
      switch (status) {
        case 0: return "Pending";
        case 1: return "InProgress"
        case 2: return "Accepted"
        case 3: return "WrongAnswer"
        case 4: return "TimeLimitExceeded"
        case 5: return "MemoryLimitExceeded"
        case 6: return "RuntimeError"
        case 7: return "CompilationError"
        default:
          return "Unknown";
      }
    },
    statusTag: function (status) {
      switch (status) {
        case 0: return "info";
        case 1: return "info"
        case 2: return "success"
        case 3: return "danger"
        case 4: return "warning"
        case 5: return "warning"
        case 6: return "danger"
        case 7: return "danger"
        default:
          return "info";
      }
    },
    formatDate: function (date) {
      return utils.formatDate(date);
    },
    showInfo: function (row) {
      this.$message({
        message: row.information.join("\n"),
        type: 'info',
        showClose: true,
      });
    },
  },
};
</script>

<style scoped>
.el-pagination {
  padding-top: 1rem;
}
</style>
