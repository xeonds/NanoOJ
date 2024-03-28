<template>
  <div>
    <h2>{{ t('nav.rank') }}</h2>
    <el-table :data="ranks">
      <el-table-column prop="index" :label="t('message.rank')"></el-table-column>
      <el-table-column prop="username" :label="t('user.name')"></el-table-column>
      <el-table-column prop="profile" :label="t('user.profile')"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { Rank } from '@/model';
import { getDataArr } from '@/utils/http';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const { data: ranks, get } = getDataArr<Rank>('/ranks');
onMounted(async () => {
  ranks.value = await get();
});
</script>
