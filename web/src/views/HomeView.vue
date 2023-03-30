<template>
  <div id="home">
    <el-row>
      <el-col :span="24">
        <el-card>
          <NotificationBoard :notifications="notifications"></NotificationBoard>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="16">
        <el-card>
          <h3>最新问题</h3>
          <ProblemList :problems="problems"></ProblemList>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <h3>排行榜</h3>
          <RankList :ranks="ranks"></RankList>
        </el-card>
        <el-card>
          <h3>近期测试</h3>
          <ContestList :contests="contests"></ContestList>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import NotificationBoard from "../components/NotificationBoard.vue";
import ProblemList from "../components/ProblemList.vue";
import RankList from "../components/RankList.vue";
import ContestList from "../components/ContestList.vue";

export default {
  name: "HomeView",
  components: {
    NotificationBoard,
    ProblemList,
    RankList,
    ContestList,
  },
  computed: {
    ...mapGetters({
      problems: "getProblems",
      notifications: "getNotifications",
      ranks: "getRanks",
      contests: "getContests",
    }),
  },
};
</script>

<style scoped>
.el-card {
  border-radius: 4px;
  margin-top: 20px;
  margin-bottom: 20px;
}

el-card {
  width: 100%;
}
</style>
