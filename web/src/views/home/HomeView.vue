<template>
  <div id="home">
    <div class="card-header">
      <h2>{{ t('nav.home') }}</h2>
      <el-button @click="refresh()" type="primary" text>{{ t('message.refresh') }}</el-button>
    </div>
    <el-row>
      <el-col :span="24">
        <el-card>
          <div v-if="notifications.length === 0">{{ t('notification.no-notification') }}</div>
          <el-carousel v-else trigger="click" :interval="5000" indicator-position="outside" arrow="never" height="auto">
            <el-carousel-item style="height: auto" v-for="item in notifications" :key="item.ID">
              <h3>{{ item.title }} </h3>
              <p v-html="item.content"></p>
              <el-divider></el-divider>
              <el-tag>{{ t('message.created-at') }} : {{ time.formatDate(item.CreatedAt) }}</el-tag> |
              <el-tag v-if="item.UpdatedAt">{{ t('message.modified-at') }} : {{ time.formatDate(item.UpdatedAt)
                }}</el-tag>
            </el-carousel-item>
          </el-carousel>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="hidden-sm-and-down">
      <el-col :span="16">
        <el-card>
          <h3>{{ t('problem.latest') }}</h3>
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
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <h3>{{ t('nav.rank') }}</h3>
          <el-table :data="ranks">
            <el-table-column prop="rank" :label="t('message.rank')"></el-table-column>
            <el-table-column prop="username" :label="t('user.name')"></el-table-column>
            <el-table-column prop="profile" :label="t('user.profile')"></el-table-column>
          </el-table>
        </el-card>
        <el-card>
          <h3>{{ t('contest.latest') }}</h3>
          <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.ID}`) }">
            <el-table-column prop="title" :label="t('message.title')"></el-table-column>
            <el-table-column :label="t('message.start-date')" :width="160">
              <template #default="{ row }"> <span type="info">{{ time.formatDate(row.start_time) }}</span> </template>
            </el-table-column>
            <el-table-column :label="t('message.end-date')" :width="160">
              <template #default="{ row }"> <span type="info">{{ time.formatDate(row.end_time) }}</span> </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="hidden-md-and-up">
      <el-col :span="24">
        <el-card>
          <h3>{{ t('problem.latest') }}</h3>
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
        </el-card>
        <el-card>
          <h3>{{ t('nav.rank') }}</h3>
          <el-table :data="ranks">
            <el-table-column prop="rank" :label="t('message.rank')"></el-table-column>
            <el-table-column prop="username" :label="t('user.name')"></el-table-column>
            <el-table-column prop="profile" :label="t('user.profile')"></el-table-column>
          </el-table>
        </el-card>
        <el-card>
          <h3>{{ t('contest.latest') }}</h3>
          <el-table :data="contests" @row-click="(row: any) => { router.push(`/contest/${row.ID}`) }">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column prop="title" :label="t('message.title')"></el-table-column>
            <el-table-column prop="start_time" :label="t('message.start-date')"></el-table-column>
            <el-table-column prop="end_time" :label="t('message.end-date')"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { Notification, Contest, Problem, Rank } from '@/model';
import { marked } from "marked";
import { getDataArr, handleArrRefresh } from '@/utils/http';
import { time } from "@/utils/datetime";
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const router = useRouter();
const colors = { 2: '#01D842', 4: '#66CCFF', 5: '#FF4040' };
const { data: notifications, get: getNotifications } = getDataArr<Notification>('/notifications');
const { data: problems, get: getProblems } = getDataArr<Problem>('/problems');
const { data: ranks, get: getRanks } = getDataArr<Rank>('/ranks');
const { data: contests, get: getContests } = getDataArr<Contest>('/contests');
const refresh = async () => {
  handleArrRefresh<Notification>(notifications, await getNotifications())
  handleArrRefresh<Problem>(problems, await getProblems())
  handleArrRefresh<Rank>(ranks, await getRanks())
  handleArrRefresh<Contest>(contests, await getContests())
}

onMounted(async () => {
  notifications.value = (await getNotifications()).map((item) => {
    item.content = marked(item.content) as string;
    return item;
  });
  problems.value = await getProblems();
  ranks.value = await getRanks();
  contests.value = await getContests();
});
</script>

<style scoped>
.notification-time {
  text-align: right;
  font-size: 12px;
  color: #999;
}
</style>
