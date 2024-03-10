<template>
  <div>
    <el-row>
      <el-col :span="24">
        <el-menu mode="horizontal" :ellipsis="false" :default-active="'1'" router>
          <p id="logo">X-OJ</p>
          <el-menu-item class="hidden-lg-and-up" v-for="item in menu" :index="item.index" :route="item.route">{{
        item.label }}</el-menu-item>
          <div class="flex-grow" />
          <el-menu-item index="0" @click="router.push('/editor')"><el-icon>
              <EditPen />
            </el-icon>在线代码编辑器</el-menu-item>
          <el-sub-menu v-if="isLogin" index="1">
            <template #title>{{ username }}, 欢迎。</template>
            <el-menu-item index="1-1" @click="router.push('/profile')">个人中心</el-menu-item>
            <el-menu-item index="1-2" @click="
        logout();
      router.push('/');
      ">登出</el-menu-item>
            <el-menu-item v-if="role > 1" index="1-3" @click="router.push('/admin')">管理</el-menu-item>
          </el-sub-menu>
          <el-menu-item v-else index="1" @click="router.push('/login')">登录/注册</el-menu-item>
        </el-menu>
      </el-col>
    </el-row>
    <el-row justify="center">
      <el-col :lg="4" class="hidden-md-and-down">
        <el-affix>
          <el-menu id="vertical-menu" default-active="1" router>
            <el-menu-item index="1" route="/"><el-icon>
                <HomeFilled />
              </el-icon>主页</el-menu-item>
            <el-menu-item index="2" route="/problem"><el-icon>
                <List />
              </el-icon>问题</el-menu-item>
            <el-menu-item index="3" route="/contest"><el-icon>
                <Flag />
              </el-icon>竞赛/作业</el-menu-item>
            <el-menu-item index="4" route="/status"><el-icon>
                <HelpFilled />
              </el-icon>状态</el-menu-item>
            <el-menu-item index="5" route="/ranklist"><el-icon>
                <Histogram />
              </el-icon>排名</el-menu-item>
          </el-menu>
        </el-affix>
      </el-col>
      <el-col :md="24" :lg="20">
        <router-view id="router"></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { getRole, getUsername, isLogin, logout } from "@/utils/login";
const router = useRouter();
const username = getUsername();
const role = parseInt(getRole());
const menu = [
  { index: "1", label: "主页", route: "/" },
  { index: "2", label: "问题", route: "/problem" },
  { index: "3", label: "竞赛/作业", route: "/contest" },
  { index: "4", label: "状态", route: "/status" },
  { index: "5", label: "排名", route: "/ranklist" },
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
</style>
