<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="100px" :rules="formRules">
                <el-form-item label="上级区域" prop="pid">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.pid"
                        :data="spaceTreeOptions"
                        clearable
                        node-key="id"
                        :props="{ label: 'name', disabled: 'disabled' }"
                        check-strictly
                        :default-expand-all="true"
                        placeholder="请选择上级区域"
                    />
                </el-form-item>
                <el-form-item label="区域名称" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入区域名称" clearable :maxlength="100" />
                </el-form-item>
                <el-form-item label="区域范围" prop="geojson">
                    <div class="w-full">
                        <div class="flex items-center gap-2">
                            <el-button type="primary" plain @click="handleOpenArea">
                                <template #icon><icon name="el-icon-MapLocation" /></template>
                                在地图上绘制区域
                            </el-button>
                            <el-tag v-if="formData.geojson" type="success">已绘制</el-tag>
                            <el-tag v-else type="info">未绘制</el-tag>
                            <el-button v-if="formData.geojson" link type="danger" @click="formData.geojson = ''">
                                清除
                            </el-button>
                        </div>
                        <div class="form-tips">用于在地图上圈定该区域的地理围栏范围（GeoJSON Polygon）</div>
                    </div>
                </el-form-item>
                <el-form-item label="排序" prop="sort">
                    <div>
                        <el-input-number v-model="formData.sort" :min="0" :max="9999" />
                        <div class="form-tips">默认为0，数值越大越排前</div>
                    </div>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入备注" clearable :maxlength="255" />
                </el-form-item>
            </el-form>
        </popup>
        <map-area-picker v-model="formData.geojson" ref="areaPickerRef" />
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    deviceSpaceAll,
    deviceSpaceAdd,
    deviceSpaceEdit,
    deviceSpaceDetail,
    type type_iot_device_space
} from '@/api/iot'
import Popup from '@/components/popup/index.vue'
import MapAreaPicker from '@/components/map-area-picker/index.vue'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const areaPickerRef = shallowRef<InstanceType<typeof MapAreaPicker>>()
const handleOpenArea = () => areaPickerRef.value?.open()
const mode = ref('add')
const popupTitle = computed(() => (mode.value == 'edit' ? '编辑区域' : '新增区域'))

const { state: formData, setState } = useReactiveWithReset<type_iot_device_space>({
    id: '',
    pid: '',
    name: '',
    sort: 0,
    geojson: '',
    remark: ''
})

const formRules = {
    name: [{ required: true, message: '请输入区域名称', trigger: ['blur'] }]
}

const flatList = ref<type_iot_device_space[]>([])
async function loadSpaceOptions() {
    flatList.value = await deviceSpaceAll()
}

const collectSelfAndDescendants = (list: any[], currentId: string): Set<string> => {
    const ids = new Set<string>()
    const findAndCollect = (nodes: any[]) => {
        for (const n of nodes) {
            if (n.id === currentId || ids.has(n.id)) {
                const collectSub = (sub: any) => {
                    ids.add(sub.id)
                    ;(sub.children || []).forEach(collectSub)
                }
                collectSub(n)
            } else if (n.children?.length) {
                findAndCollect(n.children)
            }
        }
    }
    findAndCollect(list)
    return ids
}
const markDisabled = (list: any[], disabledIds: Set<string>): any[] =>
    list.map((item) => ({
        ...item,
        disabled: disabledIds.has(item.id),
        children: item.children?.length ? markDisabled(item.children, disabledIds) : item.children
    }))

const spaceTreeOptions = computed(() => {
    if (!flatList.value.length) return []
    const tree = arrayToTree<any>(
        flatList.value.map((item) => ({ ...item })) as any[],
        ''
    )
    if (!formData.id) return tree
    const disabledIds = collectSelfAndDescendants(tree, formData.id as string)
    return markDisabled(tree, disabledIds)
})

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await deviceSpaceEdit(formData) : await deviceSpaceAdd(formData)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    loadSpaceOptions()
    popupRef.value?.open()
}

const setFormData = (data: Partial<type_iot_device_space>) => {
    setState(data)
}

const getDetail = async (row: type_iot_device_space) => {
    const data = await deviceSpaceDetail({ id: row.id as string })
    setState(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({ open, setFormData, getDetail })
</script>
