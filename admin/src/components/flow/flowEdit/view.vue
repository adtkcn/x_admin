<template>
    <div class="diagram">
        <div class="diagram-main">
            <div class="diagram-container">
                <div class="diagram-wrapper">
                    <div ref="diagram" class="lf-diagram"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import LogicFlow from '@logicflow/core'
import { BpmnElement } from '@logicflow/extension'
// import { SelectionSelect, Menu } from "@logicflow/extension";
import '@logicflow/core/dist/style/index.css'
import '@logicflow/extension/lib/style/index.css'

import { registerCustomElement } from './node'

export default {
    name: 'DiagramView',
    components: {},
    data() {
        return {
            lf: '',
            activeNodes: [],
            activeEdges: [],
            properties: {},

            timer: null,
            time: 60000
        }
    },
    watch: {
        // appStore: {
        //   handler() {
        //     if (this.lf) {
        //       // this.lf.graphModel.resize();
        //       // this.lf.fitView();
        //       // 侧栏有动画，所以要加延时
        //       setTimeout(() => {
        //         this.lf.graphModel.resize();
        //         this.lf.fitView();
        //       }, 500);
        //     }
        //   },
        //   deep: true,
        // },
    },
    mounted() {
        // this.initLogicFlow(exportInfo);
    },
    beforeUnmount() {
        if (this.timer) {
            clearInterval(this.timer)
        }
    },
    methods: {
        initLogicFlow(data) {
            // 引入框选插件
            // LogicFlow.use(SelectionSelect);
            // LogicFlow.use(Menu);
            LogicFlow.use(BpmnElement)
            const lf = new LogicFlow({
                container: this.$refs.diagram,
                overlapMode: 1,
                autoWrap: true,
                stopScrollGraph: true,
                stopZoomGraph: true,
                stopMoveGraph: true,
                metaKeyMultipleSelected: true,
                keyboard: {
                    enabled: false
                },
                isSilentMode: true,
                grid: {
                    visible: false,
                    size: 1,
                    type: 'mesh',
                    config: {
                        color: 'rgba(255,255,255,0.1)',
                        thickness: 1
                    }
                },
                background: {
                    backgroundColor: 'rgba(29, 32, 98, 1)'
                    // backgroundImage:
                    //   'url("data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDAiIGhlaWdodD0iNDAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PHBhdHRlcm4gaWQ9ImdyaWQiIHdpZHRoPSI0MCIgaGVpZ2h0PSI0MCIgcGF0dGVyblVuaXRzPSJ1c2VyU3BhY2VPblVzZSI+PHBhdGggZD0iTSAwIDEwIEwgNDAgMTAgTSAxMCAwIEwgMTAgNDAgTSAwIDIwIEwgNDAgMjAgTSAyMCAwIEwgMjAgNDAgTSAwIDMwIEwgNDAgMzAgTSAzMCAwIEwgMzAgNDAiIGZpbGw9Im5vbmUiIHN0cm9rZT0iI2QwZDBkMCIgb3BhY2l0eT0iMC4yIiBzdHJva2Utd2lkdGg9IjEiLz48cGF0aCBkPSJNIDQwIDAgTCAwIDAgMCA0MCIgZmlsbD0ibm9uZSIgc3Ryb2tlPSIjZDBkMGQwIiBzdHJva2Utd2lkdGg9IjEiLz48L3BhdHRlcm4+PC9kZWZzPjxyZWN0IHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIGZpbGw9InVybCgjZ3JpZCkiLz48L3N2Zz4=")',
                    // backgroundRepeat: "repeat",
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
            lf.fitView()

            this.lf = lf
        },
        setType(errorMap) {
            const gatwayStatus = {
                error: 0,
                success: 0
            }
            const all = this.lf.getGraphRawData()
            all.nodes.forEach((item) => {
                if (!item.type.includes('status')) {
                    return
                }
                if (
                    item.properties.deviceType == 1 &&
                    errorMap.T80ErrorArr.some((i) => i == item.properties.equipmentId)
                ) {
                    this.lf.changeNodeType(item.id, 'status_error')
                } else if (
                    item.properties.deviceType == 2 &&
                    errorMap.ZLErrorArr.some((i) => i == item.properties.equipmentId)
                ) {
                    this.lf.changeNodeType(item.id, 'status_error')
                } else if (
                    item.properties.deviceType == 3 &&
                    errorMap.TimeOutErrorArr.some((i) => i == item.properties.equipmentId)
                ) {
                    this.lf.changeNodeType(item.id, 'status_error')
                } else if (
                    item.properties.deviceType == 3 ||
                    item.properties.deviceType == 2 ||
                    item.properties.deviceType == 1
                ) {
                    this.lf.changeNodeType(item.id, 'status_success')
                } else if (item.properties.deviceType == 4) {
                    let equipments = item.properties.equipmentId.split(',')
                    equipments = equipments.map((i) => i.trim())
                    if (
                        equipments.every((equipmentId) =>
                            errorMap.TimeOutErrorArr.includes(equipmentId)
                        )
                    ) {
                        gatwayStatus.error++
                        this.lf.changeNodeType(item.id, 'status_error')
                    } else {
                        gatwayStatus.success++
                        this.lf.changeNodeType(item.id, 'status_success')
                    }
                }
            })
            this.$emit('gatewayStatus', gatwayStatus)
        },
        setTypeByDeviceID(deviceID, itemID, status) {
            const all = this.lf.getGraphRawData()
            all.nodes.forEach((item) => {
                if (item.properties.deviceType) {
                    return
                }
                if (
                    item.properties.deviceID == deviceID &&
                    item.properties.itemID == itemID &&
                    status == 1
                ) {
                    this.lf.changeNodeType(item.id, 'status_success')
                } else if (
                    item.properties.deviceID == deviceID &&
                    item.properties.itemID == itemID &&
                    status == 0
                ) {
                    this.lf.changeNodeType(item.id, 'status_error')
                }
            })
        },
        networkError() {
            const all = this.lf.getGraphRawData()
            all.nodes.forEach((item) => {
                if (item.properties.deviceType) {
                    return
                }
                if (item.properties.deviceID && item.properties.itemID) {
                    this.lf.changeNodeType(item.id, 'status_error')
                }
            })
        }
    }
}
</script>

<style scoped>
.diagram {
    width: 100%;
    height: 100%;
}

.diagram * {
    box-sizing: border-box;
}

.diagram-main {
    display: flex;
    width: 100%;
    height: 100%;
    overflow: hidden;
}

.diagram-container {
    flex: 1;
}

/* 由于背景图和gird不对齐，需要css处理一下 */
/* .diagram /deep/ .lf-background {
  left: -9px;
} */
.diagram-wrapper {
    box-sizing: border-box;
    width: 100%;
    height: 100%;
}

.lf-diagram {
    box-shadow: 0px 0px 4px #838284;
    width: 100%;
    height: 100%;
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
