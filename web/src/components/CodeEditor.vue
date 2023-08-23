<template>
    <el-row id="row-1">
        <div>
            <slot name="editor-options"></slot>
        </div>
        <el-popover placement="bottom" width="256" trigger="click" v-model:visible="visible">
            <el-form label-position="left" label-width="80px">
                <el-form-item label="主题">
                    <el-select v-model="theme" placeholder="请选择主题">
                        <el-option label="vs" value="vs"></el-option>
                        <el-option label="vs-dark" value="vs-dark"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="字体大小">
                    <el-input-number v-model="options.fontSize" :min="10" :max="30"></el-input-number>
                </el-form-item>
                <el-form-item label="编辑模式">
                    <el-switch v-model="options.vimMode" active-text="Vim" inactive-text="普通"></el-switch>
                </el-form-item>
            </el-form>
            <template #reference>
                <el-button type="primary" icon="el-icon-setting" text><el-icon>
                        <Setting />
                    </el-icon>编辑器设置</el-button>
            </template>
        </el-popover>
    </el-row>
    <el-row>
        <div id="editor" :style="{ height: height }"></div>
    </el-row>
</template>

<script>
import * as monaco from 'monaco-editor'

export default {
    props: {
        language: {
            type: String,
            default: "cpp",
        },
        callback: {
            type: Function,
            required: true,
        },
        height: {
            type: String,
            default: "42rem",
        },
    },
    data() {
        return {
            editor: null,
            code: "",
            theme: "vs-dark",
            options: {
                minimap: {
                    enabled: false,
                },
                fontSize: 14,
                vimMode: true,
            },
        }
    },
    watch: {
        fontSize(newValue) {
            this.options.fontSize = newValue;
        },
        vimMode(newValue) {
            this.options.vimMode = newValue === "true";
        },
    },
    methods: {
        initEditor() {
            // 初始化编辑器，确保dom已经渲染
            this.editor = monaco.editor.create(document.getElementById('editor'), {
                value: '',//编辑器初始显示文字
                language: this.language,//语言支持自行查阅demo
                automaticLayout: true,//自动布局
                theme: this.theme //官方自带三种主题vs, hc-black, or vs-dark
            });
        },
        getValue() {
            this.editor.getValue(); //获取编辑器中的文本
        },
        toggleTheme() {
            this.theme = this.theme === "vs-dark" ? "vs" : "vs-dark";
            this.options.theme = this.theme;
        },
        handleChangeLanguage(newLanguage) {
            this.language = newLanguage;
            this.callback(this.language, this.code);
        },
    },
    mounted() {
        this.initEditor();
    },
}
</script>

<style scoped>
#row-1 {
    display: flex;
    flex-flow: row nowrap;
    justify-content: space-between;
    padding-top: 1rem;
    padding-bottom: 1rem;
}

#editor {
    width: 100%;
    height: 100%;
}
</style>
