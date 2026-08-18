export interface ChunkUploadResult {
    md5: string
    key: string
    fileName: string
    fileSize: number
    fileHashId: string | null
    instant: boolean
    location?: string
}
