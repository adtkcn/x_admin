import { DiamondResizeModel, DiamondResizeView } from '@logicflow/extension/lib/NodeResize/node/DiamondResize'

import { getShapeStyleFunction, getTextStyleFunction } from '../getShapeStyleUtil'

// 菱形
/**
 * model控制初始化的值
 */
class DiamondModel extends DiamondResizeModel {
    initNodeData(data: any) {
        super.initNodeData(data)
        this.rx = 35
        this.ry = 35
    }
    getNodeStyle(): ReturnType<DiamondResizeModel['getNodeStyle']> {
        const style = super.getNodeStyle()
        const properties = this.properties as { backgroundColor?: string; borderColor?: string; borderWidth?: number }
        return getShapeStyleFunction(style, properties) as ReturnType<DiamondResizeModel['getNodeStyle']>
    }
    getTextStyle(): ReturnType<DiamondResizeModel['getTextStyle']> {
        const style = super.getTextStyle()
        const properties = this.properties as { color?: string; fontSize?: number }
        return getTextStyleFunction(style, properties) as ReturnType<DiamondResizeModel['getTextStyle']>
    }
    setToBottom() {
        this.zIndex = 0
    }
}

const DiamondNodeShape = {
    type: 'pro-diamond',
    view: DiamondResizeView,
    model: DiamondModel
}

export default DiamondNodeShape
