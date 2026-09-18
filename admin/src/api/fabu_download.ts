import axios from 'axios'
import { RequestCodeEnum } from '@/enums/requestEnums'

export interface type_fabu_download_resp {
    app_id: string
    version_id: string
    name: string
    platform: string
    bundle_id: string
    icon: string
    version: string
    version_code: number
    size: number
    download_url: string
    install_url: string
    download_times: number
    has_version: boolean
}

// 后端统一响应信封
interface ApiEnvelope<T> {
    code: number
    data: T
    message: string
}

// 公开接口走 /api/web 前缀（非后台 /api/admin），无需 token，故用独立 axios 调用
export async function fabuAppDownloadInfo(shortUrl: string): Promise<type_fabu_download_resp> {
    const res = await axios.get<ApiEnvelope<type_fabu_download_resp>>(
        `/api/web/fabu/app/${encodeURIComponent(shortUrl)}`
    )
    const body = res.data
    if (body?.code === RequestCodeEnum.SUCCESS) {
        return body.data
    }
    throw new Error(body?.message || '获取应用信息失败')
}
