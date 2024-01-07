<template>
    <div style="padding-bottom: 10px">
        <el-card header="条件编辑">
            <el-alert title="同一父级的网关只能有一个通过" type="warning" />

            <!-- 设置优先级 -->
            <div style="padding: 40px 0 20px">
                <el-select v-model="selectGateway" placeholder="请选择">
                    <el-option
                        v-for="item in fieldList"
                        :key="item.id"
                        :label="item.label"
                        :value="item.id"
                    />
                </el-select>
                <el-button type="primary" style="margin-left: 10px" @click="addCondition"
                    >添加条件</el-button
                >
            </div>

            <el-table fit size="small" :data="props.properties.gateway" style="width: 100%">
                <el-table-column prop="label" label="表单项">
                    <template #default="{ row }">
                        {{ getLabel(row.id) }}
                    </template>
                </el-table-column>
                <el-table-column label="权限">
                    <template #default="{ row }">
                        <el-select v-model="row.condition" placeholder="请选择判断符">
                            <el-option
                                v-for="item in conditionList"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            />
                        </el-select>
                    </template>
                </el-table-column>
                <el-table-column label="值">
                    <template #default="{ row }">
                        <el-input v-model="row.value" placeholder="请输入"></el-input>
                    </template>
                </el-table-column>
                <el-table-column width="50px">
                    <template #default="{ row, $index }">
                        <el-button :icon="Close" circle @click="removeCondition(row, $index)" />
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { Close } from '@element-plus/icons-vue'
const props = defineProps({
    node: {
        type: Object,
        default: () => ({})
    },
    fieldList: {
        type: Array,
        default: () => []
    },
    properties: {
        type: Object,
        default: () => ({})
    }
})
const conditionList = [
    {
        value: '==',
        label: '等于'
    },
    {
        value: '!=',
        label: '不等于'
    },
    {
        value: '>=',
        label: '大于等于'
    },
    {
        value: '<=',
        label: '小于等于'
    },
    {
        value: 'include',
        label: '包含'
    }
]
const selectGateway = ref('')
function getLabel(id) {
    return props.fieldList.find((item) => {
        if (item.id === id) {
            return item.label
        }
    })?.label
}
function addCondition() {
    // this.selectGateway

    props.fieldList.find((item) => {
        if (item.id === selectGateway.value) {
            props.properties.gateway.push({
                id: item.id,
                value: '',
                condition: ''
            })
        }
    })
}
function removeCondition(row, $index) {
    props.properties.gateway.splice($index, 1)
}
</script>

<style lang="scss" scoped></style>
