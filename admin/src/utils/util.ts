// import { isObject } from '@vue/shared'
import { cloneDeep } from 'lodash-es'
// import { md5 } from 'js-md5'
import MD5 from 'crypto-js/md5'
/**
 * 密码加密
 * @param {String} password 密码
 * @param {String} salt 后置盐
 * @returns {String} 加密后的密码
 */
export const encryptPassword = (password: string, salt = 'asdjioewurtjfgiopu'): string => {
    return MD5(MD5(password).toString() + salt).toString()
}
/**
 * @description 添加单位
 * @param {String | Number} value 值 100
 * @param {String} unit 单位 px em rem
 */
export const addUnit = (value: string | number, unit = 'px') => {
    return !Object.is(Number(value), NaN) ? `${value}${unit}` : value
}

/**
 * @description 是否为空
 * @param {any} value
 * @return {Boolean}
 */
export const isEmpty = (value: any) => {
    return value === '' || value === null || value === undefined
}
/**
 * 判读是否为对象
 */
export const isObject = (val: any): boolean => {
    return val !== null && typeof val === 'object'
}
/**
 * @description 树转数组，队列实现广度优先遍历
 * @param {Array} data  数据
 * @param {Object} props `{ children: 'children' }`
 */

export const treeToArray = (data: any[], props = { children: 'children' }) => {
    data = cloneDeep(data)
    const { children } = props
    const newData = []
    const queue: any[] = []
    data.forEach((child: any) => queue.push(child))
    while (queue.length) {
        const item: any = queue.shift()
        if (item[children]) {
            item[children].forEach((child: any) => queue.push(child))
            delete item[children]
        }
        newData.push(item)
    }
    return newData
}

/**
 * @description 数组转
 * @param {Array} data  数据
 * @param {Object} props `{ parent: 'pid', children: 'children' }`
 */

export function arrayToTree(arr, parentId = '') {
    const tree = []
    for (const item of arr) {
        if (item.pid == parentId) {
            const children = arrayToTree(arr, item.id)
            if (children.length > 0) {
                item.children = children
            }
            tree.push(item)
        }
    }
    return tree
}

/**
 * @description 获取正确的路经
 * @param {String} path  数据
 */
export function getNormalPath(path: string) {
    if (path.length === 0 || !path || path == 'undefined') {
        return path
    }
    const newPath = path.replace('//', '/')
    const length = newPath.length
    if (newPath[length - 1] === '/') {
        return newPath.slice(0, length - 1)
    }
    return newPath
}

/**
 * @description对象格式化为Query语法
 * @param { Object } params
 * @return {string} Query语法
 */
export function objectToQuery(params: Record<string, any>): string {
    let query = ''
    for (const props of Object.keys(params)) {
        const value = params[props]
        const part = encodeURIComponent(props) + '='
        if (!isEmpty(value)) {
            if (isObject(value)) {
                for (const key of Object.keys(value)) {
                    if (!isEmpty(value[key])) {
                        const params = props + '[' + key + ']'
                        const subPart = encodeURIComponent(params) + '='
                        query += subPart + encodeURIComponent(value[key]) + '&'
                    }
                }
            } else {
                query += part + encodeURIComponent(value) + '&'
            }
        }
    }
    return query.slice(0, -1)
}

/**
 * @description 获取不重复的id
 * @param length { Number } id的长度
 * @return { String } id
 */
export const getNonDuplicateID = (length = 8) => {
    let idStr = Date.now().toString(36)
    idStr += Math.random().toString(36).substring(3, length)
    return idStr
}

/**
 * @description 单词首字母大写
 * @param  { String } str
 * @return { String } id
 */
export const firstToUpperCase = (str = '') => {
    return str.toLowerCase().replace(/( |^)[a-z]/g, ($1) => $1.toUpperCase())
}

/**
 * @description 清空对象空值属性
 */
export const clearEmpty = (obj: Record<string, any>) => {
    for (const key of Object.keys(obj)) {
        if (isEmpty(obj[key])) {
            delete obj[key]
        }
    }
    return obj
}
