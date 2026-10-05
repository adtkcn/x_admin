<template>
    <div class="rc-diagram">
        <div class="rc-toolbar">
            <el-button-group>
                <el-button size="small" @click="zoom(true)">
                    <el-icon><ZoomIn /></el-icon>
                </el-button>
                <el-button size="small" @click="zoom(false)">
                    <el-icon><ZoomOut /></el-icon>
                </el-button>
                <el-button size="small" @click="restore">适应画布</el-button>
            </el-button-group>
            <span class="rc-toolbar__tip">双击节点配置属性，Delete 删除选中；除并行外每个节点只连一个后续节点</span>
        </div>
        <div class="rc-main">
            <DiagramSidebar class="rc-sidebar" @dragInNode="dragInNode" />
            <div class="rc-container">
                <div ref="diagramRef" class="lf-diagram"></div>
            </div>
        </div>
        <PropertyPanel
            ref="PropertyPanelRef"
            @setProperties="setProperties"
            @renameNode="renameNode"
        />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, useTemplateRef, defineAsyncComponent } from 'vue'
import { LogicFlow } from '@logicflow/core'
import { SelectionSelect, Menu, MiniMap } from '@logicflow/extension'
import { ZoomIn, ZoomOut } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

import '@logicflow/core/lib/style/index.css'
import '@logicflow/extension/lib/style/index.css'

import { registerRuleChainElement } from './node/index'
import { rcNodeKeyOf, type RcNodeTypeData } from './rule.type'

const DiagramSidebar = defineAsyncComponent(() => import('./DiagramSidebar.vue'))
const PropertyPanel = defineAsyncComponent(() => import('./PropertyPanel/index.vue'))

defineOptions({ name: 'ruleChainEdit' })

const props = defineProps<{
    /** LogicFlow 原生图数据 { nodes, edges } */
    conf?: Record<string, any>
}>()

const lf = ref<LogicFlow | null>(null)
const diagramRef = useTemplateRef<HTMLDivElement>('diagramRef')
const PropertyPanelRef = useTemplateRef<InstanceType<typeof PropertyPanel>>('PropertyPanelRef')

onMounted(() => {
    initLogicFlow(props.conf ?? {})
})
onBeforeUnmount(() => {
    if (lf.value) {
        lf.value.destroy()
        lf.value = null
    }
})

function initLogicFlow(data: Record<string, any>) {
    const instance = new LogicFlow({
        plugins: [SelectionSelect, Menu, MiniMap],
        container: diagramRef.value as HTMLDivElement,
        overlapMode: 1,
        autoWrap: true,
        adjustEdge: true,
        adjustEdgeStartAndEnd: true,
        metaKeyMultipleSelected: true,
        keyboard: { enabled: true },
        grid: { size: 10, type: 'dot' },
        // 「输入」是固定起点：键盘/框选/接口删除统一拦截（返回 false 取消删除）
        guards: {
            beforeDelete: (data: any) =>
                'sourceNodeId' in data ? true : rcNodeKeyOf(data.type) !== 'rule_trigger'
        }
    })
    instance.setTheme({
        baseEdge: { strokeWidth: 1 },
        nodeText: { overflowMode: 'ellipsis', lineHeight: 1.4, fontSize: 12 },
        edgeText: { overflowMode: 'ellipsis', lineHeight: 1.4, fontSize: 11, textWidth: 120 }
    })
    registerRuleChainElement(instance)
    instance.setDefaultEdgeType('pro-polyline')
    // 关闭双击内联编辑文本，避免与「双击打开属性面板」冲突（改名统一在面板进行）
    instance.updateEditConfig({ textEdit: false })
    // setMenuConfig 整体替换内置默认菜单（内置项自带「删除」「编辑文本」，用 add 追加会出现两个删除、
    // 且「编辑文本」与 textEdit:false 冲突）；空白/框选右键不需要菜单
    ;(instance.extension.menu as Menu).setMenuConfig({
        nodeMenu: [
            { text: '属性配置', callback(node: any) { PropertyPanelRef.value?.open(node) } },
            {
                text: '删除',
                callback(node: any) {
                    if (rcNodeKeyOf(node.type) === 'rule_trigger') {
                        return ElMessage.info('「输入」是固定起点，不可删除')
                    }
                    instance.deleteNode(node.id)
                }
            }
        ],
        edgeMenu: [
            { text: '删除', callback(edge: any) { instance.deleteEdge(edge.id) } }
        ],
        graphMenu: false
    })
    instance.render(data)
    ensureTriggerNode(instance)
    ;(instance.extension.miniMap as MiniMap).show()
    lf.value = instance

    lf.value.on('node:dbclick', (e: any) => PropertyPanelRef.value?.open(e.data))

    // 单出口约束：除「并行」外已有出边的节点再连一条时，回滚刚建的边并提示
    lf.value.on('edge:add', (e: any) => {
        const srcId = e.data?.sourceNodeId
        if (!srcId) return
        const graph = instance.getGraphData() as Record<string, any>
        const outs = ((graph.edges as any[]) || []).filter((ed) => ed.sourceNodeId === srcId)
        const srcNode = ((graph.nodes as RcNodeTypeData[]) || []).find((n) => n.id === srcId)
        if (rcNodeKeyOf(srcNode?.type) === 'rule_parallel' || outs.length <= 1) return
        // 事件派发中直接删边会干扰 LF 内部流程，延到下一拍再回滚
        setTimeout(() => instance.deleteEdge(e.data.id), 0)
        ElMessage.warning(`「${srcNode?.text?.value ?? '该节点'}」只允许连接一个后续节点，分流请使用并行节点`)
    })
}

function dragInNode(type: string, text = '') {
    lf.value!.dnd.startDrag({ type, text })
}

// 输入节点默认存在：图中缺入口时自动补放（新建空图/旧数据），落在画布视觉中心
function ensureTriggerNode(target: NonNullable<typeof lf.value>) {
    const nodes = (target.getGraphData() as Record<string, any>)?.nodes as RcNodeTypeData[] | undefined
    if ((nodes || []).some((n) => rcNodeKeyOf(n.type) === 'rule_trigger')) return
    const { width, height } = target.graphModel
    target.addNode({ type: 'rule:trigger', x: width / 2, y: height / 2, text: '输入' })
}

// 合并写回节点命名空间属性，保留该节点其他命名空间
function setProperties(node: RcNodeTypeData, patch: Record<string, any>) {
    const old = (lf.value!.getNodeModelById(node.id!)?.getProperties() ?? {}) as Record<string, any>
    lf.value!.setProperties(node.id as string, { ...old, ...patch })
}

// 重命名节点显示文字
function renameNode(id: string, text: string) {
    if (id) lf.value!.updateText(id, text ?? '')
}

function zoom(inRatio: boolean) {
    lf.value?.zoom(inRatio)
}
function restore() {
    lf.value?.resetZoom()
    lf.value?.translateCenter?.()
}

/** 导出 LogicFlow 原生有向图 { nodes, edges }，不拍树 */
async function getData(): Promise<Record<string, any>> {
    return new Promise((resolve, reject) => {
        const data = lf.value!.getGraphData() as Record<string, any>
        const nodes = (data?.nodes as RcNodeTypeData[]) || []
        const triggers = nodes.filter((n) => rcNodeKeyOf(n.type) === 'rule_trigger')
        if (!nodes.length) {
            return reject({ message: '规则链-画布为空，请至少配置一个处理节点' })
        }
        if (triggers.length !== 1) {
            return reject({ message: '规则链-必须有且仅有一个「输入」入口节点' })
        }
        resolve(data)
    })
}

/** 载入原生图数据（同样保证「输入」起点存在） */
function setData(data: Record<string, any>) {
    lf.value?.renderRawData(data as any)
    if (lf.value) ensureTriggerNode(lf.value)
}

defineExpose({ getData, setData })
</script>

<style lang="scss" scoped>
.rc-diagram {
    width: 100%;
    height: 100%;
    position: relative;
}
.rc-toolbar {
    position: absolute;
    top: 10px;
    right: 20px;
    z-index: 10;
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 6px 12px;
    background: #fff;
    border-radius: 4px;
    box-shadow: 0 0 4px rgba(0, 0, 0, 0.2);
    &__tip {
        font-size: 12px;
        color: var(--el-text-color-secondary);
    }
}
.rc-main {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
    .rc-sidebar {
        position: absolute;
        left: 20px;
        top: 16px;
        z-index: 10;
        width: 120px;
        background: #fff;
        border-radius: 4px;
        box-shadow: 0 0 4px rgba(0, 0, 0, 0.2);
    }
    .rc-container {
        height: 100%;
        .lf-diagram {
            width: 100%;
            height: 100%;
        }
    }
}
</style>
