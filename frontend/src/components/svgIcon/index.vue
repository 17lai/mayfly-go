<template>
    <i v-if="isShowIconSvg" class="el-icon icon-middle" :style="setIconSvgStyle">
        <component :is="getIconName" :style="setIconSvgStyle" />
    </i>

    <svg v-else-if="isLocalIcon()" class="el-icon local-icon icon-middle" aria-hidden="true" :style="setIconSvgStyle">
        <use :xlink:href="'#' + getLocalIconName()"></use>
    </svg>

    <div v-else-if="isShowIconImg" :style="setIconImgOutStyle">
        <img :src="getIconName" :style="setIconSvgInsStyle" />
    </div>

    <i v-else :class="getIconName" :style="setIconSvgStyle" />
</template>

<script setup lang="ts" name="svgIcon">
import { computed } from 'vue';

// 定义父组件传过来的值
const props = defineProps({
    // svg 图标组件名字
    name: {
        type: String,
    },
    // svg 大小
    size: {
        type: Number,
        default: () => 14,
    },
    // svg 颜色
    color: {
        type: String,
    },
    isEle: {
        type: Boolean,
        default: true,
    },
});

// 在线链接、本地引入地址前缀
const linesString = ['https', 'http', '/src', '/assets', 'icon ', 'data:image', import.meta.env.VITE_PUBLIC_PATH];

// 获取 icon 图标名称
const getIconName = computed(() => {
    // if (props.name?.startsWith('icon ')) {
    //     return getIcon(props?.name?.split(' ')[1] as any);
    // }
    return props?.name as any;
});

// 用于判断 element plus 自带 svg 图标的显示、隐藏。不存在 空格分隔的icon name即为element plus自带icon
const isShowIconSvg = computed(() => {
    const ss = props?.name?.split(' ');
    if (!ss) {
        return true;
    }
    return ss.length == 1;
});

const isLocalIcon = () => {
    return props?.name?.startsWith('icon ');
};

const getLocalIconName = () => {
    // icon icon-xxxx 获取icon-xxx即可
    return props?.name?.split(' ')[1];
};

// 用于判断在线链接、本地引入等图标显示、隐藏
const isShowIconImg = computed(() => {
    return linesString.find((str) => props.name?.startsWith(str));
});

// 设置图标样式
const setIconSvgStyle = computed(() => {
    return `font-size: ${props.size}px;color: ${props.color};width: ${props.size}px;height: ${props.size}px;`;
});

// 设置图片样式
const setIconImgOutStyle = computed(() => {
    return `width: ${props.size}px;height: ${props.size}px;display: inline-block;overflow: hidden;line-height:${props.size}px;vertical-align: middle;`;
});

// 设置图片样式
const setIconSvgInsStyle = computed(() => {
    if (props.color) {
        const filterStyle: string[] = [];
        const compatibles: string[] = ['-webkit', '-ms', '-o', '-moz'];
        compatibles.forEach((j) => filterStyle.push(`${j}-filter: drop-shadow(${props.color} ${props.size}px 0);`));
        return `width: ${props.size}px;height: ${props.size}px;position: relative;left: -${props.size}px;${filterStyle.join('')}`;
    }

    return `width: ${props.size}px;height: ${props.size}px;position: relative;`;
});
</script>

<style type="text/css" scoped>
.local-icon {
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
}

.icon-middle {
    display: inline-flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    vertical-align: middle;
    height: 100%; /* 确保高度与父元素一致 */
    line-height: 1; /* 确保行高与高度一致 */
}
</style>
