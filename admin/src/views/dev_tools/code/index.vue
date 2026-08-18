<template>
    <div class="code-generation">
        <el-card class="border-none!" shadow="never">
            <el-form class="mb-[-16px]" :model="formData" inline>
                <el-form-item label="表名称">
                    <el-input v-model="formData.table_name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="表描述">
                    <el-input v-model="formData.table_comment" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="border-none! mt-4" shadow="never" v-loading="pager.loading">
            <div class="flex">
                <data-table
                    v-perms="['admin:gen:importTable']"
                    class="inline-block mr-[10px]"
                    @success="getLists"
                >
                    <el-button type="primary">
                        <template #icon>
                            <icon name="el-icon-Plus" />
                        </template>
                        导入数据表
                    </el-button>
                </data-table>
                <el-button
                    v-perms="['admin:gen:delTable']"
                    :disabled="!selectData.length"
                    @click="handleDelete()"
                    type="danger"
                >
                    <template #icon>
                        <icon name="el-icon-Delete" />
                    </template>
                    删除
                </el-button>
                <!-- <el-button
                    v-perms="['admin:gen:genCode', 'admin:gen:downloadCode']"
                    :disabled="!selectData.length"
                    @click="handleGenerate(selectData)"
                >
                    生成代码
                </el-button> -->
            </div>
            <div class="mt-4">
                <vxe-table
                    ref="tableRef"
                    :data="pager.lists"
                    :row-config="{ keyField: 'id' }"
                    :checkbox-config="{ checkRowKeys: [] }"
                    @checkbox-change="selectData = $event.$table.getCheckboxRecords()"
                    @checkbox-all="selectData = $event.$table.getCheckboxRecords()"
                    :border="'inner'"
                >
                    <vxe-column type="checkbox" width="55" />
                    <vxe-column title="表名称" field="table_name" min-width="180" />
                    <vxe-column title="表描述" field="table_comment" min-width="180" />
                    <vxe-column title="创建时间" field="create_time" min-width="180" />
                    <vxe-column title="更新时间" field="update_time" min-width="180" />
                    <vxe-column title="操作" width="160" fixed="right">
                        <template #default="{ row }">
                            <div class="flex items-center">
                                <el-button
                                    v-perms="['admin:gen:previewCode']"
                                    type="primary"
                                    link
                                    @click="handlePreview(row.id)"
                                >
                                    预览
                                </el-button>

                                <el-button type="primary" link v-perms="['admin:gen:editTable']">
                                    <router-link
                                        :to="{
                                            path: 'code/edit',
                                            query: {
                                                id: row.id
                                            }
                                        }"
                                    >
                                        编辑
                                    </router-link>
                                </el-button>
                                <el-dropdown
                                    class="ml-2"
                                    @command="handleCommand($event, row)"
                                    v-perms="[
                                        'admin:gen:genCode',
                                        'admin:gen:downloadCode',
                                        'admin:gen:syncTable',
                                        'admin:gen:delTable'
                                    ]"
                                >
                                    <el-button type="primary" link>
                                        更多
                                        <icon name="el-icon-ArrowDown" :size="14" />
                                    </el-button>

                                    <template #dropdown>
                                        <el-dropdown-menu>
                                            <div
                                                v-perms="[
                                                    'admin:gen:genCode',
                                                    'admin:gen:downloadCode'
                                                ]"
                                            >
                                                <el-dropdown-item command="generate">
                                                    <el-button type="primary" link>
                                                        下载代码
                                                    </el-button>
                                                </el-dropdown-item>
                                            </div>
                                            <div v-perms="['admin:gen:syncTable']">
                                                <el-dropdown-item command="sync">
                                                    <el-button type="primary" link>
                                                        更新
                                                    </el-button>
                                                </el-dropdown-item>
                                            </div>
                                            <div v-perms="['admin:gen:delTable']">
                                                <el-dropdown-item command="delete">
                                                    <el-button type="danger" link> 删除 </el-button>
                                                </el-dropdown-item>
                                            </div>
                                        </el-dropdown-menu>
                                    </template>
                                </el-dropdown>
                            </div>
                        </template>
                    </vxe-column>
                </vxe-table>
            </div>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <CodePreview
            v-if="previewState.show"
            v-model="previewState.show"
            :code="previewState.code"
        />
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onActivated, defineAsyncComponent } from 'vue'
import {
    generateTable,
    syncColumn,
    generateDelete,
    generatePreview,
    downloadCode,
    type type_gen_table_list,
    type type_gen_table_resp
} from '@/api/tools/code'
import { usePaging } from '@/hooks/usePaging'
const DataTable = defineAsyncComponent(() => import('./components/data-table.vue'))
const CodePreview = defineAsyncComponent(() => import('./components/code-preview.vue'))
import feedback from '@/utils/feedback'
import { streamFileDownload } from '@/utils/file'
defineOptions({
    name: 'codeGenerate'
})
const formData = reactive<type_gen_table_list>({
    table_name: '',
    table_comment: ''
})

const previewState = reactive({
    show: false,
    loading: false,
    code: {}
})

const { pager, getLists, resetParams, resetPage } = usePaging({
    fetchFun: generateTable,
    params: formData
})

const tableRef = ref<any>()
const selectData = ref<type_gen_table_resp[]>([])

const handleSync = async (id: string) => {
    await feedback.confirm('确定要更新表结构？从数据库拉取最新表结构')
    await syncColumn({ id })
    feedback.msgSuccess('操作成功')
}

const handleDelete = async (ids?: string[]) => {
    if (!ids) ids = selectData.value.map(({ id }) => id)
    await feedback.confirm('确定要删除？')
    await generateDelete({ ids })
    feedback.msgSuccess('删除成功')
    getLists()
}

const handlePreview = async (id: string) => {
    const data: any = await generatePreview({ id })
    previewState.code = data
    previewState.show = true
}

const handleGenerate = async (selectData: type_gen_table_resp[]) => {
    const downloadTables = getTables(selectData)
    if (downloadTables) {
        const file = await downloadCode({ tables: downloadTables })
        streamFileDownload(file, 'code_' + downloadTables + '.zip')
    }
}

const getTables = (selectData: type_gen_table_resp[]) => {
    return selectData.map(({ table_name }) => table_name).join()
}

const handleCommand = (command: string, row: type_gen_table_resp) => {
    switch (command) {
        case 'generate':
            handleGenerate([row])
            break
        case 'sync':
            handleSync(row.id)
            break
        case 'delete':
            handleDelete([row.id])
    }
}

onActivated(() => {
    getLists()
})

getLists()
</script>
