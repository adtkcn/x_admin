// ===================== 可视化规则链：节点类型 / 属性命名空间契约 =====================
// 与后端 iot_schema/rule_chain_schema.go 的 RcNode* / RcProps* 常量逐字段对齐：
//   - 节点 type 前缀 rule:  = 后端注册名 = 执行器分发表 key
//   - 节点属性命名空间 key（snake_case）= 后端 RcProps* = properties 收纳 key
// 新增一种节点：在此加一条 RC_NODE_META + 一个属性 struct + 一个子面板，后端加一个 Handler 即可。

// LogicFlow 节点类型（rule: 前缀，与后端执行器分发表一致）
export type RcNodeType =
    | 'rule:trigger'
    | 'rule:decode'
    | 'rule:filter'
    | 'rule:parallel'
    | 'rule:bool'
    | 'rule:alarm'
    | 'rule:state'
    | 'rule:history'
    | 'rule:notify'

// 节点属性命名空间 key（= 后端 RcProps* 常量）
export type RcNodeKey =
    | 'rule_trigger'
    | 'rule_decode'
    | 'rule_filter'
    | 'rule_parallel'
    | 'rule_bool'
    | 'rule_alarm'
    | 'rule_state'
    | 'rule_history'
    | 'rule_notify'

// 图形种类：输入用圆形，并行/布尔用菱形，其余矩形
export type RcShape = 'rect' | 'circle' | 'diamond'

// 各节点类型私有属性（字段全部 snake_case，与后端 JSON tag 完全一致）
// 输入节点无需配置：后端 rule_trigger.msg_types 缺省=放行全部（当前也只有属性上报会进链）
export type RcScriptProps = { script: string }
export type RcAlarmProps = { script: string; duration: number; recover: number }
export type RcNotifyProps = {
    service_type: 'site' | 'email' | 'webhook' | ''
    service_content: string
    receiver_id: string[]
    email_to: string[]
    webhook_url: string
}

// 节点属性映射表
export type RcNodePropsMap = {
    rule_trigger: Record<string, never>
    rule_decode: RcScriptProps
    rule_filter: RcScriptProps
    rule_parallel: Record<string, never>
    rule_bool: RcScriptProps
    rule_alarm: RcAlarmProps
    rule_state: Record<string, never>
    rule_history: Record<string, never>
    rule_notify: RcNotifyProps
}

// LogicFlow 节点 properties：以命名空间 key 收纳当前类型那一份
export type RcPropertiesType = Partial<RcNodePropsMap>

export type RcNodeTypeData = {
    id?: string
    type?: RcNodeType
    text?: { value?: string; x?: number; y?: number }
    properties?: RcPropertiesType
    x?: number
    y?: number
}

// ===================== 节点元数据（侧边栏 + 注册 + 面板共用） =====================
export type RcNodeMeta = {
    type: RcNodeType
    key: RcNodeKey
    label: string
    shape: RcShape
    color: string // 主色（填充）
    desc: string
}

export const RC_NODE_META: RcNodeMeta[] = [
    { type: 'rule:trigger', key: 'rule_trigger', label: '输入', shape: 'circle', color: '#67c23a', desc: '固定起点(不可删除)：属性上报由此进入' },
    { type: 'rule:decode', key: 'rule_decode', label: '解码', shape: 'rect', color: '#409eff', desc: 'goja run(values) 产出属性' },
    { type: 'rule:filter', key: 'rule_filter', label: '数据过滤', shape: 'rect', color: '#13c2c2', desc: 'goja run(values) 布尔放行，false 丢弃整条报文' },
    { type: 'rule:parallel', key: 'rule_parallel', label: '并行', shape: 'diamond', color: '#e6a23c', desc: '所有子路径同时分流（唯一允许多个子连接的节点）' },
    { type: 'rule:bool', key: 'rule_bool', label: '布尔', shape: 'diamond', color: '#722ed1', desc: '脚本返回 true 继续 / false 仅忽略本分支' },
    { type: 'rule:alarm', key: 'rule_alarm', label: '告警', shape: 'rect', color: '#f56c6c', desc: 'goja run+持续/恢复状态机' },
    { type: 'rule:state', key: 'rule_state', label: '状态快照', shape: 'rect', color: '#eb2f96', desc: '写 Redis 最新属性，供管理端/大屏即时读' },
    { type: 'rule:history', key: 'rule_history', label: '历史存储', shape: 'rect', color: '#faad14', desc: '写 MySQL 时序历史表' },
    { type: 'rule:notify', key: 'rule_notify', label: '通知', shape: 'rect', color: '#a0d911', desc: '站内/邮件/webhook' }
]

const META_BY_TYPE: Record<string, RcNodeMeta> = RC_NODE_META.reduce(
    (acc, m) => {
        acc[m.type] = m
        return acc
    },
    {} as Record<string, RcNodeMeta>
)

export const rcMetaOf = (type?: string): RcNodeMeta | undefined =>
    type ? META_BY_TYPE[type] : undefined

// 由节点类型取得属性命名空间 key
export const rcNodeKeyOf = (type?: string): RcNodeKey | undefined =>
    rcMetaOf(type)?.key

// 各节点类型私有属性的缺省值工厂，保证子面板拿到的字段一定存在
export const rcDefaultProps = (key: RcNodeKey): any => {
    switch (key) {
        case 'rule_trigger':
            return {}
        case 'rule_decode':
        case 'rule_filter':
            return { script: '' }
        case 'rule_parallel':
            return {}
        case 'rule_bool':
            return { script: '' }
        case 'rule_alarm':
            return { script: '', duration: 0, recover: 0 }
        case 'rule_notify':
            return {
                service_type: '',
                service_content: '',
                receiver_id: [],
                email_to: [],
                webhook_url: ''
            }
        default:
            return {}
    }
}

// 各脚本节点「一键填入模板」的代码骨架（函数入口统一为 run(values, sn)，与后端契约一致）
export const RC_SCRIPT_TEMPLATES: Record<string, string> = {
    rule_decode: `function run(values, sn) {
  // values: 设备初始上报值，sn: 设备编号
  // 基于上报值转换/补全并返回属性对象；留空则透传上报的 values
  return {
    temp: 0,
    hum: 0
  }
}`,
    rule_filter: `function run(values, sn) {
  // values: 解码后的属性对象，返回 true 放行 / false 丢弃整条报文
  return true
}`,
    rule_bool: `function run(values, sn) {
  // values: 属性对象，返回 true 继续走后续节点 / false 仅忽略本分支
  return values.temp > 30
}`,
    rule_alarm: `function run(values, sn) {
  // values: 属性对象，只返数字：0=不告警/恢复，非 0 即级别；危害升序 -1=故障、1=轻微、2=严重（3+ 预留）
  // 故障危害最轻但优先于其他告警（命中即以它为准）；返其它类型按不告警处理
  if (values.temp > 80) return 2
  return 0
}`
}

// 各脚本节点的说明文案（面板顶部提示，支持 <code> 富文本）
export const RC_SCRIPT_HINTS: Partial<Record<RcNodeKey, string>> = {
    rule_decode:
        '实现 <code>run(values, sn)</code>：values 为设备初始上报值，返回属性对象；留空则透传上报的 values。',
    rule_filter:
        '实现 <code>run(values, sn)</code>：返回布尔，true 放行 / false 丢弃整条报文；留空默认放行。',
    rule_bool:
        '实现 <code>run(values, sn)</code>：返回 true 继续走后续节点 / false 仅忽略本分支（并行的其他分支不受影响）；留空默认继续。'
}

// 边无配置：除并行节点外每个节点只允许一条出边，由设计器与后端保存时共同约束

// ===================== 执行树：图 → 树（后端只消费此树，graph_data 仅供绘图） =====================
// 与后端 iot_schema.RuleTreeNode 对齐：label 取节点文字，properties 仅保留命名空间属性（剔除 width/height）。
export type RcTreeNode = {
    id: string
    type: RcNodeType
    label: string
    properties: RcPropertiesType
    children: RcTreeNode[]
}

// 规模上限与后端 iot_service.treeMaxDepth / treeMaxNodes 一致（前端先拦截，后端再防御）
const RC_TREE_MAX_DEPTH = 64
const RC_TREE_MAX_NODES = 4096

/**
 * 以「输入」为根把 LogicFlow 图展开成执行树，复刻后端旧的图展开语义：
 * 扇入节点按每条路径各留一份副本、沿路径防环、限深度与节点总数。
 * 非法（缺根 / 成环 / 悬空引用 / 超限）时 throw { message }，由调用方提示。
 */
export function buildRcTree(graph: Record<string, any>): RcTreeNode {
    const nodes = (graph?.nodes as RcNodeTypeData[]) || []
    const edges = (graph?.edges as any[]) || []
    const nodeMap = new Map<string, RcNodeTypeData>()
    let rootId = ''
    for (const n of nodes) {
        if (!n.id) continue
        nodeMap.set(n.id, n)
        if (rcNodeKeyOf(n.type) === 'rule_trigger' && !rootId) rootId = n.id
    }
    if (!rootId) throw { message: '规则链-必须有且仅有一个「输入」入口节点' }

    // 邻接表：源节点 → 目标列表（保持 edges 原序，与旧后端展开一致）
    const outMap = new Map<string, string[]>()
    for (const e of edges) {
        if (!e?.sourceNodeId || !e?.targetNodeId) continue
        const arr = outMap.get(e.sourceNodeId) || []
        arr.push(e.targetNodeId)
        outMap.set(e.sourceNodeId, arr)
    }

    let budget = RC_TREE_MAX_NODES
    const onPath = new Set<string>() // 仅标记「根→当前节点」路径，返回时回溯剔除，故扇入会展开成多副本

    const dfs = (id: string, depth: number): RcTreeNode => {
        if (budget-- <= 0) throw { message: '规则链-节点展开超限（疑似报文放大）' }
        if (depth > RC_TREE_MAX_DEPTH) throw { message: '规则链-链深度超限' }
        const node = nodeMap.get(id)
        if (!node) throw { message: `规则链-引用了不存在的节点 ${id}` }
        if (onPath.has(id)) throw { message: '规则链-存在环' }
        onPath.add(id)
        try {
            const children = (outMap.get(id) || []).map((t) => dfs(t, depth + 1))
            // 剔除 LogicFlow 节点模型的绘图尺寸（width/height），只保留命名空间属性
            const { width, height, ...props } = (node.properties as Record<string, any>) || {}
            return {
                id,
                type: node.type as RcNodeType,
                label: typeof node.text === 'string' ? node.text : node.text?.value ?? '',
                properties: props as RcPropertiesType,
                children
            }
        } finally {
            onPath.delete(id)
        }
    }
    return dfs(rootId, 1)
}
