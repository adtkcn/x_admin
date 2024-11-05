import { createApp } from 'vue'
import App from './App.vue'
import install from './install'
import './permission'
import 'element-plus/dist/index.css'
import './styles/index.scss'

import 'virtual:svg-icons-register'

import ElementPlus from 'element-plus'

import VForm3 from 'vform3-builds' //引入VForm3库

import { XErr, XErrWeb } from '../../x_err_sdk/web/index'
// { XErr, XErrWeb }
new XErr(
    {
        Dns: 'http://localhost:5174/api',
        Pid: 'e19e3be20de94f49b68fafb4c30668bc',
        Uid: '10'
    },
    new XErrWeb({
        onloadTimeOut: 3000
    })
)

const app = createApp(App)
app.use(install)
app.use(ElementPlus)
app.use(VForm3)

app.mount('#app')
