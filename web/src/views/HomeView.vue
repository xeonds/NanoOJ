<template>
  <div class="home">
    <el-row>
      <el-col :span="16">
        <el-card style="width: 100%">
          <h3>通知</h3>
          <NotificationBoard></NotificationBoard>
        </el-card>
        <el-card style="width: 100%">
          <h3>最新问题</h3>
          <ProblemList></ProblemList>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card style="width: 100%">
          <h3>排行榜</h3>
          <RankList></RankList>
        </el-card>
        <el-card style="width: 100%">
          <h3>搜索</h3>
          <SearchBox></SearchBox>
        </el-card>
        <el-card style="width: 100%">
          <h3>近期测试</h3>
          <ContestList></ContestList>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import NotificationBoard from "@/components/NotificationBoard.vue";
import ProblemList from "@/components/ProblemList.vue";
import RankList from "@/components/RankList.vue";
import SearchBox from "@/components/SearchBox.vue";
import ContestList from "@/components/ContestList.vue";

export default {
  name: "HomeView",
  components: {
    NotificationBoard,
    ProblemList,
    RankList,
    SearchBox,
    ContestList,
  },
};
</script>

<style scoped>
.home {
  width: 100%;
  max-width: 1024px;
}
</style>