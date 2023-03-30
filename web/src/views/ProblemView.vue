<template>
  <div id="problem">
    <h1>{{ problem.title }}</h1>
    <p>{{ problem.description }}</p>
    <el-form>
      <el-form-item>
        <el-input type="textarea" v-model="code" :autosize="{ minRows: 10, maxRows: 20 }"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitCode">Commit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import api from '../api';

export default {
  name: "ProblemView",
  data: function () {
    return {
      code: "",
    };
  },
  computed: {
    ...mapGetters({
      _problem: "getProblemById",
    }),
    problem() {
      return this._problem(this.$route.params.id)
    },
  },
  methods: {
    submitCode: function () {
      api.addSubmissions({ code: this.code, language: 'c', state: 'waiting', problem_id: parseInt(this.$route.params.id) })
        .then((response) => {
          if (response.status === 200) {
            this.$router.push({ name: 'SubmissionView', params: { id: response.data.id } });
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>

<style scoped></style>
