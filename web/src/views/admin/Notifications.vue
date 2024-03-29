<template>
    <h2>{{ t('nav.notifications') }}</h2>
    <el-row>
        <el-col :span="24">
        </el-col>
    </el-row>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>{{ t('contest.all') }}</span>
                <span>
                    <el-button type="primary" @click="refresh()" text>{{ t('message.refresh') }}</el-button>
                    <el-button type="primary" @click="createNotificationDialogVisible = true">{{ t('message.create')
                        }}</el-button>
                </span>
            </div>
        </template>
        <el-table :data="notifications" style="width: 100%">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column prop="title" :label="t('message.title')"></el-table-column>
            <el-table-column prop="author" :label="t('message.author')"></el-table-column>
            <el-table-column :label="t('message.created-at')">
                <template #default="{ row }"> <span type="info">{{ time.formatDate(row.CreatedAt) }}</span> </template>
            </el-table-column>
            <el-table-column :label="t('message.action')">
                <template #default="{ row }">
                    <el-button @click="editNotification(row)">{{ t('message.edit') }}</el-button>
                    <el-button @click="deleteNotification(row.ID)">{{ t('message.delete') }}</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <el-dialog v-model="createNotificationDialogVisible" :title="t('notification.create')">
        <el-form :model="newNotification" label-width="100px">
            <el-form-item :label="t('message.title')">
                <el-input v-model="newNotification.title"></el-input>
            </el-form-item>
            <el-form-item :label="t('message.content')">
                <el-input v-model="newNotification.content" type="textarea"></el-input>
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="createNotificationDialogVisible = false">{{ t('message.cancel') }}</el-button>
            <el-button type="primary"
                @click="createNotification(newNotification); createNotificationDialogVisible = false">{{
        t('message.create')
    }}</el-button>
        </div>
    </el-dialog>
    <el-dialog v-model="editNotificationDialogVisible" :title="t('message.title')">
        <el-form :model="selectedNotification" label-width="100px">
            <el-form-item :label="t('message.title')">
                <el-input v-model="selectedNotification.title"></el-input>
            </el-form-item>
            <el-form-item :label="t('message.content')">
                <el-input v-model="selectedNotification.content" type="textarea"></el-input>
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button @click="editNotificationDialogVisible = false">{{ t('message.cancel') }}</el-button>
            <el-button type="primary" @click="handleEditNotification">{{ t('message.create') }}</el-button>
        </div>
    </el-dialog>
</template>

<script lang="ts" setup>
import { getDataArr } from '@/utils/http';
import { Notification } from '@/model';
import api from '@/api';
import { useI18n } from 'vue-i18n';
import { time } from '@/utils/datetime';

const { t } = useI18n();
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
    api.updateNotification(selectedNotification.value, selectedNotification.value.ID).then(() => {
        notifications.value = notifications.value.map(n => n.ID === selectedNotification.value.ID ? selectedNotification.value : n);
        editNotificationDialogVisible.value = false;
    });
}

const deleteNotification = (id: number) => {
    api.deleteNotification(id).then(() => {
        notifications.value = notifications.value.filter(n => n.ID !== id);
    });
}

onMounted(async () => {
    notifications.value = await getNotifications();
})

const refresh = async () => {
    notifications.value = await getNotifications();
}
</script>