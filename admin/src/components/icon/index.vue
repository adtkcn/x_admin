<template>
    <div class="svg-icon-container">
        <ElIcon v-if="isElIcon" :size="size" :color="color">
            <component :is="name" style="width: 1em; height: 1em"></component>
        </ElIcon>
        <!-- </span> -->
        <ISvgIcon
            v-else-if="isLocalIcon"
            class="local-icon"
            :size="size"
            :color="color"
            :name="name"
        ></ISvgIcon>
    </div>
</template>
<script lang="ts">
// import { ElIcon } from 'element-plus'
import { EL_ICON_PREFIX, LOCAL_ICON_PREFIX } from './index'
import ISvgIcon from './svg-icon.vue'
export default defineComponent({
    name: 'Icon',
    components: {
        // ElIcon,
        ISvgIcon
    },
    props: {
        name: {
            type: String,
            required: true
        },
        size: {
            type: [String, Number],
            default: '14px'
        },
        color: {
            type: String,
            default: 'inherit'
        }
    },
    setup(props) {
        // const isElIcon = ref(false)
        const isLocalIcon = computed(() => {
            return props.name.indexOf(LOCAL_ICON_PREFIX) === 0
        })

        const isElIcon = computed(() => {
            return props.name.indexOf(EL_ICON_PREFIX) === 0
        })

        return {
            isElIcon,
            isLocalIcon
        }
    }
})
</script>
<style>
.svg-icon-container {
    display: inline-block;
}
</style>
