<template>
    <h2>Users</h2>
    <el-table :data="users" style="width: 100%">
        <el-table-column prop="id" label="ID"></el-table-column>
        <el-table-column prop="username" label="Username"></el-table-column>
        <el-table-column prop="email" label="Email"></el-table-column>
        <el-table-column prop="CreatedAt" label="CreatedAt"></el-table-column>
        <el-table-column label="Actions">
            <template #default="{ row }">
                <el-button @click="deleteUser(row.id)">Delete</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<script lang="ts" setup>
import api from '@/api';
import { User } from '@/model';
import { getDataArr } from '@/utils/http';

const { data: users, get: getUsers } = getDataArr<User>('/users');

const deleteUser = async (id: number) => {
    api.deleteUser(id).then(() => {
        users.value = users.value.filter(user => user.id !== id);
    });
}

onMounted(async () => {
    users.value = await getUsers();
})
</script>