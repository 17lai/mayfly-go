<template>
    <el-tabs v-model="activeTabName">
        <!-- <el-tab-pane :label="$t('common.basic')" :name="basic"> </el-tab-pane> -->

        <!-- <el-tab-pane :label="$t('')"> </el-tab-pane> -->
    </el-tabs>
</template>
<script lang="ts" setup>
import { ref } from 'vue';

const basic = 'basic';

const activeTabName = ref(basic);

const form = defineModel<any>('modelValue', { required: true });
</script>
