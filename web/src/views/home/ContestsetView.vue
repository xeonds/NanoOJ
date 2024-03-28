<template>
  <div id="contests">
    <h2>{{ t('nav.assignment') }}</h2>
    <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.ID}`) }">
      <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
      <el-table-column prop="title" :label="t('message.title')"></el-table-column>
      <el-table-column prop="start_time" :label="t('message.start-date')"></el-table-column>
      <el-table-column prop="end_time" :label="t('message.end-date')"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { Contest } from '@/model';
import { getDataArr } from '@/utils/http';
import { time } from '@/utils/datetime';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
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
