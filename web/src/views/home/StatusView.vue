<template>
  <div>
    <div class="card-header">
      <h2>提交记录</h2>
      <el-button @click="refresh()" type="primary" text>Refresh</el-button>
    </div>
    <el-table :data="submissions.slice((currentPage - 1) * pageSize, currentPage * pageSize)">
      <el-table-column label="提交时间">
        <template #default="{ row }">
          <span type="info">{{ time.fromString(row.CreatedAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="ID" label="提交ID"></el-table-column>
      <el-table-column prop="problem_id" label="题目id"></el-table-column>
      <el-table-column label="状态">
        <template #default="{ row }">
          <el-tag class="ml-2" :type="statusTag(row.status)">{{ row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="信息">
        <template #default="{ row }">
          <el-button type="default" @click="showInfo(row)">查看</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="rank" label="分数"></el-table-column>
      <el-table-column prop="time" label="用时"></el-table-column>
      <el-table-column prop="user_id" label="提交者"></el-table-column>
    </el-table>
    <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :small="small" :disabled="disabled"
      :background="background" layout="prev, pager, next, jumper" :total="submissions.length" />
    <el-dialog v-model="dialogVisible" title="Submission Information" width="30%">
      <el-descriptions :bordered="true" :column="1">
        <el-descriptions-item label="Submission ID">{{ dialogData.ID }}</el-descriptions-item>
        <el-descriptions-item label="Problem ID">{{ dialogData.problem_id }}</el-descriptions-item>
        <el-descriptions-item label="Status"><el-tag class="ml-2" :type="statusTag(dialogData.status)">{{
        dialogData.status
      }}</el-tag></el-descriptions-item>
        <el-descriptions-item label="Information">{{ dialogData.information ? dialogData.information.join('\n') : ''
          }}</el-descriptions-item>
        <el-descriptions-item label="Time">{{ dialogData.time }}</el-descriptions-item>
        <el-descriptions-item label="User ID">{{ dialogData.user_id }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { getDataArr } from "@/utils/http"
import { time } from "@/utils/datetime";
import { Submission } from "@/model";
import { EpPropMergeType } from "element-plus/es/utils/index.mjs";

type _EpPropMergeType = EpPropMergeType<StringConstructor, "success" | "warning" | "info" | "primary" | "danger", unknown>;

const currentPage = ref(1);
const pageSize = ref(10);
const small = ref(false);
const disabled = ref(false);
const background = ref(false);
const dialogVisible = ref(false);
const dialogData = ref({} as Submission);

const { data: submissions, get } = getDataArr<Submission>("/submissions");
const statusTag = (status: string): _EpPropMergeType => { return { "Pending": "info", "InProgress": "info", "Accepted": "success", "WrongAnswer": "danger", "TimeLimitExceeded": "warning", "MemoryLimitExceeded": "warning", "RuntimeError": "danger", "CompilationError": "danger" }[status] as _EpPropMergeType; }
const showInfo = (row: Submission) => {
  dialogData.value = row;
  dialogVisible.value = true;
};

onMounted(async () => {
  submissions.value = await get();
});

const refresh = async () => {
  submissions.value = await get();
}
</script>

<style scoped>
.el-pagination {
  padding-top: 1rem;
}
</style>