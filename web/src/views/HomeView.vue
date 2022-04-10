<template>
  <div class="home">
    <el-space wrap size="large" style="margin:16px;">
        <NotificationBoard style="width:48vw;"></NotificationBoard>
        <RankList style="width:20vw;"></RankList>
        <ProblemList style="width:48vw;"></ProblemList>
        <SearchBox style="width:20vw;"></SearchBox>
        <RecentContest style="width:20vw;"></RecentContest>
    </el-space>
  </div>
</template>

<script>
import NotificationBoard from "@/components/NotificationBoard.vue";
import ProblemList from "@/components/ProblemList.vue";
import RankList from "@/components/RankList.vue";
import SearchBox from "@/components/SearchBox.vue";
import RecentContest from "@/components/RecentContest.vue";

export default {
  name: "HomeView",
  components: {
    NotificationBoard,
    ProblemList,
    RankList,
    SearchBox,
    RecentContest,
  },
};
</script>

<style scoped>
.home {
  width: 100%;
  max-width: 1024px;
}
</style>