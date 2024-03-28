<template>
  <div>
    <div class="card-header">
      <h2>{{ t('nav.status') }}</h2>
      <el-button @click="refresh()" type="primary" text>{{ t('message.refresh') }}</el-button>
    </div>
    <el-table :data="submissions.slice((currentPage - 1) * pageSize, currentPage * pageSize)">
      <el-table-column :label="t('submission.submit-time')">
        <template #default="{ row }">
          <span type="info">{{ time.fromString(row.CreatedAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
      <el-table-column prop="problem_id" :label="t('submission.problem-id')"></el-table-column>
      <el-table-column :label="t('submission.status')">
        <template #default="{ row }">
          <el-tag class="ml-2" :type="statusTag(row.status)">{{ row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="t('submission.info')">
        <template #default="{ row }">
          <el-button type="default" @click="showInfo(row)">{{ t('submission.check') }}</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="rank" :label="t('submission.rank')"></el-table-column>
      <el-table-column prop="time" :label="t('submission.time-cost')"></el-table-column>
      <el-table-column prop="user_id" :label="t('submission.user')"></el-table-column>
    </el-table>
    <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :small="small" :disabled="disabled"
      :background="background" layout="prev, pager, next, jumper" :total="submissions.length" />
    <el-dialog v-model="dialogVisible" :title="t('submission.details')" width="30%">
      <el-descriptions :bordered="true" :column="1">
        <el-descriptions-item :label="t('message.id')">{{ dialogData.ID }}</el-descriptions-item>
        <el-descriptions-item :label="t('submission.problem-id')">{{ dialogData.problem_id }}</el-descriptions-item>
        <el-descriptions-item :label="t('submission.status')"><el-tag class="ml-2" :type="statusTag(dialogData.status)">{{
        dialogData.status
      }}</el-tag></el-descriptions-item>
        <el-descriptions-item :label="t('submission.info')">{{ dialogData.information ? dialogData.information.join('\n') : ''
          }}</el-descriptions-item>
        <el-descriptions-item :label="t('submission.time-cost')">{{ dialogData.time }}</el-descriptions-item>
        <el-descriptions-item :label="t('submission.user-id')">{{ dialogData.user_id }}</el-descriptions-item>
        <!-- TODO: commited code display line wrap fix -->
        <el-descriptions-item :label="t('submission.code')"><code>{{ dialogData.code }}</code></el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { getDataArr, handleArrRefresh } from "@/utils/http"
import { time } from "@/utils/datetime";
import { Submission } from "@/model";
import { EpPropMergeType } from "element-plus/es/utils/index.mjs";
import { useI18n } from "vue-i18n";

type _EpPropMergeType = EpPropMergeType<StringConstructor, "success" | "warning" | "info" | "primary" | "danger", unknown>;

const { t } = useI18n();
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

const refresh = async () => handleArrRefresh<Submission>(submissions, await get())
</script>

<style scoped>
.el-pagination {
  padding-top: 1rem;
}
</style>