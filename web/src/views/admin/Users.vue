<template>
    <h2>Users</h2>
    <el-card>
        <template #header>
            <div class="card-header">
                <span>All Users</span>
                <span>
                    <el-button @click="refresh()" type="primary" text>Refresh</el-button>
                </span>
            </div>
        </template>
        <el-table :data="users" style="width: 100%">
            <el-table-column prop="ID" label="ID"></el-table-column>
            <el-table-column prop="username" label="Username"></el-table-column>
            <el-table-column prop="email" label="Email"></el-table-column>
            <el-table-column prop="CreatedAt" label="CreatedAt"></el-table-column>
            <el-table-column label="Actions">
                <template #default="{ row }">
                    <el-button @click="deleteUser(row.ID)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
</template>

<script lang="ts" setup>
import api from '@/api';
import { User } from '@/model';
import { getDataArr } from '@/utils/http';

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