<template>
    <div class="header-notice">
        <el-popover
            ref="popoverRef"
            placement="bottom-end"
            :width="360"
            trigger="click"
            @show="onPopoverShow"
        >
            <template #reference>
                <div class="notice-btn-wrap">
                    <el-badge :value="unreadCount" :max="99" :hidden="unreadCount === 0">
                        <el-button class="notice-btn" text>
                            <icon name="el-icon-Bell" :size="20" />
                        </el-button>
                    </el-badge>
                </div>
            </template>

            <div class="notice-popover">
                <div class="flex justify-between items-center px-2 py-2 border-b border-br">
                    <span class="font-medium">消息通知</span>
                    <el-button
                        v-if="unreadCount > 0"
                        type="primary"
                        link
                        size="small"
                        @click="handleReadAll"
                    >
                        全部已读
                    </el-button>
                </div>

                <div class="notice-list" v-loading="loading">
                    <div
                        v-for="item in recentList"
                        :key="item.id"
                        class="notice-item"
                        :class="{ 'is-unread': item.isRead === 0 }"
                        @click="handleItemClick(item)"
                    >
                        <div class="notice-header">
                            <el-tag
                                :type="getTypeColor(item.type)"
                                size="small"
                                effect="plain"
                                round
                            >
                                {{ getTypeLabel(item.type) }}
                            </el-tag>
                            <span class="notice-time">{{ item.createTime }}</span>
                        </div>
                        <div class="notice-title mt-1">{{ item.title }}</div>
                        <div class="notice-content">{{ item.content }}</div>
                    </div>

                    <div v-if="recentList.length === 0 && !loading" class="empty-state">
                        <el-empty :image-size="60" description="暂无通知" />
                    </div>
                </div>

                <div class="border-t border-br pt-2 text-center">
                    <el-button type="primary" link size="small" @click="goNoticePage">
                        查看全部通知
                    </el-button>
                </div>
            </div>
        </el-popover>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
    noticeList,
    noticeUnreadCount,
    noticeRead,
    noticeReadAll,
    type type_notice_resp
} from '@/api/setting/notice'
import { onWsMessage } from '@/hooks/useGlobalWs'
import feedback from '@/utils/feedback'

defineOptions({
    name: 'HeaderNotice'
})

const router = useRouter()
const popoverRef = ref()
const loading = ref(false)
const unreadCount = ref(0)
const recentList = ref<type_notice_resp[]>([])

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

// 获取未读数和最近通知
async function fetchNoticeData() {
    try {
        const [countRes, listRes] = await Promise.all([
            noticeUnreadCount(),
            noticeList({ pageNo: 1, pageSize: 5, isRead: -1, type: '' })
        ])
        unreadCount.value = countRes.count
        recentList.value = listRes.lists
    } catch {}
}

// 弹出时加载数据
function onPopoverShow() {
    loading.value = true
    fetchNoticeData().finally(() => {
        loading.value = false
    })
}

// 点击通知项
async function handleItemClick(item: type_notice_resp) {
    popoverRef.value?.hide()
    if (item.isRead === 0) {
        await noticeRead({ id: item.id })
        unreadCount.value = Math.max(0, unreadCount.value - 1)
    }
    if (item.url) {
        router.push(item.url)
    }
}

// 全部已读
async function handleReadAll() {
    await noticeReadAll()
    unreadCount.value = 0
    recentList.value.forEach((item) => {
        item.isRead = 1
    })
    feedback.msgSuccess('已全部标为已读')
}

// 跳转通知页
function goNoticePage() {
    popoverRef.value?.hide()
    router.push('/system/notice')
}

// 订阅 WS 通知消息（组件卸载自动取消订阅）
onWsMessage('notice', () => {
    unreadCount.value++
    fetchNoticeData()
})

onMounted(() => {
    fetchNoticeData()
})
</script>

<style lang="scss" scoped>
.header-notice {
    .notice-btn-wrap {
        display: inline-flex;
        align-items: center;
        justify-content: center;
        width: 36px;
        height: 36px;
    }

    .notice-btn {
        font-size: 18px;
        padding: 4px;
    }
}

.notice-popover {
    .notice-list {
        min-height: 244px;
        max-height: 360px;
        overflow-y: auto;

        .notice-item {
            padding: 10px 8px;
            border-bottom: 1px solid var(--el-border-color-lighter);
            cursor: pointer;
            border-radius: 4px;

            &:last-child {
                border-bottom: none;
            }

            &:hover {
                background-color: var(--el-fill-color-light);
            }

            &.is-unread {
                background-color: var(--el-color-primary-light-9);
            }

            .notice-header {
                display: flex;
                justify-content: space-between;
                align-items: center;
            }

            .notice-time {
                font-size: 11px;
                color: var(--el-text-color-placeholder);
            }

            .notice-title {
                font-size: 13px;
                color: var(--el-text-color-primary);
                font-weight: 500;
            }

            .notice-content {
                font-size: 12px;
                color: var(--el-text-color-secondary);
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                margin-top: 2px;
            }
        }

        .empty-state {
            padding: 30px 0;
        }
    }
}
</style>
