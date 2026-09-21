import type { App } from 'vue'

// 指令（手动注册，新增指令需在此处引入并注册）
import copy from './directives/copy'
import perms from './directives/perms'

// 插件（手动注册，新增插件需在此处引入并安装）
import elIcon from './plugins/el-icon'
import pinia from './plugins/pinia'
import router from './plugins/router'
import vxeTable from './plugins/vxe-table'
import elementPlus from './plugins/element-plus'

// 安装方法，执行某一类相同操作
export function install(app: App<Element>) {
    // 使用插件
    router(app)
    pinia(app)

    elementPlus(app)
    vxeTable(app)

    elIcon(app)

    // 注册全局指令
    app.directive('copy', copy)
    app.directive('perms', perms)
}
