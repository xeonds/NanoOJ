<template>
  <div class="problem">
    <h2>问题</h2>
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
  </div>
</template>

<script lang="ts" setup>
import { Problem } from '@/model';
import { getDataArr } from '@/utils/http';

const router = useRouter();
const colors = { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' };
const { data: problems, get } = getDataArr<Problem>('/problems');

onMounted(async () => {
  problems.value = await get();
});
</script>
