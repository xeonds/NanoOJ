<template>
  <div class="block">
    <el-carousel trigger="click" :interval="5000" indicator-position="outside" arrow="never">
      <el-carousel-item v-for="(notification, index) in notifications" :key="index">
        <h3>{{ notification.title }}</h3>
        <p>{{ notification.content }}</p>
        <p>发布时间：{{ notification.CreatedAt }}</p>
        <p>更新时间：{{ notification.UpdateAt }}</p>
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
      default: () => [],
    },
  },
};
</script>

<style scoped></style>
