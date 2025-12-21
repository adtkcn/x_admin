<!-- 系统日志 -->
<template>
    <div class="journal">
        <el-card class="!border-none" shadow="never">
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
                        v-model="formData.username"
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
                        v-model:startTime="formData.startTime"
                        v-model:endTime="formData.endTime"
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

        <el-card class="!border-none mt-4" shadow="never" v-loading="pager.loading">
            <div>
                <el-table :data="pager.lists" size="large" max-height="calc(100vh - 200px)">
                    <el-table-column
                        label="序号"
                        type="index"
                        :index="handleIndex"
                        min-width="60"
                    />
                    <el-table-column label="操作" prop="title" min-width="120" />
                    <el-table-column label="管理员" prop="username" min-width="120" />
                    <el-table-column label="访问链接" prop="url" min-width="240">
                        <template #default="{ row }"> {{ row.type }}：{{ row.url }} </template>
                    </el-table-column>
                    <!-- <el-table-column label="访问方式" prop="type" min-width="100" /> -->
                    <el-table-column label="来源IP" prop="ip" min-width="160" />
                    <el-table-column label="错误信息" prop="error" min-width="200" />
                    <el-table-column label="耗时(毫秒)" prop="taskTime" min-width="100" />
                    <el-table-column label="日志时间" prop="createTime" width="170" />
                </el-table>
            </div>
            <div class="flex mt-4 justify-end">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { systemLogLists, SystemLogResp } from '@/api/setting/system'
import { usePaging } from '@/hooks/usePaging'
defineOptions({
    name: 'journal'
})
// 查询表单
const formData = ref({
    username: '',
    url: '',
    ip: '',
    type: '',
    startTime: '',
    endTime: ''
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

const { pager, getLists, resetParams, resetPage, handleIndex } = usePaging<SystemLogResp>({
    fetchFun: systemLogLists,
    params: formData.value
})

getLists()
</script>

<style lang="scss" scoped></style>
