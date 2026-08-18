<template>
    <div class="material-picker">
        <popup
            ref="popupRef"
            width="90%"
            custom-class="body-padding"
            :title="`选择${tipsText}`"
            :async="true"
            @confirm="handleConfirm"
            @close="handleClose"
        >
            <template v-if="!hiddenUpload" #trigger>
                <div class="material-select__trigger clearfix" @click.stop>
                    <draggable class="draggable" v-model="fileList" animation="300" item-key="id">
                        <template v-slot:item="{ element, index }">
                            <div
                                class="material-preview"
                                :class="{ 'is-disabled': disabled, 'is-one': limit == 1 }"
                                @click="showPopup(index)"
                            >
                                <del-wrap @close="deleteImg(index)">
                                    <FileItem
                                        :uri="excludeDomain ? getImageUrl(element) : element"
                                        :file-size="size"
                                    ></FileItem>
                                </del-wrap>
                                <div class="operation-btns text-xs text-center">
                                    <span>修改</span>
                                    |
                                    <span @click.stop="handlePreview(element)">查看</span>
                                </div>
                            </div>
                        </template>
                    </draggable>
                    <div
                        class="material-upload"
                        @click="showPopup(-1)"
                        v-show="showUpload"
                        :class="{
                            'is-disabled': disabled,
                            'is-one': limit == 1,
                            [uploadClass]: true
                        }"
                    >
                        <slot name="upload">
                            <div
                                class="text-tx-secondary box-border rounded-sm border-br border-dashed border flex flex-col justify-center items-center"
                                :style="{ width: size, height: size }"
                            >
                                <icon :size="25" name="el-icon-Plus" />
                                <span>添加</span>
                            </div>
                        </slot>
                    </div>
                </div>
            </template>

            <div class="material-wrap">
                <el-tabs v-model="activeTab" class="material-picker-tabs">
                    <el-tab-pane label="文件上传" name="upload">
                        <upload-dropzone
                            class="material-local-upload"
                            :ext="uploadExt"
                            :multiple="limit != 1"
                            :limit="materialLimit"
                            @change="handleLocalUpload"
                        />
                    </el-tab-pane>
                    <el-tab-pane label="素材选择" name="library">
                        <material
                            ref="materialRef"
                            mode="picker"
                            :default-file-type="type"
                            :ext="ext"
                            :file-size="fileSize"
                            :limit="materialLimit"
                            @change="selectChange"
                        />
                    </el-tab-pane>
                </el-tabs>
                <div class="picker-selected">
                    <div class="picker-selected__header">
                        <span>
                            已选择 {{ select.length }}
                            <template v-if="limit > 0">/{{ limit }}</template>
                        </span>
                        <el-button type="primary" link @click="clearSelect">清空</el-button>
                    </div>
                    <el-scrollbar class="picker-selected__scroll">
                        <ul v-if="select.length" class="picker-selected__list">
                            <li v-for="item in select" :key="item.id" class="picker-selected__item">
                                <del-wrap @close="removeSelect(item)">
                                    <FileItem :uri="item.uri" :file-size="'80px'" />
                                </del-wrap>
                            </li>
                        </ul>
                        <div v-else class="picker-selected__empty">暂未选择文件</div>
                    </el-scrollbar>
                </div>
            </div>
        </popup>

        <preview v-model="showPreview" :url="previewUrl" />
    </div>
</template>

<script lang="ts" name="material-picker">
import { defineComponent, ref, computed, toRefs, watch, nextTick, provide } from 'vue'
import Draggable from 'vuedraggable'
import Popup from '@/components/popup/index.vue'
import FileItem from './file.vue'
import Material from './index.vue'
import Preview from './preview.vue'
import useAppStore from '@/stores/modules/app'
import { useThrottleFn } from '@vueuse/core'
import feedback from '@/utils/feedback'
import { FileExt } from '@/enums/fileEnums'
import UploadDropzone from '@/components/upload/dropzone.vue'

export default defineComponent({
    components: {
        Popup,
        Draggable,
        FileItem,
        Material,
        Preview,
        UploadDropzone
    },
    props: {
        modelValue: { type: [String, Array], default: () => [] },
        // 选择类型：image|video|audio|office|file
        type: { type: String, default: 'image' },
        size: { type: String, default: '100px' },
        fileSize: { type: String, default: '100px' },
        limit: { type: Number, default: 1 },
        disabled: { type: Boolean, default: false },
        hiddenUpload: { type: Boolean, default: false },
        uploadClass: { type: String, default: '' },
        excludeDomain: { type: Boolean, default: false }
    },
    emits: ['change', 'update:modelValue'],
    setup(props, { emit }) {
        const popupRef = ref<InstanceType<typeof Popup>>()
        const materialRef = ref<InstanceType<typeof Material>>()
        const previewUrl = ref('')
        const showPreview = ref(false)
        const fileList = ref<any[]>([])
        const isAdd = ref(true)
        const currentIndex = ref(-1)
        const activeTab = ref('upload')
        // 本地上传（文件上传 Tab）与素材库选择分别记录，避免切换 Tab 互相覆盖
        const localFiles = ref<any[]>([])
        const materialFiles = ref<any[]>([])
        const { disabled, limit, modelValue } = toRefs(props)
        const { getImageUrl } = useAppStore()

        const ext = computed(() => FileExt[props.type as keyof typeof FileExt])
        // 文件上传 Tab 仅支持图片
        const uploadExt = computed(() => FileExt.image)
        const tipsText = computed(() => {
            switch (props.type) {
                case 'image':
                    return '图片'
                case 'video':
                    return '视频'
                case 'audio':
                    return '音频'
                case 'office':
                    return '文档'
                case 'file':
                    return '文件'
                default:
                    return ''
            }
        })
        const select = computed(() => [...localFiles.value, ...materialFiles.value])

        const showUpload = computed(() => props.limit - fileList.value.length > 0)
        const materialLimit: any = computed(() => {
            if (!isAdd.value) return 1
            if (limit.value == -1) return null
            return limit.value - fileList.value.length
        })
        const handleConfirm = useThrottleFn(
            () => {
                // 数量拦截：未选择任何文件时一律阻止关闭
                if (select.value.length === 0) {
                    feedback.msgError('请先选择文件')
                    return
                }
                // 数量限制：配置了上限时，选择数不足或超出都阻止关闭弹窗
                if (props.limit > 0) {
                    if (select.value.length < props.limit) {
                        feedback.msgError(
                            `请至少选择 ${props.limit} 个文件，当前已选 ${select.value.length} 个`
                        )
                        return
                    }
                    if (select.value.length > props.limit) {
                        feedback.msgError(
                            `最多选择 ${props.limit} 个文件，当前已选 ${select.value.length} 个`
                        )
                        return
                    }
                }
                const selectUri = select.value.map((item) =>
                    props.excludeDomain ? item.path : item.uri
                )
                if (!isAdd.value) {
                    fileList.value.splice(currentIndex.value, 1, selectUri.shift())
                } else {
                    fileList.value = [...fileList.value, ...selectUri]
                }
                popupRef.value?.close()
                handleChange()
            },
            1000,
            false
        )
        const showPopup = (index: number) => {
            if (disabled.value) return
            if (index >= 0) {
                isAdd.value = false
                currentIndex.value = index
            } else {
                isAdd.value = true
            }
            popupRef.value?.open()
        }
        const selectChange = (val: any[]) => {
            // 父组件持有素材库已选的权威数据，避免子组件翻页/刷新导致
            // selection-change 传入当前页或空数组时覆盖(清空)父已选。
            // 策略：仅将子组件主动勾选的新项增量合并(按 id 去重)；
            // 移除统一由右侧栏 removeSelect 主动触发。
            if (props.limit === 1) {
                materialFiles.value = val.length ? [val[val.length - 1]] : []
                return
            }
            val.forEach((item: any) => {
                if (!materialFiles.value.some((it: any) => it.id === item.id)) {
                    materialFiles.value.push(item)
                }
            })
        }
        // 本地上传完成：归一为 { uri, path } 累加进选择，不自动关闭以便继续添加
        const handleLocalUpload = (fileLists: any[]) => {
            const mapped = (fileLists || []).map((item) => ({
                uri: item.response?.data?.url,
                path: item.response?.data?.file_path
            }))
            localFiles.value = [...localFiles.value, ...mapped]
            activeTab.value = 'upload'
        }
        // 从已选择列表中移除某一项（本地上传与素材库分别处理）
        const removeSelect = (item: any) => {
            localFiles.value = localFiles.value.filter((it: any) => it.path !== item.path)
            materialRef.value?.cancelSelect(item.id)
        }
        // 清空所有已选择
        const clearSelect = () => {
            localFiles.value = []
            materialRef.value?.clearSelect()
        }
        const handleChange = () => {
            const valueImg = limit.value != 1 ? fileList.value : fileList.value[0] || ''
            emit('update:modelValue', valueImg)
            emit('change', valueImg)
            handleClose()
        }
        const deleteImg = (index: number) => {
            fileList.value.splice(index, 1)
            handleChange()
        }
        const handlePreview = (url: string) => {
            previewUrl.value = url
            showPreview.value = true
        }
        const handleClose = () => {
            nextTick(() => {
                if (props.hiddenUpload) fileList.value = []
                localFiles.value = []
                materialFiles.value = []
                materialRef.value?.clearSelect()
            })
        }
        watch(
            modelValue,
            (val: any[] | string) => {
                fileList.value = Array.isArray(val) ? val : val == '' ? [] : [val]
            },
            { immediate: true }
        )
        provide('limit', props.limit)
        provide('hiddenUpload', props.hiddenUpload)
        return {
            ext,
            uploadExt,
            popupRef,
            materialRef,
            fileList,
            tipsText,
            activeTab,
            localFiles,
            select,
            handleConfirm,
            handleLocalUpload,
            removeSelect,
            clearSelect,
            materialLimit,
            showUpload,
            showPopup,
            selectChange,
            deleteImg,
            previewUrl,
            showPreview,
            handlePreview,
            handleClose,
            getImageUrl
        }
    }
})
</script>

<style scoped lang="scss">
.material-select__trigger {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    .material-upload,
    .material-preview {
        margin: 10px;
    }
}
.material-preview {
    box-sizing: content-box;
    position: relative;
    border-radius: 4px;
    cursor: pointer;
    .operation-btns {
        position: absolute;
        bottom: 0;
        left: 0;
        width: 100%;
        color: #fff;
        background-color: rgba(#000, 0.5);
        border-radius: 0 0 4px 4px;
        > span {
            cursor: pointer;
        }
    }
    &.is-disabled {
        cursor: not-allowed;
    }
    &.is-one {
        margin-left: 0;
    }
}
.material-upload {
    flex: none;
    cursor: pointer;
    &.is-disabled {
        cursor: not-allowed;
    }
    &.is-one {
        margin-left: 0;
    }
}
.material-wrap {
    height: calc(100vh - 170px);
    display: flex;
    flex-direction: row;
    gap: 16px;
}
.picker-selected {
    flex: none;
    width: 220px;
    border-left: 1px solid #ebeef5;
    padding-left: 16px;
    display: flex;
    flex-direction: column;
    min-height: 0;
    &__header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        font-size: 13px;
        color: #909399;
        margin-bottom: 12px;
    }
    &__scroll {
        flex: 1;
        min-height: 0;
    }
    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: 12px;
        padding: 0;
        margin: 0;
        list-style: none;
    }
    &__item {
        position: relative;
        border: 1px solid #ebeef5;
        border-radius: 4px;
        overflow: hidden;
    }
    &__empty {
        color: #c0c4cc;
        font-size: 13px;
        text-align: center;
        padding-top: 40px;
    }
}
.material-picker-tabs {
    flex: 1;
    min-width: 0;
    min-height: 0;
    display: flex;
    flex-direction: column;
    :deep(.el-tabs__content) {
        height: calc(100% - 55px);
        overflow: hidden;
    }
    :deep(.el-tab-pane) {
        height: 100%;
    }
    // 素材库组件根节点撑满 tab-pane，保证内部 vxe-table height="auto" 能取到确定父高
    // 注意：.material 自身是横向 flex（左侧分类 + 右侧表格），此处只补高度，勿改 flex-direction
    :deep(.material) {
        height: 100%;
    }
}
</style>
