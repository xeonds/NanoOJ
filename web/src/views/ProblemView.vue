<template>
  <div id="problem">
    <h1>{{ problem.ProblemTitle }}</h1>
    <p>{{ problem.ProblemDescription }}</p>
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

export default {
  name: "ProblemView",
  components: {},
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
      this.axios
        .post("/submissions", { code: this.code })
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>

<style scoped></style>
