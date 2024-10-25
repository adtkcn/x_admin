<template>
    <div class="logo">
        <image-contain :width="size" :height="size" :src="config.webLogo" />
        <transition name="title-width">
            <div
                v-show="showTitle"
                class="logo-title overflow-hidden whitespace-nowrap"
                :class="{ 'text-white': theme == ThemeEnum.DARK }"
                :style="{ left: `${size + 16}px` }"
            >
                <overflow-tooltip
                    :content="title || config.webName"
                    :teleported="true"
                    placement="bottom"
                    overflow-type="unset"
                >
                </overflow-tooltip>
            </div>
        </transition>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import useAppStore from '@/stores/modules/app'
import { ThemeEnum } from '@/enums/appEnums'
defineOptions({
    name: 'SideLogo'
})
defineProps({
    size: { type: Number, default: 34 },
    title: { type: String },
    theme: { type: String },
    showTitle: { type: Boolean, default: true }
})
const appStore = useAppStore()
const config = computed(() => appStore.config)
</script>
<style lang="scss" scoped>
.logo {
    height: var(--navbar-height);
    overflow: hidden;
    box-sizing: border-box;
    position: relative;
    display: flex;
    align-items: center;
    padding: 8px;

    .logo-title {
        width: 70%;
        position: absolute;
        font-size: var(--el-font-size-large);
    }

    .title-width-enter-active {
        opacity: 0;
        transition: all 0.3s ease-out;
    }

    .title-width-leave-active {
        transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
    }

    .title-width-enter-from,
    .title-width-leave-to {
        width: 0;
        opacity: 0;
    }
}
</style>
