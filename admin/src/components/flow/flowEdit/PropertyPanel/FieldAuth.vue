<template>
    <div style="padding-bottom: 10px">
        <el-card header="表单项权限">
            <el-table
                v-if="['bpmn:startEvent', 'bpmn:userTask'].includes(props.node.type)"
                fit
                size="small"
                :data="props.fieldList"
                :show-header="false"
                style="width: 100%"
            >
                <el-table-column prop="label" label=""></el-table-column>
                <el-table-column label="">
                    <template #default="{ row }">
                        <el-radio-group v-model="row.auth">
                            <el-radio :label="1">读写</el-radio>
                            <el-radio :label="2">可读</el-radio>
                            <el-radio :label="3">隐藏</el-radio>
                        </el-radio-group>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
    </div>
</template>

<script setup>
const props = defineProps({
    node: {
        type: Object,
        default: () => ({})
    },
    fieldList: {
        type: Array,
        default: () => []
    }
})
</script>

<style lang="scss" scoped></style>
