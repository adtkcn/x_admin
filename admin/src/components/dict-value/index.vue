<template>
    <template v-for="(item, index) in getOptions" :key="index">
        <span :style="{ color: item.color }"
            >{{ index != 0 ? '、' : '' }}{{ item[props.labelKey] }}</span
        >
    </template>
</template>
<script lang="ts" setup>
import { computed } from 'vue'
defineOptions({
    name: 'dict-value'
})
const props = withDefaults(
    defineProps<{
        options: any[]
        value: any
        labelKey?: string
        valueKey?: string
    }>(),
    {
        options: () => [],
        value: null,
        labelKey: 'name',
        valueKey: 'value'
    }
)

const values = computed(() => {
    if (props.value !== null && typeof props.value !== 'undefined') {
        return Array.isArray(props.value) ? props.value : String(props.value).split(',')
    } else {
        return []
    }
})

const getOptions = computed(() => {
    return props.options.filter((item) => values.value.includes(item[props.valueKey]))
})
</script>
