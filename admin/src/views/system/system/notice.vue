<!-- 消息通知 -->
<template>
    <div class="notice-page">
        <el-card class="border-none!" shadow="never">
            <div class="flex justify-between items-center">
                <div class="flex items-center gap-3">
                    <el-radio-group v-model="formData.isRead" @change="resetPage">
                        <el-radio-button :value="-1">全部</el-radio-button>
                        <el-radio-button :value="0">未读</el-radio-button>
                        <el-radio-button :value="1">已读</el-radio-button>
                    </el-radio-group>
                    <el-select
                        v-model="formData.type"
                        placeholder="通知类型"
                        clearable
                        style="width: 140px"
                        @change="resetPage"
                    >
                        <el-option label="全部类型" value="" />
                        <el-option label="审批通过" value="flow_pass" />
                        <el-option label="审批驳回" value="flow_back" />
                        <el-option label="新审批任务" value="flow_new" />
                        <el-option label="审批完成" value="flow_finish" />
                        <el-option label="系统公告" value="system" />
                    </el-select>
                </div>
                <div class="flex gap-2">
                    <el-button type="primary" @click="handleReadAll">全部标为已读</el-button>
                    <el-button @click="handleSetting">通知设置</el-button>
                </div>
            </div>
        </el-card>

        <el-card v-loading="pager.loading" class="mt-4 border-none!" shadow="never">
            <div class="notice-list">
                <div
                    v-for="item in pager.lists"
                    :key="item.id"
                    class="notice-item"
                    :class="{ 'is-unread': item.isRead === 0 }"
                    @click="handleClick(item)"
                >
                    <div class="flex items-start gap-3">
                        <div class="notice-icon">
                            <el-tag
                                :type="getTypeColor(item.type)"
                                size="small"
                                effect="dark"
                                round
                            >
                                {{ getTypeLabel(item.type) }}
                            </el-tag>
                        </div>
                        <div class="flex-1 min-w-0">
                            <div class="flex justify-between items-center">
                                <span
                                    class="notice-title"
                                    :class="{ 'font-bold': item.isRead === 0 }"
                                >
                                    {{ item.title }}
                                </span>
                                <span class="notice-time">{{ item.createTime }}</span>
                            </div>
                            <div class="notice-content mt-1">{{ item.content }}</div>
                        </div>
                        <div class="notice-actions flex gap-1">
                            <el-button
                                v-if="item.isRead === 0"
                                type="primary"
                                link
                                size="small"
                                @click.stop="handleRead(item.id)"
                            >
                                标为已读
                            </el-button>
                            <el-button
                                type="danger"
                                link
                                size="small"
                                @click.stop="handleDelete(item.id)"
                            >
                                删除
                            </el-button>
                        </div>
                    </div>
                </div>
                <div v-if="pager.lists.length === 0 && !pager.loading" class="empty-state">
                    <el-empty description="暂无通知" />
                </div>
            </div>
            <div class="flex mt-4 justify-end">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>

        <!-- 通知设置弹窗 -->
        <el-dialog
            v-model="settingVisible"
            title="通知设置"
            width="420px"
            :close-on-click-modal="false"
        >
            <el-form label-width="100px">
                <el-form-item label="站内消息">
                    <el-switch
                        v-model="settingForm.siteEnabled"
                        :active-value="1"
                        :inactive-value="0"
                    />
                </el-form-item>
                <el-form-item label="邮件通知">
                    <el-switch
                        v-model="settingForm.emailEnabled"
                        :active-value="1"
                        :inactive-value="0"
                    />
                </el-form-item>
            </el-form>
            <template #footer>
                <el-button @click="settingVisible = false">取消</el-button>
                <el-button type="primary" @click="handleSaveSetting">保存</el-button>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
    noticeList,
    noticeRead,
    noticeReadAll,
    noticeDel,
    noticeGetSetting,
    noticeSaveSetting,
    type type_notice_list,
    type type_notice_resp
} from '@/api/setting/notice'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'

defineOptions({
    name: 'SystemNotice'
})

const router = useRouter()
const formData = reactive<type_notice_list>({
    type: '',
    isRead: -1
})

const { pager, getLists, resetPage } = usePaging<type_notice_resp>({
    fetchFun: noticeList,
    params: formData
})

// 类型映射
const typeMap: Record<string, { label: string; color: string }> = {
    flow_pass: { label: '审批通过', color: 'success' },
    flow_back: { label: '审批驳回', color: 'danger' },
    flow_new: { label: '新审批', color: 'primary' },
    flow_finish: { label: '审批完成', color: 'info' },
    system: { label: '系统公告', color: 'warning' }
}

function getTypeLabel(type: string) {
    return typeMap[type]?.label || type
}

function getTypeColor(type: string) {
    return typeMap[type]?.color || ''
}

function handleClick(item: type_notice_resp) {
    // 未读的先标已读再跳转
    if (item.isRead === 0) {
        handleRead(item.id)
    }
    if (item.url) {
        router.push(item.url)
    }
}

async function handleRead(id: string) {
    await noticeRead({ id })
    getLists()
}

async function handleReadAll() {
    try {
        await feedback.confirm('确定将全部通知标为已读？')
        await noticeReadAll()
        feedback.msgSuccess('已全部标为已读')
        getLists()
    } catch {}
}

async function handleDelete(id: string) {
    try {
        await feedback.confirm('确定删除该通知？')
        await noticeDel({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch {}
}

// 通知设置
const settingVisible = ref(false)
const settingForm = reactive({
    siteEnabled: 1 as number,
    emailEnabled: 1 as number
})

async function handleSetting() {
    const res = await noticeGetSetting()
    settingForm.siteEnabled = res.siteEnabled
    settingForm.emailEnabled = res.emailEnabled
    settingVisible.value = true
}

async function handleSaveSetting() {
    await noticeSaveSetting({ ...settingForm })
    feedback.msgSuccess('设置已保存')
    settingVisible.value = false
}

onMounted(() => {
    getLists()
})
</script>

<style lang="scss" scoped>
.notice-page {
    .notice-list {
        .notice-item {
            padding: 16px 0;
            border-bottom: 1px solid var(--el-border-color-lighter);
            cursor: pointer;
            transition: background-color 0.2s;

            &:last-child {
                border-bottom: none;
            }

            &:hover {
                background-color: var(--el-fill-color-light);
                margin: 0 -20px;
                padding-left: 20px;
                padding-right: 20px;
                border-radius: 6px;
            }

            &.is-unread {
                .notice-title {
                    &::before {
                        content: '';
                        display: inline-block;
                        width: 8px;
                        height: 8px;
                        background-color: var(--el-color-primary);
                        border-radius: 50%;
                        margin-right: 8px;
                        vertical-align: middle;
                    }
                }
            }

            .notice-title {
                font-size: 15px;
                color: var(--el-text-color-primary);
            }

            .notice-content {
                font-size: 13px;
                color: var(--el-text-color-secondary);
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }

            .notice-time {
                font-size: 12px;
                color: var(--el-text-color-placeholder);
                white-space: nowrap;
                margin-left: 12px;
            }

            .notice-actions {
                opacity: 0;
            }

            &:hover .notice-actions {
                opacity: 1;
            }
        }
    }

    .empty-state {
        padding: 60px 0;
    }
}
</style>
