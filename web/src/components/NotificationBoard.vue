<template>
  <div class="block">
    <el-carousel trigger="click" :interval="5000" indicator-position="outside" arrow="never">
      <el-carousel-item v-for="(notification, index) in Notifications" :key="index">
        <div class="notification-content">
          <h3>{{ notification.title }} </h3>
          <p v-html="notification.content"></p>
          <el-divider></el-divider>
          <el-tag>创建于：{{ formatDate(notification.CreatedAt) }}</el-tag>
          <el-tag v-if="notification.UpdateAt">修改于：{{ formatDate(notification.UpdateAt) }}</el-tag>
        </div>
      </el-carousel-item>
    </el-carousel>
    <div v-if="notifications.length === 0">暂无公告</div>
  </div>
</template>

<script>
import utils from "../utils";
import { marked } from "marked";
export default {
  name: "NotificationBoard",
  props: {
    notifications: {
      type: Array,
      default: () => [],
    },
  },
  computed: {
	Notifications() {
	  return this.notifications.map((notification) => {
		notification.content = marked(notification.content);
		return notification;
	  });
	},
  },
  methods: {
	formatDate(date) {
	  return utils.formatDate(date);
	},
  },
};
</script>

<style scoped>

.notification-content {
  margin-bottom: 10px;
}

.notification-time {
  text-align: right;
  font-size: 12px;
  color: #999;
}

</style>
