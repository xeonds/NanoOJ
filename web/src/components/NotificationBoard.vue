<template>
  <div class="block">
    <el-carousel
      trigger="click"
      :interval="5000"
      indicator-position="outside"
      arrow="never"
    >
      <el-carousel-item
        v-for="(notification, index) in notifications"
        :key="index"
        ><div v-html="notification"></div>
      </el-carousel-item>
    </el-carousel>
    <div v-if="notifications.length === 0">暂无公告</div>
  </div>
</template>

<script>
export default {
  name: "NotificationBoard",
  props: {
    notifications: {
      type: Array,
      default: () => [
        "<h2>测试公告1</h2> <p>Hello, world!</p>",
        "<h2>测试公告2</h2> <p>Hello, world!</p>",
      ],
    },
  },
};
</script>

<style scoped>
</style>