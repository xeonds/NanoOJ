<template>
  <div id="problem">
    <h1>{{ problemTitle }}</h1>
    <p>{{ problemNote }}</p>
    <el-form>
      <el-form-item>
        <el-input
          type="textarea"
          v-model="code"
          :autosize="{ minRows: 10, maxRows: 20 }"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitCode">Commit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "ProblemView",
  components: {},
  computed: {
    problemData() {
      return this.$store.state.problemData;
    },
    problemTitle() {
      const id = this.$route.params.id;
      return this.problemData[id].title;
    },
    problemNote() {
      const id = this.$route.params.id;
      return this.problemData[id].note;
    },
  },
  setup() {
    const code = "";
    const submitCode = () => {
      axios
        .post("/api/v1/judge", { code: code.value })
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    };
    return {
      code,
      submitCode,
    };
  },
};
</script>

<style scoped>
</style>
