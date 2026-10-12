<template>
    <el-dialog
        v-model="dialogVisible"
        append-to-body
        :show-close="false"
        :fullscreen="true"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        top="1px"
        style="height: 100vh"
        class="rc-config-dialog"
    >
        <template #header>
            <header class="rc-header">
                <div class="rc-header__title">{{ title }}</div>
                <el-form :model="formData" inline class="rc-header__form">
                    <el-form-item label="名称" prop="name">
                        <el-input v-model="formData.name" placeholder="规则链名称" style="width: 180px" />
                    </el-form-item>
                    <el-form-item label="备注" prop="remark">
                        <el-input v-model="formData.remark" placeholder="备注" style="width: 220px" />
                    </el-form-item>
                    <el-form-item label="停用" prop="is_stop">
                        <el-switch v-model="formData.is_stop" :active-value="1" :inactive-value="0" />
                    </el-form-item>
                </el-form>
                <div class="rc-header__actions">
                    <el-button @click="exit">关闭</el-button>
                    <el-button type="primary" @click="publish">保存</el-button>
                </div>
            </header>
        </template>

        <div class="rc-content" v-if="dialogVisible">
            <RuleChainDesigner ref="designerRef" :conf="graphConf" />
        </div>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, useTemplateRef, defineAsyncComponent } from 'vue'
import { ruleChainAdd, ruleChainEdit, type type_rule_chain } from '@/api/iot/rule_chain'
import feedback from '@/utils/feedback'

const RuleChainDesigner = defineAsyncComponent(
    () => import('@/components/rule-chain/index.vue')
)

defineOptions({ name: 'ruleChainEditPopup' })
const emit = defineEmits(['success'])

const designerRef = useTemplateRef<InstanceType<typeof RuleChainDesigner>>('designerRef')
const dialogVisible = ref(false)
const title = ref('新增规则链')

const formData = ref<type_rule_chain>({
    id: '',
    name: '',
    remark: '',
    is_stop: 0
})
// 传给设计器的 LogicFlow 原生图；每次打开前重置
const graphConf = ref<Record<string, any>>({ nodes: [], edges: [] })

const mode = ref<'add' | 'edit'>('add')

function reset() {
    formData.value = { id: '', name: '', remark: '', is_stop: 0 }
    graphConf.value = { nodes: [], edges: [] }
}

async function open(row?: type_rule_chain) {
    reset()
    if (row && row.id) {
        mode.value = 'edit'
        title.value = '编辑规则链'
        formData.value = { ...row }
        try {
            graphConf.value = row.graph_data ? JSON.parse(row.graph_data) : { nodes: [], edges: [] }
        } catch {
            graphConf.value = { nodes: [], edges: [] }
        }
    } else {
        mode.value = 'add'
        title.value = '新增规则链'
    }
    dialogVisible.value = true
}

async function publish() {
    if (!formData.value.name?.trim()) {
        return feedback.msgError('请输入规则链名称')
    }
    let result: { graph: Record<string, any>; tree: Record<string, any> }
    try {
        result = await designerRef.value!.getData()
    } catch (err: any) {
        return feedback.msgError(err?.message ?? '规则链校验未通过')
    }
    const payload: type_rule_chain = {
        ...formData.value,
        graph_data: JSON.stringify(result.graph),
        tree_data: JSON.stringify(result.tree)
    }
    if (mode.value === 'edit') {
        await ruleChainEdit(payload)
        feedback.msgSuccess('修改成功')
    } else {
        await ruleChainAdd(payload)
        feedback.msgSuccess('新增成功')
    }
    dialogVisible.value = false
    emit('success')
}

function exit() {
    feedback
        .confirm('离开此页面您的修改将会丢失，是否继续？')
        .then(() => {
            dialogVisible.value = false
        })
        .catch(() => {})
}

defineExpose({ open })
</script>

<style>
.rc-config-dialog {
    padding: 0;
}
.rc-config-dialog .el-dialog__header {
    padding: 0;
    margin: 0;
}
</style>
<style lang="scss" scoped>
$height: 54px;
.rc-header {
    position: fixed;
    z-index: 100;
    width: 100%;
    height: $height;
    display: flex;
    align-items: center;
    gap: 20px;
    box-sizing: border-box;
    padding: 0 20px;
    color: #fff;
    background: #3296fa;
    font-size: 14px;
    user-select: none;
    &__title {
        font-weight: 600;
    }
    &__form {
        flex: 1;
        :deep(.el-form-item) {
            margin-bottom: 0;
        }
        :deep(.el-form-item__label) {
            color: #fff;
        }
    }
    &__actions {
        display: flex;
        gap: 10px;
    }
}
.rc-content {
    padding-top: $height;
    height: 100%;
    box-sizing: border-box;
}
</style>
