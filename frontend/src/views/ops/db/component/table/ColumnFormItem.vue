<template>
    <div class="string-input-container !w-full" v-if="dataType == DataType.String || dataType == DataType.Number">
        <el-input
            :ref="
                (el: any) => {
                    nextTick(() => {
                        focus && el?.focus();
                    });
                }
            "
            :disabled="disabled"
            @blur="handleBlur"
            :class="`!w-full !mb-1 ${showEditorIcon ? 'string-input-container-show-icon' : ''}`"
            size="small"
            v-model="itemValue"
            :placeholder="placeholder ?? $t('common.pleaseInput')"
        />
        <SvgIcon v-if="showEditorIcon" @mousedown="openEditor" class="string-input-container-icon" name="FullScreen" :size="10" />
    </div>

    <el-date-picker
        v-else-if="dataType == DataType.Date"
        :ref="
            (el: any) => {
                nextTick(() => {
                    focus && el?.focus();
                });
            }
        "
        :disabled="disabled"
        @change="handleBlur"
        @blur="handleBlur"
        class="edit-time-picker !mb-1"
        popper-class="edit-time-picker-popper"
        size="small"
        v-model="itemValue"
        :clearable="false"
        type="date"
        value-format="YYYY-MM-DD"
        :placeholder="`date-${placeholder ?? $t('common.pleaseSelect')}`"
    />

    <el-date-picker
        v-else-if="dataType == DataType.DateTime"
        :ref="
            (el: any) => {
                nextTick(() => {
                    focus && el?.focus();
                });
            }
        "
        :disabled="disabled"
        @change="handleBlur"
        @blur="handleBlur"
        class="edit-time-picker !mb-1"
        popper-class="edit-time-picker-popper"
        size="small"
        v-model="itemValue"
        :clearable="false"
        type="datetime"
        value-format="YYYY-MM-DD HH:mm:ss"
        :placeholder="`datetime-${placeholder ?? $t('common.pleaseSelect')}`"
    />

    <el-time-picker
        v-else-if="dataType == DataType.Time"
        :ref="
            (el: any) => {
                nextTick(() => {
                    focus && el?.focus();
                });
            }
        "
        :disabled="disabled"
        @change="handleBlur"
        @blur="handleBlur"
        class="edit-time-picker !mb-1"
        popper-class="edit-time-picker-popper"
        size="small"
        v-model="itemValue"
        :clearable="false"
        value-format="HH:mm:ss"
        :placeholder="`time-${placeholder ?? $t('common.pleaseSelect')}`"
    />
</template>

<script lang="ts" setup>
import { computed, nextTick, ref, Ref } from 'vue';
import { ElInput, ElMessage } from 'element-plus';
import { DataType } from '../../dialect/index';
import SvgIcon from '@/components/svgIcon/index.vue';
import MonacoEditorBox from '@/components/monaco/MonacoEditorBox';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

export interface ColumnFormItemProps {
    modelValue: string | number; // 绑定的值
    dataType: DataType; // 数据类型
    focus?: boolean; // 是否获取焦点
    placeholder?: string;
    columnName?: string;
    disabled?: boolean;
}

const props = withDefaults(defineProps<ColumnFormItemProps>(), {
    focus: false,
    dataType: DataType.String,
    disabled: false,
});

const emit = defineEmits(['update:modelValue', 'blur']);

const itemValue: Ref<any> = ref(props.modelValue);

const showEditorIcon = computed(() => {
    return typeof itemValue.value === 'string' && itemValue.value.length > 50;
});

const editorOpening = ref(false);

const openEditor = () => {
    editorOpening.value = true;
    // 编辑器语言，如：json、html、text
    let editorLang = getEditorLangByValue(itemValue.value);
    MonacoEditorBox({
        content: itemValue.value,
        title: `${t('db.editField')} [${props.columnName}]`,
        language: editorLang,
        confirmFn: (newVal: any) => {
            itemValue.value = newVal;
            closeEditorDialog();
        },
        closeFn: closeEditorDialog,
    });
};

const closeEditorDialog = () => {
    editorOpening.value = false;
    handleBlur();
};

const handleBlur = () => {
    if (editorOpening.value) {
        return;
    }
    if (props.dataType == DataType.Number && itemValue.value && !/^-?\d*\.?\d+$/.test(itemValue.value)) {
        ElMessage.error(t('db.valueTypeNoMatch'));
        return;
    }
    emit('update:modelValue', itemValue.value);
    emit('blur');
};

const getEditorLangByValue = (value: any) => {
    // 判断是否是json
    try {
        if (typeof JSON.parse(value) === 'object') {
            return 'json';
        }
    } catch (e) {
        /* empty */
    }

    // 判断是否是html
    try {
        const doc = new DOMParser().parseFromString(value, 'text/html');
        if (Array.from(doc.body.childNodes).some((node) => node.nodeType === 1)) {
            return 'html';
        }
    } catch (e) {
        /* empty */
    }

    return 'text';
};
</script>

<style lang="scss">
.string-input-container {
    position: relative;

    .el-input__wrapper {
        padding: 1px 3px;
    }
}
.string-input-container-show-icon {
    .el-input__inner {
        padding-right: 10px;
    }
}
.string-input-container-icon {
    position: absolute;
    top: 5px; /* 调整图标的垂直位置 */
    right: 3px; /* 调整图标的水平位置 */
    color: var(--el-color-primary);
}
.string-input-container-icon:hover {
    color: var(--el-color-success);
}

.edit-time-picker {
    height: 26px;
    width: 100% !important;
    .el-input__prefix {
        display: none;
    }

    .el-input__wrapper {
        padding: 1px 3px;
    }
}

.edit-time-picker-popper {
    .el-date-picker {
        width: 250px !important;
        .el-date-picker__header {
            margin: 0 5px;
        }
        .el-picker-panel__content {
            width: unset;
            margin: 0 5px;
        }
        .el-date-picker__header-label {
            font-size: 13px;
        }
        .el-picker-panel__footer {
            padding: 0 5px;
            button {
                font-size: 11px;
                padding: 5px 6px;
                height: 20px;
            }
        }
    }
    .el-date-table {
        th {
            font-size: 10px;
            font-weight: 600;
            padding: 0;
        }
        td {
            padding: 0;
        }
    }
    .el-time-panel {
        width: 100px;

        .el-time-spinner__list {
            &::after,
            &::before {
                height: 10px;
            }
            .el-time-spinner__item {
                height: 20px;
                line-height: 20px;
            }
        }
    }
}
</style>
