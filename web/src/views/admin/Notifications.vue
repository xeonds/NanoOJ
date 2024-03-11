<template>
    <h2>Notifications</h2>
    <el-row>
        <el-col :span="24">
        </el-col>
    </el-row>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>All Notifications</span>
                <el-button type="primary" @click="createNotificationDialogVisible = true">Create
                    Notification</el-button>
            </div>
        </template>
        <el-table :data="notifications" style="width: 100%">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="title" label="通知标题"></el-table-column>
            <el-table-column prop="author" label="发布者"></el-table-column>
            <el-table-column prop="CreatedAt" label="发布时间"></el-table-column>
            <el-table-column label="Actions">
                <template #default="{ row }">
                    <el-button @click="editNotification(row)">Edit</el-button>
                    <el-button @click="deleteNotification(row.id)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <el-dialog v-model="createNotificationDialogVisible" title="Create Notification">
        <el-form :model="newNotification" label-width="100px">
            <el-form-item label="Title">
                <el-input v-model="newNotification.title"></el-input>
            </el-form-item>
            <el-form-item label="Content">
                <el-input v-model="newNotification.content" type="textarea"></el-input>
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="createNotificationDialogVisible = false">Cancel</el-button>
            <el-button type="primary"
                @click="createNotification(newNotification); createNotificationDialogVisible = false">Create</el-button>
        </div>
    </el-dialog>
    <el-dialog v-model="editNotificationDialogVisible" title="Edit Notification">
        <el-form :model="selectedNotification" label-width="100px">
            <el-form-item label="Title">
                <el-input v-model="selectedNotification.title"></el-input>
            </el-form-item>
            <el-form-item label="Content">
                <el-input v-model="selectedNotification.content" type="textarea"></el-input>
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="editNotificationDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="handleEditNotification">Create</el-button>
        </div>
    </el-dialog>
</template>

<script lang="ts" setup>
import { getDataArr } from '@/utils/http';
import { Notification } from '@/model';
import api from '@/api';

const { data: notifications, get: getNotifications } = getDataArr<Notification>('/notifications');

const newNotification = ref<Notification>({} as Notification);
const createNotificationDialogVisible = ref(false);
const createNotification = (notification: Notification) => {
    api.addNotifications(notification).then(() => {
        notifications.value.push(notification);
    });
}

const selectedNotification = ref<Notification>({} as Notification);
const editNotificationDialogVisible = ref(false);
const editNotification = (notification: Notification) => {
    selectedNotification.value = notification;
    editNotificationDialogVisible.value = true;
}
const handleEditNotification = () => {
    api.updateNotification(selectedNotification.value, selectedNotification.value.id).then(() => {
        notifications.value = notifications.value.map(n => n.id === selectedNotification.value.id ? selectedNotification.value : n);
        editNotificationDialogVisible.value = false;
    });
}

const deleteNotification = (id: number) => {
    api.deleteNotification(id).then(() => {
        notifications.value = notifications.value.filter(n => n.id !== id);
    });
}

onMounted(async () => {
    notifications.value = await getNotifications();
})
</script>