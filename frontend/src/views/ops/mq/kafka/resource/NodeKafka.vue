<template>
    <BaseTreeNode v-bind="$attrs">
        <template #prefix="{ data }">
            <el-popover :show-after="500" placement="right-start" :title="$t('common.detail')" trigger="hover" :width="250">
                <template #reference>
                    <SvgIcon :name="ResourceTypeEnum.MqKafka.extra.icon" :color="ResourceTypeEnum.MqKafka.extra.iconColor" :size="13" />
                </template>

                <template #default>
                    <el-descriptions :column="1" size="small">
                        <el-descriptions-item :label="$t('common.name')">
                            {{ data.params.name }}
                        </el-descriptions-item>
                        <el-descriptions-item label="hosts">
                            {{ data.params.hosts }}
                        </el-descriptions-item>
                        <el-descriptions-item :label="$t('common.remark')" label-align="right">
                            {{ data.params.remark }}
                        </el-descriptions-item>
                    </el-descriptions>
                </template>
            </el-popover>
        </template>
    </BaseTreeNode>
</template>

<script lang="ts" setup>
import { ResourceTypeEnum } from '@/common/commonEnum';
import BaseTreeNode from '@/views/ops/resource/BaseTreeNode.vue';
</script>

<style lang="scss"></style>
