type NodeShapeStyle = Record<string, unknown>

export function getShapeStyleFunction(style: NodeShapeStyle, properties: Record<string, any>): NodeShapeStyle {
    const { backgroundColor, borderColor, borderWidth } = properties
    const newStyle: NodeShapeStyle = { ...style }
    if (backgroundColor) {
        newStyle.fill = backgroundColor
    }
    if (borderColor) {
        newStyle.stroke = borderColor
    }
    if (borderWidth) {
        newStyle.strokeWidth = borderWidth
    }
    return newStyle
}

export function getTextStyleFunction(style: NodeShapeStyle, properties: Record<string, any>): NodeShapeStyle {
    const { color, fontSize } = properties
    const newStyle: NodeShapeStyle = { ...style }
    if (color) {
        newStyle.color = color
    }
    if (fontSize) {
        newStyle.fontSize = fontSize
    }
    return newStyle
}
