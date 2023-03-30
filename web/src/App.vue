<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-menu default-active="1" mode="horizontal" :router="true" :ellipsis="false">
          <el-menu-item id="logo" index="0" route="/">X-OJ</el-menu-item>
          <el-menu-item index="1" route="/">主页</el-menu-item>
          <el-menu-item index="2" route="/problemset">问题</el-menu-item>
          <el-menu-item index="3" route="/contest">竞赛/作业</el-menu-item>
          <el-menu-item index="4" route="/status">状态</el-menu-item>
          <el-menu-item index="5" route="/ranklist">排名</el-menu-item>
          <!--<el-menu-item index="6" route="/about">常见问答</el-menu-item>-->
          <div class="flex-grow" />
          <el-sub-menu v-if="isLogin" index="6">
            <template #title>{{ username }}, 欢迎。</template>
            <div>
              <el-menu-item index="6-1" route="/profile">个人中心</el-menu-item>
              <el-menu-item index="6-2" route="/" @click="logout">登出</el-menu-item>
              <el-menu-item v-if="userPermission > 1" index="6-3" route="/admin">管理</el-menu-item>
            </div>
          </el-sub-menu>
          <el-menu-item v-else index="6" route="/login">登录/注册</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <router-view> </router-view>
      </el-main>
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
    </el-container>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import FooterBox from "./components/FooterBox.vue";

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
#logo {
  padding-inline: 2rem;
  color: #409eff;
  align-self: center;
  font-family: Exo 2;
  font-size: 1.5em;
  font-weight: 600;
  min-width: fit-content;
}

body {
  margin: 0px;
}

.el-main>* {
  max-width: 1300px;
  width: calc(100% - 4rem);
}

.el-main {
  display: flex !important;
  flex-flow: column;
  align-items: center;
  margin-top: 2rem !important;
  margin-bottom: 2rem !important;
}

.flex-grow {
  flex-grow: 1;
}

.el-header,
.el-main,
.el-footer {
  padding: 0px !important;
}

.wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-end;
  overflow: hidden;
}

.waves {
  position: relative;
  width: 100%;
  margin-bottom: -7px;
  /*Fix for safari gap*/
  min-height: 100px;
  max-height: 150px;
}

/* Animation */

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
