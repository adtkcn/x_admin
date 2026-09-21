import { BezierEdge, BezierEdgeModel } from '@logicflow/core'

import { getShapeStyleFunction, getTextStyleFunction } from '../getShapeStyleUtil'

// 贝塞尔曲线
class CustomBezierModel extends BezierEdgeModel {
    constructor(data: any, graphModel: any) {
        super(data, graphModel)
        this.strokeWidth = 1
    }
    setAttributes() {
        super.setAttributes()
    }
    getTextStyle(): ReturnType<BezierEdgeModel['getTextStyle']> {
        const style = super.getTextStyle()
        const properties = this.properties as { color?: string; fontSize?: number }
        return getTextStyleFunction(style, properties) as ReturnType<BezierEdgeModel['getTextStyle']>
    }
    getEdgeStyle(): Record<string, any> {
        const attributes = super.getEdgeStyle()
        const properties = this.properties as Record<string, any>
        const style = getShapeStyleFunction(attributes, properties) as Record<string, any>
        style.stroke = 'rgb(24, 125, 255)'
        return { ...style, fill: 'none' }
    }
}

const BezierEdgeShape = {
    type: 'pro-bezier',
    view: BezierEdge,
    model: CustomBezierModel
}

export default BezierEdgeShape
