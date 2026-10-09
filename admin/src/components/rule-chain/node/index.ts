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

// 主色混色工具：向白/黑线性插值，得到浅底填充与深色文字
const WHITE: [number, number, number] = [255, 255, 255]
const BLACK: [number, number, number] = [0, 0, 0]
function mix(hex: string, target: [number, number, number], ratio: number) {
    const h = hex.replace('#', '')
    const full = h.length === 3 ? h.split('').map((c) => c + c).join('') : h
    const n = parseInt(full, 16)
    const m = (a: number, t: number) => Math.round(a + (t - a) * ratio)
    return `rgb(${m((n >> 16) & 255, target[0])}, ${m((n >> 8) & 255, target[1])}, ${m(n & 255, target[2])})`
}

// 浅底卡片风格：淡色填充 + 主色描边，主色作为识别色
function applyColor(style: Record<string, any>, meta: RcNodeMeta) {
    style.fill = mix(meta.color, WHITE, 0.88)
    style.stroke = meta.color
    style.strokeWidth = 1.5
    return style
}

// 文字取略深的主色，保证在浅底上清晰可读
function applyTextColor(style: any, meta: RcNodeMeta) {
    style.color = mix(meta.color, BLACK, 0.2)
    style.fontWeight = 600
    return style
}

function makeRectNode(meta: RcNodeMeta) {
    class Model extends RectNodeModel {
        initNodeData(data: any) {
            super.initNodeData(data)
            this.width = 150
            this.height = 52
            this.radius = 8
        }
        getNodeStyle() {
            const style = super.getNodeStyle()
            return applyColor(style as Record<string, any>, meta)
        }
        getTextStyle() {
            return applyTextColor(super.getTextStyle(), meta)
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
            return applyTextColor(super.getTextStyle(), meta)
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
            return applyTextColor(super.getTextStyle(), meta)
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
