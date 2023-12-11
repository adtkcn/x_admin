import { PolylineEdge, PolylineEdgeModel } from '@logicflow/core'
import { getShapeStyleFuction, getTextStyleFunction } from '../getShapeStyleUtil'

 // 折线
class Model extends PolylineEdgeModel {
  constructor (data, graphModel) {
    super(data, graphModel)
    this.strokeWidth = 1
  }
  setAttributes() {
    this.isAnimation = true;
  }
  getEdgeAnimationStyle() {
    const style = super.getEdgeAnimationStyle();
    style.strokeDasharray = "50 5";
    style.animationDuration = "10s";
    style.stroke = 'rgb(24, 125, 255)'
    return style;
  }
  getTextStyle () {
    const style = super.getTextStyle()
    return getTextStyleFunction(style, this.properties)
  }

  getEdgeStyle () {
    const attributes = super.getEdgeStyle()
    const properties = this.properties;
    const style = getShapeStyleFuction(attributes, properties)
    style.stroke = 'rgb(24, 125, 255)'
    return { ...style, fill: 'none' }
  }
}
export default {
  type: 'pro-polyline',
  view: PolylineEdge,
  model: Model
}
