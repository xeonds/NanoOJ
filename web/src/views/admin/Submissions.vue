<template>
    <h2>{{ t('nav.submissions') }}</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>{{ t('submission.all') }}</span>
                <span>
                    <el-button @click="refresh()" type="primary" text>{{ t('message.refresh') }}</el-button>
                </span>
            </div>
        </template>
        <el-table :data="submissions" style="width: 100%">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column prop="problem_id" :label="t('nav.problems')"></el-table-column>
            <el-table-column prop="user_id" :label="t('message.user-id')"></el-table-column>
            <el-table-column prop="language" :label="t('message.language')"></el-table-column>
            <el-table-column prop="status" :label="t('nav.status')"></el-table-column>
            <el-table-column prop="CreatedAt" :label="t('message.created-at')"></el-table-column>
            <el-table-column :label="t('message.action')">
                <template #default="{ row }">
                    <el-button @click="handleDeleteSubmission(row.ID)">{{ t('message.delete') }}</el-button>
                    <el-button @click="handleReRunJudge(row.ID)">{{ t('submission.re-run') }}</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
</template>

<script lang="ts" setup>
import { Submission } from '@/model';
import { getDataArr } from '@/utils/http';
import api from '@/api'
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const { data: submissions, get: getSubmissions } = getDataArr<Submission>('/submissions');
const handleDeleteSubmission = async (id: number) => {
    const { err } = await api.deleteSubmission(id);
    if (err.value != null ) ElMessage.error(err.value.message);
    else {
        ElMessage.success('Submission deleted successfully');
        submissions.value = submissions.value.filter(submission => submission.ID !== id);
    }
}
const handleReRunJudge = async (id: number) => {
    const { err } = await api.reRunSubmission(id);
    if (err.value != null ) ElMessage.error(err.value.message);
    else {
        ElMessage.success('Judge re-run successfully');
        submissions.value = await getSubmissions();
    }
}

onMounted(async () => {
    submissions.value = await getSubmissions();
})

const refresh = async () => {
    submissions.value = await getSubmissions();
}
</script>