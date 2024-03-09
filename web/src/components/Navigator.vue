<template>
    <el-menu mode="horizontal" :ellipsis="false">
        <p id="logo">X-OJ</p>
        <el-menu-item class="hidden-lg-and-up" index="1" @click="router.push('/')">主页</el-menu-item>
        <el-menu-item class="hidden-lg-and-up" index="2" @click="router.push('/problem')">问题</el-menu-item>
        <el-menu-item class="hidden-lg-and-up" index="3" @click="router.push('/contest')">竞赛/作业</el-menu-item>
        <el-menu-item class="hidden-lg-and-up" index="4" @click="router.push('/status')">状态</el-menu-item>
        <el-menu-item class="hidden-lg-and-up" index="5" @click="router.push('/ranklist')">排名</el-menu-item>
        <div class="flex-grow" />
        <el-sub-menu v-if="isLogin" index="1">
            <template #title>{{ username }}, 欢迎。</template>
            <el-menu-item index="1-1" @click="router.push('/profile')">个人中心</el-menu-item>
            <el-menu-item index="1-2" @click="logout(); router.push('/')">登出</el-menu-item>
            <el-menu-item v-if="permission > 1" index="1-3" @click="router.push('/admin')">管理</el-menu-item>
        </el-sub-menu>
        <el-menu-item v-else index="1" @click="router.push('/login')">登录/注册</el-menu-item>
    </el-menu>
</template>

<script lang="ts" setup>
import { getRole, getUsername, isLogin, logout } from '@/utils/login';

const router = useRouter();
const username = getUsername();
const permission = parseInt(getRole());
</script>