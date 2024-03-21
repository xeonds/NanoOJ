<template>
    <div>
        <el-row :gutter="20">
            <h2>Profile</h2>
            <el-col :span="12">
                <el-card>
                    <p>Username: {{ username }}</p>
                    <p>Rank: {{ rank }}</p>
                </el-card>
            </el-col>
        </el-row>
        <el-row>
            <h4>My Submissions</h4>
            <el-table :data="submissions">
                <el-table-column prop="ID" label="ID"></el-table-column>
                <el-table-column prop="problem_id" label="Problem ID"></el-table-column>
                <el-table-column prop="language" label="Language"></el-table-column>
                <el-table-column label="Code">
                    <template #default="{ row }">
                        <el-link :href="row.code" target="_blank">View</el-link>
                    </template>
                </el-table-column>
                <el-table-column prop="information" label="Information"></el-table-column>
                <el-table-column prop="status" label="Status"></el-table-column>
            </el-table>
        </el-row>
    </div>
</template>

<script lang="ts" setup>
import { Rank, Submission } from '@/model';
import { getDataArr, getData } from '@/utils/http';
import { getUsername } from '@/utils/login';

const username = getUsername();
// TODO: use userid instead of username
const { data: submissions, get: getSubmissions } = getDataArr<Submission>('/submissions/' + username);
const { data: rank, get: getRank } = getData<Rank>('/ranks/' + username);

onMounted(async () => {
    submissions.value = await getSubmissions();
    rank.value = await getRank();
});
</script>
