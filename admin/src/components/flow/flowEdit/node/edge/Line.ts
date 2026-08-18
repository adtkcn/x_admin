import { LineEdge, LineEdgeModel } from '@logicflow/core'

import { getShapeStyleFunction, getTextStyleFunction } from '../getShapeStyleUtil'

// 直线
class CustomLineModel extends LineEdgeModel {
    constructor(data: any, graphModel: any) {
        super(data, graphModel)
        this.strokeWidth = 1
    }
    getTextStyle(): ReturnType<LineEdgeModel['getTextStyle']> {
        const style = super.getTextStyle()
        const properties = this.properties as { color?: string; fontSize?: number }
        return getTextStyleFunction(style, properties) as ReturnType<LineEdgeModel['getTextStyle']>
    }
    getEdgeStyle(): Record<string, any> {
        const attributes = super.getEdgeStyle()
        const properties = this.properties as Record<string, any>
        const style = getShapeStyleFunction(attributes, properties) as Record<string, any>
        style.stroke = 'rgb(24, 125, 255)'
        return { ...style, fill: 'none' }
    }
}

const LineEdgeShape = {
    type: 'pro-line',
    view: LineEdge,
    model: CustomLineModel
}

export default LineEdgeShape
