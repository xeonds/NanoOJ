<template>
    <h2>{{ t('nav.system') }}</h2>
    <h3>{{ t('message.danger-zone') }}</h3>
    <el-row>
        <el-col :span="24">
            <el-button type="danger" @click="clearCache">{{ t('system.clear-cache') }}</el-button>
            <el-button type="danger" @click="shutdown">{{ t('system.shutdown') }}</el-button>
            <el-button type="danger" @click="restart">{{ t('system.restart') }}</el-button>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import { useI18n } from 'vue-i18n';

const { t } = useI18n()
const clearCache = () => {
    ElMessage.info(t('message.not-implemented'))
}
const shutdown = () => {
    ElMessage.info(t('message.not-implemented'))
}
const restart = () => {
    ElMessage.info(t('message.not-implemented'))
}
</script>