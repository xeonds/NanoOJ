<template>
    <CodeEditor :language="language" :height="'calc(100vh - 215px)'" id="editor">
        <template #editor-options>
            <div id="header-items">
                <el-page-header @back="goBack">
                    <template #content>
                        <el-text type="primary" class="title">X-OJ</el-text>
                        <el-text class="subtitle"> | Playground</el-text>
                    </template>
                </el-page-header>
                <el-select v-model="language" placeholder="Select Language">
                    <el-option v-for="lang in languages" :key="lang" :label="lang" :value="lang"></el-option>
                </el-select>
                <el-button type="primary" @click="submitCode" text>Commit</el-button>
                <el-button type="primary" @click="show_output = true" text>Show Output</el-button>
            </div>
        </template>
    </CodeEditor>
    <el-drawer v-model="show_output" :show-close="false" direction="btt">
        <template #header="">
            <div>
                <el-text class="title">Running result</el-text>
                <el-text class="subtitle"> | Output</el-text>
            </div>
            <el-button type="success" @click="copyToClipboard(output)" text>
                <el-icon class="el-icon--left">
                    <DocumentCopy />
                </el-icon>
                Copy Output
            </el-button>
        </template>
        <div id="console">
            <!-- 控制台风格的窗格 -->
            <pre>{{ output }}</pre>
            <!-- 输出结果 -->
        </div>
    </el-drawer>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import CodeEditor from '@/components/CodeEditor.vue';
import * as monaco from 'monaco-editor';
import api from '@/api';

const code = ref('Hello, World');
const language = ref('c');
const show_output = ref(false);
const output = ref('Hello, World!');
const languages: Ref<string[]> = ref([]);
const router = useRouter();

const submitCode = () => {
    api
        .addSubmissions({ code: code.value, language: language.value, state: 'waiting', mode: 'playground' })
        .then((response: any) => {
            if (response.status === 200) {
                console.log('Code submitted successfully');
            }
        })
        .catch((error) => {
            console.error('Failed to submit code:', error);
        });
};
const goBack = () => { router.go(-1) };
const copyToClipboard = async (text: string) => {
    try {
        await navigator.clipboard.writeText(text);
        console.log('Copied to clipboard');
    } catch (err) {
        console.error('Failed to copy to clipboard');
    }
};

onMounted(() => {
    languages.value = monaco.languages.getLanguages().map((lang) => lang.id);
    console.log(languages.value);
});

// pass ref:code to CodeEditor component
// thus the code in CodeEditor component can be passed back to here
provide('code', code)
</script>

<style scoped>
.el-page-header {
    padding-left: 1rem;
    text-align: center;
}

.title {
    font-size: 1.2rem;
    font-weight: 600;
}

.subtitle {
    font-size: 1.2rem;
    font-weight: 100;
}

#header-items {
    width: 64rem;
    display: flex;
    align-items: center;
}

#header-items>* {
    margin: 0 0.5rem;
}

#console {
    width: 100%;
    height: 100%;
    padding: 1rem;
    background-color: black;
    color: white;
    font-family: monospace;
    overflow-y: scroll;
}
</style>