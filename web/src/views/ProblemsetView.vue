<template>
  <div class="problem">
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
