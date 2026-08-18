import { h } from '@logicflow/core'
import { RectResizeModel, RectResizeView } from '@logicflow/extension/lib/NodeResize/node/RectResize'

import { getShapeStyleFunction, getTextStyleFunction } from '../getShapeStyleUtil'

class CustomRectRadiusModel extends RectResizeModel {
    constructor(data: any, graphModel: any) {
        super(data, graphModel)
        this.zIndex = 10
    }
    getNodeStyle(): ReturnType<RectResizeModel['getNodeStyle']> {
        const style = super.getNodeStyle()
        const properties = this.properties as { backgroundColor?: string; borderColor?: string; borderWidth?: number }
        return getShapeStyleFunction(style, properties) as ReturnType<RectResizeModel['getNodeStyle']>
    }
    getTextStyle(): ReturnType<RectResizeModel['getTextStyle']> {
        const style = super.getTextStyle()
        const properties = this.properties as { color?: string; fontSize?: number }
        return getTextStyleFunction(style, properties) as ReturnType<RectResizeModel['getTextStyle']>
    }
}

class CustomRectRadiusView extends RectResizeView {
    getShape() {
        const { model } = this.props
        const style = model.getNodeStyle()
        const { x, y, width, height } = model
        return h('rect', {
            ...style,
            x: x - width / 2,
            y: y - height / 2,
            width,
            height,
            rx: 10,
            ry: 10
        })
    }
}

const RectRadiusNodeShape = {
    type: 'rect-radius',
    view: CustomRectRadiusView,
    model: CustomRectRadiusModel
}

export default RectRadiusNodeShape
