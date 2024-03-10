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

<script lang="ts" setup>
import { getDataArr } from "@/utils/http"

const currentPage = ref(1);
const pageSize = ref(10);
const small = ref(false);
const disabled = ref(false);
const background = ref(false);

const { data: submissions, get } = getDataArr<Submission>("submissions");

const status = (status) => ["Pending", "InProgress", "Accepted", "WrongAnswer", "TimeLimitExceeded", "MemoryLimitExceeded", "RuntimeError", "CompilationError"][status] || "Unknown";
const statusTag = (status) => ["info", "info", "success", "danger", "warning", "warning", "danger", "danger"][status] || "info";

const showInfo = (row) => {
  ElMessage({
    message: row.information.join("\n"),
    type: 'info',
    showClose: true,
  });
};


const formattedSubmissions = () => submissions.value.slice((currentPage.value - 1) * pageSize.value, currentPage.value * pageSize.value);
</script>

<style scoped>
.el-pagination {
  padding-top: 1rem;
}
</style>