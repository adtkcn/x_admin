import { VxeTable, VxeColumn } from 'vxe-table'
import 'vxe-table/lib/style.css'
import type { App } from 'vue'

export default (app: App<Element>) => {
    app.use(VxeTable)
    app.use(VxeColumn)
}
