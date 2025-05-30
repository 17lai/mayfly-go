<template>
    <div class="format-viewer-container">
        <div class="mb-1 flex justify-end">
            <el-select v-model="selectedView" class="format-selector" size="small" placeholder="Text">
                <template #prefix>
                    <SvgIcon name="view" />
                </template>
                <el-option v-for="item of Object.keys(viewers)" :key="item" :label="item" :value="item"> </el-option>
            </el-select>
            <el-tag type="primary" :disable-transitions="true" class="ml-2">Size: {{ formatByteSize(state.contentSize) }}</el-tag>
        </div>

        <component ref="viewerRef" :is="components[viewerComponent]" :content="state.content" :name="selectedView"> </component>
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive, computed, shallowReactive, watch, toRefs, onMounted } from 'vue';
import ViewerText from './ViewerText.vue';
import ViewerJson from './ViewerJson.vue';
import { formatByteSize } from '@/common/utils/format';

const props = defineProps({
    content: {
        type: String,
    },
    height: {
        type: String,
        default: '0px',
    },
});

const components: any = shallowReactive({
    ViewerText,
    ViewerJson,
});
const viewerRef: any = ref(null);

const state = reactive({
    content: '',
    contentSize: 0,
    selectedView: 'Text',
});

const viewers: any = {
    Text: {
        value: 'ViewerText',
    },

    Json: {
        value: 'ViewerJson',
    },
};

const { selectedView } = toRefs(state);

const viewerComponent = computed(() => {
    return viewers[state.selectedView].value;
});

watch(
    () => props.content,
    (val: any) => {
        setContent(val);
    }
);

onMounted(() => {
    setContent(props.content as any);
});

const setContent = (content: string) => {
    state.content = content;
    state.contentSize = new Blob([content]).size;
    try {
        JSON.parse(content);
        state.selectedView = 'Json';
    } catch (e) {
        state.selectedView = 'Text';
    }
};

const getContent = () => {
    return viewerRef.value.getContent();
};

defineExpose({ getContent });
</script>

<style lang="scss">
.format-selector {
    width: 130px;
}

.format-selector .el-input__inner {
    height: 22px !important;
}

/*outline same with text viewer's .el-textarea__inner*/
.format-viewer-container .text-formated-container {
    border: 1px solid var(--el-border-color-light, #ebeef5);
    padding: 5px 10px;
    border-radius: 4px;
    clear: both;
}

.format-viewer-container .formater-binary-tag {
    font-size: 80%;
}

// 默认文本框样式

.format-viewer-container .el-textarea textarea {
    font-size: 14px;
    height: calc(100vh - 550px + v-bind(height));
}

.format-viewer-container .monaco-editor-content {
    height: calc(100vh - 565px + v-bind(height)) !important;
}
</style>
