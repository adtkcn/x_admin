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
 * @param {unknown} value
 * @return {Boolean}
 */
export const isEmpty = (value: unknown): boolean => {
    return value === '' || value === null || value === undefined
}
/**
 * 判读是否为对象
 */
export const isObject = (val: unknown): boolean => {
    return val !== null && typeof val === 'object'
}
/**
 * @description 树转数组，队列实现广度优先遍历
 * @param {Array} data  数据
 * @param {Object} props `{ children: 'children' }`
 */
export const treeToArray = <T extends Record<string, any>>(
    data: T[],
    props = { children: 'children' }
): T[] => {
    data = cloneDeep(data)
    const children = props.children
    const newData: T[] = []
    const queue: T[] = []
    data.forEach((child) => queue.push(child))
    while (queue.length) {
        const item = queue.shift()!
        if (item[children]) {
            ;(item[children] as T[]).forEach((child) => queue.push(child))
            delete item[children]
        }
        newData.push(item)
    }
    return newData
}

/**
 * @description 数组转树
 * @param {Array} arr  数据
 * @param {String} parentId 父节点ID
 */

export function arrayToTree<
    T extends { id: string | number; pid: string | number; children?: T[] }
>(arr: T[], parentId: string | number = ''): T[] {
    const tree: T[] = []
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
                    const val = value[key]
                    if (!isEmpty(val)) {
                        const paramsStr = props + '[' + key + ']'
                        const subPart = encodeURIComponent(paramsStr) + '='
                        query += subPart + encodeURIComponent(String(val)) + '&'
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
export const clearEmpty = (obj?: Record<string, any>) => {
    if (!obj || typeof obj !== 'object') {
        return {}
    }
    for (const key of Object.keys(obj)) {
        if (isEmpty(obj[key])) {
            delete obj[key]
        }
    }
    return obj
}

/**
 * @description 格式化时间戳
 * @param {Number} timestamp 时间戳（秒或毫秒）
 * @param {String} format 格式化模板
 * @return {String} 格式化后的时间字符串
 */
export const formatTime = (
    timestamp: number | string | null | undefined,
    format = 'YYYY-MM-DD HH:mm:ss'
): string => {
    if (!timestamp) return '-'
    const ts = typeof timestamp === 'string' ? parseInt(timestamp, 10) : timestamp
    // 判断是秒还是毫秒
    const date = new Date(ts < 1e12 ? ts * 1000 : ts)
    if (isNaN(date.getTime())) return '-'

    const pad = (n: number) => n.toString().padStart(2, '0')
    const year = date.getFullYear()
    const month = pad(date.getMonth() + 1)
    const day = pad(date.getDate())
    const hours = pad(date.getHours())
    const minutes = pad(date.getMinutes())
    const seconds = pad(date.getSeconds())

    return format
        .replace('YYYY', year.toString())
        .replace('MM', month)
        .replace('DD', day)
        .replace('HH', hours)
        .replace('mm', minutes)
        .replace('ss', seconds)
        .replace('hh', hours)
}
