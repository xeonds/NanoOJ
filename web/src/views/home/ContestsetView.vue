<template>
  <div id="contests">
    <h2>测试</h2>
    <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.ID}`) }">
      <el-table-column prop="ID" label="ID"></el-table-column>
      <el-table-column prop="title" label="标题"></el-table-column>
      <el-table-column prop="start_time" label="开始时间"></el-table-column>
      <el-table-column prop="end_time" label="结束时间"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { Contest } from '@/model';
import { getDataArr } from '@/utils/http';
import { time } from '@/utils/datetime';

const router = useRouter();
const { data: contests, get } = getDataArr<Contest>('/contests');

onMounted(async () => {
  contests.value = (await get()).map((contest) => {
    contest.start_time = time.formatDate(contest.start_time);
    contest.end_time = time.formatDate(contest.end_time);
    return contest;
  });
});
</script>
