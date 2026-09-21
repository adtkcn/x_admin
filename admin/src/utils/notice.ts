/**
 * 通知消息公共逻辑：消息严重等级（与后端 NoticePayload.Type 取值约定一致）
 * 统一维护等级 → 展示样式的映射，避免各组件重复定义。
 */

// 消息严重等级
export type NoticeType = 'success' | 'danger' | 'primary' | 'info' | 'warning'

// 严重等级 → 圆点/标识颜色（element 主题色），未知值兜底 info
const levelColorMap: Record<NoticeType, string> = {
    success: 'var(--el-color-success)',
    danger: 'var(--el-color-danger)',
    primary: 'var(--el-color-primary)',
    info: 'var(--el-color-info)',
    warning: 'var(--el-color-warning)'
}
export function getNoticeLevelColor(type?: string): string {
    return levelColorMap[type as NoticeType] || levelColorMap.info
}

// 严重等级 → ElNotification 类型；其不支持 danger，映射为 error
const notifyTypeMap: Record<NoticeType, 'success' | 'error' | 'primary' | 'info' | 'warning'> = {
    success: 'success',
    danger: 'error',
    primary: 'primary',
    info: 'info',
    warning: 'warning'
}
export function getNoticeNotifyType(type?: string) {
    return notifyTypeMap[type as NoticeType] || 'info'
}
