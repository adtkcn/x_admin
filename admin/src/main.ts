import { createApp } from 'vue'
import App from './App.vue'
import install from './install'
import './permission'
import 'element-plus/dist/index.css'
import './styles/index.scss'

import 'virtual:svg-icons-register'

import ElementPlus from 'element-plus'
import { VxeTable, VxeColumn } from 'vxe-table'
import 'vxe-table/lib/style.css'

// import VForm3 from 'vform3-builds' //引入VForm3库

// import { Base, Web } from '../../x_err_sdk/web/index'
// const xErr = new Base(
//     {
//         Dns: `${location.origin}/api`,
//         Pid: 'e19e3be20de94f49b68fafb4c30668bc',
//         Uid: ''
//     },
//     new Web()
// )
// xErr.SetUid(1) //设置用户ID

const app = createApp(App)
app.use(install)
app.use(ElementPlus)
// app.use(VForm3)
app.use(VxeTable)
app.use(VxeColumn)

app.mount('#app')
