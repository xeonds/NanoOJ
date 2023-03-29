<template>
  <div id="contest">
    <h1>测试</h1>
    <ContestList></ContestList>
  </div>
</template>

<script>
import ContestList from "../components/ContestList.vue";

export default {
  name: "ContestView",
  components: {
    ContestList,
  },
};
</script>

<style scoped>
</style>