<template>
  <div class="footer">
    <div class="ui center aligned container">
      <div>
        >&copy;2022 {{ msg }} | Powered by
        <a target="_blank" href="http://www.jiujiuer.xyz:3000/xeonds/NanoOJ"
          >NanoOJ</a
        >,
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Footer",
  props: {
    msg: "X-OJ",
  },
};
</script>

<style scoped>
.footer {
  line-height: 1.4285em;
  font-family: "Lato", "Noto Sans CJK SC", "Source Han Sans SC", "PingFang SC",
    "Hiragino Sans GB", "Microsoft Yahei", "WenQuanYi Micro Hei",
    "Droid Sans Fallback", "sans-serif";
  box-sizing: inherit;
  padding: 0 !important;
  border: none !important;
  color: #888;
  font-size: 1rem;
  margin: 35px 0 14px !important;
  position: relative;
  width: 100%;
  bottom: 0;
  background: none transparent;
  border-radius: 0;
  box-shadow: none;
}
</style>