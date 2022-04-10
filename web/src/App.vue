<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-menu default-active="1" mode="horizontal" router="true">
          <span
            class="logo"
            style="
              color: #409eff;
              align-self: center;
              font-family: 'Exo 2';
              font-size: 1.5em;
              font-weight: 600;
            "
            >X-OJ</span
          >
          <el-menu-item index="1" route="/">
            <span>主页</span>
          </el-menu-item>
          <el-menu-item index="2" route="/problemset">问题</el-menu-item>
          <el-menu-item index="3" route="/category">来源/分类</el-menu-item>
          <el-menu-item index="4" route="/contest">竞赛/作业</el-menu-item>
          <el-menu-item index="5" route="/status">状态</el-menu-item>
          <el-menu-item index="6" route="/ranklist">排名</el-menu-item>
          <el-menu-item index="7" route="/about">常见问答</el-menu-item>
          <el-menu-item index="8" route="/account">Username</el-menu-item>
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
  <!--  -->
</template>

<script>
import FooterBox from "@/components/FooterBox.vue";

export default {
  components: {
    FooterBox,
  },
};
</script>
 
<style>
body {
  margin: 0px;
}

.logo {
  padding-inline: 2rem;
}

.el-main {
  display: flex !important;
  flex-flow: column;
  align-items: center;
  margin-top: 2rem !important;
  margin-bottom: 2rem !important;
}

.el-header,
.el-main,
.el-footer {
  padding: 0px !important;
}
</style>