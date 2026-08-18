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

const app = createApp(App)
app.use(install)
app.use(ElementPlus)
// app.use(VForm3)
app.use(VxeTable)
app.use(VxeColumn)

app.mount('#app')
