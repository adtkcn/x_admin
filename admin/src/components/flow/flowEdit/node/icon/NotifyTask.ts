import IconNode from './IconNode'
import icon from './notifyicon.svg?url'

// 左上角ICON为通知（铃铛）的节点，BPMN 类型使用 bpmn:notifyTask
class NotifyTaskNode extends IconNode.view {
    getImageHref() {
        // 内联 SVG 铃铛图标，避免外链依赖
        const svg = icon
        return svg
    }
}

export default {
    name: '通知',
    icon: icon,
    type: 'bpmn:notifyTask',
    view: NotifyTaskNode,
    model: IconNode.model
}
