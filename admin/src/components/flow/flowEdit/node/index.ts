// 基础图形

import CircleNode from './basic/CircleNode'
import RectNode from './basic/RectNode'
import RectRadiusNode from './basic/RectRadiusNode'
import EllipseNode from './basic/EllipseNode'
import TextNode from './basic/TextNode'
import DiamondNode from './basic/DiamondNode'

// image绘制左上角icon节点
// 注册边
import type LogicFlow from '@logicflow/core'

import Ployline from './edge/Polyline'
import Line from './edge/Line'
import Bezier from './edge/Bezier'
import NotifyTask from './icon/NotifyTask'

type RegisterElement = LogicFlow.RegisterConfig & {
    name?: string
    icon?: any
}

export const List: RegisterElement[] = [
    // CircleNode,
    // RectNode,
    // RectRadiusNode,
    // EllipseNode,
    // TextNode,
    // DiamondNode
]

export const registerCustomElement = (lf: LogicFlow) => {
    // 注册基础图形
    // lf.register(CircleNode);
    // lf.register(RectNode);
    // lf.register(RectRadiusNode);
    // lf.register(EllipseNode);
    // lf.register(DiamondNode);
    // lf.register(TextNode);

    List.forEach((item) => {
        lf.register(item as LogicFlow.RegisterConfig)
    })
    // // 注册边
    lf.register(Ployline as LogicFlow.RegisterConfig)
    lf.register(Line as LogicFlow.RegisterConfig)
    lf.register(Bezier as LogicFlow.RegisterConfig)
    // 注册通知节点（原系统任务节点）
    lf.register(NotifyTask as LogicFlow.RegisterConfig)
}
