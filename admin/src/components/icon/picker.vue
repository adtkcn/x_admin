<template>
    <div class="icon-select">
        <el-popover trigger="click" :width="500">
            <div>
                <div>
                    <div class="flex justify-between">
                        <div class="mb-3">请选择图标</div>
                        <div>
                            <span
                                v-for="(item, index) in iconTabsMap"
                                :key="index"
                                class="cursor-pointer text-sm ml-2"
                                :class="{
                                    'text-primary': index == tabIndex
                                }"
                                @click="tabIndex = index"
                            >
                                {{ item.name }}
                            </span>
                        </div>
                    </div>

                    <div class="h-[280px]">
                        <el-scrollbar>
                            <div class="flex flex-wrap">
                                <div v-for="item in iconNamesFilter" :key="item" class="m-1">
                                    <el-button @click="handleSelect(item)">
                                        <icon :name="item" :size="18" />
                                    </el-button>
                                </div>
                            </div>
                        </el-scrollbar>
                    </div>
                </div>
            </div>
            <template #reference>
                <el-input
                    ref="inputRef"
                    v-model.trim="state.inputValue"
                    placeholder="搜索图标"
                    :disabled="disabled"
                    clearable
                >
                    <template #prepend>
                        <div class="flex items-center" v-if="modelValue">
                            <el-tooltip class="flex-1 w-20" :content="modelValue" placement="top">
                                <icon
                                    class="mr-1"
                                    :key="modelValue"
                                    :name="modelValue"
                                    :size="16"
                                />
                            </el-tooltip>
                        </div>

                        <template v-else>无</template>
                    </template>
                    <template #append>
                        <el-button @click.stop="handleClear">
                            <icon name="el-icon-Close" :size="18" />
                        </el-button>
                    </template>
                </el-input>
            </template>
        </el-popover>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { computed, reactive, shallowRef } from 'vue'

import { ElInput } from 'element-plus'
import { getElementPlusIconNames, getLocalIconNames } from './index'
interface Props {
    modelValue: string
    disabled?: boolean
}
withDefaults(defineProps<Props>(), {
    modelValue: '',
    disabled: false
})

const emits = defineEmits<{
    (e: 'update:modelValue', value: string): void
    (e: 'change', value: string): void
}>()

const tabIndex = ref(0)
const iconTabsMap = [
    {
        name: 'element图标',
        icons: getElementPlusIconNames()
    },
    {
        name: '本地图标',
        icons: getLocalIconNames()
    }
]

const inputRef = shallowRef<InstanceType<typeof ElInput>>()

const state = reactive({
    inputValue: ''
})
// 选中图标
const handleSelect = (icon: string) => {
    emits('update:modelValue', icon)
    emits('change', icon)
}
//取消选中
const handleClear = () => {
    emits('update:modelValue', '')
    emits('change', '')
}

//根据输入框内容塞选
const iconNamesFilter = computed(() => {
    const iconNames = iconTabsMap[tabIndex.value]?.icons ?? []
    if (!state.inputValue) {
        return iconNames
    }
    const inputValue = state.inputValue.toLowerCase()
    return iconNames.filter((icon: string) => {
        if (icon.toLowerCase().indexOf(inputValue) !== -1) {
            return icon
        }
    })
})
</script>
