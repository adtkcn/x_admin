<template>
    <div class="color-picker flex flex-1">
        <el-color-picker v-model="color" :predefine="predefineColors" />
        <el-input v-model="color" class="mx-[10px] flex-1" type="text" readonly />
        <el-button type="text" @click="reset">重置</el-button>
    </div>
</template>
<script lang="ts" setup>
const props = defineProps({
    modelValue: {
        type: String
    },
    defaultColor: {
        type: String
    }
})

const color = defineModel({
    type: String,
    default: () => {
        return props.defaultColor
    }
})

const predefineColors = ['#409EFF', '#28C76F', '#EA5455', '#FF9F43', '#01CFE8', '#4A5DFF']
const reset = () => {
    color.value = props.defaultColor
}
</script>
