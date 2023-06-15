<template>
  <el-scrollbar height="100vh" id="app" class="dark">
    <router-view id="root-router"></router-view>
    <FooterBox msg="X-OJ"></FooterBox>
  </el-scrollbar>
</template>
 
<script>
import FooterBox from "./components/FooterBox.vue";

export default {
  name: "App",
  components: {
    FooterBox,
  },
};
</script>

<style>
body {
  margin: 0px;
}

#root-router {
  min-height: calc(100vh - 151px);
}
</style>
