<template>
    <h2>{{ t('nav.contests') }}</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>{{ t('contest.all') }}</span>
                <span>
                    <el-button @click="refresh()" type="primary" text>{{ t('message.refresh') }}</el-button>
                    <el-button @click="createContestDialogVisible = true" type="primary">{{ t('message.create')
                        }}</el-button>
                </span>
            </div>
        </template>
        <el-table :data="contests" style="width: 100%">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column prop="title" :label="t('message.title')"></el-table-column>
            <el-table-column :label="t('message.created-at')">
                <template #default="{ row }"> <span type="info">{{ time.formatDate(row.CreatedAt) }}</span> </template>
            </el-table-column>
            <el-table-column :label="t('message.action')">
                <template #default="{ row }">
                    <el-button @click="editContest(row)">{{ t('message.edit') }}</el-button>
                    <el-button @click="deleteContest(row.ID)">{{ t('message.delete') }}</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <el-dialog v-model="createContestDialogVisible" :title="t('message.create')">
        <el-form :model="newContest" label-width="128px">
            <el-form-item :label="t('message.title')">
                <el-input v-model="newContest.title"></el-input>
            </el-form-item>
            <el-form-item :label="t('message.description')">
                <el-input v-model="newContest.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item :label="t('nav.problems')">
                <el-select v-model="newContest.problems" multiple :placeholder="t('message.please-select')">
                    <el-option v-for="problem in problems" :key="problem.ID" :label="problem.title" :value="problem.ID">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item :label="t('message.time-duration')">
                <el-date-picker v-model="newContestTimePeriod" type="datetimerange" :range-separator="t('message.to')"
                    :start-placeholder="t('message.start-date')" :end-placeholder="t('message.end-date')" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="createContestDialogVisible = false">{{ t('message.cancel') }}</el-button>
            <el-button type="primary" @click="handleCreateContest">{{ t('message.create') }}</el-button>
        </div>
    </el-dialog>
    <el-dialog v-model="editContestDialogVisible" :title="t('message.edit')">
        <el-form :model="selectedContest" label-width="128px">
            <el-form-item :label="t('message.title')">
                <el-input v-model="selectedContest.title"></el-input>
            </el-form-item>
            <el-form-item :label="t('message.description')">
                <el-input v-model="selectedContest.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item :label="t('contest.problem')">
                <el-select v-model="selectedContest.problems" multiple placeholder="Please select">
                    <el-option v-for="problem in problems" :key="problem.ID" :label="problem.title" :value="problem.ID">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item :label="t('message.time-duration')">
                <el-date-picker v-model="selectedContestTimePeriod" type="datetimerange"
                    :range-separator="t('message.to')" :start-placeholder="t('message.start-date')"
                    :end-placeholder="t('message.end-date')" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="editContestDialogVisible = false">{{ t('message.cancel') }}</el-button>
            <el-button type="primary" @click="handleEditContest">{{ t('message.update') }}</el-button>
        </div>
    </el-dialog>
</template>

<script lang="ts" setup>
import api from '@/api';
import { time } from '@/utils/datetime';
import { Contest, Problem } from '@/model';
import { getDataArr, handleHttp } from '@/utils/http';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const createContestDialogVisible = ref(false);
const editContestDialogVisible = ref(false);
const { data: contests, get: getContests } = getDataArr<Contest>('/contests');
const { data: problems, get: getProblems } = getDataArr<Problem>('/problems');

const newContest: Ref<Contest> = ref({} as Contest);
const newContestTimePeriod = ref('');
const handleCreateContest = () => {
    newContest.value.start_time = newContestTimePeriod.value[0];
    newContest.value.end_time = newContestTimePeriod.value[1];
    newContest.value.problems = problems.value.filter(problem => newContest.value.problems.includes(problem.ID as any));
    api.addContests(newContest.value).then(() => {
        contests.value.push(newContest.value);
        createContestDialogVisible.value = false;
    });
};

const selectedContest: Ref<Contest> = ref({} as Contest);
const selectedContestTimePeriod = ref('');
const editContest = (contest: Contest) => {
    selectedContest.value = contest;
    editContestDialogVisible.value = true;
};
const handleEditContest = () => {
    selectedContest.value.start_time = selectedContestTimePeriod.value[0];
    selectedContest.value.end_time = selectedContestTimePeriod.value[1];
    selectedContest.value.problems = problems.value.filter(problem => selectedContest.value.problems.includes(problem.ID as any));
    api.updateContests(selectedContest.value, selectedContest.value.ID).then(() => {
        contests.value = contests.value.map(contest => {
            if (contest.ID === selectedContest.value.ID) {
                return selectedContest.value;
            }
            return contest;
        });
        editContestDialogVisible.value = false;
    });
};

const deleteContest = async (id: number) => {
    handleHttp(await api.deleteContest(id), () => {
        contests.value = contests.value.filter(contest => contest.ID !== id);
    });
}

onMounted(async () => {
    contests.value = await getContests();
    problems.value = await getProblems();
});

const refresh = async () => {
    contests.value = await getContests();
    problems.value = await getProblems();
};
</script>