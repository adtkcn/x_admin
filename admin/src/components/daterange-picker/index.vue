<template>
    <el-date-picker
        v-model="content"
        type="daterange"
        range-separator="-"
        format="YYYY-MM-DD"
        value-format="YYYY-MM-DD"
        start-placeholder="开始时间"
        end-placeholder="结束时间"
        clearable
        @change="changeDate"
    ></el-date-picker>
</template>

<script lang="ts" setup>
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
// const content = computed<any>({
//     get: () => {
//         return [props.startTime, props.endTime]
//     },
//     set: (value: Event | any) => {
//         console.log('change', value)

//         if (value === null) {
//             emit('update:startTime', '')
//             emit('update:endTime', '')
//         } else {
//             emit('update:startTime', value[0])
//             emit('update:endTime', value[1])
//         }
//     }
// })
</script>
