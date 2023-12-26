<template>
    <div class="diagram">
        <diagram-toolbar
            v-if="lf"
            class="diagram-toolbar"
            :lf="lf"
            :active-edges="activeEdges"
            @saveGraph="$_saveGraph"
            @importData="importData"
        />
        <div class="diagram-main">
            <diagram-sidebar class="diagram-sidebar" @dragInNode="$_dragInNode" />
            <div ref="container" class="diagram-container">
                <div class="diagram-wrapper">
                    <div ref="diagram" class="lf-diagram"></div>
                </div>
            </div>
        </div>
        <!-- 右侧属性面板 -->
        <PropertyPanel ref="panelRef" @setProperties="$setProperties" />
    </div>
</template>

<script>
import LogicFlow from '@logicflow/core'
import { SelectionSelect, Menu, BpmnElement } from '@logicflow/extension'

import '@logicflow/core/dist/style/index.css'
import '@logicflow/extension/lib/style/index.css'
import DiagramToolbar from './DiagramToolbar.vue'
import DiagramSidebar from './DiagramSidebar.vue'
import PropertyPanel from './PropertyPanel.vue'
import { registerCustomElement } from './node'

// import { Control } from "@logicflow/extension";

export default {
    name: 'Diagram',
    components: {
        DiagramToolbar,
        DiagramSidebar,
        PropertyPanel
    },
    props: {
        fieldList: {
            type: Array,
            default: () => {
                return []
            }
        },
        conf: {
            type: Object,
            default: () => {
                return {}
            }
        }
    },
    data() {
        return {
            sidebarWidth: 200,
            diagramWidth: 0,
            diagramHeight: 0,
            lf: '',
            filename: '',
            activeEdges: []
        }
    },
    mounted() {
        // const data = ''

        this.filename = 'export.json'
        // const d = window.sessionStorage.getItem(this.filename)
        // if (d) {
        //     data = JSON.parse(d)
        // }

        this.initLogicFlow(this.conf)
    },
    methods: {
        initLogicFlow(data) {
            // 引入框选插件
            LogicFlow.use(SelectionSelect)
            LogicFlow.use(Menu)
            LogicFlow.use(BpmnElement)

            const lf = new LogicFlow({
                container: this.$refs.diagram,
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
            lf.setTheme({
                baseEdge: { strokeWidth: 1 },
                baseNode: { strokeWidth: 1 },
                nodeText: { overflowMode: 'autoWrap', lineHeight: 1.5 },
                edgeText: { overflowMode: 'autoWrap', lineHeight: 1.5 }
            })
            // 注册自定义元素
            registerCustomElement(lf)
            lf.setDefaultEdgeType('pro-polyline')
            lf.render(data)
            this.lf = lf

            this.lf.on('node:click', (e) => {
                console.log('点击节点', e.data)
                this.$refs.panelRef.open(e.data, this.fieldList)
            })

            // lf.renderRawData( )
        },

        $_dragInNode(type, text = '') {
            this.lf.dnd.startDrag({
                type,
                text
            })
        },
        // 获取可以进行设置的属性
        $_getProperty() {
            // this.lf.getProperties(node.id)
        },
        $setProperties(node, item) {
            this.lf.setProperties(node.id, item)
        },
        $_setZIndex(node, type) {
            this.lf.setElementZIndex(node.id, type)
        },
        importData(text) {
            this.lf.renderRawData(text)
        },
        $_saveGraph() {
            const data = this.lf.getGraphData()
            this.download(this.filename, JSON.stringify(data))
        },
        download(filename, text) {
            window.sessionStorage.setItem(filename, text)
            const element = document.createElement('a')
            element.setAttribute(
                'href',
                'data:text/plain;charset=utf-8,' + encodeURIComponent(text)
            )
            element.setAttribute('download', filename)
            element.style.display = 'none'
            document.body.appendChild(element)
            element.click()
            document.body.removeChild(element)
        },
        getData() {
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
            return new Promise((resolve, reject) => {
                const data = this.lf.getGraphData()

                const nodes = data.nodes
                const edges = data.edges

                let haveMoreChildNode = false
                const sourceNodeIdSum = {} //节点id->[子节点id]

                edges.forEach((edge) => {
                    const targetNode = nodes.find((item) => item.id == edge.targetNodeId)
                    if (sourceNodeIdSum[edge.sourceNodeId]) {
                        sourceNodeIdSum[edge.sourceNodeId].push(targetNode)

                        for (const n of sourceNodeIdSum[edge.sourceNodeId]) {
                            if (n.type != 'bpmn:exclusiveGateway') {
                                haveMoreChildNode = true
                                break
                            }
                        }
                    } else {
                        sourceNodeIdSum[edge.sourceNodeId] = [targetNode]
                    }
                })
                console.log('sourceNodeIdSum', sourceNodeIdSum)
                if (haveMoreChildNode) {
                    // 如果都是网关就可以,等优化
                    return reject({
                        msg: '流程设计-一个节点只能有一个子节点，可以有多个网关'
                    })
                }

                // if (data.nodes.length != data.edges.length + 1) {
                //   return reject({
                //     msg: "流程设计-节点之间必须连线，可以清理多余的节点",//不是很重要
                //   });
                // }
                // 检查开始节点和结束节点是否存在
                const findStartNode = nodes.filter((item) => item.type == 'bpmn:startEvent')
                const findEndNode = nodes.filter((item) => item.type == 'bpmn:endEvent')
                if (findStartNode.length != 1 || findEndNode.length != 1) {
                    return reject({
                        msg: '流程设计-流程必须有且只有一个开始节点和结束节点'
                    })
                }

                function handel(arr, pid) {
                    const newArr = []
                    arr.forEach((node) => {
                        // console.log('sourceNodeIdSum', sourceNodeIdSum, node.id)
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
                            postId: node?.properties?.postId || 0
                        }
                        if (sourceNodeIdSum[node.id]) {
                            newNode.children = handel(sourceNodeIdSum[node.id], node.id)
                        }
                        newArr.push(newNode)
                    })
                    return newArr
                }
                const TreeNode = handel(findStartNode, 0)

                // tree转list
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

                console.log('TreeNode', TreeNode)
                console.log('treeToList', treeToList(TreeNode))
                // 检查连线方向是否正确;
                resolve({ formData: data, treeToList: treeToList(TreeNode) })
            })
        }
    }
}
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
