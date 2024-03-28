<template>
  <div id="contest">
    <el-row :gutter="20" class="hidden-sm-and-down">
      <el-col :span="16">
        <el-card>
          <template #header>
            <el-text type="primary">#{{ contest.ID }}.{{ contest.title }}</el-text>
          </template>
          <h4>{{ t('message.time') }}</h4>
          <el-row>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">{{ t('message.start-date') }}</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.start_time) }}</div>
              </div>
            </el-col>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">{{ t('message.end-date') }}</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.end_time) }}</div>
              </div>
            </el-col>
            <el-col :span="8">
              <el-countdown :title="t('message.countdown')" format="DD [days] HH:mm:ss" :value="time.fromString(contest.end_time)" />
            </el-col>
          </el-row>
          <h4>{{ t('message.description') }}</h4>
          <p v-html="description"></p>
        </el-card>
        <el-card>
          <template #header>
            <el-text type="primary">#{{ t('problem.list') }}</el-text>
          </template>
          <el-table :data="contest.problems" @row-click="(row) => { router.push(`/problem/${row.ID}`) }">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column prop="title" :label="t('message.title')"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <el-text type="primary">#{{ t('rank.list') }}</el-text>
          </template>
          <el-table :data="rank">
            <el-table-column prop="user_id" :label="t('user.id')"></el-table-column>
            <el-table-column prop="rank" :label="t('message.rank')"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row class="hidden-md-and-up">
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-text type="primary">#{{ contest.ID }}.{{ contest.title }}</el-text>
          </template>
          <h4>{{ t('message.time') }}</h4>
          <el-row>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">{{ t('message.start-date') }}</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.start_time) }}</div>
              </div>
            </el-col>
            <el-col :span="8">
              <div class="el-statistic">
                <div class="el-statistic__head">{{ t('message.end-date') }}</div>
                <div class="el-statistic__content">{{ time.formatDate(contest.end_time) }}</div>
              </div>
            </el-col>
            <el-col :span="8">
              <el-countdown :title="t('message.countdown')" format="DD [days] HH:mm:ss" :value="time.fromString(contest.end_time)" />
            </el-col>
          </el-row>
          <h4>{{ t('message.description') }}</h4>
          <p v-html="description"></p>
          <el-divider>OwO</el-divider>
          <h4>{{ t('problem.list') }}</h4>
          <el-table :data="contest.problems" @row-click="(row) => { router.push(`/problem/${row.ID}`) }">
            <el-table-column prop="index" :label="t('message.id')"></el-table-column>
            <el-table-column prop="title" :label="t('message.title')"></el-table-column>
            <el-table-column prop="score" :label="t('nav.status')"></el-table-column>
          </el-table>
        </el-card>
        <el-card style="">
          <template #header>
            <el-text type="primary">#{{ t('rank.list') }}</el-text>
          </template>
          <el-table :data="rank">
            <el-table-column prop="user_id" :label="t('user.id')"></el-table-column>
            <el-table-column prop="rank" :label="t('message.rank')"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { marked } from "marked";
import { time } from "@/utils/datetime";
import { Contest, Rank } from "@/model";
import { getData, getDataArr } from "@/utils/http";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const id = Number(route.params.id);
const description = computed(() => marked(contest.value.description ?? 'loading...'));
const { data: contest, get: getContest } = getData<Contest>(`/contests/${id}`);
const { data: rank, get: getRank } = getDataArr<Rank>(`/contests/${id}/rank`);

onMounted(async () => {
  contest.value = await getContest();
  rank.value = await getRank();
});
</script>