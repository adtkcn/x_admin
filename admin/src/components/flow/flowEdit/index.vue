<template>
    <div class="diagram">
        <diagram-toolbar
            v-if="lf"
            class="diagram-toolbar"
            :lf="lf"
            :active-edges="activeEdges"
            @saveGraph="saveGraph"
            @importData="importData"
        />
        <div class="diagram-main">
            <diagram-sidebar class="diagram-sidebar" @dragInNode="dragInNode" />
            <div class="diagram-container">
                <div class="diagram-wrapper">
                    <div ref="diagramRef" class="lf-diagram"></div>
                </div>
            </div>
        </div>
        <!-- Right-side property panel -->
        <PropertyPanel ref="PropertyPanelRef" @setProperties="setProperties" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, useTemplateRef, defineAsyncComponent } from 'vue'
import { LogicFlow } from '@logicflow/core'

import { SelectionSelect, Menu, BpmnElement, MiniMap } from '@logicflow/extension'
import type { NodeType, PropertiesType, FormFieldListType } from './PropertyPanel/property.type'
import { nodeKeyOf, defaultProps } from './PropertyPanel/property.type'

import '@logicflow/core/lib/style/index.css'
import '@logicflow/extension/lib/style/index.css'

// import DiagramToolbar from './DiagramToolbar.vue'
// import DiagramSidebar from './DiagramSidebar.vue'
// import PropertyPanel from './PropertyPanel/index.vue'
const DiagramToolbar = defineAsyncComponent(() => import('./DiagramToolbar.vue'))
const DiagramSidebar = defineAsyncComponent(() => import('./DiagramSidebar.vue'))
const PropertyPanel = defineAsyncComponent(() => import('./PropertyPanel/index.vue'))
import { registerCustomElement } from './node/index'

defineOptions({
    name: 'flowEdit'
})
// Define component props
const props = defineProps<{
    tabName: string
    fieldList: FormFieldListType[]
    /** LogicFlow 渲染数据（节点/边等内部结构复杂，保持宽松） */
    conf: Record<string, any>
}>()

/** LogicFlow 边数据（仅取使用到的字段） */
interface EdgeItem {
    sourceNodeId: string
    targetNodeId: string
}

/** 流程节点树节点（由 handel 构造，用于前后端交互） */
interface FlowTreeNode {
    id: string
    pid: string | number
    label?: string
    type?: string
    // 节点私有属性：取出对应命名空间 key 的内容，平铺到 props
    // 新增任何字段都不需要改这里，子面板与后端 NodeProps 已约定好
    props: Record<string, any>
    children: FlowTreeNode[] | null
}

/** getData 返回结果 */
interface FlowGraphData {
    /** 原始 LogicFlow 图形数据 */
    formData: Record<string, any>
    /** 流程节点树拍平后的列表 */
    treeToList: FlowTreeNode[]
}

// Define refs for reactive data and component references
const lf = ref<LogicFlow | null>(null) // Reference to LogicFlow instance
const activeEdges = ref<EdgeItem[]>([]) // Reactive array for active edges
const diagramRef = useTemplateRef<HTMLDivElement>('diagramRef') // Reference to the diagram container
const PropertyPanelRef = useTemplateRef<InstanceType<typeof PropertyPanel>>('PropertyPanelRef') // Reference to the PropertyPanel component

// Lifecycle hook to initialize LogicFlow when the component is mounted
onMounted(() => {
    initLogicFlow(props.conf)
})
// Lifecycle hook to clean up LogicFlow when the component is unmounted
onBeforeUnmount(() => {
    if (lf.value) {
        console.log('卸载LogicFlow')
        lf.value.destroy()
        lf.value = null
    }
})
// Function to initialize LogicFlow
function initLogicFlow(data: Record<string, any>) {
    const logicFlowInstance = new LogicFlow({
        plugins: [SelectionSelect, Menu, MiniMap, BpmnElement],
        container: diagramRef.value as HTMLDivElement,
        overlapMode: 1,
        // allowResize: true,
        autoWrap: true,
        adjustEdge: true,
        adjustEdgeStartAndEnd: true,
        metaKeyMultipleSelected: true,
        keyboard: {
            enabled: true
        },
        grid: {
            size: 10,
            type: 'dot'
        }
    })

    // Setting theme for LogicFlow
    logicFlowInstance.setTheme({
        baseEdge: { strokeWidth: 1 },
        baseNode: { strokeWidth: 1 },
        nodeText: { overflowMode: 'autoWrap', lineHeight: 1.5, fontSize: 12 },
        edgeText: { overflowMode: 'autoWrap', lineHeight: 1.5, fontSize: 12, textWidth: 100 }
    })

    // Registering custom elements for LogicFlow
    registerCustomElement(logicFlowInstance)

    // Setting default edge type and rendering initial data
    logicFlowInstance.setDefaultEdgeType('pro-polyline')
    ;(logicFlowInstance.extension.menu as Menu).addMenuConfig({
        nodeMenu: [
            {
                text: '属性配置',
                callback(node) {
                    PropertyPanelRef.value?.open(node as NodeType, props.fieldList)
                }
            }
        ]
    })
    logicFlowInstance.render(data)
    ;(logicFlowInstance.extension.miniMap as MiniMap).show()
    // Assigning the LogicFlow instance to the 'lf' ref
    lf.value = logicFlowInstance

    // Event listener for node clicks
    lf.value.on('node:dbclick', (e) => {
        console.log('dbclick on node', e.data, props.fieldList)
        PropertyPanelRef.value?.open(e.data as NodeType, props.fieldList)
    })
}

// Function to handle dragging nodes into the diagram
function dragInNode(type: string, text = '') {
    lf.value!.dnd.startDrag({
        type,
        text
    })
}

// Function to set properties of a node
// 合并写入：保留该节点已有的其他命名空间属性，仅覆盖本次 patch 的 key
function setProperties(node: NodeType, item: PropertiesType) {
    const old = (lf.value!.getNodeModelById(node.id!)?.getProperties() ?? {}) as PropertiesType
    lf.value!.setProperties(node.id as string, { ...old, ...item })
}

function importData(text: string) {
    lf.value!.renderRawData(text as any)
}

// Function to save the graph data
function saveGraph() {
    const data = lf.value!.getGraphData()
    download(`export.${Date.now()}.json`, JSON.stringify(data))
}

// Function to download the graph data as a file
function download(filename: string, text: string) {
    window.sessionStorage.setItem(filename, text)
    const element = document.createElement('a')
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text))
    element.setAttribute('download', filename)
    element.style.display = 'none'
    document.body.appendChild(element)
    element.click()
    document.body.removeChild(element)
}
async function getData(): Promise<FlowGraphData> {
    /**
     * 校验目标
     * 1. 必须存在开始节点和结束节点
     * 2. 连线方向正确
     * 3. 多余的节点（不重要）
     * 4. 所有分支必须结束节点
     * 5. 检查审批节点设置情况
     * 6. 一个节点可以有多个子网关，子网关只能通过一个，不能有多个普通子节点
     *
     */
    return new Promise<FlowGraphData>((resolve, reject) => {
        const data: Record<string, any> = lf.value!.getGraphData() as Record<string, any>
        const nodes = (data?.nodes as NodeType[]) || []
        const edges = (data?.edges as EdgeItem[]) || []

        let haveMoreChildNode = false
        const sourceNodeIdSum: Record<string, NodeType[]> = {} // Node ID -> child nodes mapping

        edges.forEach((edge) => {
            const targetNode = nodes.find((item) => item.id === edge.targetNodeId)
            if (sourceNodeIdSum[edge.sourceNodeId]) {
                sourceNodeIdSum[edge.sourceNodeId].push(targetNode!)
                for (const n of sourceNodeIdSum[edge.sourceNodeId]) {
                    if (n.type !== 'bpmn:exclusiveGateway') {
                        haveMoreChildNode = true
                        break
                    }
                }
            } else {
                sourceNodeIdSum[edge.sourceNodeId] = [targetNode!]
            }
        })

        if (haveMoreChildNode) {
            return reject({
                target: props.tabName,
                message: '流程设计-一个节点只能有一个子节点，可以有多个网关'
            })
        }
        // 检查开始节点和结束节点是否存在
        const findStartNode = nodes.filter((item) => item.type === 'bpmn:startEvent')
        const findEndNode = nodes.filter((item) => item.type === 'bpmn:endEvent')
        if (findStartNode.length !== 1 || findEndNode.length !== 1) {
            return reject({
                target: props.tabName,
                message: '流程设计-流程必须有且只有一个开始节点和结束节点'
            })
        }

        // 仅从开始节点递归构造树，避免把全部节点当顶层导致每个节点重复出现
        const startNodes = nodes.filter((item) => item.type === 'bpmn:startEvent')
        const TreeNode = handel(startNodes, 0)

        function handel(arr: NodeType[], pid: string | number): FlowTreeNode[] {
            const newArr: FlowTreeNode[] = []
            arr.forEach((node) => {
                // 取出当前节点类型对应的属性命名空间，按命名空间包裹到 props，
                // 与后端 NodeProps（按 key 命名空间拆分）一致；既有缺省补齐又新增字段无需改动。
                const key = nodeKeyOf(node.type)
                const newNode: FlowTreeNode = {
                    id: node.id as string,
                    pid: pid,
                    label: node?.text?.value,
                    type: node.type,
                    props: key
                        ? { [key]: { ...defaultProps(key), ...(node.properties?.[key] ?? {}) } }
                        : {},
                    children: null
                }
                if (sourceNodeIdSum[node.id as string]) {
                    newNode.children = handel(sourceNodeIdSum[node.id as string], node.id as string)
                }
                newArr.push(newNode)
            })
            return newArr
        }

        function treeToList(tree: FlowTreeNode[]): FlowTreeNode[] {
            const arr: FlowTreeNode[] = []
            tree.forEach((item) => {
                arr.push(item)
                if (item.children) {
                    arr.push(...treeToList(item.children))
                }
            })
            return arr
        }
        // 检查连线方向是否正确;
        resolve({ formData: data, treeToList: treeToList(TreeNode) })
    })
}

defineExpose({
    getData
})
</script>

<style lang="scss" scoped>
.diagram {
    width: 100%;
    height: 100%;
    position: relative;
}

.diagram-toolbar {
    position: absolute;
    top: 10px;
    right: 20px;
    height: 40px;
    padding: 0 10px;
    /* width: 310px; */
    display: flex;
    align-items: center;
    /* border-bottom: 1px solid #e5e5e5; */
    z-index: 10;
    background: #fff;
    box-shadow: 0px 0px 4px rgba($color: #000000, $alpha: 0.5);
}
.diagram-main {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;

    .diagram-sidebar {
        position: absolute;
        left: 20px;
        z-index: 10;
        top: 50px;
        width: 60px;
        background: #fff;
        box-shadow: 0px 0px 4px rgba($color: #000000, $alpha: 0.5);
    }
    .diagram-container {
        height: 100%;
        .diagram-wrapper {
            box-sizing: border-box;
            width: 100%;
            height: 100%;

            .lf-diagram {
                // box-shadow: 0px 0px 4px #838284;
                width: 100%;
                height: 100%;
            }
        }
    }
    /* 由于背景图和gird不对齐，需要css处理一下 */
    .diagram-container :deep(.lf-background) {
        left: -9px;
    }
}

::-webkit-scrollbar {
    width: 9px;
    height: 9px;
    background: white;
    border-left: 1px solid #e8e8e8;
}
::-webkit-scrollbar-thumb {
    border-width: 1px;
    border-style: solid;
    border-color: #fff;
    border-radius: 6px;
    background: #c9c9c9;
}
::-webkit-scrollbar-thumb:hover {
    background: #b5b5b5;
}
</style>
