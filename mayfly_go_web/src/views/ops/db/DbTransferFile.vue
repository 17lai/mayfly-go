<template>
    <div class="db-transfer-file">
        <el-dialog :title="title" v-model="dialogVisible" :close-on-click-modal="false" :destroy-on-close="true" width="1000px">
            <page-table
                ref="pageTableRef"
                :data="state.tableData"
                v-model:query-form="state.query"
                :show-selection="true"
                v-model:selection-data="state.selectionData"
                :columns="columns"
                @page-num-change="
                    (args) => {
                        state.query.pageNum = args.pageNum;
                        search();
                    }
                "
                @page-size-change="
                    (args) => {
                        state.query.pageSize = args.pageNum;
                        search();
                    }
                "
            >
                <template #tableHeader>
                    <el-button v-auth="perms.del" :disabled="state.selectionData.length < 1" @click="del()" type="danger" icon="delete">删除</el-button>
                </template>

                <template #fileKey="{ data }">
                    <FileInfo :fileKey="data.fileKey" :canDownload="actionBtns[perms.down] && data.status === 2" />
                </template>

                <template #fileDbType="{ data }">
                    <span>
                        <SvgIcon :name="getDbDialect(data.fileDbType).getInfo().icon" :size="18" />
                        {{ data.fileDbType }}
                    </span>
                </template>

                <template #action="{ data }">
                    <el-button v-if="actionBtns[perms.run] && data.status === DbTransferFileStatusEnum.Success.value" @click="openRun(data)" type="primary" link
                        >执行</el-button
                    >
                    <el-button v-if="data.logId" @click="openLog(data)" type="success" link>日志</el-button>
                </template>
            </page-table>
            <TerminalLog v-model:log-id="state.logsDialog.logId" v-model:visible="state.logsDialog.visible" :title="state.logsDialog.title" />
        </el-dialog>

        <el-dialog :title="state.runDialog.title" v-model="state.runDialog.visible" :destroy-on-close="true" width="600px">
            <el-form :model="state.runDialog.runForm" ref="runFormRef" label-width="auto" :rules="state.runDialog.formRules">
                <el-form-item label="文件数据库类型" prop="dbType">
                    <SvgIcon :name="getDbDialect(state.runDialog.runForm.dbType).getInfo().icon" :size="18" /> {{ state.runDialog.runForm.dbType }}
                </el-form-item>
                <el-form-item label="选择目标数据库" prop="targetDbId" required>
                    <db-select-tree
                        placeholder="请选择目标数据库"
                        v-model:db-id="state.runDialog.runForm.targetDbId"
                        v-model:inst-name="state.runDialog.runForm.targetInstName"
                        v-model:db-name="state.runDialog.runForm.targetDbName"
                        v-model:tag-path="state.runDialog.runForm.targetTagPath"
                        v-model:db-type="state.runDialog.runForm.targetDbType"
                        @select-db="state.runDialog.onSelectRunTargetDb"
                    />
                </el-form-item>
            </el-form>

            <template #footer>
                <div>
                    <el-button @click="state.runDialog.cancel()">取 消</el-button>
                    <el-button type="primary" :loading="state.runDialog.loading" @click="state.runDialog.btnOk">确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, reactive, Ref, ref, watch } from 'vue';
import { dbApi } from '@/views/ops/db/api';
import { getDbDialect } from '@/views/ops/db/dialect';
import PageTable from '@/components/pagetable/PageTable.vue';
import { TableColumn } from '@/components/pagetable';
import { ElMessage, ElMessageBox } from 'element-plus';
import { hasPerms } from '@/components/auth/auth';
import TerminalLog from '@/components/terminal/TerminalLog.vue';
import DbSelectTree from '@/views/ops/db/component/DbSelectTree.vue';
import { getClientId } from '@/common/utils/storage';
import FileInfo from '@/components/file/FileInfo.vue';
import { DbTransferFileStatusEnum } from './enums';
const props = defineProps({
    data: {
        type: [Object],
    },
    title: {
        type: String,
    },
});

const dialogVisible = defineModel<boolean>('visible', { default: false });

const columns = ref([
    TableColumn.new('fileKey', '文件').setMinWidth(280).isSlot(),
    TableColumn.new('createTime', '执行时间').setMinWidth(180).isTime(),
    TableColumn.new('fileDbType', 'sql语言').setMinWidth(90).isSlot(),
    TableColumn.new('status', '状态').typeTag(DbTransferFileStatusEnum),
]);

const perms = {
    del: 'db:transfer:files:del',
    down: 'db:transfer:files:down',
    run: 'db:transfer:files:run',
};

const actionBtns = hasPerms([perms.del, perms.down, perms.run]);

const actionWidth = ((actionBtns[perms.run] ? 1 : 0) + 1) * 55;

const actionColumn = TableColumn.new('action', '操作').isSlot().setMinWidth(actionWidth).fixedRight().alignCenter();

onMounted(async () => {
    if (Object.keys(actionBtns).length > 0) {
        columns.value.push(actionColumn);
    }
});

const runFormRef: any = ref(null);

const state = reactive({
    query: {
        taskId: props.data?.id,
        name: null,
        pageNum: 1,
        pageSize: 10,
    },
    logsDialog: {
        logId: 0,
        title: '数据库迁移日志',
        visible: false,
        data: null as any,
        running: false,
    },
    runDialog: {
        title: '指定数据库执行sql文件',
        visible: false,
        data: null as any,
        formRules: {
            targetDbId: [
                {
                    required: true,
                    message: '请选择目标数据库',
                    trigger: ['change', 'blur'],
                },
            ],
        },
        runForm: {
            id: 0,
            dbType: '',
            clientId: '',
            targetDbId: 0,
            targetDbName: '',
            targetTagPath: '',
            targetInstName: '',
            targetDbType: '',
        },
        loading: false,
        cancel: function () {
            state.runDialog.visible = false;
            state.runDialog.runForm = {} as any;
        },
        btnOk: function () {
            runFormRef.value.validate(async (valid: boolean) => {
                if (!valid) {
                    ElMessage.error('请正确填写信息');
                    return false;
                }
                console.log(state.runDialog.runForm);
                if (state.runDialog.runForm.targetDbType !== state.runDialog.runForm.dbType) {
                    ElMessage.warning(`请选择[${state.runDialog.runForm.dbType}]数据库`);
                    return false;
                }
                state.runDialog.runForm.clientId = getClientId();
                await dbApi.dbTransferFileRun.request(state.runDialog.runForm);
                ElMessage.success('保存成功');
                state.runDialog.cancel();
                await search();
            });
        },
        onSelectRunTargetDb: function (param: any) {
            if (param.type !== state.runDialog.runForm.dbType) {
                ElMessage.warning(`请选择[${state.runDialog.runForm.dbType}]数据库`);
            }
        },
    },
    selectionData: [], // 选中的数据
    tableData: [],
});

const search = async () => {
    const { total, list } = await dbApi.dbTransferFileList.request(state.query);
    state.tableData = list;
    pageTableRef.value.total = total;
};

const pageTableRef: Ref<any> = ref(null);

const del = async function () {
    try {
        await ElMessageBox.confirm(`将会删除sql文件，确定删除?`, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        });
        await dbApi.dbTransferFileDel.request({ fileId: state.selectionData.map((x: any) => x.id).join(',') });
        ElMessage.success('删除成功');
        await search();
    } catch (err) {
        //
    }
};

const openLog = function (data: any) {
    state.logsDialog.logId = data.logId;
    state.logsDialog.visible = true;
    state.logsDialog.title = '数据库迁移日志';
    state.logsDialog.running = data.state === 1;
};

// 运行sql，弹出选择需要运行的库，默认运行当前数据库，需要保证数据库类型与sql文件一致
const openRun = function (data: any) {
    console.log(data);
    state.runDialog.runForm = { id: data.id, dbType: data.fileDbType } as any;
    console.log(state.runDialog.runForm);
    state.runDialog.visible = true;
};

watch(dialogVisible, async (newValue: boolean) => {
    if (!newValue) {
        return;
    }
    state.query.taskId = props.data?.id;
    state.query.pageNum = 1;
    state.query.pageSize = 10;

    await search();
});
</script>
<style lang="scss"></style>
