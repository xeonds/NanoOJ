<template>
  <el-scrollbar height="100vh" id="app" class="dark">
    <el-row>
      <el-col :span="24">
        <el-menu mode="horizontal" :ellipsis="false">
          <p id="logo">X-OJ</p>
          <el-menu-item class="hidden-lg-and-up" index="1" @click="this.$router.push('/')">主页</el-menu-item>
          <el-menu-item class="hidden-lg-and-up" index="2" @click="this.$router.push('/problemset')">问题</el-menu-item>
          <el-menu-item class="hidden-lg-and-up" index="3" @click="this.$router.push('/contest')">竞赛/作业</el-menu-item>
          <el-menu-item class="hidden-lg-and-up" index="4" @click="this.$router.push('/status')">状态</el-menu-item>
          <el-menu-item class="hidden-lg-and-up" index="5" @click="this.$router.push('/ranklist')">排名</el-menu-item>
          <div class="flex-grow" />
          <el-sub-menu v-if="isLogin" index="1">
            <template #title>{{ username }}, 欢迎。</template>
            <el-menu-item index="1-1" @click="this.$router.push('/profile')">个人中心</el-menu-item>
            <el-menu-item index="1-2" @click="logout(); this.$router.push('/')">登出</el-menu-item>
            <el-menu-item v-if="userPermission > 1" index="1-3" @click="this.$router.push('/admin')">管理</el-menu-item>
          </el-sub-menu>
          <el-menu-item v-else index="1" @click="this.$router.push('/login')">登录/注册</el-menu-item>
        </el-menu>
      </el-col>
    </el-row>
    <el-row justify="center">
      <el-col :lg="4" class="hidden-md-and-down">
        <el-affix>
          <el-menu id="vertical-menu" default-active="1" :router="true">
            <el-menu-item index="1" route="/">主页</el-menu-item>
            <el-menu-item index="2" route="/problem">问题</el-menu-item>
            <el-menu-item index="3" route="/contest">竞赛/作业</el-menu-item>
            <el-menu-item index="4" route="/status">状态</el-menu-item>
            <el-menu-item index="5" route="/ranklist">排名</el-menu-item>
          </el-menu>
        </el-affix>
      </el-col>
      <el-col :md="24" :lg="20">
        <!-- < height="800px"> -->
        <router-view id="router-view"></router-view>
        <!-- </> -->
      </el-col>
    </el-row>
    <el-footer>
      <FooterBox msg="X-OJ"></FooterBox>
    </el-footer>
    <div class="wrapper">
      <svg class="waves" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
        viewBox="0 24 150 28" preserveAspectRatio="none" shape-rendering="auto">
        <defs>
          <path id="gentle-wave" d="M-160 44c30 0 58-18 88-18s 58 18 88 18 58-18 88-18 58 18 88 18 v44h-352z" />
        </defs>
        <g class="parallax">
          <use xlink:href="#gentle-wave" x="48" y="0" fill="rgba(64,158,255,0.9)" />
          <use xlink:href="#gentle-wave" x="48" y="3" fill="rgba(64,158,255,0.5)" />
          <use xlink:href="#gentle-wave" x="48" y="5" fill="rgba(64,158,255,0.3)" />
          <use xlink:href="#gentle-wave" x="48" y="7" fill="rgba(64,158,255,0.1)" />
        </g>
      </svg>
    </div>
  </el-scrollbar>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import FooterBox from "./components/FooterBox.vue";
import { useDark, useToggle } from "@vueuse/core";


export default {
  components: {
    FooterBox,
  },
  created() {
    this.$store.dispatch("init");
  },
  computed: {
    ...mapGetters({
      isLogin: "getLoggedInStatus",
      userInfo: "getUserInfo"
    }),
    username: function () {
      return this.userInfo.username;
    },
    userPermission: function () {
      return this.userInfo.permission;
    },
  },
  methods: {
    ...mapActions({
      _logout: "logout"
    })
    , logout: function () {
      this._logout();
      this.$router.push("/");
    }
  }
};
</script>
 
<style>
#vertical-menu {
  min-width: 12rem;
  min-height: calc(100vh);
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

#router-view {
  padding: 2rem;
  min-height: calc(100vh - 14rem);
}

.flex-grow {
  flex-grow: 1;
}

.el-header,
.el-main,
.el-footer {
  padding: 0px !important;
}

/* Animation */
.wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-end;
  overflow: hidden;
}

.waves {
  z-index: 100;
  position: relative;
  width: 100%;
  margin-bottom: -7px;
  /*Fix for safari gap*/
  min-height: 100px;
  max-height: 150px;
}

.parallax>use {
  animation: move-forever 25s cubic-bezier(0.55, 0.5, 0.45, 0.5) infinite;
}

.parallax>use:nth-child(1) {
  animation-delay: -2s;
  animation-duration: 7s;
}

.parallax>use:nth-child(2) {
  animation-delay: -3s;
  animation-duration: 10s;
}

.parallax>use:nth-child(3) {
  animation-delay: -4s;
  animation-duration: 13s;
}

.parallax>use:nth-child(4) {
  animation-delay: -5s;
  animation-duration: 20s;
}

@keyframes move-forever {
  0% {
    transform: translate3d(-90px, 0, 0);
  }

  100% {
    transform: translate3d(85px, 0, 0);
  }
}
</style>
