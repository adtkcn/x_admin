import SparkMD5 from 'spark-md5'
import axios from 'axios'

export interface FileUploaderOptions {
    /**
     * 分块大小
     */
    chunkSize?: number
    onSuccess?: (filePath: string) => void
    onError?: (error: Error) => void
    onCalculateMD5Progress?: (percent: number) => void
    onUploadProgress?: (
        chunkIndex: number,
        chunkCount: number,
        chunkLoaded: number,
        chunkTotal: number,
        chunkPercent: number
    ) => void
    onChunkSuccess?: (chunkIndex: number) => void
    onChunkError?: (chunkIndex: number, error: Error) => void
}

export default class FileUploader {
    file: File
    fileName: string
    fileMd5: string
    fileSize: number
    chunkSize: number = 1024 * 1024 // 1MB
    chunkCount: number = 0

    uploading = false

    private abortControllers: AbortController[] = []
    private startChunkIndex = -1

    // 上传完成
    success(filePath: string) {
        this.uploading = false
        this.startChunkIndex = -1
        this.onSuccess(filePath)
    }
    error(error: Error) {
        this.uploading = false
        this.startChunkIndex = -1
        this.onError(error)
    }
    calculateMD5Progress(chunkIndex: number, chunkTotal: number) {
        const chunkPercent = Math.floor((chunkIndex / chunkTotal) * 100)
        this.onCalculateMD5Progress(chunkPercent)
    }
    uploadProgress(chunkIndex: number, chunkLoaded: number, chunkTotal: number) {
        // 计算百分比
        const chunkPercent = Math.floor((chunkLoaded / chunkTotal) * 100)
        this.onUploadProgress(chunkIndex, this.chunkCount, chunkLoaded, chunkTotal, chunkPercent)
    }
    onSuccess: FileUploaderOptions['onSuccess'] = function () {}
    onChunkSuccess: FileUploaderOptions['onChunkSuccess'] = function () {}
    onChunkError: FileUploaderOptions['onChunkError'] = function () {}
    onError: FileUploaderOptions['onError'] = function () {}
    onCalculateMD5Progress: FileUploaderOptions['onCalculateMD5Progress'] = function () {}
    onUploadProgress: FileUploaderOptions['onUploadProgress'] = function () {}

    /**
     * 构造函数
     * @param options
     * @param file 上传的文件
     */
    constructor(options: FileUploaderOptions, file?: File) {
        if (options?.chunkSize) {
            this.chunkSize = options.chunkSize
        }
        if (options?.onSuccess) {
            this.onSuccess = options.onSuccess
        }
        if (options?.onError) {
            this.onError = options.onError
        }
        if (options?.onUploadProgress) {
            this.onUploadProgress = options.onUploadProgress
        }
        if (options?.onChunkSuccess) {
            this.onChunkSuccess = options.onChunkSuccess
        }
        if (options?.onChunkError) {
            this.onChunkError = options.onChunkError
        }
        if (file) {
            this.loadFile(file)
        }
    }
    /**
     * 加载文件
     * @param file
     * @param fileName
     * @returns
     */
    public loadFile(file: File, fileName?: string) {
        if (this.uploading) {
            this.error(new Error('请等待上一个文件上传完成'))
            return
        }

        this.file = file
        if (fileName) {
            this.fileName = fileName
        } else {
            this.fileName = file.name
        }
        this.fileSize = file.size
        this.chunkCount = Math.ceil(this.file.size / this.chunkSize)
    }
    // readerFile(file: File): Promise<ArrayBuffer> {
    //     return new Promise((resolve, reject) => {
    //         if (!file) {
    //             return reject(new Error('读取文件失败'))
    //         }
    //         const reader = new FileReader()
    //         reader.onload = (e) => {
    //             resolve(e.target?.result as ArrayBuffer)
    //         }
    //         reader.onerror = (e) => {
    //             reject(e)
    //         }
    //         reader.readAsArrayBuffer(file)
    //     })
    // }
    getAbortControllerSignal() {
        const controller = new AbortController()
        this.abortControllers.push(controller)
        return controller.signal
    }
    /**
     * 取消上传
     */
    public cancel() {
        this.uploading = false
        this.abortControllers.forEach((controller) => controller.abort())
        this.abortControllers.length = 0 // 清空数组

        this.startChunkIndex = -1
    }
    /**
     * 开始上传
     */
    public async start() {
        try {
            if (!this.file) {
                this.error(new Error('请选择文件后上传'))
                return
            }
            if (this.uploading) {
                this.error(new Error('正在上传中'))
                return
            }
            this.uploading = true

            // const arrayBuffer = await this.readerFile(this.file)
            console.time('SparkMD5')
            const fileMd5 = await this.calculateMD5(this.file)
            console.timeEnd('SparkMD5')
            this.fileMd5 = fileMd5 + '_' + this.fileSize
            const isExistFilePath = await this.checkFileExist()

            if (isExistFilePath) {
                this.success(isExistFilePath)
                return
            }
            const hasChunk = await this.getHasChunk()
            this.startChunkIndex = hasChunk && hasChunk.length ? Math.max(...hasChunk) : -1
            console.log('hasChunk', hasChunk)

            await this.splitChunksAndUpload()
            await this.mergeChunk()
        } catch (error) {
            this.error(error)
        }
    }
    /**
     * 计算文件的MD5值
     * @param file
     * @returns
     */
    async calculateMD5(file: File) {
        return new Promise((resolve, reject) => {
            const spark = new SparkMD5.ArrayBuffer()
            const reader = new FileReader()
            const chunkSize = 10 * 1024 * 1024 // 10MB 分块
            let currentChunk = 0
            const chunks = Math.ceil(file.size / chunkSize)

            reader.onload = function (e) {
                spark.append(e.target.result) // 添加数组缓冲区
                currentChunk++

                if (currentChunk < chunks) {
                    loadNext()
                } else {
                    resolve(spark.end())
                }
            }

            reader.onerror = function () {
                reject('文件读取错误')
            }

            function loadNext() {
                console.log('md5进度', currentChunk, chunks)
                this.calculateMD5Progress(currentChunk, chunks)

                const start = currentChunk * chunkSize
                const end = Math.min(start + chunkSize, file.size)
                reader.readAsArrayBuffer(file.slice(start, end))
            }

            loadNext()
        })
    }
    // 计算文件的MD5值
    // getMd5(arrayBuffer: ArrayBuffer): string {
    //     if (!this.uploading) {
    //         return
    //     }
    //     console.time('SparkMD5')
    //     const spark = new SparkMD5.ArrayBuffer()
    //     spark.append(arrayBuffer)
    //     const hash = spark.end()
    //     console.timeEnd('SparkMD5')
    //     return hash
    // }
    // 检查文件是否存在,可实现秒传
    async checkFileExist(): Promise<string> {
        if (!this.uploading) {
            return
        }
        try {
            /* 检查文件是否存在 */
            const res = await axios.get('/api/admin/common/uploadChunk/CheckFileExist', {
                signal: this.getAbortControllerSignal(),
                params: {
                    fileMd5: this.fileMd5,
                    fileName: this.fileName
                }
            })
            if (res.data?.code === 200) {
                return res.data.data
            }
            throw new Error(res.data.message)
        } catch (error) {
            if (axios.isCancel(error)) {
                return
            }
            throw error
        }
    }
    async getHasChunk(): Promise<number[]> {
        if (!this.uploading) {
            return
        }
        try {
            const hasChunkRes = await axios.get('/api/admin/common/uploadChunk/HasChunk', {
                signal: this.getAbortControllerSignal(),
                params: {
                    fileMd5: this.fileMd5,
                    chunkSize: this.chunkSize,
                    fileName: this.fileName
                }
            })
            console.log('HasChunk', hasChunkRes)

            if (hasChunkRes.data.code === 200) {
                return hasChunkRes.data.data || []
            }
            throw new Error(hasChunkRes.data.message)
        } catch (error) {
            if (axios.isCancel(error)) {
                return
            }
            throw error
        }
    }

    async splitChunksAndUpload() {
        for (let index = this.startChunkIndex + 1; index < this.chunkCount; index++) {
            if (!this.uploading) {
                break
            }
            const chunkStart = index * this.chunkSize
            const chunkEnd = Math.min(chunkStart + this.chunkSize, this.file.size)
            const chunk = this.file.slice(chunkStart, chunkEnd)
            await this.uploadChunk(this.fileMd5, index, chunk)
        }
    }
    async uploadChunk(fileMd5: string, index: number, chunk: Blob) {
        if (!this.uploading) {
            return
        }
        try {
            const formData = new FormData()
            formData.append('fileMd5', fileMd5)
            formData.append('chunk', chunk)
            formData.append('chunkSize', String(this.chunkSize))
            formData.append('index', String(index))

            const result = await axios.post('/api/admin/common/uploadChunk/UploadChunk', formData, {
                signal: this.getAbortControllerSignal(),
                onUploadProgress: (progressEvent) => {
                    this.uploadProgress(index + 1, progressEvent.loaded, progressEvent.total)
                }
            })
            console.log('result', result)

            if (result.data.code === 200) {
                this.onChunkSuccess?.(index + 1)
            } else {
                this.onChunkError?.(index + 1, new Error(result.data.message))
                this.error(new Error(result.data.message))
            }
        } catch (error) {
            if (axios.isCancel(error)) {
                return
            }
            this.onChunkError?.(index + 1, error)
            this.error(error)
        } finally {
            chunk = null
        }
    }
    async mergeChunk() {
        if (!this.uploading) {
            return
        }
        try {
            const res = await axios.post(
                '/api/admin/common/uploadChunk/MergeChunk',
                {
                    fileMd5: this.fileMd5,
                    fileName: this.fileName,
                    chunkCount: this.chunkCount,
                    chunkSize: this.chunkSize
                },
                {
                    signal: this.getAbortControllerSignal()
                }
            )
            if (res.data.code === 200) {
                console.log('合并分片成功')
                this.success(res.data.data)
            } else {
                console.log('MergeChunk', res)
                this.error(new Error('合并分片失败'))
            }
        } catch (error) {
            if (axios.isCancel(error)) {
                return
            }
            console.error(`合并分片失败: ${error}`)
            this.error(error)
        }
    }
}
