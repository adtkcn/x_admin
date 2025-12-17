//引入 echarts 核心模块，核心模块提供了 echarts 使用必须要的接口。
import * as echarts from 'echarts/core'
import type {
    // 系列类型的定义后缀都为 SeriesOption
    BarSeriesOption,
    LineSeriesOption
} from 'echarts/charts'
import type {
    // 组件类型的定义后缀都为 ComponentOption
    TitleComponentOption,
    TooltipComponentOption,
    GridComponentOption,
    DatasetComponentOption
} from 'echarts/components'

import type { ComposeOption } from 'echarts/core'

// 通过 ComposeOption 来组合出一个只有必须组件和图表的 Option 类型
export type ECOption = ComposeOption<
    | BarSeriesOption
    | LineSeriesOption
    | TitleComponentOption
    | TooltipComponentOption
    | GridComponentOption
    | DatasetComponentOption
>

//引入柱状图图表，图表后缀都为 Chart
import {
    BarChart,
    LineChart,
    PieChart,
    GaugeChart
    // MapChart,//地图
    // PictorialBarChart,//象形柱状图
    // RadarChart,//雷达图
    // ScatterChart,//散点图
    // GaugeChart,//仪表盘
} from 'echarts/charts'
// 引入提示框，标题，直角坐标系，数据集，内置数据转换器组件，组件后缀都为 Component
import {
    TitleComponent,
    TooltipComponent,
    GridComponent,
    // PolarComponent,//极坐标系
    // AriaComponent, //无障碍访问组件
    // ParallelComponent,//平行坐标系
    LegendComponent,
    // RadarComponent,//雷达图
    // ToolboxComponent,
    DataZoomComponent
    // VisualMapComponent,//视觉映射组件
    // TimelineComponent,//时间线组件
    // CalendarComponent,//日历组件
    // GraphicComponent,//图形组件
} from 'echarts/components'

//引入 Canvas 渲染器，注意引入 CanvasRenderer 或者 SVGRenderer 是必须的一步
import { CanvasRenderer } from 'echarts/renderers'
//标签自动布局，全局过渡动画等特性
import { LabelLayout, UniversalTransition } from 'echarts/features'

// 注册必须的组件
echarts.use([
    LegendComponent,
    TitleComponent,
    TooltipComponent,
    GridComponent,
    // PolarComponent,
    // AriaComponent,
    // ParallelComponent,
    BarChart,
    LineChart,
    PieChart,
    GaugeChart,
    // MapChart,
    // RadarChart,
    // PictorialBarChart,
    // RadarComponent,
    // ToolboxComponent,
    DataZoomComponent,
    // VisualMapComponent,
    // TimelineComponent,
    // CalendarComponent,
    // GraphicComponent,
    // ScatterChart,
    CanvasRenderer,
    LabelLayout,
    UniversalTransition
    // GaugeChart
])

export { echarts }
