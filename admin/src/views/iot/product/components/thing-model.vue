<template>
    <el-dialog
        v-model="visible"
        title="物模型设计"
        width="1080px"
        append-to-body
        draggable
        destroy-on-close
        top="5px"
        @closed="handleClosed"
    >
        <div class="thing-model">
            <el-alert
                v-if="structureError"
                class="thing-model__alert"
                type="warning"
                show-icon
                :closable="false"
                :title="structureError"
            />
            <div class="thing-model__bar">
                <el-radio-group v-model="activeTab" size="small" @change="handleTabChange">
                    <el-radio-button value="table">属性列表</el-radio-button>
                    <el-radio-button value="json">JSON 源码</el-radio-button>
                </el-radio-group>
                <template v-if="activeTab === 'table'">
                    <el-button type="primary" plain size="small" @click="handleAdd">
                        新增属性
                    </el-button>
                    <span class="thing-model__tips">
                        属性标识即设备上报的字段名；未在此声明的属性，处理层按白名单丢弃
                    </span>
                </template>
                <template v-else>
                    <el-button size="small" @click="handleFormat">格式化</el-button>
                    <span class="thing-model__tips">JSON 模式按原文保存，可手工编写扩展字段</span>
                </template>
            </div>

            <el-table
                v-if="activeTab === 'table'"
                :data="properties"
                size="small"
                border
                max-height="440"
                empty-text="尚未配置属性，点击上方「新增属性」开始"
            >
                <el-table-column label="属性标识" min-width="150">
                    <template #default="{ row }">
                        <el-input v-model="row.identifier" placeholder="如 temperature" clearable />
                    </template>
                </el-table-column>
                <el-table-column label="显示名称" min-width="130">
                    <template #default="{ row }">
                        <el-input v-model="row.name" placeholder="留空则同标识" clearable />
                    </template>
                </el-table-column>
                <el-table-column label="数据类型" width="120">
                    <template #default="{ row }">
                        <el-select v-model="row.dataType" class="w-full">
                            <el-option
                                v-for="opt in ThingDataTypeOptions"
                                :key="opt.value"
                                :label="opt.label"
                                :value="opt.value"
                            />
                        </el-select>
                    </template>
                </el-table-column>
                <el-table-column label="单位" width="90">
                    <template #default="{ row }">
                        <el-input v-model="row.unit" placeholder="如 ℃" clearable />
                    </template>
                </el-table-column>
                <el-table-column label="取值范围" width="180">
                    <template #default="{ row }">
                        <div v-if="isNumericType(row.dataType)" class="thing-model__range">
                            <el-input-number
                                v-model="row.min"
                                :controls="false"
                                placeholder="最小"
                                class="flex-1"
                            />
                            <span>~</span>
                            <el-input-number
                                v-model="row.max"
                                :controls="false"
                                placeholder="最大"
                                class="flex-1"
                            />
                        </div>
                        <span v-else class="thing-model__muted">—</span>
                    </template>
                </el-table-column>
                <el-table-column label="描述" min-width="140">
                    <template #default="{ row }">
                        <el-input v-model="row.desc" placeholder="选填" clearable />
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="150" align="center">
                    <template #default="{ $index }">
                        <el-button link type="primary" :disabled="$index === 0" @click="handleMove($index, -1)">
                            上移
                        </el-button>
                        <el-button
                            link
                            type="primary"
                            :disabled="$index === properties.length - 1"
                            @click="handleMove($index, 1)"
                        >
                            下移
                        </el-button>
                        <el-button link type="danger" @click="handleRemove($index)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-input
                v-else
                v-model="jsonText"
                type="textarea"
                :rows="18"
                spellcheck="false"
                placeholder='[ { "identifier": "temperature", "name": "温度", "dataType": "float" } ]'
            />
        </div>
        <template #footer>
            <el-button @click="visible = false">取消</el-button>
            <el-button type="primary" @click="handleConfirm">保存</el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import {
    ThingDataTypeOptions,
    emptyProperty,
    isNumericType,
    parseThingModel,
    serializeThingModel,
    validateProperties,
    type ThingProperty
} from './thing-model'

// v-model 绑定产品物模型 JSON 原文（空串=未物模型化）
const props = defineProps<{ modelValue?: string }>()

const emit = defineEmits<{
    (e: 'update:modelValue', value: string): void
    (e: 'confirm', value: string): void
}>()

const visible = ref(false)
const activeTab = ref<'table' | 'json'>('table')
const properties = ref<ThingProperty[]>([])
const jsonText = ref('')
// 原文能解析成 JSON 但认不出属性结构时置此值：不丢用户数据，退到 JSON 模式手工改
const structureError = ref('')

// syncFromTable 把当前属性行渲染进 JSON 文本框
const syncFromTable = () => {
    jsonText.value = serializeThingModel(properties.value)
}

const handleAdd = () => {
    properties.value.push(emptyProperty())
    nextTick(() => {
        const body = document.querySelector('.thing-model .el-table__body-wrapper')
        if (body) body.scrollTop = body.scrollHeight
    })
}

const handleRemove = (index: number) => {
    properties.value.splice(index, 1)
}

// handleMove 先割后插：不重建整个数组，避开 el-table 整表重渲染
const handleMove = (index: number, offset: number) => {
    const target = index + offset
    const list = properties.value
    if (target < 0 || target >= list.length) return
    list.splice(target, 0, list.splice(index, 1)[0])
}

const handleFormat = () => {
    if (!jsonText.value.trim()) return
    try {
        jsonText.value = JSON.stringify(JSON.parse(jsonText.value), null, 2)
        structureError.value = parseThingModel(jsonText.value).error || ''
    } catch (e: any) {
        ElMessage.error(`JSON 格式错误：${e?.message || e}`)
    }
}

// loadFromRaw 打开弹窗时从 JSON 原文重建编辑态：认不出结构就强绑 JSON 模式并给出提示
const loadFromRaw = (raw?: string) => {
    const res = parseThingModel(raw || '')
    structureError.value = res.error || ''
    activeTab.value = res.error ? 'json' : 'table'
    properties.value = res.properties
    jsonText.value = res.error ? (raw || '') : serializeThingModel(res.properties)
}

// 切换视图：radio-group 的 v-model 改值即生效，只能在变更后把不一致的那侧校正
// （形参按 EP 的 change 签名 string|number|boolean|undefined，避开与窄类型的逆变不兼容）
const handleTabChange = (tab?: string | number | boolean) => {
    if (tab === 'json') {
        // 表格转 JSON：以列表为准重新生成，扩展字段已在行内保留
        if (!structureError.value) syncFromTable()
        return
    }
    // JSON 转表格：解析失败就留在 JSON 模式，不把用户的原文冲掉
    const res = parseThingModel(jsonText.value)
    if (res.error) {
        ElMessage.error(res.error)
        activeTab.value = 'json'
        return
    }
    structureError.value = ''
    properties.value = res.properties
}

const handleConfirm = () => {
    if (activeTab.value === 'json') {
        const raw = jsonText.value.trim()
        if (raw) {
            try {
                JSON.parse(raw)
            } catch (e: any) {
                ElMessage.error(`JSON 格式错误：${e?.message || e}`)
                return
            }
        }
        emitValue(raw)
        return
    }
    const problems = validateProperties(properties.value)
    if (problems.length) {
        ElMessage.error(problems[0])
        return
    }
    emitValue(serializeThingModel(properties.value))
}

const emitValue = (value: string) => {
    emit('update:modelValue', value)
    emit('confirm', value)
    visible.value = false
}

const handleClosed = () => {
    structureError.value = ''
}

const open = () => {
    loadFromRaw(props.modelValue)
    visible.value = true
}

defineExpose({ open })
</script>

<style scoped lang="scss">
.thing-model {
    .thing-model__alert {
        margin-bottom: 10px;
    }

    .thing-model__bar {
        display: flex;
        align-items: center;
        gap: 10px;
        margin-bottom: 10px;
    }

    .thing-model__tips {
        color: var(--el-text-color-secondary);
        font-size: 12px;
    }

    .thing-model__range {
        display: flex;
        align-items: center;
        gap: 4px;
    }

    .thing-model__muted {
        color: var(--el-text-color-placeholder);
    }

    :deep(.el-table__cell) {
        padding: 4px 0;
    }

    :deep(.el-input-number .el-input__inner) {
        text-align: left;
    }
}
</style>
