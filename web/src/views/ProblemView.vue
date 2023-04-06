<template>
  <div id="problem">
    <h1>{{ problem.title }}</h1>
    <p v-html="description"></p>
    <el-form>
      <el-form-item label="Code">
      </el-form-item>
      <codemirror
        v-model="code"
        placeholder="Code goes here..."
        :style="{ height: '400px' }"
        :indent-with-tab="true"
        :tab-size="4"
        :extensions="extensions"
      /><br />
      <el-form-item>
        <el-button type="primary" @click="submitCode">Commit</el-button>
      </el-form-item>
    </el-form>
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
    ]
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
