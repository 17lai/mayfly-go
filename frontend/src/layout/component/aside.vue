<template>
    <el-aside class="layout-aside" :class="setCollapseWidth" v-if="clientWidth > 1000">
        <Logo v-if="setShowLogo" />
        <el-scrollbar class="flex-auto" ref="layoutAsideScrollbarRef">
            <Vertical :menuList="state.menuList" :class="setCollapseWidth" />
        </el-scrollbar>
    </el-aside>
    <el-drawer v-model="themeConfig.isCollapse" :with-header="false" direction="ltr" size="220px" v-else>
        <el-aside class="layout-aside !w-full !h-full">
            <Logo v-if="setShowLogo" />
            <el-scrollbar class="flex-auto" ref="layoutAsideScrollbarRef">
                <Vertical :menuList="state.menuList" />
            </el-scrollbar>
        </el-aside>
    </el-drawer>
</template>

<script lang="ts" setup name="layoutAside">
import { reactive, computed, watch, getCurrentInstance, onBeforeMount, onUnmounted, inject } from 'vue';
import pinia from '@/store/index';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '@/store/themeConfig';
import { useRoutesList } from '@/store/routesList';
import Logo from '@/layout/logo/index.vue';
import Vertical from '@/layout/navMenu/vertical.vue';
import { useWindowSize } from '@vueuse/core';

const { proxy } = getCurrentInstance() as any;

const { themeConfig } = storeToRefs(useThemeConfig());
const { routesList } = storeToRefs(useRoutesList());

const { width: clientWidth } = useWindowSize();

const state = reactive({
    menuList: [] as any[],
});

// 注入 菜单数据
const columnsMenuData: any = inject('columnsMenuData', null);
const classicMenuData: any = inject('classicMenuData', null);

// 设置菜单展开/收起时的宽度
const setCollapseWidth = computed(() => {
    let { layout, isCollapse, menuBar } = themeConfig.value;
    let asideBrColor = menuBar === '#FFFFFF' || menuBar === '#FFF' || menuBar === '#fff' || menuBar === '#ffffff' ? 'layout-el-aside-br-color' : '';
    if (layout === 'columns') {
        // 分栏布局，菜单收起时宽度给 1px
        if (isCollapse) {
            return ['layout-aside-width1', asideBrColor];
        } else {
            return ['layout-aside-width-default', asideBrColor];
        }
    } else {
        // 其它布局给 64px
        if (isCollapse) {
            return ['layout-aside-width64', asideBrColor];
        } else {
            return ['layout-aside-width-default', asideBrColor];
        }
    }
});

// 设置显示/隐藏 logo
const setShowLogo = computed(() => {
    let { layout, isShowLogo } = themeConfig.value;
    return (isShowLogo && layout === 'defaults') || (isShowLogo && layout === 'columns');
});

// 设置/过滤路由（非静态路由/是否显示在菜单中）
const setFilterRoutes = () => {
    if (themeConfig.value.layout === 'columns') {
        return false;
    }
    state.menuList = filterRoutesFun(routesList.value);
};

// 路由过滤递归函数
const filterRoutesFun = (arr: Array<object>) => {
    return arr
        .filter((item: any) => !item.meta.isHide)
        .map((item: any) => {
            item = Object.assign({}, item);
            if (item.children) item.children = filterRoutesFun(item.children);
            return item;
        });
};

// 监听 themeConfig 配置文件的变化，更新菜单 el-scrollbar 的高度
watch(themeConfig.value, (val) => {
    if (val.isShowLogoChange !== val.isShowLogo) {
        if (!proxy.$refs.layoutAsideScrollbarRef) {
            return false;
        }
        proxy.$refs.layoutAsideScrollbarRef.update();
    }
});

// 监听路由的变化，动态赋值给菜单中
watch(pinia.state, (val) => {
    if (val.routesList.routesList.length === state.menuList.length) {
        return false;
    }
    let { layout, isClassicSplitMenu } = val.themeConfig.themeConfig;
    if (layout === 'classic' && isClassicSplitMenu) {
        return;
    }
    setFilterRoutes();
});

// 监听经典布局分割菜单的变化
watch(
    () => themeConfig.value.isClassicSplitMenu,
    () => {
        // 当经典布局分割菜单选项变化时，重新设置过滤路由
        setFilterRoutes();
    }
);

// 页面加载前
onBeforeMount(() => {
    setFilterRoutes();

    if (columnsMenuData) {
        watch(columnsMenuData, (newVal) => {
            if (newVal) {
                state.menuList = newVal.children;
            }
        });
    }

    if (classicMenuData) {
        watch(classicMenuData, (newVal) => {
            let { layout, isClassicSplitMenu } = themeConfig.value;
            if (newVal && layout === 'classic' && isClassicSplitMenu) {
                state.menuList = [];
                state.menuList = newVal.children;
            }
        });
    }
});
</script>
