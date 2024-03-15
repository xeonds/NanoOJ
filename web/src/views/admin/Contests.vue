<template>
    <h2>Contests</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>All Contests</span>
                <el-button @click="createContestDialogVisible = true" type="primary">Create Contest</el-button>
            </div>
        </template>
        <el-table :data="contests" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="title" label="Title"></el-table-column>
            <el-table-column prop="CreatedAt" label="Created At"></el-table-column>
            <el-table-column label="Actions">
                <template #default="{ row }">
                    <el-button @click="editContest(row)">Edit</el-button>
                    <el-button @click="deleteContest(row.id)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <el-dialog v-model="createContestDialogVisible" title="Create Contest">
        <el-form :model="newContest" label-width="128px">
            <el-form-item label="Title">
                <el-input v-model="newContest.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
                <el-input v-model="newContest.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Problems">
                <el-select v-model="newContest.problems" multiple placeholder="Please select">
                    <el-option v-for="problem in problems" :key="problem.id" :label="problem.title" :value="problem.id">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="Time Duration">
                <el-date-picker v-model="newContestTimePeriod" type="datetimerange" range-separator="To"
                    start-placeholder="Start date" end-placeholder="End date" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="createContestDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="handleCreateContest">Create</el-button>
        </div>
    </el-dialog>
    <el-dialog v-model="editContestDialogVisible" title="Edit Contest">
        <el-form :model="selectedContest" label-width="128px">
            <el-form-item label="Title">
                <el-input v-model="selectedContest.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
                <el-input v-model="selectedContest.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Problems">
                <el-select v-model="selectedContest.problems" multiple placeholder="Please select">
                    <el-option v-for="problem in problems" :key="problem.id" :label="problem.title" :value="problem.id">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="Time Duration">
                <el-date-picker v-model="selectedContestTimePeriod" type="datetimerange" range-separator="To"
                    start-placeholder="Start date" end-placeholder="End date" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="editContestDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="handleEditContest">Update</el-button>
        </div>
    </el-dialog>
</template>

<script lang="ts" setup>
import api from '@/api';
import { Contest, Problem } from '@/model';
import { getDataArr } from '@/utils/http';

const createContestDialogVisible = ref(false);
const editContestDialogVisible = ref(false);
const { data: contests, get: getContests } = getDataArr<Contest>('/contests');
const { data: problems, get: getProblems } = getDataArr<Problem>('/problems');

const newContest: Ref<Contest> = ref({} as Contest);
const newContestTimePeriod = ref('');
const handleCreateContest = () => {
    newContest.value.start_time = newContestTimePeriod.value[0];
    newContest.value.end_time = newContestTimePeriod.value[1];
    newContest.value.problems = problems.value.filter(problem => newContest.value.problems.includes(problem.id as any));
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
    selectedContest.value.problems = problems.value.filter(problem => selectedContest.value.problems.includes(problem.id as any));
    api.updateContests(selectedContest.value, selectedContest.value.id).then(() => {
        contests.value = contests.value.map(contest => {
            if (contest.id === selectedContest.value.id) {
                return selectedContest.value;
            }
            return contest;
        });
        editContestDialogVisible.value = false;
    });
};

const deleteContest = (id: number) => {
    api.deleteContest(id).then(() => {
        contests.value = contests.value.filter(contest => contest.id !== id);
    });
};

onMounted(async () => {
    contests.value = await getContests();
    problems.value = await getProblems();
});
</script>