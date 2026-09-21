<template>
    <el-date-picker
        class="x-date-picker"
        v-model="content"
        :type="props.type"
        range-separator="-"
        :value-format="props.valueFormat"
        start-placeholder="开始时间"
        end-placeholder="结束时间"
        clearable
        :default-time="props.defaultTime"
        @change="changeDate"
    ></el-date-picker>
</template>

<script lang="ts" setup>
import { computed, ref, watch } from 'vue'
type Props = {
    type?: 'datetimerange' | 'daterange'
    startTime?: string
    endTime?: string
    defaultTime?: [Date, Date]
    valueFormat?: string
}
const props = withDefaults(defineProps<Props>(), {
    type: 'datetimerange',
    startTime: '',
    endTime: '',
    defaultTime: () =>
        [new Date(2000, 1, 1, 0, 0, 0), new Date(2000, 2, 1, 23, 59, 59)] as [Date, Date],
    valueFormat: 'YYYY-MM-DD HH:mm:ss'
})

const emit = defineEmits(['update:startTime', 'update:endTime'])
const content = ref<[string, string]>([props.startTime, props.endTime])

function changeDate(value: any) {
    console.log('change', value)
    if (value === null) {
        emit('update:startTime', '')
        emit('update:endTime', '')
    } else {
        emit('update:startTime', value[0])
        emit('update:endTime', value[1])
    }
}

watch([() => props.startTime, () => props.endTime], () => {
    console.log('watch', props)
    content.value = [props.startTime, props.endTime]
})
</script>
<style lang="scss">
.x-date-picker {
    & > .el-range__icon:first-child {
        display: none !important;
    }
    & > .el-range-input {
        width: 45%;
    }
}
</style>
