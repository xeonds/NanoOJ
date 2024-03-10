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
                <el-button type="primary" icon="el-icon-setting" text>
                    <el-icon>
                        <Setting />
                    </el-icon>编辑器设置
                </el-button>
            </template>
        </el-popover>
    </el-row>
    <el-row>
        <div id="editor" :style="{ height: height }"></div>
    </el-row>
</template>

<script lang="ts" setup>
import { reactive, onMounted } from 'vue';
import * as monaco from 'monaco-editor';
const props = defineProps<{
    language: string,
    height: string,
}>();
const code = inject('code') as Ref<string>;
const visible = ref(false);
const theme = ref('vs-dark');
const options = reactive({
    minimap: {
        enabled: true,
    },
    fontSize: 14,
    vimMode: true,
    theme: theme.value,
});
let editorInstance: monaco.editor.IStandaloneCodeEditor;

onMounted(() => {
    editorInstance = monaco.editor.create(document.getElementById("editor")!, {
        value: code.value, // Initial text to display in the editor
        language: props.language, // Language support, refer to the documentation for available options
        automaticLayout: true, // Automatic layout
        theme: theme.value, // Officially supported themes: vs, hc-black, or vs-dark
    });

    //bind value of editor to ref:code from parent
    editorInstance.onDidChangeModelContent(() => {
        code.value = editorInstance.getValue();
    });
});

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
