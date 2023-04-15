<template>
  <div id="problem">
    <el-card>
      <template #header>
        <el-text type="primary" size="large">#{{ problem.id }}.{{ problem.title }}</el-text>
      </template>
      <el-descriptions title="Problem Info">
        <el-descriptions-item label="Time Limitation">114</el-descriptions-item>
        <el-descriptions-item label="Memory Limitation">514</el-descriptions-item>
        <el-descriptions-item label="Test Cases">1919</el-descriptions-item>
        <el-descriptions-item label="Submission">8/10</el-descriptions-item>
        <el-descriptions-item label="Pass Rate">
          <el-progress :percentage="42"></el-progress>
        </el-descriptions-item>
      </el-descriptions>
      <h4 type="primary" size="large">Problem Description</h4>
      <p v-html="description"></p><br>
      <el-divider>OwO</el-divider>
      <h4>Code</h4>
      <el-row>
        <el-col :span="8">
          <el-select v-model="language" placeholder="Select Language">
            <el-option label="C" value="c"></el-option>
            <el-option label="C++" value="cpp"></el-option>
            <el-option label="Java" value="java"></el-option>
            <el-option label="Python" value="python"></el-option>
          </el-select>
        </el-col>
      </el-row><br>
      <codemirror v-model="code" placeholder="Submit your solution..." :style="{ height: '400px' }"
        :indent-with-tab="true" :tab-size="4" :extensions="extensions" /><br />
      <el-form-item>
        <el-button type="primary" @click="submitCode">Commit</el-button>
      </el-form-item>
    </el-card>
  </div>
</template>

<script>
import { Codemirror } from 'vue-codemirror'
import { javascript } from '@codemirror/lang-javascript'
import { oneDark } from '@codemirror/theme-one-dark'
import { mapGetters } from "vuex";
import { marked } from "marked";
import api from '../api';

export default {
  name: "ProblemView",
  components: {
    Codemirror
  },
  data: () => ({
    code: '',
    extensions: [
      javascript(),
      oneDark
    ],
    language: 'c'
  }),
  computed: {
    ...mapGetters({
      _problem: "getProblemById",
    }),
    problem() {
      return this._problem(this.$route.params.id);
    },
    description() {
      return marked(this.problem.description);
    },
  },
  methods: {
    submitCode: function () {
      api.addSubmissions({ code: this.code, language: 'c', state: 'waiting', problem_id: parseInt(this.$route.params.id) })
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
  },
};
</script>

<style scoped></style>
