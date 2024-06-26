<template>
    <div>
        <el-descriptions :column="3" border>
            <el-descriptions-item :span="3" label="标签"><TagCodePath :path="redis.codePaths" /></el-descriptions-item>

            <el-descriptions-item :span="2" label="编号">{{ redis?.code }}</el-descriptions-item>
            <el-descriptions-item :span="1" label="名称">{{ redis?.name }}</el-descriptions-item>

            <el-descriptions-item :span="1" label="主机">{{ `${redis?.host}` }}</el-descriptions-item>
            <el-descriptions-item :span="1" label="库">{{ state.db }}</el-descriptions-item>
            <el-descriptions-item :span="1" label="mode">
                {{ redis.mode }}
            </el-descriptions-item>

            <el-descriptions-item :span="3" label="执行Cmd">
                <el-input type="textarea" disabled v-model="cmd" rows="5" />
            </el-descriptions-item>
        </el-descriptions>
    </div>
</template>

<script lang="ts" setup>
import { toRefs, reactive, watch, onMounted } from 'vue';
import { redisApi } from '@/views/ops/redis/api';
import TagCodePath from '@/views/ops/component/TagCodePath.vue';
import { tagApi } from '@/views/ops/tag/api';
import { TagResourceTypeEnum } from '@/common/commonEnum';

const props = defineProps({
    // 业务表单
    bizForm: {
        type: [String],
        default: '',
    },
});

const state = reactive({
    cmd: '',
    db: 0,
    redis: {} as any,
});

const { cmd, redis } = toRefs(state);

onMounted(() => {
    parseRunCmdForm(props.bizForm);
});

watch(
    () => props.bizForm,
    (newValue: any) => {
        parseRunCmdForm(newValue);
    }
);

const parseRunCmdForm = async (bizForm: string) => {
    if (!bizForm) {
        return;
    }
    const form = JSON.parse(bizForm);

    const cmds = form.cmd.map((item: any, index: number) => {
        if (index === 0) {
            return item; // 第一个元素直接返回原值
        }
        if (typeof item === 'string') {
            return `'${item}'`; // 字符串加单引号
        }
        return item; // 其他类型直接返回
    });
    state.cmd = cmds.join('  ');
    state.db = form.db;

    const res = await redisApi.redisList.request({ id: form.id });
    if (!res.list) {
        return;
    }
    state.redis = res.list?.[0];

    tagApi.listByQuery.request({ type: TagResourceTypeEnum.Redis.value, codes: state.redis.code }).then((res) => {
        state.redis.codePaths = res.map((item: any) => item.codePath);
    });
};
</script>
<style lang="scss"></style>
