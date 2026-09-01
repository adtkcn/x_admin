<template>
    <!-- v-loading="pager.loading" -->
    <div class="material" v-loading="pager.loading">
        <div class="material__left">
            <div class="flex-1 min-h-0">
                <el-scrollbar>
                    <div class="material-left__content pt-4 pb-4">
                        <el-tree
                            ref="treeRef"
                            node-key="id"
                            :data="cateLists"
                            empty-text=""
                            :highlight-current="true"
                            :expand-on-click-node="false"
                            :current-node-key="cateId"
                            @node-click="handleCatSelect"
                        >
                            <template v-slot="{ data }">
                                <div class="flex flex-1 items-center min-w-0 pr-4">
                                    <img
                                        class="w-[20px] h-[16px] mr-3"
                                        src="@/assets/images/icon_folder.png"
                                    />
                                    <span class="flex-1 truncate mr-2">
                                        <overflow-tooltip :content="data.name" />
                                    </span>
                                    <span v-if="data.id === ''">
                                        <popover-input
                                            v-perms="['admin:common:album:cateAdd']"
                                            @confirm="handleAddCate"
                                            size="default"
                                            width="500px"
                                            :limit="20"
                                            show-limit
                                            teleported
                                            placeholder="新分组名称"
                                        >
                                            <span class="p-1 mr-[5px]">+</span>
                                        </popover-input>
                                    </span>
                                    <el-dropdown
                                        v-perms="[
                                            'admin:common:album:cateRename',
                                            'admin:common:album:cateDel'
                                        ]"
                                        v-if="data.id != 0"
                                        :hide-on-click="false"
                                    >
                                        <span class="p-1 mr-[5px]">···</span>
                                        <template #dropdown>
                                            <el-dropdown-menu>
                                                <popover-input
                                                    v-perms="['admin:common:album:cateRename']"
                                                    @confirm="handleEditCate($event, data.id)"
                                                    size="default"
                                                    :value="data.name"
                                                    width="400px"
                                                    :limit="20"
                                                    show-limit
                                                    teleported
                                                >
                                                    <div>
                                                        <el-dropdown-item>
                                                            命名分组
                                                        </el-dropdown-item>
                                                    </div>
                                                </popover-input>
                                                <div
                                                    v-perms="['admin:common:album:cateDel']"
                                                    @click="handleDeleteCate(data.id)"
                                                >
                                                    <el-dropdown-item>删除分组</el-dropdown-item>
                                                </div>
                                            </el-dropdown-menu>
                                        </template>
                                    </el-dropdown>
                                </div>
                            </template>
                        </el-tree>
                    </div>
                </el-scrollbar>
            </div>
        </div>
        <div class="material__center flex flex-col">
            <el-tabs
                v-if="defaultFileType == 'all'"
                v-model="activeFileType"
                @tab-change="handleTabChange"
            >
                <el-tab-pane
                    v-for="item in FileTabsMap"
                    :label="item.name"
                    :name="item.fileType"
                    :key="item.fileType"
                >
                </el-tab-pane>
            </el-tabs>
            <div class="flex flex-col">
                <div class="operate-btn flex">
                    <div class="flex-1"></div>
                    <el-form ref="formRef" class="mb-[-20px]" :inline="true" label-position="right">
                        <!-- <el-form-item class="w-[270px]">
                            <el-select
                                v-model="activeFileType"
                                clearable
                                :empty-values="[null, undefined]"
                                @change="handleTabChange"
                            >
                                <el-option label="全部" value="" />
                                <el-option
                                    v-for="(item, index) in FileTabsMap"
                                    :key="index"
                                    :label="item.name"
                                    :value="item.fileType"
                                />
                            </el-select>
                        </el-form-item> -->
                        <el-form-item>
                            <el-input
                                placeholder="请输入名称"
                                v-model="fileParams.name"
                                @keyup.enter="refresh"
                            >
                                <template #append>
                                    <el-button @click="refresh">
                                        <template #icon>
                                            <icon name="el-icon-Search" />
                                        </template>
                                    </el-button>
                                </template>
                            </el-input>
                        </el-form-item>
                        <el-form-item>
                            <Upload :ext="ext" :show-progress="true" @change="handleUpload">
                                <el-button type="primary">本地上传</el-button>
                            </Upload>
                        </el-form-item>
                    </el-form>
                </div>
            </div>

            <div class="material-center__content flex flex-col flex-1 mb-1 min-h-0 overflow-hidden">
                <vxe-table
                    ref="tableRef"
                    class="mt-4"
                    :data="pager.lists"
                    height="auto"
                    :row-config="{ keyField: 'id' }"
                    :checkbox-config="{ checkRowKeys: [] }"
                    @checkbox-change="selectItems($event.$table.getCheckboxRecords())"
                    @checkbox-all="selectItems($event.$table.getCheckboxRecords())"
                    :border="'inner'"
                >
                    <vxe-column type="checkbox" width="55" />
                    <vxe-column title="图片" width="100">
                        <template #default="{ row }">
                            <FileItem
                                :uri="row.uri + '?quality=80&scale_width=200&scale_height=200'"
                                file-size="50px"
                                @click.stop="handlePreview(row.uri)"
                            ></FileItem>
                        </template>
                    </vxe-column>
                    <vxe-column title="名称" min-width="100" show-overflow>
                        <template #default="{ row }">
                            <el-link @click.stop="handlePreview(row.uri)" underline="never">
                                {{ row.name }}
                            </el-link>
                        </template>
                    </vxe-column>
                    <vxe-column title="大小" field="size" width="120" v-if="mode == 'page'">
                    </vxe-column>
                    <!-- <vxe-column
                        title="格式"
                        field="ext"
                        width="100"
                        v-if="mode == 'page'"
                    ></vxe-column> -->

                    <vxe-column
                        field="create_time"
                        title="上传时间"
                        width="170"
                        v-if="mode == 'page'"
                    />
                    <vxe-column title="操作" width="150" fixed="right">
                        <template #default="{ row }">
                            <div
                                class="inline-block mr-2"
                                v-perms="['admin:common:album:albumRename']"
                            >
                                <popover-input
                                    @confirm="handleFileRename($event, row.id)"
                                    size="default"
                                    :value="row.name"
                                    width="500px"
                                    :limit="50"
                                    show-limit
                                    teleported
                                >
                                    <el-link type="primary" link> 重命名 </el-link>
                                </popover-input>
                            </div>
                            <!-- <div class="inline-block mr-2">
                                <el-link type="primary" link @click.stop="handlePreview(row.uri)">
                                    查看
                                </el-link>
                            </div> -->

                            <div
                                class="inline-block mr-2"
                                v-perms="['admin:common:album:albumDel']"
                            >
                                <el-link
                                    type="primary"
                                    link
                                    @click.stop="batchFileDelete([row.id])"
                                >
                                    删除
                                </el-link>
                            </div>
                            <div class="inline-block mr-2" v-if="mode == 'page'">
                                <el-link type="primary" :href="row.uri" download>下载</el-link>
                            </div>
                        </template>
                    </vxe-column>
                </vxe-table>
            </div>
            <div class="material-center__footer flex justify-between items-center mt-2">
                <div class="flex">
                    <template v-if="mode == 'page'">
                        <el-button
                            v-perms="['admin:common:album:albumDel']"
                            :disabled="!select.length"
                            @click="batchFileDelete()"
                        >
                            删除
                        </el-button>
                        <popup
                            v-perms="['admin:common:album:albumMove']"
                            class="ml-3 inline"
                            :disabled="!select.length"
                            title="移动文件"
                            @confirm="batchFileMove"
                        >
                            <template #trigger>
                                <el-button :disabled="!select.length">移动</el-button>
                            </template>

                            <el-select v-model="moveId" placeholder="请选择">
                                <template v-for="item in cateLists" :key="item.id">
                                    <el-option
                                        v-if="item.id !== ''"
                                        :label="item.name"
                                        :value="item.id"
                                    ></el-option>
                                </template>
                            </el-select>
                        </popup>
                    </template>
                </div>
                <pagination
                    v-model="pager"
                    @change="getFileList"
                    layout="total, prev, pager, next, jumper"
                />
            </div>
        </div>
        <Preview v-model="showPreview" :url="previewUrl" />
    </div>
</template>

<script lang="ts" setup>
import { FileTabsMap } from '@/enums/fileEnums'
import { onMounted, ref, watch, computed, defineAsyncComponent } from 'vue'
import { FileExt } from '@/enums/fileEnums'

import { useCate, useFile } from './hook'
import { albumAddFromFile } from '@/api/file'
import feedback from '@/utils/feedback'
import FileItem from './file.vue'
// import Preview from './preview.vue'
const Preview = defineAsyncComponent(() => import('./preview.vue'))
const props = defineProps({
    limit: {
        type: Number,
        default: 1
    },
    defaultFileType: {
        type: String,
        default: 'all'
    },
    mode: {
        type: String,
        default: 'picker'
    },
    pageSize: {
        type: Number,
        default: 15
    }
})
const emit = defineEmits(['change'])
// const { limit } = toRefs(props)

const previewUrl = ref('')
const showPreview = ref(false)
const {
    treeRef,
    cateId,
    cateLists,
    handleAddCate,
    handleEditCate,
    handleDeleteCate,
    getCateLists,
    handleCatSelect
} = useCate()

const activeFileType = ref<keyof typeof FileExt>(props.defaultFileType as keyof typeof FileExt)

const ext = computed(() => {
    if (activeFileType.value) {
        return FileExt[activeFileType.value]
    }
    return []
})
const listExt = computed(() => {
    if (activeFileType.value == 'all') {
        return []
    }
    if (activeFileType.value) {
        return FileExt[activeFileType.value]
    }
    return []
})
const {
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

    clearSelect,
    cancelSelect,

    selectItems,
    handleFileRename
} = useFile(cateId, listExt, props.limit, props.pageSize)
const handleTabChange = () => {
    // activeFileType.value = tab
    console.log('handleTabChange', activeFileType.value)

    getFileList()
}
const getData = async () => {
    try {
        await getCateLists()
        treeRef.value?.setCurrentKey(cateId.value)
        getFileList()
    } catch (error) {
        console.error('素材分类获取失败:', error)
    }
}

const handlePreview = (url: string) => {
    previewUrl.value = url
    showPreview.value = true
}

// 上传成功后把文件以挂载方式写入相册（不耦合到通用上传组件）
const handleUpload = async (fileLists: any[]) => {
    if (fileLists?.length) {
        const cid = cateId.value
        const tasks = fileLists
            .filter((item) => item.response?.data?.file_hash_id)
            .map((item) =>
                albumAddFromFile({
                    file_hash_id: item.response.data.file_hash_id,
                    cid: String(cid),
                    file_name: item.name
                })
            )
        try {
            await Promise.all(tasks)
        } catch (e: any) {
            feedback.msgError(e?.message || '加入相册失败')
        }
    }
    refresh()
}

watch(cateId, () => {
    fileParams.name = ''
    refresh()
})

watch(
    select,
    (val: any[]) => {
        emit('change', val)
        if (val.length == pager.lists.length && val.length !== 0) {
            isIndeterminate.value = false
            isCheckAll.value = true
            return
        }
        if (val.length > 0) {
            isIndeterminate.value = true
        } else {
            isCheckAll.value = false
            isIndeterminate.value = false
        }
    },
    {
        deep: true
    }
)

onMounted(() => {
    // props.mode == 'page' && getData()
    getData()
})

defineExpose({
    clearSelect,
    cancelSelect
})
</script>

<style scoped lang="scss">
.material {
    height: 100%;
    min-height: 0px;
    display: flex;
    flex: 1 1 0%;
    &__left {
        border-right-width: 1px;
        border-color: var(--el-border-color);
        display: flex;
        flex-direction: column;
        width: 200px;
        :deep(.el-tree-node__content) {
            height: 36px;
        }
    }
    &__center {
        flex: 1;
        min-width: 0;
        min-height: 0;
        padding: 16px 16px 0;
        width: 100%;
        .list-icon {
            border-radius: 3px;
            display: flex;
            padding: 5px;
            cursor: pointer;
            &.select {
                color: var(--el-color-primary);
                background-color: var(--el-color-primary-light-8);
            }
        }
        .file-list {
            .file-item-wrap {
                margin-right: 16px;
                line-height: 1.3;
                cursor: pointer;
                .item-selected {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    position: absolute;
                    top: 0;
                    left: 0;
                    width: 100%;
                    height: 100%;
                    border-radius: 4px;
                    background-color: rgba(0, 0, 0, 0.5);
                    box-sizing: border-box;
                }
                .operation-btns {
                    height: 28px;
                    visibility: hidden;
                }
                &:hover .operation-btns {
                    visibility: visible;
                }
            }
        }
    }
}
</style>
