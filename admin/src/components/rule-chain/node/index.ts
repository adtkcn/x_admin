// 规则链自定义节点注册：按 RC_NODE_META 批量生成不同形状/配色的节点，
// 再加一条折线边 pro-polyline，注册逻辑与审批流 flowEdit 一致。
import {
    RectNode,
    RectNodeModel,
    CircleNode,
    CircleNodeModel,
    DiamondNode,
    DiamondNodeModel,
    PolylineEdge,
    PolylineEdgeModel
} from '@logicflow/core'
import type LogicFlow from '@logicflow/core'
import { RC_NODE_META, type RcNodeMeta } from '../rule.type'

// 读取节点上自定义的 rcColor（注册时按类型注入），决定填充/描边
function applyColor(style: Record<string, any>, meta: RcNodeMeta) {
    style.fill = meta.color
    style.stroke = meta.color
    return style
}

function makeRectNode(meta: RcNodeMeta) {
    class Model extends RectNodeModel {
        initNodeData(data: any) {
            super.initNodeData(data)
            this.width = 150
            this.height = 52
            this.radius = 6
        }
        getNodeStyle() {
            const style = super.getNodeStyle()
            return applyColor(style as Record<string, any>, meta)
        }
        getTextStyle() {
            const style: any = super.getTextStyle()
            style.color = '#fff'
            return style
        }
    }
    return { type: meta.type, view: RectNode, model: Model }
}

function makeCircleNode(meta: RcNodeMeta) {
    class Model extends CircleNodeModel {
        initNodeData(data: any) {
            super.initNodeData(data)
            this.r = 38
        }
        getNodeStyle() {
            const style = super.getNodeStyle()
            return applyColor(style as Record<string, any>, meta)
        }
        getTextStyle() {
            const style: any = super.getTextStyle()
            style.color = '#fff'
            return style
        }
    }
    return { type: meta.type, view: CircleNode, model: Model }
}

function makeDiamondNode(meta: RcNodeMeta) {
    class Model extends DiamondNodeModel {
        initNodeData(data: any) {
            super.initNodeData(data)
            this.rx = 60
            this.ry = 42
        }
        getNodeStyle() {
            const style = super.getNodeStyle()
            return applyColor(style as Record<string, any>, meta)
        }
        getTextStyle() {
            const style: any = super.getTextStyle()
            style.color = '#fff'
            return style
        }
    }
    return { type: meta.type, view: DiamondNode, model: Model }
}

function makeNodeConfig(meta: RcNodeMeta) {
    switch (meta.shape) {
        case 'circle':
            return makeCircleNode(meta)
        case 'diamond':
            return makeDiamondNode(meta)
        default:
            return makeRectNode(meta)
    }
}

// 折线边：统一蓝色，样式与审批流一致
class CustomPolylineModel extends PolylineEdgeModel {
    constructor(data: any, graphModel: any) {
        super(data, graphModel)
        this.strokeWidth = 1.5
    }
    getEdgeStyle() {
        const style: any = super.getEdgeStyle()
        style.stroke = '#187bffa6'
        style.fill = 'none'
        return style
    }
    getTextStyle() {
        const style: any = super.getTextStyle()
        style.fontSize = 11
        return style
    }
}

const ProPolyline = { type: 'pro-polyline', view: PolylineEdge, model: CustomPolylineModel }

export const registerRuleChainElement = (lf: LogicFlow) => {
    RC_NODE_META.forEach((meta) => {
        lf.register(makeNodeConfig(meta) as any)
    })
    lf.register(ProPolyline as any)
}
