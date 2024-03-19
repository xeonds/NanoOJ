<template>
  <div class="admin-view">
    <el-affix>
      <el-page-header @back="goBack">
        <template #content>
          <el-text type="primary" id="title">X-OJ</el-text><el-text id="subtitle"> | Admin Center</el-text>
        </template>
      </el-page-header>
    </el-affix>
    <el-tabs tab-position="left">
      <el-tab-pane label="Problems">
        <Problems />
      </el-tab-pane>
      <el-tab-pane label="Contests">
        <Contests />
      </el-tab-pane>
      <el-tab-pane label="Notifications">
        <Notifications />
      </el-tab-pane>
      <el-tab-pane label="Submissions">
        <Submisssions />
      </el-tab-pane>
      <el-tab-pane label="Users">
        <Users />
      </el-tab-pane>
      <el-tab-pane label="System">
        <System />
      </el-tab-pane>
      <el-tab-pane label="Manual">
        <Manual />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts" setup>
import Problems from '@/views/admin/Problems.vue';
import Contests from '@/views/admin/Contests.vue';
import Notifications from '@/views/admin/Notifications.vue';
import Submisssions from '@/views/admin/Submissions.vue';
import Users from '@/views/admin/Users.vue';
import System from '@/views/admin/System.vue';
import Manual from '@/views/admin/Manual.vue';

const router = useRouter();
const goBack = () => {
  router.go(-1);
}
</script>

<style scoped>
.el-page-header {
  padding: 1rem;
  border-bottom: 1px solid #ebeef5;
  text-align: center;
}

#title {
  font-size: 1.2rem;
  font-weight: 600;
}

#subtitle {
  font-size: 1.2rem;
  font-weight: 100;
}

.el-tabs {
  padding: 2rem;
}

.el-tab-pane {
  padding: 2rem;
}
</style>
