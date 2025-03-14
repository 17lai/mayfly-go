<template>
    <div class="machine-file-content">
        <el-dialog
            destroy-on-close
            :before-close="handleClose"
            :title="title || path"
            v-model="dialogVisible"
            :close-on-click-modal="false"
            top="5vh"
            width="65%"
        >
            <div>
                <monaco-editor :can-change-mode="true" v-model="content" :language="fileType" />
            </div>

            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="handleClose">{{ $t('common.cancel') }}</el-button>
                    <el-button v-auth="'machine:file:write'" type="primary" @click="updateContent">{{ $t('common.save') }}</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { reactive, toRefs, watch } from 'vue';
import { machineApi } from '../api';
import MonacoEditor from '@/components/monaco/MonacoEditor.vue';
import { useI18nSaveSuccessMsg } from '@/hooks/useI18n';

const props = defineProps({
    visible: { type: Boolean, default: false },
    protocol: { type: Number, default: 1 },
    title: { type: String, default: '' },
    machineId: { type: Number },
    authCertName: { type: String },
    fileId: { type: Number, default: 0 },
    path: { type: String, default: '' },
});

const emit = defineEmits(['update:visible', 'cancel', 'update:machineId']);

const updateFileContent = machineApi.updateFileContent;

const state = reactive({
    dialogVisible: false,
    content: '',
    fileType: '',
});

const { dialogVisible, content, fileType } = toRefs(state);

watch(props, async (newValue) => {
    if (newValue.visible) {
        await getFileContent();
    }
    state.dialogVisible = newValue.visible;
});

const getFileContent = async () => {
    const path = props.path;
    const res = await machineApi.fileContent.request({
        fileId: props.fileId,
        path,
        machineId: props.machineId,
        authCertName: props.authCertName,
        protocol: props.protocol,
    });
    state.fileType = getFileType(path);
    state.content = res;
};

const handleClose = () => {
    state.dialogVisible = false;
    emit('update:visible', false);
};

const updateContent = async () => {
    await updateFileContent.request({
        content: state.content,
        id: props.fileId,
        path: props.path,
        machineId: props.machineId,
        authCertName: props.authCertName,
        protocol: props.protocol,
    });
    useI18nSaveSuccessMsg();
    handleClose();
    state.content = '';
};

const getFileType = (path: string) => {
    if (path.endsWith('.sh')) {
        return 'shell';
    }
    if (path.endsWith('js')) {
        return 'javascript';
    }
    if (path.endsWith('json')) {
        return 'json';
    }
    if (path.endsWith('Dockerfile')) {
        return 'dockerfile';
    }
    if (path.endsWith('nginx.conf')) {
        return 'shell';
    }
    if (path.endsWith('sql')) {
        return 'sql';
    }
    if (path.endsWith('yaml') || path.endsWith('yml')) {
        return 'yaml';
    }
    if (path.endsWith('xml') || path.endsWith('html')) {
        return 'html';
    }
    if (path.endsWith('py')) {
        return 'python';
    }
    return 'text';
};
</script>
<style lang="scss"></style>
