<template>
  <div class="problem">
    <h2>{{ t('nav.problems') }}</h2>
    <el-table :data="problems" @row-click="(row) => { router.push(`/problem/${row.ID}`) }">
      <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
      <el-table-column prop="title" :label="t('message.title')"></el-table-column>
      <el-table-column :label="t('problem.difficulty')">
        <template v-slot="{ row }">
          <el-rate v-model="row.difficulty" :colors="colors" disabled />
        </template>
      </el-table-column>
      <el-table-column prop="passRate" :label="t('problem.pass-rate')"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { Problem } from '@/model';
import { getDataArr } from '@/utils/http';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const router = useRouter();
const colors = { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' };
const { data: problems, get } = getDataArr<Problem>('/problems');

onMounted(async () => {
  problems.value = await get();
});
</script>
