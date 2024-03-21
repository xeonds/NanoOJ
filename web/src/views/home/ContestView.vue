<template>
  <div id="contest">
    <el-card>
      <template #header>
        <el-text type="primary">#{{ contest.ID }}.{{ contest.title }}</el-text>
      </template>
      <h4>时间</h4>
      <el-row>
        <el-col :span="8">
          <h2>Start Time</h2>
          <span>{{ time.formatDate(contest.start_time) }}</span>
        </el-col>
        <el-col :span="8">
          <h2>End Time</h2>
          <span>{{ time.formatDate(contest.end_time) }}</span>
        </el-col>
        <el-col :span="8">
          <el-countdown title="Countdown" format="DD [days] HH:mm:ss" :value="time.fromString(contest.end_time)" />
        </el-col>
      </el-row>
      <h4>描述</h4>
      <p v-html="description"></p>
      <el-divider>OwO</el-divider>
      <h4>问题列表</h4>
      <el-table :data="contest.problems" @row-click="(row) => { router.push(`/problem/${row.ID}`) }">
        <el-table-column prop="index" label="ID"></el-table-column>
        <el-table-column prop="title" label="Title"></el-table-column>
        <el-table-column prop="score" label="Result"></el-table-column>
      </el-table>
      <h4>排名</h4>
      <el-table :data="rank">
        <el-table-column prop="user_id" label="User ID"></el-table-column>
        <el-table-column prop="rank" label="Rank"></el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { marked } from "marked";
import { time } from "@/utils/datetime";
import { Contest, Rank } from "@/model";
import { getData, getDataArr } from "@/utils/http";

const route = useRoute();
const router = useRouter();
const id = Number(route.params.id);
const description = computed(() => marked(contest.value.description ?? 'loading...'));
const { data: contest, get: getContest } = getData<Contest>(`/contests/${id}`);
const { data: rank, get: getRank } = getDataArr<Rank>(`/contests/${id}/rank`);

onMounted(async () => {
  contest.value = await getContest();
  rank.value = await getRank();
});
</script>
