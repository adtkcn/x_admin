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
                        :class="{ 'is-unread': item.is_read === 0 }"
                        @click="handleItemClick(item)"
                    >
                        <!-- 等级圆点前置，右侧标题与描述同列对齐 -->
                        <span
                            class="notice-level-dot"
                            :style="{ background: getNoticeLevelColor(item.type) }"
                        ></span>
                        <div class="notice-item-main">
                            <div class="notice-title-text mt-1">
                                <span class="notice-title">{{ item.title }}</span>
                                <span class="notice-time">{{ item.create_time }}</span>
                            </div>

                            <div class="notice-content">{{ item.content }}</div>
                        </div>
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
import { ElNotification } from 'element-plus'
import { useRouter } from 'vue-router'
import {
    noticeList,
    noticeUnreadCount,
    noticeRead,
    noticeReadAll,
    type type_notice_resp
} from '@/api/setting/notice'
import { onWsMessage } from '@/hooks/useGlobalWs'
import { getNoticeLevelColor, getNoticeNotifyType } from '@/utils/notice'
import feedback from '@/utils/feedback'

defineOptions({
    name: 'HeaderNotice'
})

const router = useRouter()
const popoverRef = ref()
const loading = ref(false)
const unreadCount = ref(0)
const recentList = ref<type_notice_resp[]>([])

// 获取未读数（角标，初始化及 WS 推送时刷新）
async function fetchUnreadCount() {
    try {
        unreadCount.value = (await noticeUnreadCount()).count
    } catch (error) {
        console.error('未读数获取失败:', error)
    }
}

// 获取最近通知列表（仅在弹层展开时调用）
async function fetchNoticeList() {
    try {
        const listRes = await noticeList({ pageNo: 1, pageSize: 5, is_read: 0, type: '' })
        recentList.value = listRes.lists
    } catch (error) {
        console.error('通知列表获取失败:', error)
    }
}

// 弹出时加载列表数据，并同步一次未读数
function onPopoverShow() {
    loading.value = true
    Promise.all([fetchNoticeList(), fetchUnreadCount()]).finally(() => {
        loading.value = false
    })
}

// 点击通知项
async function handleItemClick(item: type_notice_resp) {
    try {
        popoverRef.value?.hide()
        if (item.is_read === 0) {
            await noticeRead({ id: item.id })
            unreadCount.value = Math.max(0, unreadCount.value - 1)
        }
        if (item.url) {
            router.push(item.url)
        }
    } catch (error) {
        console.error('通知标记已读失败:', error)
    }
}

// 全部已读
async function handleReadAll() {
    try {
        await noticeReadAll()
        unreadCount.value = 0
        recentList.value.forEach((item) => {
            item.is_read = 1
        })
        feedback.msgSuccess('已全部标为已读')
    } catch (error) {
        console.error('通知全部已读失败:', error)
    }
}

// 跳转通知页
function goNoticePage() {
    popoverRef.value?.hide()
    router.push('/system/notice')
}

// 订阅 WS 通知消息（组件卸载自动取消订阅）；只刷新角标，列表待展开弹层时再拉取
onWsMessage('notice', (msg) => {
    fetchUnreadCount()
    const data = msg.data || {}

    ElNotification({
        title: data.title || '新消息通知',
        message: data.content || '',
        // 按消息严重等级区分弹窗样式
        type: getNoticeNotifyType(data.type),

        duration: 4500,
        onClick: () => {
            if (data.url) {
                router.push(data.url)
            }
        }
    })
})

onMounted(() => {
    fetchUnreadCount()
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
            display: flex;
            align-items: flex-start;
            padding: 10px 8px;
            border-bottom: 1px solid var(--el-border-color-lighter);
            cursor: pointer;
            border-radius: 4px;

            .notice-level-dot {
                flex-shrink: 0;
                width: 8px;
                height: 8px;
                // 下移与标题行垂直居中：mt-1(4px) + 半行高(约10px) - 半圆点(4px)
                margin-top: 10px;
                margin-right: 8px;
                border-radius: 50%;
            }

            .notice-item-main {
                flex: 1;
                min-width: 0;
            }

            &:last-child {
                border-bottom: none;
            }

            &:hover {
                background-color: var(--el-fill-color-light);
            }

            &.is-unread {
                background-color: var(--el-color-primary-light-9);
            }

            .notice-title-text {
                display: flex;
                justify-content: space-between;
                align-items: center;
                font-size: 13px;
                color: var(--el-text-color-primary);
                font-weight: 500;

                .notice-title {
                    flex: 1;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }
            }

            .notice-time {
                font-size: 11px;
                color: var(--el-text-color-placeholder);
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
