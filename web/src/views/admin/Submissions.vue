<template>
    <h2>Submissions</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>All Submissions</span>
                <span>
                    <el-button @click="refresh()" type="primary" text>Refresh</el-button>
                </span>
            </div>
        </template>
        <el-table :data="submissions" style="width: 100%">
            <el-table-column prop="ID" label="ID"></el-table-column>
            <el-table-column prop="problem_id" label="Problem ID"></el-table-column>
            <el-table-column prop="user_id" label="User ID"></el-table-column>
            <el-table-column prop="language" label="Language"></el-table-column>
            <el-table-column prop="status" label="Status"></el-table-column>
            <el-table-column prop="CreatedAt" label="CreatedAt"></el-table-column>
            <el-table-column label="Actions">
                <template #default="{ row }">
                    <el-button @click="handleDeleteSubmission(row.ID)">Delete</el-button>
                    <el-button @click="handleReRunJudge(row.ID)">Re-Run Judge</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
</template>

<script lang="ts" setup>
import { Submission } from '@/model';
import { getDataArr } from '@/utils/http';
import api from '@/api'

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