<!-- 系统日志 -->
<template>
    <div class="journal">
        <el-card class="border-none!" shadow="never">
            <el-form
                class="ls-form mb-[-16px]"
                :model="formData"
                inline
                label-width="90px"
                label-position="left"
            >
                <el-form-item label="管理员" class="w-[360px]">
                    <el-input
                        placeholder="请输入"
                        v-model="formData.email"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>

                <el-form-item label="访问方式" class="w-[360px]">
                    <el-select
                        v-model="formData.type"
                        placeholder="请选择"
                        :empty-values="[null, undefined]"
                    >
                        <el-option
                            v-for="(item, index) in visitType"
                            :key="index"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="来源IP" class="w-[360px]">
                    <el-input
                        placeholder="请输入"
                        v-model="formData.ip"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>

                <el-form-item label="访问时间" class="w-[360px]">
                    <daterange-picker
                        v-model:startTime="formData.start_time"
                        v-model:endTime="formData.end_time"
                    />
                </el-form-item>

                <el-form-item label="访问链接" class="w-[360px]">
                    <el-input
                        placeholder="请输入"
                        v-model="formData.url"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>

        <el-card class="border-none! mt-4" shadow="never" v-loading="pager.loading">
            <div style="height: calc(100vh - 350px)">
                <vxe-table
                    :data="pager.lists"
                    height="100%"
                    :row-config="{ keyField: 'id' }"
                    :border="'inner'"
                >
                    <vxe-column type="seq" title="序号" min-width="60" />
                    <vxe-column title="操作" field="title" width="160" />
                    <vxe-column title="管理员" field="email" min-width="120" />
                    <vxe-column title="访问链接" field="url" min-width="240">
                        <template #default="{ row }"> {{ row.type }}：{{ row.url }} </template>
                    </vxe-column>
                    <!-- <vxe-column title="访问方式" field="type" min-width="100" /> -->
                    <vxe-column title="来源IP" field="ip" min-width="160" />
                    <vxe-column title="错误信息" field="error" width="100" />
                    <vxe-column title="耗时(毫秒)" field="task_time" min-width="100" />
                    <vxe-column title="日志时间" field="create_time" width="170" />
                </vxe-table>
            </div>
            <div class="flex mt-4 justify-end">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { systemLogLists, type_system_log_resp } from '@/api/setting/system'
import { usePaging } from '@/hooks/usePaging'
defineOptions({
    name: 'journal'
})
// 查询表单
const formData = ref({
    email: '',
    url: '',
    ip: '',
    type: '',
    start_time: '',
    end_time: ''
})

// 访问方式
const visitType = ref<Array<any>>([
    {
        label: '全部',
        value: ''
    },
    {
        label: 'get',
        value: 'GET'
    },
    {
        label: 'post',
        value: 'POST'
    }
])

const { pager, getLists, resetParams, resetPage, handleIndex } = usePaging<type_system_log_resp>({
    fetchFun: systemLogLists,
    params: formData.value
})

getLists()
</script>

<style lang="scss" scoped></style>
