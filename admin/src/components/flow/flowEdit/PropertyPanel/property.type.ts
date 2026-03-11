export type PropertiesType = {
    userType?: number | null
    userId?: string | number | null
    deptId?: string | number | null
    postId?: string | number | null
    fieldAuth?: {
        [key: string]: number
    }
    gateway?: {
        id: string
        value: string
        condition: string
    }[]
}
export type NodeType = {
    id?: string
    type?:
        | 'bpmn:startEvent'
        | 'bpmn:userTask'
        | 'bpmn:serviceTask'
        | 'bpmn:exclusiveGateway'
        | 'bpmn:endEvent'
    text?: {
        value?: string
        x?: number
        y?: number
    }
    properties?: PropertiesType
    x?: number
    y?: number
}
// Form读取读取德列表结构
export type FormFieldListType = {
    id: string
    name: string
}

export type FieldListType = {
    id?: string
    name?: string
    auth?: number
}
