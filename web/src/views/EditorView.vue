<template>
    <CodeEditor :language="language" :callback="submitCode" :height="'calc(100vh - 215px)'" id="editor">
        <template #editor-options>
            <div id="header-items">
                <el-page-header @back="goBack">
                    <template #content>
                        <el-text type="primary" class="title">X-OJ</el-text><el-text class="subtitle"> |
                            Playground</el-text>
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
        <template #header="{ close, titleId, titleClass }">
            <div>
                <el-text class="title">Running result</el-text><el-text class="subtitle"> | Output</el-text>
            </div>
            <el-button type="success" @click="copyToClipboard(output)" text>
                <el-icon class="el-icon--left">
                    <DocumentCopy />
                </el-icon>
                Copy Output
            </el-button>
        </template>
        <div id="console"> <!-- 控制台风格的窗格 -->
            <pre>{{ output }}</pre> <!-- 输出结果 -->
        </div>
    </el-drawer>
</template>

<script>
import CodeEditor from '../components/CodeEditor.vue';
import * as monaco from 'monaco-editor'
import api from '../api';

export default {
    components: {
        CodeEditor,
    },
    data() {
        return {
            problem: { description: '' },
            languages: [],
            language: 'c',
            show_output: false,
            output: 'Hello, World!',
        }
    },
    created() {
        this.languages = monaco.languages.getLanguages().map(function (lang) { return lang.id; });
        console.log(this.languages);
    },
    methods: {
        submitCode: function (language, code) {
            api.addSubmissions({ code: code, language: 'c', state: 'waiting', mode: 'playground' })
                .then((response) => {
                    if (response.status === 200) {
                        this.$message({
                            message: 'Code submitted successfully',
                            type: 'success'
                        });
                    }
                })
                .catch((error) => {
                    this.$message({ type: 'error', message: 'Failed to submit code: ' + error });
                });
        },
        goBack() {
            this.$router.go(-1);
        },
        async copyToClipboard(text) {
            try {
                await navigator.clipboard.writeText(text);
                this.$message({
                    message: 'Copied to clipboard',
                    type: 'success'
                });
            } catch (err) {
                this.$message({
                    message: 'Failed to copy to clipboard',
                    type: 'error'
                });
            }
        }

    },
};
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