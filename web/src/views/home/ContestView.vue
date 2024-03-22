<template>
  <div id="contest">
    <el-row :gutter="20" class="hidden-sm-and-down">
      <el-col :span="16">
        <el-card>
          <template #header>
            <el-text type="primary">#{{ contest.ID }}.{{ contest.title }}</el-text>
          </template>
          <h4>时间</h4>
          <el-row>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">Start Time</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.start_time) }}</div>
              </div>
            </el-col>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">End Time</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.end_time) }}</div>
              </div>
            </el-col>
            <el-col :span="8">
              <el-countdown title="Countdown" format="DD [days] HH:mm:ss" :value="time.fromString(contest.end_time)" />
            </el-col>
          </el-row>
          <h4>描述</h4>
          <p v-html="description"></p>
        </el-card>
        <el-card>
          <template #header>
            <el-text type="primary">#ProblemList</el-text>
          </template>
          <el-table :data="contest.problems" @row-click="(row) => { router.push(`/problem/${row.ID}`) }">
            <el-table-column prop="index" label="ID"></el-table-column>
            <el-table-column prop="title" label="Title"></el-table-column>
            <el-table-column prop="score" label="Result"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <el-text type="primary">#RankList</el-text>
          </template>
          <el-table :data="rank">
            <el-table-column prop="user_id" label="User ID"></el-table-column>
            <el-table-column prop="rank" label="Rank"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row class="hidden-md-and-up">
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-text type="primary">#{{ contest.ID }}.{{ contest.title }}</el-text>
          </template>
          <h4>时间</h4>
          <el-row>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">Start Time</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.start_time) }}</div>
              </div>
            </el-col>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">End Time</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.end_time) }}</div>
              </div>
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
        </el-card>
        <el-card style="">
          <template #header>
            <el-text type="primary">#RankList</el-text>
          </template>
          <el-table :data="rank">
            <el-table-column prop="user_id" label="User ID"></el-table-column>
            <el-table-column prop="rank" label="Rank"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
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