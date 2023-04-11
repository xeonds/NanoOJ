<template>
  <div id="contest">
    <h1>测试</h1>
    <ContestList :contests="contests"></ContestList>
  </div>
</template>

<script>
import ContestList from "../components/ContestList.vue";
import { mapGetters } from "vuex";

export default {
  name: "ContestView",
  components: {
    ContestList,
  },
  computed: {
    ...mapGetters({
      contests: "getContests",
    }),
  },
};
</script>

<style scoped></style>