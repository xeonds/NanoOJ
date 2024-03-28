<template>
  <div id="problem">
    <el-row :gutter="20" class="hidden-sm-and-down">
      <el-col :span="16">
        <el-card>
          <template #header>
            <el-text type="primary" size="large">#{{ problem.ID }}.{{ problem.title }}</el-text>
          </template>
          <h4 type="primary" size="large">{{ t('problem.description') }}</h4>
          <p v-html="description"></p><br>
          <h4 type="primary" size="large">{{ t('problem.demo-test-case') }}</h4>
          <template v-for="(_, index) in problem.inputs" :key="index">
            <el-row>
              <el-col :span="12">
                <h5>{{ t('problem.input') }}</h5>
                <pre>{{ problem.inputs[index] }}</pre>
              </el-col>
              <el-col :span="12">
                <h5>{{ t('problem.output') }}</h5>
                <pre>{{ problem.outputs[index] }}</pre>
              </el-col>
            </el-row>
          </template>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <el-text type="primary">{{ t('problem.info') }}</el-text>
          </template>
          <el-descriptions direction="vertical" :column="2">
            <el-descriptions-item :label="t('problem.time-limit')">{{ problem.time_limit }}</el-descriptions-item>
            <el-descriptions-item :label="t('problem.mem-limit')">514</el-descriptions-item>
            <el-descriptions-item :label="t('problem.test-case')">{{ problem.inputs?.length }}</el-descriptions-item>
            <el-descriptions-item :label="t('problem.submit-status')">{{ commits.length }}</el-descriptions-item>
            <el-descriptions-item :label="t('problem.pass-rate')" :span="2">
              <el-progress :percentage="42"></el-progress>
            </el-descriptions-item>
          </el-descriptions>
        </el-card>
        <el-card>
          <template #header>
            <span class="card-header">
              <el-text type="primary">{{ t('problem.commit-history') }}</el-text>
              <el-button type="text" @click="getCommits()">{{ t('message.refresh') }}</el-button>
            </span>
          </template>
          <el-table :data="commits" @row-click="showInfo">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column :label="t('problem.submit-status')">
              <template #default="{ row }">
                <el-tag class="ml-2" :type="statusTag(row.status)">{{ row.status }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="language" :label="t('message.language')"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row class="hidden-md-and-up">
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-text type="primary" size="large">#{{ problem.ID }}.{{ problem.title }}</el-text>
          </template>
          <h4 type="primary" size="large">{{ t('problem.description') }}</h4>
          <p v-html="description"></p><br>
          <h4 type="primary" size="large">{{ t('problem.demo-test-case') }}</h4>
          <template v-for="(_, index) in problem.inputs" :key="index">
            <el-row>
              <el-col :span="12">
                <h5>{{ t('problem.input') }}</h5>
                <pre>{{ problem.inputs[index] }}</pre>
              </el-col>
              <el-col :span="12">
                <h5>{{ t('problem.output') }}</h5>
                <pre>{{ problem.outputs[index] }}</pre>
              </el-col>
            </el-row>
          </template>
        </el-card>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <CodeEditor :language="language" :height="'24rem'" id="editor">
            <template #editor-options>
              <div id="buttons">
                <el-form-item :label="t('message.language')">
                  <el-select v-model="language" :placeholder="t('message.select-lang')" style="width: 100px">
                    <el-option label="C" value="c"></el-option>
                    <el-option label="C++" value="cpp"></el-option>
                    <el-option label="Java" value="java"></el-option>
                    <el-option label="Python" value="python"></el-option>
                  </el-select>
                </el-form-item>
                <el-button type="primary" @click="submitCode">{{ t('message.submit') }}</el-button>
              </div>
            </template>
          </CodeEditor>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="hidden-md-and-up">
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-text type="primary">{{ t('problem.info') }}</el-text>
          </template>
          <el-descriptions direction="vertical" :column="4">
            <el-descriptions-item :label="t('problem.time-limit')">{{ problem.time_limit }}</el-descriptions-item>
            <el-descriptions-item :label="t('problem.mem-limit')">514</el-descriptions-item>
            <el-descriptions-item :label="t('problem.test-case')">{{ problem.inputs?.length }}</el-descriptions-item>
            <el-descriptions-item :label="t('problem.submit-status')">{{ commits.length }}</el-descriptions-item>
            <el-descriptions-item :label="t('problem.pass-rate')" :span="2">
              <el-progress :percentage="42"></el-progress>
            </el-descriptions-item>
          </el-descriptions>
        </el-card>
        <el-card>
          <template #header>
            <span class="card-header">
              <el-text type="primary">{{ t('problem.commit-history') }}</el-text>
              <el-button type="text" @click="getCommits()">{{ t('message.refresh') }}</el-button>
            </span>
          </template>
          <el-table :data="commits" @row-click="showInfo">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column :label="t('problem.submit-status')">
              <template #default="{ row }">
                <el-tag class="ml-2" :type="statusTag(row.status)">{{ row.status }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="language" :label="t('message.language')"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog v-model="dialogVisible" :title="t('submission.info')" width="30%">
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
        <el-descriptions-item :label="t('submission.code')"><code>{{ dialogData.code }}</code></el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { marked } from "marked";
import { onMounted } from "vue";
import CodeEditor from '@/components/CodeEditor.vue';
import { applyData, getData, getDataArr, handleHttp } from "@/utils/http";
import { Problem, Submission } from "@/model";
import api from "@/api";
import { EpPropMergeType } from "element-plus/es/utils/index.mjs";
import { useI18n } from "vue-i18n";

type _EpPropMergeType = EpPropMergeType<StringConstructor, "success" | "warning" | "info" | "primary" | "danger", unknown>;

const { t } = useI18n();
const route = useRoute();
const id = parseInt(route.params.id as string);
const language = ref('c');
const code = ref('');
const { data: problem, get: getProblemInfo } = getData<Problem>(`/problems/${id}`);
const description = computed(() => marked(problem.value.description ?? 'loading...'));
const { data: commits, get: _getCommits } = getDataArr<Submission>(`/problems/${id}/submissions`);
const dialogVisible = ref(false);
const dialogData = ref({} as Submission);

const statusTag = (status: string): _EpPropMergeType => { return { "Pending": "info", "InProgress": "info", "Accepted": "success", "WrongAnswer": "danger", "TimeLimitExceeded": "warning", "MemoryLimitExceeded": "warning", "RuntimeError": "danger", "CompilationError": "danger" }[status] as _EpPropMergeType; }
const showInfo = (row: Submission) => {
  dialogData.value = row;
  dialogVisible.value = true;
};
const submitCode = async () => {
  handleHttp(await api.addSubmission({ code: code.value, language: language.value, status: 'Pending', problem_id: id } as Submission),
    () => ElMessage({ message: 'Code submitted successfully', type: 'success' }),
    (err: any) => ElMessage({ message: 'Code submission failed: ' + err, type: 'error' }));
}
const getCommits = async () => applyData(commits, _getCommits, []);

onMounted(async () => {
  problem.value = await getProblemInfo();
  commits.value = await _getCommits() ?? [];
});

provide('code', code)
</script>

<style scoped>
#buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;

}

#buttons>* {
  margin: 0 20px 0 0;
}
</style>
