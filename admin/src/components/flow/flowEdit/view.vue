<template>
    <div class="diagram">
        <div class="diagram-main">
            <div class="diagram-container">
                <div class="diagram-wrapper">
                    <div ref="diagramRef" class="lf-diagram"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import LogicFlow from '@logicflow/core'
import { BpmnElement } from '@logicflow/extension'

import '@logicflow/core/lib/style/index.css'
import '@logicflow/extension/lib/style/index.css'

import { registerCustomElement } from './node'

const diagramRef = ref<HTMLElement | null>(null)
const lf = ref<any>(null)
const activeNodes = ref<any[]>([])
const activeEdges = ref<any[]>([])
const properties = ref<Record<string, any>>({})

const emit = defineEmits<{
    (e: 'gatewayStatus', status: { error: number; success: number }): void
}>()

function initLogicFlow(data: any) {
    // 引入框选插件
    // LogicFlow.use(SelectionSelect);
    // LogicFlow.use(Menu);
    LogicFlow.use(BpmnElement)
    const instance = new LogicFlow({
        container: diagramRef.value as HTMLElement,
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

    instance.setTheme({
        baseEdge: { strokeWidth: 1 },
        baseNode: { strokeWidth: 1 },
        nodeText: { overflowMode: 'autoWrap', lineHeight: 1.5, fontSize: 12 },
        edgeText: { overflowMode: 'autoWrap', lineHeight: 1.5, textWidth: 80, fontSize: 12 }
    })
    // 注册自定义元素
    registerCustomElement(instance)
    instance.setDefaultEdgeType('pro-polyline')
    instance.render(data)
    instance.fitView()

    lf.value = instance
}

function setType(errorMap: any) {
    const gatwayStatus = {
        error: 0,
        success: 0
    }
    const all = lf.value!.getGraphRawData()
    all.nodes.forEach((item: any) => {
        if (!item.type.includes('status')) {
            return
        }
        if (
            item.properties.deviceType == 1 &&
            errorMap.T80ErrorArr.some((i: any) => i == item.properties.equipmentId)
        ) {
            lf.value!.changeNodeType(item.id, 'status_error')
        } else if (
            item.properties.deviceType == 2 &&
            errorMap.ZLErrorArr.some((i: any) => i == item.properties.equipmentId)
        ) {
            lf.value!.changeNodeType(item.id, 'status_error')
        } else if (
            item.properties.deviceType == 3 &&
            errorMap.TimeOutErrorArr.some((i: any) => i == item.properties.equipmentId)
        ) {
            lf.value!.changeNodeType(item.id, 'status_error')
        } else if (
            item.properties.deviceType == 3 ||
            item.properties.deviceType == 2 ||
            item.properties.deviceType == 1
        ) {
            lf.value!.changeNodeType(item.id, 'status_success')
        } else if (item.properties.deviceType == 4) {
            let equipments = item.properties.equipmentId.split(',')
            equipments = equipments.map((i: string) => i.trim())
            if (
                equipments.every((equipmentId: string) =>
                    errorMap.TimeOutErrorArr.includes(equipmentId)
                )
            ) {
                gatwayStatus.error++
                lf.value!.changeNodeType(item.id, 'status_error')
            } else {
                gatwayStatus.success++
                lf.value!.changeNodeType(item.id, 'status_success')
            }
        }
    })
    emit('gatewayStatus', gatwayStatus)
}

function setTypeByDeviceID(deviceID: any, itemID: any, status: number) {
    const all = lf.value!.getGraphRawData()
    all.nodes.forEach((item: any) => {
        if (item.properties.deviceType) {
            return
        }
        if (
            item.properties.deviceID == deviceID &&
            item.properties.itemID == itemID &&
            status == 1
        ) {
            lf.value!.changeNodeType(item.id, 'status_success')
        } else if (
            item.properties.deviceID == deviceID &&
            item.properties.itemID == itemID &&
            status == 0
        ) {
            lf.value!.changeNodeType(item.id, 'status_error')
        }
    })
}

function networkError() {
    const all = lf.value!.getGraphRawData()
    all.nodes.forEach((item: any) => {
        if (item.properties.deviceType) {
            return
        }
        if (item.properties.deviceID && item.properties.itemID) {
            lf.value!.changeNodeType(item.id, 'status_error')
        }
    })
}

defineExpose({
    initLogicFlow,
    setType,
    setTypeByDeviceID,
    networkError
})
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
