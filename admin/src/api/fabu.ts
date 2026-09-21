import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// ===================== 应用 App =====================
export type type_fabu_app_list = {
    keyword?: string
    pageNo?: number
    pageSize?: number
}
export type type_fabu_app_resp = {
    id: string
    name: string
    platform: string
    bundle_id: string
    bundle_name: string
    short_url: string
    icon: string
    download_times: number
}
export function fabuAppLists(params: type_fabu_app_list) {
    return request.get<Pages<type_fabu_app_resp>>({ url: '/fabu/app/list', params })
}
export function fabuAppDel(data: { id: string }) {
    return request.post({ url: '/fabu/app/del', data })
}

// ===================== 版本 Version =====================
export type type_fabu_version_list = {
    app_id: string
    pageNo?: number
    pageSize?: number
}
export type type_fabu_version_resp = {
    id: string
    app_id: string
    version: string
    version_code: number
    size: number
    md5: string
    download_url: string
    install_url: string
    released: boolean
    update_mode: number
    gray: boolean
    download_times: number
    create_time: string
}
export function fabuVersionLists(params: type_fabu_version_list) {
    return request.get<Pages<type_fabu_version_resp>>({ url: '/fabu/version/list', params })
}
export function fabuVersionUpload(data: { file_hash_id: string; file_name: string }) {
    return request.post({ url: '/fabu/version/upload', data })
}
export function fabuVersionRelease(data: { app_id: string; id: string }) {
    return request.post({ url: '/fabu/version/release', data })
}
export function fabuVersionCancel(data: { app_id: string; id: string }) {
    return request.post({ url: '/fabu/version/cancel', data })
}
export function fabuVersionGray(data: { app_id: string; id: string; gray: boolean }) {
    return request.post({ url: '/fabu/version/gray', data })
}
export function fabuVersionUpdateMode(data: { app_id: string; id: string; update_mode: number }) {
    return request.post({ url: '/fabu/version/updateMode', data })
}
export function fabuVersionDel(data: { app_id: string; id: string }) {
    return request.post({ url: '/fabu/version/del', data })
}

// ===================== 热更新包 Wgt（版本子表） =====================
export type type_fabu_wgt_list = {
    version_id: string
    pageNo?: number
    pageSize?: number
}
export type type_fabu_wgt_resp = {
    id: string
    app_id: string
    version_id: string
    version: string
    version_code: number
    download_url: string
    md5: string
    size: number
    released: boolean
    create_time: string
}
export function fabuWgtLists(params: type_fabu_wgt_list) {
    return request.get<Pages<type_fabu_wgt_resp>>({ url: '/fabu/wgt/list', params })
}
export function fabuWgtUpload(data: {
    version_id: string
    file_hash_id: string
    file_name: string
}) {
    return request.post({ url: '/fabu/wgt/upload', data })
}
export function fabuWgtRelease(data: { id: string; released: boolean }) {
    return request.post({ url: '/fabu/wgt/release', data })
}
export function fabuWgtDel(data: { id: string }) {
    return request.post({ url: '/fabu/wgt/del', data })
}
