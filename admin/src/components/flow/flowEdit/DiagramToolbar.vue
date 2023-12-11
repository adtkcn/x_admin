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
            <el-select v-model="linetype" @change="$_changeLineType">
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

<script>
// import { Sketch } from 'vue-color'
// import ColorFill from './icon/ColorFill.vue'
// import ColorText from './icon/ColorText.vue'
// import IconFont from './icon/Font.vue'
// import IconBlod from './icon/Blod.vue'
// import IconLine from './icon/Line.vue'
import ZoomIn from './icon/ZoomIn.vue'
import ZoomOut from './icon/ZoomOut.vue'
import StepBack from './icon/StepBack.vue'
import StepFoward from './icon/StepFoward.vue'
import AreaSelect from './icon/AreaSelect.vue'

let fileHandle

async function getFile() {
    ;[fileHandle] = await window.showOpenFilePicker()
    console.log('fileHandle', fileHandle)
}

async function getText() {
    const file = await fileHandle.getFile()
    const text = await file.text()
    console.log(text)
    return text
}

async function writeFile() {
    const writable = await fileHandle.createWritable()
    await writable.write('测试一下')
    await writable.close()
}

export default {
    props: {
        lf: Object,
        activeEdges: Array,
        fillColor: {
            type: String,
            default: ''
        }
    },
    data() {
        return {
            selectionOpened: false,
            undoAble: false,
            redoAble: false,
            colors: '#345678',
            linetype: 'pro-polyline',
            lineOptions: [
                {
                    value: 'pro-polyline',
                    label: '折线'
                },
                {
                    value: 'pro-line',
                    label: '直线'
                },
                {
                    value: 'pro-bezier',
                    label: '曲线'
                }
            ]
        }
    },
    mounted() {
        this.$props.lf.on('history:change', ({ data: { undoAble, redoAble } }) => {
            this.$data.redoAble = redoAble
            this.$data.undoAble = undoAble
        })
    },
    methods: {
        async $import() {
            try {
                await getFile()
                const text = await getText()
                if (JSON.parse(text)) {
                    this.$emit('importData', JSON.parse(text))
                }
            } catch (error) {
                this.$message.error('文件读取错误')
            }
        },
        $_changeFillColor(val) {
            this.$emit('changeNodeFillColor', val.hex)
        },
        $_saveGraph() {
            this.$emit('saveGraph')
        },
        $_zoomIn() {
            this.$props.lf.zoom(true)
        },
        $_zoomOut() {
            this.$props.lf.zoom(false)
        },
        $_undo() {
            if (this.$data.undoAble) {
                this.$props.lf.undo()
            }
        },
        $_redo() {
            if (this.$data.redoAble) {
                this.$props.lf.redo()
            }
        },
        $_selectionSelect() {
            this.selectionOpened = !this.selectionOpened
            if (this.selectionOpened) {
                this.lf.extension.selectionSelect.openSelectionSelect()
            } else {
                this.lf.extension.selectionSelect.closeSelectionSelect()
            }
        },
        $_changeLineType(value) {
            const { lf, activeEdges } = this.$props
            const { graphModel } = lf
            lf.setDefaultEdgeType(value)
            if (activeEdges && activeEdges.length > 0) {
                activeEdges.forEach((edge) => {
                    graphModel.changeEdgeType(edge.id, value)
                })
            }
        }
    },
    components: {
        // ColorFill,
        // ColorText,
        // IconFont,
        // IconBlod,
        // IconLine,
        ZoomIn,
        ZoomOut,
        StepBack,
        StepFoward,
        AreaSelect
        // SketchPicker: Sketch
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
