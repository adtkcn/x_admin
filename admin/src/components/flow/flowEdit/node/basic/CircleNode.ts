import { EllipseResizeModel, EllipseResizeView } from '@logicflow/extension/lib/NodeResize/node/EllipseResize'

import { getShapeStyleFunction, getTextStyleFunction } from '../getShapeStyleUtil'

// 圆形
class CircleNewModel extends EllipseResizeModel {
    initNodeData(data: any) {
        super.initNodeData(data)
        this.rx = 35
        this.ry = 35
    }
    setToBottom() {
        this.zIndex = 0
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

const CircleNodeShape = {
    type: 'pro-circle',
    view: EllipseResizeView,
    model: CircleNewModel
}

export default CircleNodeShape
