<template>
  <div>
    <el-carousel :interval="5000" arrow="never">
      <el-carousel-item v-for="(notification, index) in notifications" :key="index">
        {{ notification }}
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
      default: () => ["测试公告\nHello, world!","测试公告2\nHello, world!"],
    },
  },
};
</script>

<style scoped>
</style>