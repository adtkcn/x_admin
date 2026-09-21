<template>
    <div>
        <div
            class="toolbar-item"
            :class="{ 'selection-active': selectionOpened }"
            @click="$_selectionSelect()"
        >
            <area-select size="18" />
        </div>
        <div class="toolbar-item" @click="$_zoomIn()">
            <zoom-in size="18" />
        </div>
        <div class="toolbar-item" @click="$_zoomOut()">
            <zoom-out size="18" />
        </div>
        <div class="toolbar-item" :class="{ disabled: !undoAble }" @click="$_undo()">
            <step-back size="18" />
        </div>
        <div class="toolbar-item" :class="{ disabled: !redoAble }" @click="$_redo()">
            <step-foward size="18" />
        </div>
        <div class="toolbar-item" @click="$import">导入</div>

        <div class="toolbar-item" @click="$_saveGraph">导出</div>
        <div>
            <el-select v-model="linetype" @change="$_changeLineType" style="width: 80px">
                <el-option
                    v-for="item in lineOptions"
                    :key="item.value"
                    :value="item.value"
                    :label="item.label"
                ></el-option>
            </el-select>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import ZoomIn from './icon/ZoomIn.vue'
import ZoomOut from './icon/ZoomOut.vue'
import StepBack from './icon/StepBack.vue'
import StepFoward from './icon/StepFoward.vue'
import AreaSelect from './icon/AreaSelect.vue'

const props = defineProps<{
    lf: any
    activeEdges: any[]
    fillColor?: string
}>()

const emit = defineEmits<{
    (e: 'importData', data: any): void
    (e: 'saveGraph'): void
}>()

const selectionOpened = ref(false)
const undoAble = ref(false)
const redoAble = ref(false)
const linetype = ref('pro-polyline')
const lineOptions = [
    { value: 'pro-polyline', label: '折线' },
    { value: 'pro-line', label: '直线' },
    { value: 'pro-bezier', label: '曲线' }
]

onMounted(() => {
    props.lf.on('history:change', ({ data: { undoAble: u, redoAble: r } }: any) => {
        undoAble.value = u
        redoAble.value = r
    })
})

async function $import() {
    try {
        const [handle] = await (window as any).showOpenFilePicker()
        console.log('fileHandle', handle)
        const file = await handle.getFile()
        const text = await file.text()
        const data = JSON.parse(text)
        if (data) {
            emit('importData', data)
        }
    } catch (error) {
        ElMessage.error('文件读取错误')
    }
}

function $_saveGraph() {
    emit('saveGraph')
}

function $_zoomIn() {
    props.lf.zoom(true)
}

function $_zoomOut() {
    props.lf.zoom(false)
}

function $_undo() {
    if (undoAble.value) {
        props.lf.undo()
    }
}

function $_redo() {
    if (redoAble.value) {
        props.lf.redo()
    }
}

function $_selectionSelect() {
    selectionOpened.value = !selectionOpened.value
    if (selectionOpened.value) {
        props.lf.extension.selectionSelect.openSelectionSelect()
    } else {
        props.lf.extension.selectionSelect.closeSelectionSelect()
    }
}

function $_changeLineType(value: any) {
    console.log('value', value)
    const { lf, activeEdges } = props
    const { graphModel } = lf
    lf.setDefaultEdgeType(value)
    if (activeEdges && activeEdges.length > 0) {
        activeEdges.forEach((edge: any) => {
            graphModel.changeEdgeType(edge.id, value)
        })
    }
}
</script>

<style scoped>
.toolbar-item {
    /* width: 18px;
  height: 18px; */
    float: left;
    margin: 12px 6px;
    padding: 5px;
    cursor: pointer;
}
.selection-active {
    background: rgba(0, 0, 0, 0.2);
}
</style>
