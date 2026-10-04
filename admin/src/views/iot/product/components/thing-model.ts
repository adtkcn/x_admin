// 物模型可视化编辑器的数据模型与 JSON 互转。
//
// 后端只用物模型做属性白名单（app/iot/pipeline/thingfilter.go 递归收集 identifier），
// 未固定 schema，故这里定一个最扁平的结构（顶层就是属性数组，不再包 properties）：
//     [ { "identifier": "temperature", "name": "温度", "dataType": "float" } ]
// 解析时仍兼容旧写法的多种形态：{properties:[...]}、properties 为对象映射（键名即标识）、
// 阿里云/腾讯云的 dataType 对象形态与 specs 嵌套；未被编辑器识别的字段原样保留并回写。

export type ThingDataType = 'int' | 'float' | 'string' | 'bool' | 'date' | 'array' | 'object'

export const ThingDataTypeOptions: { label: string; value: ThingDataType }[] = [
    { label: '整型 int', value: 'int' },
    { label: '浮点 float', value: 'float' },
    { label: '字符串 string', value: 'string' },
    { label: '布尔 bool', value: 'bool' },
    { label: '时间 date', value: 'date' },
    { label: '数组 array', value: 'array' },
    { label: '对象 object', value: 'object' }
]

export type ThingProperty = {
    identifier: string
    name: string
    dataType: ThingDataType
    unit: string
    // 数值型属性的取值范围，非数值型不使用（null 表示不限）
    min: number | null
    max: number | null
    desc: string
    // 原 JSON 中本编辑器未管理的字段，保存时原样带回，避免手写的扩展信息被吞掉
    extra?: Record<string, any>
}

// 属性标识即设备上报的字段名，白名单按它匹配，故限制为字母开头的标识符
export const IDENTIFIER_PATTERN = /^[a-zA-Z][a-zA-Z0-9_]{0,63}$/

export function isNumericType(dataType: ThingDataType) {
    return dataType === 'int' || dataType === 'float'
}

export function emptyProperty(): ThingProperty {
    return {
        identifier: '',
        name: '',
        dataType: 'int',
        unit: '',
        min: null,
        max: null,
        desc: ''
    }
}

type ParseResult = {
    properties: ThingProperty[]
    // 非空表示结构无法识别：调用方应退回纯 JSON 文本编辑，不要把原文丢弃
    error?: string
}

// pickDataType 兼容 "int" 与 { "type": "int" } / { "value": "int" } 两种写法
function pickDataType(v: any): ThingDataType {
    const raw = (typeof v === 'object' && v !== null ? v.type || v.value : v) || ''
    const hit = ThingDataTypeOptions.find((o) => o.value === String(raw).toLowerCase())
    return hit ? hit.value : 'string'
}

// toProperty 把一条原始属性定义转成编辑器行；已识别字段被消费，其余进 extra
function toProperty(item: any, fallbackId: string): ThingProperty {
    const src = item && typeof item === 'object' ? { ...item } : {}
    const identifier = String(src.identifier ?? src.key ?? src.id ?? fallbackId ?? '')
    const specs = src.specs && typeof src.specs === 'object' ? src.specs : {}
    const num = (v: any) => {
        const n = Number(v)
        return v === null || v === undefined || v === '' || Number.isNaN(n) ? null : n
    }
    const prop: ThingProperty = {
        identifier,
        name: String(src.name ?? src.title ?? identifier),
        dataType: pickDataType(src.dataType ?? src.data_type),
        unit: String(src.unit ?? specs.unit ?? ''),
        min: num(src.min ?? specs.min),
        max: num(src.max ?? specs.max),
        desc: String(src.desc ?? src.description ?? '')
    }
    // 交给编辑器的字段从原始对象里剔除，剩下的才是需要原样保留的扩展
    // rw/accessMode 一并剔除：属性读写对本项目无意义（只做上报白名单），旧数据里的该字段不再回写
    ;[
        'identifier',
        'key',
        'id',
        'name',
        'title',
        'dataType',
        'data_type',
        'rw',
        'accessMode',
        'unit',
        'min',
        'max',
        'desc',
        'description',
        'specs'
    ].forEach((k) => delete src[k])
    // specs 里除 min/max/unit 之外仍有内容（如 enum/step）时整体保留，云端物模型常这么写
    const restSpecs = { ...specs }
    ;['unit', 'min', 'max'].forEach((k) => delete restSpecs[k])
    if (Object.keys(restSpecs).length) src.specs = restSpecs
    prop.extra = Object.keys(src).length ? src : undefined
    return prop
}

// collectProperties 从任意物模型 JSON 里取属性列表，支持数组与对象映射两种承载
function collectProperties(node: any): ThingProperty[] | null {
    if (Array.isArray(node)) {
        return node.map((item, i) => toProperty(item, `property_${i + 1}`))
    }
    if (!node || typeof node !== 'object') return null

    const holder = node.properties ?? node.props ?? node.thingModelProperties
    if (Array.isArray(holder)) {
        return holder.map((item, i) => toProperty(item, `property_${i + 1}`))
    }
    if (holder && typeof holder === 'object') {
        return Object.entries(holder).map(([k, v]) => toProperty(v, k))
    }
    // 没有 properties 包装：{"temperature": {"name":"温度"}} 这种纯对象映射也认下来，
    // 但要求值必须是“属性形状”的对象，否则 {"events":[...]} 之类会被误当成属性
    const entries = Object.entries(node)
    const looksLikeProps =
        entries.length > 0 &&
        entries.every(
            ([, v]) =>
                v !== null &&
                typeof v === 'object' &&
                !Array.isArray(v) &&
                ('identifier' in v || 'name' in v || 'dataType' in v || 'data_type' in v || 'specs' in v)
        )
    return looksLikeProps ? entries.map(([k, v]) => toProperty(v, k)) : null
}

// parseThingModel 物模型 JSON -> 编辑器行；空串视为未配置（合法）
export function parseThingModel(raw: string): ParseResult {
    if (!raw || !raw.trim()) return { properties: [] }
    let node: any
    try {
        node = JSON.parse(raw)
    } catch (e: any) {
        return { properties: [], error: `JSON 解析失败：${e?.message || e}` }
    }
    const properties = collectProperties(node)
    if (!properties) {
        return {
            properties: [],
            error: '未识别到属性定义，请在 properties 中配置，或继续在 JSON 模式下手工编辑'
        }
    }
    return { properties }
}

// serializeThingModel 编辑器行 -> 物模型 JSON（顶层属性数组）；无有效属性时返回空串
// （空串才等于“未物模型化”，存 [] 会让处理层拿到空白名单而丢全部属性）
export function serializeThingModel(properties: ThingProperty[]): string {
    const list = properties
        .map((p) => ({ ...p, identifier: p.identifier.trim() }))
        .filter((p) => p.identifier)
    if (!list.length) return ''
    const body = list.map((p) => {
        const item: Record<string, any> = {
            identifier: p.identifier,
            name: p.name.trim() || p.identifier,
            dataType: p.dataType
        }
        if (p.unit) item.unit = p.unit
        if (isNumericType(p.dataType)) {
            if (p.min !== null) item.min = p.min
            if (p.max !== null) item.max = p.max
        }
        if (p.desc.trim()) item.desc = p.desc.trim()
        // 手工扩展字段先铺，编辑器管理的字段优先级更高
        return { ...(p.extra || {}), ...item }
    })
    return JSON.stringify(body, null, 2)
}

// validateProperties 保存前校验：标识必填、字符合法、不重复
export function validateProperties(properties: ThingProperty[]): string[] {
    const problems: string[] = []
    const seen = new Map<string, number>()
    properties.forEach((p, i) => {
        const id = p.identifier.trim()
        const row = i + 1
        if (!id) {
            problems.push(`第 ${row} 行缺少属性标识`)
            return
        }
        if (!IDENTIFIER_PATTERN.test(id)) {
            problems.push(`第 ${row} 行「${id}」标识不合法：字母开头，仅可含字母/数字/下划线`)
        }
        if (seen.has(id)) {
            problems.push(`第 ${row} 行「${id}」与第 ${seen.get(id)! + 1} 行重复`)
        } else {
            seen.set(id, i)
        }
        if (isNumericType(p.dataType) && p.min !== null && p.max !== null && p.min > p.max) {
            problems.push(`第 ${row} 行「${id}」最小值大于最大值`)
        }
    })
    return problems
}

// countProperties 统计已配置的属性条数（供表单项摘要展示），解析不出来的按 0
export function countProperties(raw?: string): number {
    if (!raw) return 0
    return parseThingModel(raw).properties.length
}
