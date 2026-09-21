import { h, TextNode, TextNodeModel } from '@logicflow/core'

import { getTextStyleFunction } from '../getShapeStyleUtil'

class CustomTextModel extends TextNodeModel {
    constructor(data: any, graphModel: any) {
        super(data, graphModel)
        this.zIndex = 10
    }
    getTextStyle(): ReturnType<TextNodeModel['getTextStyle']> {
        const style = super.getTextStyle()
        const properties = this.properties as { color?: string; fontSize?: number }
        return getTextStyleFunction(style, properties) as ReturnType<TextNodeModel['getTextStyle']>
    }
}

class CustomTextView extends TextNode {
    getShape() {
        const { model } = this.props
        const style = model.getTextStyle()
        const { x, y } = model
        return h('text', {
            ...style,
            x,
            y
        })
    }
}

const TextNodeShape = {
    type: 'text',
    view: CustomTextView,
    model: CustomTextModel
}

export default TextNodeShape
