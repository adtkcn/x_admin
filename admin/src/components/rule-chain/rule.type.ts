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
    { type: 'rule:decode', key: 'rule_decode', label: '解码', shape: 'rect', color: '#409eff', desc: 'goja decode(values) 产出属性' },
    { type: 'rule:filter', key: 'rule_filter', label: '数据过滤', shape: 'rect', color: '#13c2c2', desc: 'goja filter(values) 布尔放行，false 丢弃整条报文' },
    { type: 'rule:parallel', key: 'rule_parallel', label: '并行', shape: 'diamond', color: '#e6a23c', desc: '所有子路径同时分流（唯一允许多个子连接的节点）' },
    { type: 'rule:bool', key: 'rule_bool', label: '布尔', shape: 'diamond', color: '#722ed1', desc: '脚本返回 true 继续 / false 仅忽略本分支' },
    { type: 'rule:alarm', key: 'rule_alarm', label: '告警', shape: 'rect', color: '#f56c6c', desc: 'goja alert+持续/恢复状态机' },
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

// 各脚本节点「一键填入模板」的代码骨架（与后端函数签名契约一致）
export const RC_SCRIPT_TEMPLATES: Record<string, string> = {
    rule_decode: `function decode(values, sn) {
  // values: 设备初始上报值，sn: 设备编号
  // 基于上报值转换/补全并返回属性对象；留空则透传上报的 values
  return {
    temp: 0,
    hum: 0
  }
}`,
    rule_filter: `function filter(values, sn) {
  // values: 解码后的属性对象，返回 true 放行 / false 丢弃整条报文
  return true
}`,
    rule_bool: `function filter(values, sn) {
  // values: 属性对象，返回 true 继续走后续节点 / false 仅忽略本分支
  return values.temp > 30
}`,
    rule_alarm: `function alert(values, sn) {
  // values: 属性对象，只返数字：0=不告警/恢复，非 0 即级别；危害升序 -1=故障、1=轻微、2=严重（3+ 预留）
  // 故障危害最轻但优先于其他告警（命中即以它为准）；返其它类型按不告警处理
  if (values.temp > 80) return 2
  return 0
}`
}

// 各脚本节点的说明文案（面板顶部提示，支持 <code> 富文本）
export const RC_SCRIPT_HINTS: Partial<Record<RcNodeKey, string>> = {
    rule_decode:
        '实现 <code>decode(values, sn)</code>：values 为设备初始上报值，返回属性对象；留空则透传上报的 values。',
    rule_filter:
        '实现 <code>filter(values, sn)</code>：返回布尔，true 放行 / false 丢弃整条报文；留空默认放行。',
    rule_bool:
        '实现 <code>filter(values, sn)</code>：返回 true 继续走后续节点 / false 仅忽略本分支（并行的其他分支不受影响）；留空默认继续。'
}

// 边无配置：除并行节点外每个节点只允许一条出边，由设计器与后端保存时共同约束
