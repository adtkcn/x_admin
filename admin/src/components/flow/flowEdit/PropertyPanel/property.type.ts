export type PropertiesType = {
    userType?: number
    userId?: string | number
    deptId?: string | number
    postId?: string | number
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
    name?: string
    type?: string
    field?: {
        id: string

        options?: {
            name: string
            label: string
            defaultValue: ''
        }
        type?: string //'textarea'
    }
}

export type FieldListType = {
    id?: string
    label?: string
    auth?: number
}
