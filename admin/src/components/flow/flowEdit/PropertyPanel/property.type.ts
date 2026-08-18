// ===================== 节点类型 =====================
// LogicFlow / BPMN 标准节点类型，由 LogicFlow 注册时决定，不做改动
export type BpmnNodeType =
    | 'bpmn:startEvent'
    | 'bpmn:userTask'
    | 'bpmn:notifyTask'
    | 'bpmn:exclusiveGateway'
    | 'bpmn:endEvent'

// 节点属性命名空间 key（snake_case 短名），与后端 flow_schema.NodeProps* 常量一一对应。
// 前端 LogicFlow 节点 properties 以本 key 收纳各节点类型的私有数据：
//   properties = { "user_task": { "user_type": 3, "user_id": "1", ... } }
export type NodeKey =
    'start_event' | 'user_task' | 'notify_task' | 'exclusive_gateway' | 'end_event'

// ===================== 各节点类型私有属性 =====================
// 字段全部 snake_case，与后端 JSON tag、下发态 props 完全一致，不再有命名转换
export type FieldAuthMap = { [key: string]: number }

export type GatewayCondition = {
    id: string // 表单项 id
    condition: string // 判断符：== != >= <= include
    value: string // 比较值
}

export type StartEventProps = {
    field_auth: FieldAuthMap
}

export type UserTaskProps = {
    user_type: number | undefined // 1指定部门、岗位 2用户部门负责人 3指定审批人
    user_id: string
    dept_id: string
    post_id: string
    field_auth: FieldAuthMap
}

export type NotifyTaskProps = {
    service_type: 'site' | 'email' | 'webhook' | '' // site=站内消息 email=邮件 webhook=回调
    service_content: string // 消息内容

    // 站内消息
    receiver_id: string[] // 接收人 admin_id 列表，为空则默认通知申请人

    // 邮件：收件邮箱地址数组（可直接填写，也可从用户下拉追加其邮箱）
    email_to: string[]

    // webhook
    webhook_url: string // 回调地址
}

export type GatewayProps = {
    gateway: GatewayCondition[]
}

export type EndEventProps = Record<string, never>

// 节点类型 -> 私有属性 映射表
export type NodePropsMap = {
    start_event: StartEventProps
    user_task: UserTaskProps
    notify_task: NotifyTaskProps
    exclusive_gateway: GatewayProps
    end_event: EndEventProps
}

// LogicFlow 节点 properties：以节点 key 收纳各类型私有数据（仅存在当前类型的那一份）
export type PropertiesType = Partial<NodePropsMap>

// ===================== 节点基础结构 =====================
export type NodeType = {
    id?: string
    type?: BpmnNodeType
    text?: {
        value?: string
        x?: number
        y?: number
    }
    properties?: PropertiesType
    x?: number
    y?: number
}

// ===================== 表单字段结构 =====================
// Form 读取的字段列表结构
export type FormFieldListType = {
    id: string
    name: string
}

export type FieldListType = {
    id: string
    name: string
    auth: number
}

// ===================== 类型 / key 映射与缺省值 =====================
// node.type(bpmn:xxx) -> 属性命名空间 key(短名)
export const NODE_KEY_MAP: Record<BpmnNodeType, NodeKey> = {
    'bpmn:startEvent': 'start_event',
    'bpmn:userTask': 'user_task',
    'bpmn:notifyTask': 'notify_task',
    'bpmn:exclusiveGateway': 'exclusive_gateway',
    'bpmn:endEvent': 'end_event'
}

// 由节点类型取得属性命名空间 key
export const nodeKeyOf = (type?: string): NodeKey | undefined => NODE_KEY_MAP[type as BpmnNodeType]

// 各节点类型私有属性的缺省值工厂，保证子面板拿到的字段一定存在
export const defaultProps = (key: NodeKey): any => {
    switch (key) {
        case 'start_event':
            return { field_auth: {} }
        case 'user_task':
            return {
                user_type: undefined,
                user_id: '',
                dept_id: '',
                post_id: '',
                field_auth: {}
            }
        case 'notify_task':
            return {
                service_type: '',
                service_content: '',
                receiver_id: [],
                email_to: [],
                webhook_url: ''
            }
        case 'exclusive_gateway':
            return { gateway: [] }
        default:
            return {}
    }
}
