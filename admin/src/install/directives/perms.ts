/**
 * perm 操作权限处理,数组同时满足才会显示
 * 指令用法：
 *  <el-button v-perms="['admin:auth:menu:edit']">编辑</el-button>
 */

import useUserStore from '@/stores/modules/user'
export default {
    mounted: (el: HTMLElement, binding: any) => {
        const { value } = binding
        const userStore = useUserStore()
        const permissions = userStore.perms
        const all_permission = '*' // 具有所有权限
        if (Array.isArray(value)) {
            if (value.length > 0) {
                // const hasPermission = permissions.some((key: string) => {
                //     return all_permission == key || value.includes(key)
                // })
                // debugger
                let hasPermission = false
                // 查找是不是有所有权限
                for (let i = 0; i < permissions.length; i++) {
                    if (all_permission == permissions[i]) {
                        hasPermission = true
                        break
                    }
                }
                // 没有全部权限则匹配按钮权限，需满足所有条件
                if (!hasPermission) {
                    let flag = true
                    for (let i = 0; i < value.length; i++) {
                        const item = value[i]
                        const find = permissions.includes(item)
                        if (!find) {
                            //没有权限
                            flag = false
                            break
                        }
                    }
                    hasPermission = flag
                }

                if (!hasPermission) {
                    el.parentNode && el.parentNode.removeChild(el)
                }
            }
        } else {
            throw new Error(`like v-perms="['auth:menu:edit']"`)
        }
    }
}
