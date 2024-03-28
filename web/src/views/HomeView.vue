<template>
  <div>
    <el-row>
      <el-col :span="24">
        <el-menu mode="horizontal" :default-active="'1'" router>
          <p id="logo">X-OJ</p>
          <el-menu-item class="hidden-lg-and-up" v-for="item in menu" :index="item.index" :route="item.route">{{
            item.label }}</el-menu-item>
          <div class="flex-grow"></div>
          <el-menu-item index="0" @click="router.push('/editor')">
            <el-icon>
              <EditPen />
            </el-icon>{{ t('nav.editor') }}
          </el-menu-item>
          <el-sub-menu v-if="isLogin()" index="1">
            <template #title>{{ username }}, {{ t('message.welcome') }}</template>
            <el-menu-item index="1-1" @click="router.push('/profile')">{{ t('nav.personal-center') }}</el-menu-item>
            <el-menu-item index="1-2" @click="logout()">{{ t('nav.logout') }}</el-menu-item>
            <el-menu-item v-if="isAdmin()" index="1-3" @click="router.push('/admin')">{{ t('nav.admin') }}</el-menu-item>
          </el-sub-menu>
          <el-menu-item v-else index="1" @click="router.push('/login')">{{ t('nav.login-register') }}</el-menu-item>
        </el-menu>
      </el-col>
    </el-row>
    <el-row justify="center">
      <el-col :lg="4" class="hidden-md-and-down">
        <el-affix>
          <el-menu id="vertical-menu" default-active="1" router>
            <el-menu-item index="1" route="/"><el-icon>
                <HomeFilled />
              </el-icon>{{ t('nav.home') }}</el-menu-item>
            <el-menu-item index="2" route="/problem"><el-icon>
                <List />
              </el-icon>{{ t('nav.problem') }}</el-menu-item>
            <el-menu-item index="3" route="/contest"><el-icon>
                <Flag />
              </el-icon>{{ t('nav.assignment') }}</el-menu-item>
            <el-menu-item index="4" route="/status"><el-icon>
                <HelpFilled />
              </el-icon>{{ t('nav.status') }}</el-menu-item>
            <el-menu-item index="5" route="/ranklist"><el-icon>
                <Histogram />
              </el-icon>{{ t('nav.rank') }}</el-menu-item>
          </el-menu>
        </el-affix>
      </el-col>
      <el-col :md="24" :lg="20">
        <router-view id="router"></router-view>
      </el-col>
    </el-row>
    <FooterBox :msg="'X-OJ'" />
  </div>
</template>

<script lang="ts" setup>
import { getUsername, isAdmin, isLogin, logout } from "@/utils/login";
import FooterBox from "@/components/FooterBox.vue";
import { useI18n } from "vue-i18n";
const { t } = useI18n();
const router = useRouter();
const username = getUsername();
const menu = [
  { index: "1", label: t('nav.home'), route: "/" },
  { index: "2", label: t('nav.problem'), route: "/problem" },
  { index: "3", label: t('nav.assignment'), route: "/contest" },
  { index: "4", label: t('nav.status'), route: "/status" },
  { index: "5", label: t('nav.rank'), route: "/ranklist" },
];
</script>

<style>
#vertical-menu {
  min-width: 12rem;
  min-height: 100vh;
}

#logo {
  color: #409eff;
  align-self: center;
  font-family: Exo 2;
  font-size: 1.5em;
  font-weight: 600;
  padding-inline: 2rem;
  margin: 0;
  white-space: nowrap;
}

body {
  margin: 0px;
}

#router {
  padding: 2rem;
}

.flex-grow {
  flex-grow: 1;
}

.el-header,
.el-main,
.el-footer {
  padding: 0px !important;
}

.el-card {
  border-radius: 4px;
  margin-top: 20px;
  margin-bottom: 20px;
}
</style>
