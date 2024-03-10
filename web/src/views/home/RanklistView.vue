<template>
  <div>
    <h1>排行榜</h1>
    <el-table :data="ranks">
      <el-table-column prop="index" label="Rank"></el-table-column>
      <el-table-column prop="username" label="用户名"></el-table-column>
      <el-table-column prop="profile" label="简介"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { getDataArr } from '@/utils/http';
const { data: ranks, get } = getDataArr('/ranks');
onMounted(async () => {
  ranks.value = await get();
});
</script>
