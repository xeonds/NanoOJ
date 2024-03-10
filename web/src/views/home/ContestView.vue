<template>
  <div id="contest">
    <el-card>
      <template #header>
        <el-text type="primary">#{{ props.contest.id }}.{{ props.contest.title }}</el-text>
      </template>
      <h4>时间</h4>
      <el-row>
        <el-col :span="8">
          <el-countdown title="Start Time" format="YYYY.MM.DD HH:mm:ss"
            :value="time.fromString(props.contest.start_time)" />
        </el-col>
        <el-col :span="8">
          <el-countdown title="End Time" format="YYYY.MM.DD HH:mm:ss"
            :value="time.fromString(props.contest.end_time)" />
        </el-col>
        <el-col :span="8">
          <el-countdown title="Countdown" format="DD [days] HH:mm:ss"
            :value="time.fromString(props.contest.end_time) - time.fromString(props.contest.start_time)" />
        </el-col>
      </el-row>
      <h4>描述</h4>
      <p v-html="description"></p>
      <el-divider>OwO</el-divider>
      <h4>问题列表</h4>
      <el-table :data="props.contest.problems">
        <el-table-column prop="index" label="ID"></el-table-column>
        <el-table-column prop="title" label="Title"></el-table-column>
        <el-table-column prop="score" label="Result"></el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { marked } from "marked";
import { time } from "@/utils/datetime";
const props = defineProps<{ contest: Contest }>();
</script>
