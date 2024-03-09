<template>
  <div id="contest">
    <el-card>
      <template #header>
        <el-text type="primary">#{{ contest.id }}.{{ contest.title }}</el-text>
      </template>
      <h4>Time</h4>
      <el-row>
        <el-col :span="8">
          <el-countdown title="Start Time" format="YYYY.MM.DD HH:mm:ss"
            :value="(new Date(contest.start_time)).getTime()" />
        </el-col>
        <el-col :span="8">
          <el-countdown title="End Time" format="YYYY.MM.DD HH:mm:ss" :value="(new Date(contest.end_time)).getTime()" />
        </el-col>
        <el-col :span="8">
          <el-countdown title="Countdown" format="DD [days] HH:mm:ss"
            :value="(new Date(contest.end_time)).getTime() - (new Date(contest.start_time)).getTime()" />
        </el-col>
      </el-row>
      <h4>Description</h4>
      <p v-html="description"></p>
      <el-divider>OwO</el-divider>
      <h4>Problems</h4>
      <el-table :data="contest.problems">
        <el-table-column prop="index" label="ID"></el-table-column>
        <el-table-column prop="title" label="Title"></el-table-column>
        <el-table-column prop="score" label="Result"></el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import { marked } from "marked";

export default {
  name: "ContestView",
};
</script>

<style scoped></style>
