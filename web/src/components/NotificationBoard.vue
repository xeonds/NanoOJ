<template>
  <div class="block">
    <el-carousel trigger="click" :interval="5000" indicator-position="outside" arrow="never">
      <el-carousel-item v-for="item in notifications" :key="item.id">
        <div class="notification-content">
          <h3>{{ item.title }} </h3>
          <p v-html="item.content"></p>
          <el-divider></el-divider>
          <el-tag>创建于：{{ formatDate(item.CreatedAt) }}</el-tag>
          <el-tag v-if="item.UpdatedAt">修改于：{{ formatDate(item.UpdatedAt) }}</el-tag>
        </div>
      </el-carousel-item>
    </el-carousel>
    <div v-if="notifications.length === 0">暂无公告</div>
  </div>
</template>

<script lang="ts" setup>
import { formatDate } from "@/utils/date.ts";
import { marked } from "marked";
import { Notification } from "@/assets/types";
import { http } from "@/utils/http";

const notifications: Ref<Notification[]> = ref([]);

const getNotifications = async (): Promise<Notification[]> => {
  const { data, err } = await http.get("/notifications");
  if (err) {
    console.error(err);
    return [];
  }
  return (data as Ref<Notification[]>).value.map((notification: any) => {
    notification.content = marked(notification.content);
    return notification;
  });
};

onMounted(async () => {
  notifications.value = await getNotifications();
});
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
