import { h } from '@logicflow/core'
import { EllipseResizeModel, EllipseResizeView } from '@logicflow/extension/lib/NodeResize/node/EllipseResize'

import { getShapeStyleFunction, getTextStyleFunction } from '../getShapeStyleUtil'

class CustomEllipseModel extends EllipseResizeModel {
    constructor(data: any, graphModel: any) {
        super(data, graphModel)
        this.zIndex = 10
    }
    getNodeStyle(): ReturnType<EllipseResizeModel['getNodeStyle']> {
        const style = super.getNodeStyle()
        const properties = this.properties as { backgroundColor?: string; borderColor?: string; borderWidth?: number }
        return getShapeStyleFunction(style, properties) as ReturnType<EllipseResizeModel['getNodeStyle']>
    }
    getTextStyle(): ReturnType<EllipseResizeModel['getTextStyle']> {
        const style = super.getTextStyle()
        const properties = this.properties as { color?: string; fontSize?: number }
        return getTextStyleFunction(style, properties) as ReturnType<EllipseResizeModel['getTextStyle']>
    }
}

class CustomEllipseView extends EllipseResizeView {
    getShape() {
        const { model } = this.props
        const style = model.getNodeStyle()
        const { x, y, rx, ry } = model
        return h('ellipse', {
            ...style,
            cx: x,
            cy: y,
            rx,
            ry
        })
    }
}

const EllipseNodeShape = {
    type: 'ellipse',
    view: CustomEllipseView,
    model: CustomEllipseModel
}

export default EllipseNodeShape
