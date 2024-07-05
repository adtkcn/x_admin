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
            <div ref="container" class="diagram-container">
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
// Importing necessary functions and components
import { ref, onMounted } from 'vue'
import LogicFlow from '@logicflow/core'
import { SelectionSelect, Menu, BpmnElement } from '@logicflow/extension'

import '@logicflow/core/dist/style/index.css'
import '@logicflow/extension/lib/style/index.css'
import DiagramToolbar from './DiagramToolbar.vue'
import DiagramSidebar from './DiagramSidebar.vue'
import PropertyPanel from './PropertyPanel/PropertyPanel.vue'
import { registerCustomElement } from './node'

defineOptions({
    name: 'flowEdit'
})
// Define component props
const props = defineProps({
    tabName: {
        type: String,
        default: ''
    },
    fieldList: {
        type: Array,
        default: () => []
    },
    conf: {
        type: Object,
        default: () => ({})
    }
})

// Define refs for reactive data and component references
const lf = ref(null) // Reference to LogicFlow instance
const activeEdges = ref([]) // Reactive array for active edges
const diagramRef = ref(null) // Reference to the diagram container
const PropertyPanelRef = ref(null) // Reference to the PropertyPanel component

// Lifecycle hook to initialize LogicFlow when the component is mounted
onMounted(() => {
    initLogicFlow(props.conf)
})

// Function to initialize LogicFlow
function initLogicFlow(data) {
    // 引入框选插件
    LogicFlow.use(SelectionSelect)
    LogicFlow.use(Menu)
    LogicFlow.use(BpmnElement)

    // Creating a new LogicFlow instance
    const logicFlowInstance = new LogicFlow({
        container: diagramRef.value, // Setting the container where LogicFlow will be rendered
        overlapMode: 1,
        autoWrap: true,
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
        nodeText: { overflowMode: 'autoWrap', lineHeight: 1.5 },
        edgeText: { overflowMode: 'autoWrap', lineHeight: 1.5 }
    })

    // Registering custom elements for LogicFlow
    registerCustomElement(logicFlowInstance)

    // Setting default edge type and rendering initial data
    logicFlowInstance.setDefaultEdgeType('pro-polyline')
    logicFlowInstance.render(data)

    // Assigning the LogicFlow instance to the 'lf' ref
    lf.value = logicFlowInstance

    // Event listener for node clicks
    lf.value.on('node:click', (e) => {
        console.log('Click on node', e.data)
        PropertyPanelRef.value.open(e.data, props.fieldList)
    })
}

// Function to handle dragging nodes into the diagram
function dragInNode(type, text = '') {
    lf.value.dnd.startDrag({
        type,
        text
    })
}

// Function to set properties of a node
function setProperties(node, item) {
    lf.value.setProperties(node.id, item)
}
// function setZIndex(node, type) {
//     lf.value.setElementZIndex(node.id, type)
// }
// Function to import data into the LogicFlow instance
function importData(text) {
    lf.value.renderRawData(text)
}

// Function to save the graph data
function saveGraph() {
    const data = lf.value.getGraphData()
    download('export.json', JSON.stringify(data))
}

// Function to download the graph data as a file
function download(filename, text) {
    window.sessionStorage.setItem(filename, text)
    const element = document.createElement('a')
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text))
    element.setAttribute('download', filename)
    element.style.display = 'none'
    document.body.appendChild(element)
    element.click()
    document.body.removeChild(element)
}
async function getData() {
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
    return new Promise<{
        formData: any
        treeToList: any
    }>((resolve, reject) => {
        const data = lf.value.getGraphData()
        const nodes = data.nodes
        const edges = data.edges

        let haveMoreChildNode = false
        const sourceNodeIdSum = {} // Node ID -> child nodes mapping

        edges.forEach((edge) => {
            const targetNode = nodes.find((item) => item.id === edge.targetNodeId)
            if (sourceNodeIdSum[edge.sourceNodeId]) {
                sourceNodeIdSum[edge.sourceNodeId].push(targetNode)
                for (const n of sourceNodeIdSum[edge.sourceNodeId]) {
                    if (n.type !== 'bpmn:exclusiveGateway') {
                        haveMoreChildNode = true
                        break
                    }
                }
            } else {
                sourceNodeIdSum[edge.sourceNodeId] = [targetNode]
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

        const TreeNode = handel(nodes, 0)

        function handel(arr, pid) {
            const newArr = []
            arr.forEach((node) => {
                const newNode = {
                    id: node.id,
                    pid: pid,
                    label: node?.text?.value,
                    type: node.type,
                    fieldAuth: node?.properties?.fieldAuth,
                    gateway: node?.properties?.gateway,

                    userType: node?.properties?.userType || 0,
                    userId: node?.properties?.userId || 0,
                    deptId: node?.properties?.deptId || 0,
                    postId: node?.properties?.postId || 0,
                    children: null
                }
                if (sourceNodeIdSum[node.id]) {
                    newNode.children = handel(sourceNodeIdSum[node.id], node.id)
                }
                newArr.push(newNode)
            })
            return newArr
        }

        function treeToList(tree) {
            const arr = []
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

<style lang="scss">
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
    .diagram-container :v-deep .lf-background {
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
