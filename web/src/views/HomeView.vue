<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to NanoOJ" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";
import NotificationBoard from "@/components/NotificationBoard.vue";
import ProblemList from "@/components/ProblemList.vue";
import RankList from "@/components/RankList.vue";

export default {
  name: "HomeView",
  components: {
    HelloWorld,
    NotificationBoard,
    ProblemList,
    RankList,
  },
};
</script>
