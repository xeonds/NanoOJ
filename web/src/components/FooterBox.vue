<template>
  <div class="footer">
    &copy;2022 {{ msg }} | Powered by
    <el-link
      type="primary"
      target="_blank"
      href="http://www.jiujiuer.xyz:3000/xeonds/NanoOJ"
      >NanoOJ</el-link>
  </div>
</template>

<script>
export default {
  name: "FooterBox",
  props: {
    msg: String,
  },
  data() {
    return {};
  },
};
</script>

<style scoped>
.footer {
  padding: 2rem;
  border-top: 1px solid #c8c9cc;
  color: #606266;
  text-align: center;
}
</style>