<template>
  <div id="home">
    <el-row>
      <el-col :span="24">
        <el-card>
          <NotificationBoard></NotificationBoard>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="16">
        <el-card>
          <h3>最新问题</h3>
          <el-table :data="problems" @row-click="(row) => { router.push(`/problem/${row.id}`) }">
            <el-table-column prop="id" label="问题ID"></el-table-column>
            <el-table-column prop="title" label="问题标题"></el-table-column>
            <el-table-column label="题目难度">
              <template v-slot="{ row }">
                <el-rate v-model="row.difficulty" :colors="colors" disabled />
              </template>
            </el-table-column>
            <el-table-column prop="passRate" label="通过率"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <h3>排行榜</h3>
          <el-table :data="ranks">
            <el-table-column prop="index" label="Rank"></el-table-column>
            <el-table-column prop="username" label="用户名"></el-table-column>
            <el-table-column prop="profile" label="简介"></el-table-column>
          </el-table>
        </el-card>
        <el-card>
          <h3>近期测试</h3>
          <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.id}`) }">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="title" label="标题"></el-table-column>
            <el-table-column prop="start_time" label="开始时间"></el-table-column>
            <el-table-column prop="end_time" label="结束时间"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import NotificationBoard from "@/components/NotificationBoard.vue";
const router = useRouter();
const colors = { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' };
import { Contest, Problem, Rank } from '@/model';
import { getDataArr } from '@/utils/http';

const { data: problems, get: getProblems } = getDataArr<Problem>('/problems');
const { data: ranks, get: getRanks } = getDataArr<Rank>('/ranks');
const { data: contests, get: getContests } = getDataArr<Contest>('/contests');

onMounted(async () => {
  problems.value = await getProblems();
  ranks.value = await getRanks();
  contests.value = await getContests();
});
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
