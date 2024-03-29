<template>
    <h2>{{ t('message.user') }}</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>{{ t('user.all') }}</span>
                <span>
                    <el-button @click="refresh()" type="primary" text>{{ t('message.refresh') }}</el-button>
                </span>
            </div>
        </template>
        <el-table :data="users" style="width: 100%">
            <el-table-column prop="ID" :label="t('message.id')"></el-table-column>
            <el-table-column prop="username" :label="t('user.name')"></el-table-column>
            <el-table-column prop="email" :label="t('user.email')"></el-table-column>
            <el-table-column :label="t('message.created-at')">
                <template #default="{ row }"> <span type="info">{{ time.formatDate(row.CreatedAt) }}</span> </template>
            </el-table-column>
            <el-table-column :label="t('message.action')">
                <template #default="{ row }">
                    <el-button @click="deleteUser(row.ID)">{{ t('message.delete') }}</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
</template>

<script lang="ts" setup>
import api from '@/api';
import { User } from '@/model';
import { time } from '@/utils/datetime';
import { getDataArr } from '@/utils/http';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const { data: users, get: getUsers } = getDataArr<User>('/users');
const deleteUser = async (id: number) => {
    api.deleteUser(id).then(() => {
        users.value = users.value.filter(user => user.ID !== id);
    });
}

onMounted(async () => {
    users.value = await getUsers();
})

const refresh = async () => {
    users.value = await getUsers();
}
</script>