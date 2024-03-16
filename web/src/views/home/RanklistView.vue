<template>
  <div>
    <h2>排行榜</h2>
    <el-table :data="ranks">
      <el-table-column prop="index" label="Rank"></el-table-column>
      <el-table-column prop="username" label="用户名"></el-table-column>
      <el-table-column prop="profile" label="简介"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { Rank } from '@/model';
import { getDataArr } from '@/utils/http';
const { data: ranks, get } = getDataArr<Rank>('/ranks');
onMounted(async () => {
  ranks.value = await get();
});
</script>
