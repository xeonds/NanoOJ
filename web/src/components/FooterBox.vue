<template>
  <div class="footer">
    <div id="footer-main">
      &copy;2022 {{ msg }} | Powered by
      <el-link type="primary" target="_blank" href="https://github.com/xeonds/NanoOJ">NanoOJ</el-link>
    </div>
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
#footer-main {
  padding-top: 2rem;
  border-top: 1px solid #c8c9cc;
  color: #606266;
  text-align: center;
  z-index: 110 !important;
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
  z-index: 100 !important;
  position: relative;
  width: 100%;
  min-height: 64px;
  max-height: 96px;
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