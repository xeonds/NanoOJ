<template>
  <div class="problem">
    <h1>问题</h1>
    <ProblemList :problems="problems"></ProblemList>
  </div>
</template>

<script>
import ProblemList from "../components/ProblemList.vue";
import { mapGetters } from "vuex";

export default {
  name: "ProblemsetView",
  computed: {
    ...mapGetters({ problems: "getProblems" }),
  },
  components: {
    ProblemList,
  },
};
</script>
