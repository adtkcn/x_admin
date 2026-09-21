<template>
    <div style="padding-bottom: 10px">
        <el-card header="表单项权限">
            <!-- 仅在 start_event / user_task 节点渲染（由父组件控制），本组件只负责 fieldList 权限编辑 -->
            <vxe-table
                size="small"
                :data="props.fieldList"
                :show-header="false"
                :row-config="{ keyField: 'name' }"
                :border="'inner'"
                style="width: 100%"
            >
                <vxe-column field="name" title=""></vxe-column>
                <vxe-column title="">
                    <template #default="{ row }">
                        <el-radio-group v-model="row.auth">
                            <el-radio :value="1">读写</el-radio>
                            <el-radio :value="2">可读</el-radio>
                            <el-radio :value="3">隐藏</el-radio>
                        </el-radio-group>
                    </template>
                </vxe-column>
            </vxe-table>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import type { FieldListType } from './property.type'
const props = withDefaults(
    defineProps<{
        fieldList?: FieldListType[]
    }>(),
    {
        fieldList: () => []
    }
)
</script>

<style lang="scss" scoped></style>
