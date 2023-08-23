<template>
  <div>
    <el-row>
      <el-col :span="24">
        <el-menu
          mode="horizontal"
          :ellipsis="false"
          :default-active="'1'"
          router
        >
          <p id="logo">X-OJ</p>
          <el-menu-item class="hidden-lg-and-up" index="1" route="/"
            >主页</el-menu-item
          >
          <el-menu-item class="hidden-lg-and-up" index="2" route="/problem"
            >问题</el-menu-item
          >
          <el-menu-item class="hidden-lg-and-up" index="3" route="/contest"
            >竞赛/作业</el-menu-item
          >
          <el-menu-item class="hidden-lg-and-up" index="4" route="/status"
            >状态</el-menu-item
          >
          <el-menu-item class="hidden-lg-and-up" index="5" route="/ranklist"
            >排名</el-menu-item
          >
          <div class="flex-grow" />
          <el-menu-item index="0" @click="this.$router.push('/editor')"
            ><el-icon>
              <EditPen /> </el-icon
            >在线代码编辑器</el-menu-item
          >
          <el-sub-menu v-if="isLogin" index="1">
            <template #title>{{ username }}, 欢迎。</template>
            <el-menu-item index="1-1" @click="this.$router.push('/profile')"
              >个人中心</el-menu-item
            >
            <el-menu-item
              index="1-2"
              @click="
                logout();
                this.$router.push('/');
              "
              >登出</el-menu-item
            >
            <el-menu-item
              v-if="userPermission > 1"
              index="1-3"
              @click="this.$router.push('/admin')"
              >管理</el-menu-item
            >
          </el-sub-menu>
          <el-menu-item v-else index="1" @click="this.$router.push('/login')"
            >登录/注册</el-menu-item
          >
        </el-menu>
      </el-col>
    </el-row>
    <el-row justify="center">
      <el-col :lg="4" class="hidden-md-and-down">
        <el-affix>
          <el-menu id="vertical-menu" default-active="1" router>
            <el-menu-item index="1" route="/"
              ><el-icon>
                <HomeFilled /> </el-icon
              >主页</el-menu-item
            >
            <el-menu-item index="2" route="/problem"
              ><el-icon>
                <List /> </el-icon
              >问题</el-menu-item
            >
            <el-menu-item index="3" route="/contest"
              ><el-icon>
                <Flag /> </el-icon
              >竞赛/作业</el-menu-item
            >
            <el-menu-item index="4" route="/status"
              ><el-icon>
                <HelpFilled /> </el-icon
              >状态</el-menu-item
            >
            <el-menu-item index="5" route="/ranklist"
              ><el-icon>
                <Histogram /> </el-icon
              >排名</el-menu-item
            >
          </el-menu>
        </el-affix>
      </el-col>
      <el-col :md="24" :lg="20">
        <router-view id="router"></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
  created() {
    this.$store.dispatch("init");
  },
  computed: {
    ...mapGetters({
      isLogin: "getLoggedInStatus",
      userInfo: "getUserInfo",
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
      _logout: "logout",
    }),
    logout: function () {
      this._logout();
      this.$router.push("/");
    },
  },
};
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
