<template>
  <div class="problem">
    <h1>问题</h1>
    <ProblemList style="width: 80vw"></ProblemList>
  </div>
</template>

<script>
import ProblemList from "@/components/ProblemList.vue";

export default {
  name: "ProblemsetView",
  components: {
    ProblemList,
  },
};
</script>
