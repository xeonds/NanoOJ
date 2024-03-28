<template>
    <div>
        <h2>{{ t('nav.profile') }}</h2>
        <el-row :gutter="20">
            <el-col :span="12">
                <el-card>
                    <p>{{ t('user.name') }}: {{ user.username }}</p>
                    <p>{{ t('user.rank') }}: {{ rank }}</p>
                </el-card>
            </el-col>
        </el-row>
        <el-row>
            <h4>{{ t('profile.my-submissions') }}</h4>
            <el-table :data="submissions">
                <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
                <el-table-column prop="problem_id" :label="t('submission.problem-id')"></el-table-column>
                <el-table-column prop="language" :label="t('message.language')"></el-table-column>
                <el-table-column :label="t('submission.code')">
                    <template #default="{ row }">
                        <el-link :href="row.code" target="_blank">{{ t('message.view') }}</el-link>
                    </template>
                </el-table-column>
                <el-table-column prop="information" :label="t('submission.info')"></el-table-column>
                <el-table-column prop="status" :label="t('submission.status')"></el-table-column>
            </el-table>
        </el-row>
    </div>
</template>

<script lang="ts" setup>
import { Rank, Submission, User } from '@/model';
import { getDataArr, getData } from '@/utils/http';
import { getUsername } from '@/utils/login';
import { useI18n } from 'vue-i18n';

const {t } = useI18n();
const username = getUsername();
const { data: user, get: getUser } = getData<User>('/users/' + username);
const { data: submissions, get: getSubmissions } = getDataArr<Submission>('/submissions/' + username);
const { data: rank, get: getRank } = getData<Rank>('/ranks/' + username);

onMounted(async () => {
    submissions.value = await getSubmissions();
    rank.value = await getRank();
    user.value = await getUser();
});
</script>
