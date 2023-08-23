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
          <el-divider>Code here</el-divider>
          <CodeEditor language="c" :callback="submitCode">
            <template #editor-options>
              <div id="buttons">
                <el-select v-model="problem.language" placeholder="Select Language">
                  <el-option label="C" value="c"></el-option>
                  <el-option label="C++" value="cpp"></el-option>
                  <el-option label="Java" value="java"></el-option>
                  <el-option label="Python" value="python"></el-option>
                </el-select>
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

<script>
import { marked } from "marked";
import CodeEditor from '../../components/CodeEditor.vue';
import api from '../../api';

export default {
  name: "ProblemView",
  components: {
    CodeEditor,
  },
  data() {
    return {
      problem: { description: '' },
    }
  },
  created() {
    api.getProblemInfo(this.$route.params.id).then((response) => {
      if (response.status === 200) {
        this.problem = response.data;
      }
    }).catch((error) => {
      this.$message({ type: 'error', message: 'Failed to get problem info: ' + error });
    });
  },
  computed: {
    description() {
      return marked(this.problem.description);
    },
  },
  methods: {
    submitCode: function (language, code) {
      api.addSubmissions({ code: code, language: 'c', state: 'waiting', problem_id: parseInt(this.$route.params.id) })
        .then((response) => {
          if (response.status === 200) {
            this.$message({
              message: 'Code submitted successfully',
              type: 'success'
            });
            this.$router.push('/status');
          }
        })
        .catch((error) => {
        });
    },
  }
};
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
