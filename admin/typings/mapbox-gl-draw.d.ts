/**
 * @mapbox/mapbox-gl-draw 无内置类型，npm 上也无 @types/mapbox-gl-draw，
 * 这里按 src/api.js 实际暴露的方法做最小声明，仅覆盖本项目用到的部分。
 */
declare module '@mapbox/mapbox-gl-draw' {
    export type DrawMode = 'simple_select' | 'direct_select' | 'draw_polygon' | 'draw_line_string' | 'draw_point'

    export interface MapboxDrawControls {
        point?: boolean
        line_string?: boolean
        polygon?: boolean
        trash?: boolean
        combine_features?: boolean
        uncombine_features?: boolean
    }

    export interface MapboxDrawOptions {
        defaultMode?: DrawMode
        keybindings?: boolean
        touchEnabled?: boolean
        clickBuffer?: number
        touchBuffer?: number
        boxSelect?: boolean
        displayControlsDefault?: boolean
        controls?: MapboxDrawControls
        styles?: unknown[]
        modes?: Record<string, unknown>
        userProperties?: boolean
    }

    export interface DrawFeature {
        type: 'Feature'
        id?: string | number
        properties: Record<string, any>
        geometry: { type: string; coordinates: any }
    }

    export interface DrawFeatureCollection {
        type: 'FeatureCollection'
        features: DrawFeature[]
    }

    export interface DrawLayerStyle {
        id: string
        type: string
        filter?: any
        layout?: Record<string, any>
        paint?: Record<string, any>
        [key: string]: any
    }

    export default class MapboxDraw {
        static constants: {
            types: Record<string, string>
            classes: Record<string, string>
            meta: Record<string, string>
            mode: Record<string, DrawMode>
            [key: string]: any
        }
        static modes: Record<string, unknown>
        static lib: {
            // 图层样式表，可在 addControl 之前修改以适配 maplibre 的 style-spec
            theme: DrawLayerStyle[]
            [key: string]: unknown
        }

        constructor(options?: MapboxDrawOptions)

        add(features: DrawFeature | DrawFeatureCollection | any): string[]
        set(features: DrawFeatureCollection): string[]
        get(id: string): DrawFeature | undefined
        getAll(): DrawFeatureCollection
        getSelected(): DrawFeatureCollection
        getSelectedIds(): string[]
        getFeatureIdsAt(point: [number, number]): string[]
        delete(ids?: string | string[]): string[]
        deleteAll(): this
        changeMode(mode: DrawMode, options?: Record<string, any>): this
        getMode(): { mode: DrawMode }
        trash(): this
        onAdd(map: any): HTMLElement
        onRemove(map: any): void
    }
}
