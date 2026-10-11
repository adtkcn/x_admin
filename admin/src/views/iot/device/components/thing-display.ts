// 设备详情侧：把产品物模型属性用于展示（名称 / 单位 / 布尔），
// 供「实时状态 / 历史趋势 / 告警记录」三个组件复用，避免各处重复映射逻辑。
import type { ThingProperty } from '@/views/iot/product/components/thing-model'

export type ThingNameMap = Record<string, ThingProperty>

// 按 identifier 建索引，供 O(1) 取名称 / 单位
export function buildThingMap(props: ThingProperty[]): ThingNameMap {
    const map: ThingNameMap = {}
    for (const p of props) map[p.identifier] = p
    return map
}

// 单个属性值展示：布尔转是 / 否，数值型补单位，空值转 '-'
export function displayValue(val: any, prop?: ThingProperty): string {
    if (val === null || val === undefined || val === '') return '-'
    if (prop?.dataType === 'bool') {
        return val === true || val === 1 || val === '1' || val === 'true' ? '是' : '否'
    }
    if (prop && (prop.dataType === 'int' || prop.dataType === 'float') && prop.unit) {
        const num = Number(val)
        if (Number.isFinite(num)) return `${val} ${prop.unit}`
    }
    return String(val)
}

// 属性快照文本：有物模型时把 identifier 映射成中文名，否则原样拼接
export function snapText(vals?: Record<string, any>, map?: ThingNameMap): string {
    if (!vals) return '-'
    const entries = Object.entries(vals)
    if (!entries.length) return '-'
    return entries.map(([k, v]) => `${map?.[k]?.name || k}=${displayValue(v, map?.[k])}`).join('; ')
}
