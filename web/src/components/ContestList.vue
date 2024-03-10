<template>
  <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.id}`) }">
    <el-table-column prop="id" label="ID"></el-table-column>
    <el-table-column prop="title" label="标题"></el-table-column>
    <el-table-column prop="start_time" label="开始时间"></el-table-column>
    <el-table-column prop="end_time" label="结束时间"></el-table-column>
  </el-table>
</template>

<script lang="ts" setup>
import { Contest } from '@/model';
import { getDataArr } from '@/utils/http';
const router = useRouter();
const { data: contests, get } = getDataArr<Contest>('/contests');

onMounted(async () => {
  contests.value = await get();
});
</script>
