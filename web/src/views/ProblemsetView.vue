<template>
  <div class="problem">
    <h1>问题</h1>
    <ProblemList :problems="problem"></ProblemList>
  </div>
</template>

<script>
import ProblemList from "@/components/ProblemList.vue";
import { mapGetters } from "vuex";

export default {
  name: "ProblemsetView",
  computed: {
    ...mapGetters("problem", ["getProblemById"]),
    problem() {
      return this.getProblemById(this.$route.params.id);
    },
  },
  components: {
    ProblemList,
  },
};
</script>
