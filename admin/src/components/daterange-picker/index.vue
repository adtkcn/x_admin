<template>
    <el-date-picker
        v-model="content"
        type="datetimerange"
        range-separator="-"
        value-format="YYYY-MM-DD HH:mm:ss"
        start-placeholder="开始时间"
        end-placeholder="结束时间"
        clearable
        :default-time="defaultTime"
        @change="changeDate"
    ></el-date-picker>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
/* Props S */
const props = defineProps({
    startTime: {
        type: String,
        default: ''
    },
    endTime: {
        type: String,
        default: ''
    }
})
const defaultTime: [Date, Date] = [new Date(2000, 1, 1, 0, 0, 0), new Date(2000, 2, 1, 23, 59, 59)] // '12:00:00', '08:00:00'

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
