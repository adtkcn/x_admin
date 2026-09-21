import {
    fileCateAdd,
    fileCateDelete,
    fileCateEdit,
    fileCateLists,
    fileDelete,
    fileList,
    fileMove,
    fileRename,
    type type_file_cate
} from '@/api/file'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import { ElMessage, ElTree, type CheckboxValueType } from 'element-plus'
import { shallowRef, ref, reactive } from 'vue'
import type { Ref } from 'vue'

// 左侧分组的钩子函数
export function useCate() {
    const treeRef = shallowRef<InstanceType<typeof ElTree>>()
    // 分组列表
    const cateLists = ref<type_file_cate[]>([])

    // 选中的分组id
    const cateId = ref<string>('')

    // 获取分组列表
    const getCateLists = async () => {
        const data = await fileCateLists({})
        const item: type_file_cate[] = [
            {
                name: '全部',
                id: ''
            }
        ]
        cateLists.value = data
        cateLists.value.unshift(...item)
        setTimeout(() => {
            treeRef.value?.setCurrentKey(cateId.value)
        }, 0)
    }

    // 添加分组
    const handleAddCate = async (value: string) => {
        await fileCateAdd({
            name: value,
            pid: ''
        })
        getCateLists()
    }

    // 编辑分组
    const handleEditCate = async (value: string, id: string) => {
        await fileCateEdit({
            id,
            name: value
        })
        getCateLists()
    }

    // 删除分组
    const handleDeleteCate = async (id: string) => {
        await feedback.confirm('确定要删除？')
        await fileCateDelete({ id })
        cateId.value = ''
        getCateLists()
    }

    //选中分类
    const handleCatSelect = (item: any) => {
        cateId.value = item.id
    }

    return {
        treeRef,
        cateId,
        cateLists,
        handleAddCate,
        handleEditCate,
        handleDeleteCate,
        getCateLists,
        handleCatSelect
    }
}

// 处理文件的钩子函数
export function useFile(
    cateId: Ref<string | number>,
    ext: Ref<string[]>,
    limit: number,
    size: number
) {
    const tableRef = shallowRef()
    const moveId = ref('')
    const select = ref<any[]>([])
    const isCheckAll = ref(false)
    const isIndeterminate = ref(false)
    const fileParams = reactive({
        name: '',
        ext: ext,
        cid: cateId
    })

    const { pager, getLists, resetPage } = usePaging({
        fetchFun: fileList,
        params: fileParams,
        firstLoading: true,
        size
    })

    const getFileList = () => {
        getLists()
    }
    const refresh = () => {
        resetPage()
    }

    const batchFileDelete = async (id?: number[]) => {
        try {
            await feedback.confirm('从相册中删除，但不会删除文件')
            const ids = id ? id : select.value.map((item: any) => item.id)
            await fileDelete({ ids })
            getFileList()
            clearSelect()
        } catch (error) {
            console.log(error)
        }
    }

    const batchFileMove = async () => {
        const ids = select.value.map((item: any) => item.id)
        await fileMove({ ids, cid: moveId.value })
        moveId.value = ''
        getFileList()
        clearSelect()
    }

    const selectFile = (item: any) => {
        const index = select.value.findIndex((items: any) => items.id == item.id)
        if (index != -1) {
            select.value.splice(index, 1)
            return
        }
        if (select.value.length == limit) {
            if (limit == 1) {
                select.value = []
                select.value.push(item)
                return
            }
            ElMessage.warning('已达到选择上限')
            return
        }
        select.value.push(item)
    }

    const clearSelect = () => {
        select.value = []
    }

    const cancelSelect = (id: string) => {
        select.value = select.value.filter((item: any) => item.id != id)
    }
    const selectItems = (items: any[]) => {
        select.value = items
    }
    const selectAll = (value: CheckboxValueType) => {
        isIndeterminate.value = false
        tableRef.value?.toggleAllSelection()
        if (value) {
            select.value = [...pager.lists]
            return
        }
        clearSelect()
    }

    const handleFileRename = async (value: string, id: string) => {
        await fileRename({
            id,
            name: value
        })
        getFileList()
    }
    return {
        tableRef,
        moveId,
        pager,
        fileParams,
        select,
        isCheckAll,
        isIndeterminate,
        getFileList,
        refresh,
        batchFileDelete,
        batchFileMove,
        selectFile,
        clearSelect,
        cancelSelect,
        selectAll,
        selectItems,
        handleFileRename
    }
}
