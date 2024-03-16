<template>
    <h2>System</h2>
    <h3>Danger Zone</h3>
    <el-row>
        <el-col :span="24">
            <el-button type="danger" @click="clearCache">Clear Cache</el-button>
            <el-button type="danger" @click="shutdown">Shutdown</el-button>
            <el-button type="danger" @click="restart">Restart</el-button>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
const clearCache = () => {
    ElMessage.info('Not implemented yet')
}
const shutdown = () => {
    ElMessage.info('Not implemented yet')
}
const restart = () => {
    ElMessage.info('Not implemented yet')
}
</script>