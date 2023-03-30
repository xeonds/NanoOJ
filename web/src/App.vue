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
</style>
