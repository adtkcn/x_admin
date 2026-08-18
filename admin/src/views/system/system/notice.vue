<template>
    <div class="notice">
        <el-card class="border-none!" shadow="never">
            <div class="flex items-center justify-between mb-2">
                <div class="flex items-center gap-2">
                    <el-tag type="danger" effect="dark" v-if="unreadCount"
                        >{{ unreadCount }} 条未读</el-tag
                    >
                    <!-- <el-tag type="info" effect="plain">共 {{ total }} 条</el-tag> -->
                </div>
                <div class="flex items-center gap-2">
                    <el-radio-group v-model="activeTab" @change="handleTabChange">
                        <el-radio-button :value="-1">全部</el-radio-button>
                        <el-radio-button :value="0">未读</el-radio-button>
                        <el-radio-button :value="1">已读</el-radio-button>
                    </el-radio-group>
                    <el-button type="primary" :disabled="!unreadCount" @click="handleReadAll">
                        全部已读
                    </el-button>
                    <el-button @click="handleOpenSetting">通知设置</el-button>
                </div>
            </div>

            <vxe-table
                class="mt-2"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :scroll-y="{ enabled: false }"
                :border="'inner'"
                @cell-click="handleCellClick"
            >
                <vxe-column title="状态" width="90" align="center">
                    <template #default="{ row }">
                        <span
                            class="inline-block w-[8px] h-[8px] rounded-full align-middle"
                            :class="
                                row.is_read === 0 ? 'bg-primary' : 'bg-[var(--el-border-color)]'
                            "
                        ></span>
                        <span class="ml-1">{{ row.is_read === 0 ? '未读' : '已读' }}</span>
                    </template>
                </vxe-column>
                <vxe-column title="通知内容" min-width="400">
                    <template #default="{ row }">
                        <div class="font-bold" :class="{ 'text-primary': row.is_read === 0 }">
                            {{ row.title }}
                        </div>
                        <div
                            class="text-[13px] text-[var(--el-text-color-secondary)] truncate mt-1"
                        >
                            {{ row.content }}
                        </div>
                    </template>
                </vxe-column>
                <vxe-column title="通知时间" field="create_time" width="180" />
                <vxe-column title="操作" width="140" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-if="row.is_read === 0"
                            type="primary"
                            link
                            @click.stop="handleRead(row.id)"
                        >
                            标为已读
                        </el-button>
                        <el-button type="danger" link @click.stop="handleDelete(row.id)">
                            删除
                        </el-button>
                    </template>
                </vxe-column>
                <template #empty>
                    <span class="text-[var(--el-text-color-secondary)]">暂无通知</span>
                </template>
            </vxe-table>

            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <!-- 通知设置 -->
        <el-dialog v-model="settingVisible" title="通知设置" width="460px">
            <el-form label-width="100px">
                <el-form-item
                    v-for="item in settingForm.channels"
                    :key="item.key"
                    :label="item.label"
                >
                    <el-switch v-model="item.enabled" :active-value="1" :inactive-value="0" />
                </el-form-item>
            </el-form>
            <template #footer>
                <el-button @click="settingVisible = false">取消</el-button>
                <el-button type="primary" @click="handleSaveSetting">确定</el-button>
            </template>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { computed, ref, reactive } from 'vue'
import {
    noticeList,
    noticeRead,
    noticeReadAll,
    noticeDel,
    noticeUnreadCount,
    noticeGetSetting,
    noticeSaveSetting,
    type type_notice_resp,
    type type_notice_list,
    type type_notice_setting
} from '@/api/setting/notice'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'

defineOptions({
    name: 'systemNotice'
})

const activeTab = ref<number>(-1)
const queryParams = reactive<type_notice_list>({ is_read: -1 })

const { pager, getLists, resetPage } = usePaging<type_notice_resp>({
    fetchFun: noticeList,
    params: queryParams
})

const unreadCount = ref(0)
const total = computed(() => pager.count ?? pager.lists.length)

const getUnreadCount = () => {
    noticeUnreadCount().then((res) => {
        unreadCount.value = res.count ?? 0
    })
}

// 通知设置
const settingVisible = ref(false)
const settingForm = reactive<type_notice_setting>({
    channels: []
})

const handleOpenSetting = () => {
    noticeGetSetting().then((res) => {
        settingForm.channels = res.channels ?? []
        settingVisible.value = true
    })
}

const handleSaveSetting = () => {
    const settings: Record<string, number> = {}
    settingForm.channels.forEach((item) => {
        settings[item.key] = item.enabled
    })
    noticeSaveSetting({ settings }).then(() => {
        feedback.msgSuccess('设置成功')
        settingVisible.value = false
    })
}

const handleTabChange = () => {
    queryParams.is_read = activeTab.value
    resetPage()
}

const handleCellClick = ({ row, column }: { row: type_notice_resp; column: any }) => {
    // 点击非操作列时，未读通知直接标为已读
    if (column?.type !== 'seq' && column?.field !== '操作' && row.is_read === 0) {
        handleRead(row.id, false)
    }
}

const handleRead = (id: string, refresh = true) => {
    noticeRead({ id }).then(() => {
        feedback.msgSuccess('已标为已读')
        getUnreadCount()
        if (refresh) {
            getLists()
        } else {
            const item = pager.lists.find((i) => i.id === id)
            if (item) item.is_read = 1
        }
    })
}

const handleReadAll = async () => {
    try {
        await feedback.confirm('确定全部标为已读？')
        await noticeReadAll()
        feedback.msgSuccess('已全部标为已读')
        getUnreadCount()
        getLists()
    } catch (error) {
        console.error('全部已读失败:', error)
    }
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除该通知？')
        await noticeDel({ id })
        feedback.msgSuccess('删除成功')
        getUnreadCount()
        getLists()
    } catch (error) {
        console.error('通知删除失败:', error)
    }
}

getLists()
getUnreadCount()
</script>

<style lang="scss" scoped>
.notice {
    :deep(.bg-primary) {
        background-color: var(--el-color-primary);
    }
}
</style>
