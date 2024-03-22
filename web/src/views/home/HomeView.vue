<template>
  <div id="home">
    <h2>主页</h2>
    <el-row>
      <el-col :span="24">
        <el-card>
          <div v-if="notifications.length === 0">暂无公告</div>
          <el-carousel v-else trigger="click" :interval="5000" indicator-position="outside" arrow="never" height="auto">
            <el-carousel-item style="height: auto" v-for="item in notifications" :key="item.ID">
              <h3>{{ item.title }} </h3>
              <p v-html="item.content"></p>
              <el-divider></el-divider>
              <el-tag>创建于：{{ time.formatDate(item.CreatedAt) }}</el-tag> |
              <el-tag v-if="item.UpdatedAt">修改于：{{ time.formatDate(item.UpdatedAt) }}</el-tag>
            </el-carousel-item>
          </el-carousel>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="hidden-sm-and-down">
      <el-col :span="16">
        <el-card>
          <h3>最新问题</h3>
          <el-table :data="problems" @row-click="(row) => { router.push(`/problem/${row.ID}`) }">
            <el-table-column prop="ID" label="问题ID"></el-table-column>
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
          <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.ID}`) }">
            <el-table-column prop="ID" label="ID"></el-table-column>
            <el-table-column prop="title" label="标题"></el-table-column>
            <el-table-column prop="start_time" label="开始时间"></el-table-column>
            <el-table-column prop="end_time" label="结束时间"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="hidden-md-and-up">
      <el-col :span="24">
        <el-card>
          <h3>最新问题</h3>
          <el-table :data="problems" @row-click="(row) => { router.push(`/problem/${row.ID}`) }">
            <el-table-column prop="ID" label="问题ID"></el-table-column>
            <el-table-column prop="title" label="问题标题"></el-table-column>
            <el-table-column label="题目难度">
              <template v-slot="{ row }">
                <el-rate v-model="row.difficulty" :colors="colors" disabled />
              </template>
            </el-table-column>
            <el-table-column prop="passRate" label="通过率"></el-table-column>
          </el-table>
        </el-card>
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
          <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.ID}`) }">
            <el-table-column prop="ID" label="ID"></el-table-column>
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
import { Notification, Contest, Problem, Rank } from '@/model';
import { marked } from "marked";
import { getDataArr } from '@/utils/http';
import { time } from "@/utils/datetime";

const router = useRouter();
const colors = { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' };
const { data: notifications, get } = getDataArr<Notification>('/notifications');
const { data: problems, get: getProblems } = getDataArr<Problem>('/problems');
const { data: ranks, get: getRanks } = getDataArr<Rank>('/ranks');
const { data: contests, get: getContests } = getDataArr<Contest>('/contests');

onMounted(async () => {
  notifications.value = (await get()).map((item) => {
    item.content = marked(item.content) as string;
    return item;
  });
  problems.value = await getProblems();
  ranks.value = await getRanks();
  contests.value = await getContests();
});
</script>

<style scoped>
.notification-time {
  text-align: right;
  font-size: 12px;
  color: #999;
}
</style>
