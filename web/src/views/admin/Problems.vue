<template>
    <h2>{{ t('nav.problem') }}</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>{{ t('problem.all') }}</span>
                <span>
                    <el-button @click="refresh()" type="primary" text>{{ t('message.refresh') }}</el-button>
                    <el-button @click="handleImportProblems()" type="primary">{{ t('problem.import') }}</el-button>
                    <el-button @click="handleExportProblems()" type="primary">{{ t('problem.export') }}</el-button>
                    <el-button @click="createProblemDialogVisible = true" type="primary">{{ t('problem.create') }}</el-button>
                </span>
            </div>
        </template>
        <el-table :data="problems" style="width: 100%">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column prop="title" :label="t('message.title')"></el-table-column>
            <el-table-column :label="t('problem.difficulty')">
                <template #default="{ row }">
                    <el-rate v-model="row.difficulty" :colors="colors" disabled />
                </template>
            </el-table-column>
            <el-table-column prop="CreatedAt" :label="t('message.created-at')"></el-table-column>
            <el-table-column :label="t('message.action')">
                <template #default="{ row }">
                    <el-button @click="editProblem(row)">{{ t('message.edit') }}</el-button>
                    <el-button @click="deleteProblem(row.ID)">{{ t('message.delete') }}</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <el-dialog v-model="createProblemDialogVisible" :title="t('message.create')">
        <el-form :model="newProblem" label-width="100px">
            <el-form-item :label="t('message.title')">
                <el-input v-model="newProblem.title"></el-input>
            </el-form-item>
            <el-form-item :label="t('message.description')">
                <el-input v-model="newProblem.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item v-for="(_, index) in newProblem.inputs" :key="index" :label="`${t('problem.test-case')} ${index + 1}`">
                <el-row :gutter="10">
                    <el-col :span="20">
                        <el-form-item :label="t('problem.input')">
                            <el-input v-model="newProblem.inputs[index]" type="textarea" />
                        </el-form-item>
                        <el-form-item :label="t('problem.output')">
                            <el-input v-model="newProblem.outputs[index]" type="textarea" />
                        </el-form-item>
                        <el-form-item :label="t('problem.score')">
                            <el-input v-model.number="newProblem.ranks[index]" type="number" :min="0" :max="100" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="4">
                        <el-button @click="newProblem.inputs.splice(index, 1); newProblem.outputs.splice(index, 1)"
                            type="danger">{{ t('message.delete') }}</el-button>
                    </el-col>
                </el-row>
            </el-form-item>
            <el-button @click="newProblem.inputs.push(''); newProblem.outputs.push('')">{{ t('problem.add-test-case') }}</el-button>
            <el-form-item :label="t('problem.difficulty')">
                <el-rate v-model="newProblem.difficulty" :colors="colors" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="createProblemDialogVisible = false">{{ t('message.cancel') }}</el-button>
            <el-button type="primary" @click="createProblem">{{ t('message.create') }}</el-button>
        </div>
    </el-dialog>
    <el-dialog v-model="editProblemDialogVisible" :title="t('message.edit')">
        <el-form :model="selectedProblem" label-width="100px">
            <el-form-item :label="t('message.title')">
                <el-input v-model="selectedProblem.title"></el-input>
            </el-form-item>
            <el-form-item :label="t('message.description')">
                <el-input v-model="selectedProblem.description" type="textarea"></el-input>
            </el-form-item>
            <el-form-item v-for="(_, index) in selectedProblem.inputs" :key="index" :label="`${t('problem.test-case')} ${index + 1}`">
                <el-row :gutter="10">
                    <el-col :span="20">
                        <el-form-item :label="t('problem.input')">
                            <el-input v-model="selectedProblem.inputs[index]" type="textarea" />
                        </el-form-item>
                        <el-form-item :label="t('problem.output')">
                            <el-input v-model="selectedProblem.outputs[index]" type="textarea" />
                        </el-form-item>
                        <el-form-item :label="t('problem.score')">
                            <el-input v-model.number="selectedProblem.ranks[index]" type="number" :min="0" :max="100" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="4">
                        <el-button @click="selectedProblem.inputs.splice(index, 1); selectedProblem.outputs.splice(index, 1)"
                            type="danger">{{ t('message.delete') }}</el-button>
                    </el-col>
                </el-row>
            </el-form-item>
            <el-button @click="selectedProblem.inputs.push(''); selectedProblem.outputs.push('')">{{ t('problem.add-test-case') }}</el-button>
            <el-form-item :label="t('problem.difficulty')">
                <el-rate v-model="selectedProblem.difficulty" :colors="colors" />
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="editProblemDialogVisible = false">{{ t('message.cancel') }}</el-button>
            <el-button type="primary" @click="handleEditProblem">{{ t('message.update') }}</el-button>
        </div>
    </el-dialog>
</template>

<script lang="ts" setup>
import { Problem } from '@/model';
import api from '@/api';
import { getDataArr, http } from '@/utils/http';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const createProblemDialogVisible = ref(false);
const editProblemDialogVisible = ref(false);
const { data: problems, get: getProblems } = getDataArr<Problem>('/problems');
const colors = { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' }

const newProblem: Ref<Problem> = ref({
    title: '',
    description: '',
    inputs: [''],
    outputs: [''],
    ranks: [100],
    difficulty: 2
} as Problem);
const createProblem = () => {
    api.addProblems(newProblem.value as Problem).then(() => {
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
    api.updateProblem(selectedProblem.value, selectedProblem.value.ID).then(() => {
        problems.value = problems.value.map(p => p.ID === selectedProblem.value.ID ? selectedProblem.value : p);
        editProblemDialogVisible.value = false;
    });
}

const deleteProblem = (id: number) => {
    api.deleteProblem(id).then(() => {
        problems.value = problems.value.filter(p => p.ID !== id);
    });
}

const refresh = async () => {
    problems.value = await getProblems();
}

const handleImportProblems = async () => {
    try {
        const fileInput = document.createElement('input');
        fileInput.type = 'file';
        fileInput.accept = 'application/json';

        fileInput.addEventListener('change', async (event) => {
            const file = (event.target as HTMLInputElement)?.files?.[0];
            if (file) {
                const formData = new FormData();
                formData.append('file', file);
                const response = await http.post('/admin/problems/import', formData);
                ElMessage({ message: response.data.value, type: "info" })
            }
        });
        fileInput.click();
    } catch (error) {
        ElMessage({ message: t('message.import-problem-fail') + error, type: "error" })
    }
}

const handleExportProblems = async () => {
    http.get('/admin/problems/export')
        .then((res) => {
            if (res.err.value != null) {
                ElMessage({ message: t('message.export-problem-fail') + res.err.value, type: "error" })
            } else {
                const downloadLink = window.document.createElement('a')
                downloadLink.href = window.URL.createObjectURL(
                    new Blob([res.data.value], { type: 'application/json' })
                )
                downloadLink.download = 'export.json'
                document.body.appendChild(downloadLink)
                downloadLink.click()
                document.body.removeChild(downloadLink)
                ElMessage({ message: t('message.export-problem-success'), type: "info" })
            }
        })
}

onMounted(async () => {
    problems.value = await getProblems();
});
</script>