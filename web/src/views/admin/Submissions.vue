<template>
    <h2>Submissions</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>All Submissions</span>
            </div>
        </template>
        <el-table :data="submissions" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="problem_id" label="Problem ID"></el-table-column>
            <el-table-column prop="user_id" label="User ID"></el-table-column>
            <el-table-column prop="language" label="Language"></el-table-column>
            <el-table-column prop="status" label="Status"></el-table-column>
            <el-table-column prop="CreatedAt" label="CreatedAt"></el-table-column>
            <el-table-column label="Actions">
                <template #default="{ row }">
                    <el-button @click="deleteSubmission(row.id)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
</template>

<script lang="ts" setup>
import { Submission } from '@/model';
import { getDataArr } from '@/utils/http';

const { data: submissions, get: getSubmissions } = getDataArr<Submission>('submissions');

const deleteSubmission = async (id: number) => {
    await fetch(`/api/submissions/${id}`, {
        method: 'DELETE',
    });
    submissions.value = submissions.value.filter(submission => submission.id !== id);
}

onMounted(async () => {
    submissions.value = await getSubmissions();
})
</script>