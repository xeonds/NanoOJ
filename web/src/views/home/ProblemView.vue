<template>
  <div id="problem">
    <el-row :gutter="20">
      <el-col :span="16">
        <el-card>
          <template #header>
            <el-text type="primary" size="large">#{{ problem.id }}.{{ problem.title }}</el-text>
          </template>
          <h4 type="primary" size="large">Problem Description</h4>
          <p v-html="description"></p><br>
          <h4 type="primary" size="large">Test Cases</h4>
          <template v-for="(_, index) in problem.inputs" :key="index">
            <el-row>
              <el-col :span="12">
                <h5>Input:</h5>
                <pre>{{ problem.inputs[index] }}</pre>
              </el-col>
              <el-col :span="12">
                <h5>Output:</h5>
                <pre>{{ problem.outputs[index] }}</pre>
              </el-col>
            </el-row>
          </template>
          <el-divider>Code here</el-divider>
          <CodeEditor :language="language" :height="'24rem'" id="editor">
            <template #editor-options>
              <div id="buttons">
                <el-form-item label="language">
                  <el-select v-model="language" placeholder="Select Language" style="width: 100px">
                    <el-option label="C" value="c"></el-option>
                    <el-option label="C++" value="cpp"></el-option>
                    <el-option label="Java" value="java"></el-option>
                    <el-option label="Python" value="python"></el-option>
                  </el-select>
                </el-form-item>
                <el-button type="primary" @click="submitCode">Commit</el-button>
              </div>
            </template>
          </CodeEditor>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <el-text type="primary">Problem Info</el-text>
          </template>
          <el-descriptions direction="vertical" :column="2">
            <el-descriptions-item label="Time Limit">114</el-descriptions-item>
            <el-descriptions-item label="Memory Limit">514</el-descriptions-item>
            <el-descriptions-item label="Test Cases">1919</el-descriptions-item>
            <el-descriptions-item label="Submission">8/10</el-descriptions-item>
            <el-descriptions-item label="Pass Rate" :span="2">
              <el-progress :percentage="42"></el-progress>
            </el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { marked } from "marked";
import { onMounted } from "vue";
import CodeEditor from '@/components/CodeEditor.vue';
import { getData, http } from "@/utils/http";
import { Problem } from "@/model";

const route = useRoute();
const router = useRouter();
const id = parseInt(route.params.id as string);
const language = ref('c');
const code = ref('');
const { data: problem, get: getProblemInfo } = getData<Problem>(`/problems/${id}`);
const description = computed(() => marked(problem.value.description ?? 'loading...'));

const submitCode = async () => {
  const { err } = await http.post('/actions/submit', { code: code.value, language: language.value, state: 'waiting', problem_id: id })
  if (err.value !== null) {
    console.error(err);
  } else {
    ElMessage({ message: 'Code submitted successfully', type: 'success' });
    router.push('/status');
  }
}

onMounted(async () => {
  problem.value = await getProblemInfo();
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
