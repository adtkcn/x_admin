import { PolylineEdge, PolylineEdgeModel } from '@logicflow/core'

import { getShapeStyleFunction, getTextStyleFunction } from '../getShapeStyleUtil'

// 折线
class CustomPolylineModel extends PolylineEdgeModel {
    constructor(data: any, graphModel: any) {
        super(data, graphModel)
        this.strokeWidth = 1
    }
    setAttributes() {
        this.isAnimation = true
    }
    getEdgeAnimationStyle(): Record<string, any> {
        const style = super.getEdgeAnimationStyle()
        style.strokeDasharray = '50 5'
        style.animationDuration = '10s'
        style.stroke = 'rgb(24, 125, 255)'
        return style
    }
    getEdgeStyle(): Record<string, any> {
        const attributes = super.getEdgeStyle()
        const properties = this.properties as Record<string, any>
        const style = getShapeStyleFunction(attributes, properties) as Record<string, any>
        style.stroke = 'rgb(24, 125, 255)'
        return { ...style, fill: 'none' }
    }
    getTextStyle(): ReturnType<PolylineEdgeModel['getTextStyle']> {
        const style = super.getTextStyle()
        const properties = this.properties as { color?: string; fontSize?: number }
        return getTextStyleFunction(style, properties) as ReturnType<PolylineEdgeModel['getTextStyle']>
    }
}

const PolylineEdgeShape = {
    type: 'pro-polyline',
    view: PolylineEdge,
    model: CustomPolylineModel
}

export default PolylineEdgeShape
