import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// ===================== 告警历史（处理层告警状态机自动写入，管理端只读/清理） =====================
// 告警级别：由产品 goja 告警脚本 alert(values) 的数字返回决定（0=不告警/恢复，非 0 即告警）。
// 级别数值本身就是危害升序：故障(-1) < 轻微(1) < 严重(2)，正数 3 及以上预留给未来高级告警。
// 故障特殊：危害档最低（故用冷色/info），但优先于其他告警——设备失能时的阈值告警不可信，以后端为准。
export const SeverityMap: Record<number, string> = {
    '-1': '故障',
    1: '轻微',
    2: '严重'
}
// 遍历一律用这张表，不依赖对象键序（负数键不是整型索引键，不参与数字排序）
export const SeverityLevels: number[] = [-1, 1, 2]
// tag 颜色：危害越高越醒目；未录入的高级别(3+)由调用处兜底成 danger
export const SeverityTagType: Record<number, 'success' | 'warning' | 'danger' | 'info'> = {
    '-1': 'info',
    1: 'warning',
    2: 'danger'
}
// 大屏点位/图例用色：故障取冷色低一档，轻微橙、严重红；更高（3+）兜底成顶格红
export const SeverityColor: Record<number, string> = {
    '-1': '#22d3ee',
    1: '#e6a23c',
    2: '#f56c6c'
}
export type type_iot_alert_list = {
    device_id?: string
    sn?: string
    product_id?: string
    space_id?: string
    status?: number
    severity?: number
    start_time?: string
    end_time?: string
    pageNo?: number
    pageSize?: number
}
export type type_iot_alert = {
    id?: string
    device_id?: string
    sn?: string
    product_id?: string
    product_name?: string
    space_id?: string
    space_name?: string
    status?: number
    severity?: number
    values?: Record<string, any>
    triggered_at?: string
    recovered_at?: string
}
export function alertLists(params: type_iot_alert_list) {
    return request.get<Pages<type_iot_alert>>({ url: '/iot/alert/list', params })
}
export function alertDel(data: { ids: string[] }) {
    return request.post({ url: '/iot/alert/del', data })
}
