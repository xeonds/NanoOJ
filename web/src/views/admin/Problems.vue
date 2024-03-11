<template>
    <h2>Problems</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>All Problems</span>
                <el-button @click="createProblemDialogVisible = true" type="primary">Create Problem</el-button>
            </div>
        </template>
        <el-table :data="problems" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="title" label="Title"></el-table-column>
            <el-table-column label="Difficulty">
                <template #default="{ row }">
                    <el-rate v-model="row.difficulty" :colors="colors" disabled />
                </template>
            </el-table-column>
            <el-table-column prop="CreatedAt" label="Created At"></el-table-column>
            <el-table-column label="Actions">
                <template #default="{ row }">
                    <el-button @click="editProblem(row)">Edit</el-button>
                    <el-button @click="deleteProblem(row.id)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <el-dialog v-model="createProblemDialogVisible" title="Create Problem">
        <el-form :model="newProblem" label-width="100px">
            <el-form-item label="Title">
                <el-input v-model="newProblem.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
                <el-input v-model="newProblem.description" type="textarea"></el-input>
            </el-form-item>
            <!-- TODO: 增加多项表单 -->
            <el-form-item label="Inputs">
                <el-input v-model="newProblem.inputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Outputs">
                <el-input v-model="newProblem.outputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Difficulty">
                <el-rate v-model="newProblem.difficulty" :colors="colors" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="createProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="createProblem">Create</el-button>
        </div>
    </el-dialog>
    <el-dialog v-model="editProblemDialogVisible" title="Edit Problem">
        <el-form :model="selectedProblem" label-width="100px">
            <el-form-item label="Title">
                <el-input v-model="selectedProblem.title"></el-input>
            </el-form-item>
            <el-form-item label="Description">
                <el-input v-model="selectedProblem.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="Inputs">
                <el-input v-model="selectedProblem.inputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Outputs">
                <el-input v-model="selectedProblem.outputs" type="textarea" />
            </el-form-item>
            <el-form-item label="Difficulty">
                <el-rate v-model="selectedProblem.difficulty" :colors="colors" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="editProblemDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="handleEditProblem">Update</el-button>
        </div>
    </el-dialog>
</template>

<script lang="ts" setup>
import { Problem } from '@/model';
import api from '@/api';
import { getDataArr } from '@/utils/http';

const createProblemDialogVisible = ref(false);
const editProblemDialogVisible = ref(false);
const { data: problems, get: getProblems } = getDataArr<Problem>('/problems');
const colors = { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' }

const newProblem: Ref<Problem> = ref({} as Problem);
const createProblem = () => {
    api.addProblems(newProblem.value).then(() => {
        problems.value.push(newProblem.value);
        createProblemDialogVisible.value = false;
    });
}

const selectedProblem: Ref<Problem> = ref({} as Problem);
const editProblem = (problem: Problem) => {
    selectedProblem.value = problem;
    editProblemDialogVisible.value = true;
}
const handleEditProblem = () => {
    api.updateProblem(selectedProblem.value, selectedProblem.value.id).then(() => {
        problems.value = problems.value.map(p => p.id === selectedProblem.value.id ? selectedProblem.value : p);
        editProblemDialogVisible.value = false;
    });
}

const deleteProblem = (id: number) => {
    api.deleteProblem(id).then(() => {
        problems.value = problems.value.filter(p => p.id !== id);
    });
}

onMounted(async () => {
    problems.value = await getProblems();
});
</script>