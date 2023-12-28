<template>
    <i class="el-icon" :style="styles">
        <svg aria-hidden="true" style="width: 1em; height: 1em">
            <use :xlink:href="symbolId" fill="currentColor" />
        </svg>
    </i>
</template>

<script lang="ts">
import { addUnit } from '@/utils/util'
import type { CSSProperties } from 'vue'

export default defineComponent({
    name: 'SvgIcon',
    props: {
        name: {
            type: String,
            required: true
        },
        size: {
            type: [Number, String],
            default: 16
        },
        color: {
            type: String,
            default: 'inherit'
        }
    },
    setup(props) {
        const symbolId = computed(() => `#${props.name}`)
        const styles = computed<CSSProperties>(() => {
            return {
                fontSize: addUnit(props.size),
                color: props.color
            }
        })
        return { symbolId, styles }
    }
})
</script>
