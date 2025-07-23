import SparkMD5 from 'spark-md5'
import axios from 'axios'

export interface FileUploaderOptions {
    chunkSize?: number
    onSuccess?: (filePath: string) => void
    onError?: (error: Error) => void
}

export default class FileUploader {
    file: File
    fileMd5: string
    fileName: string
    fileSize: number
    chunkSize: number = 1024 * 1024 // 1MB
    chunkCount: number = 0

    private uploading = false
    private startChunkIndex = -1

    // 上传完成
    success(filePath: string) {
        this.uploading = false
        this.startChunkIndex = -1
        this.onSuccess(filePath)
    }
    onSuccess: FileUploaderOptions['onSuccess'] = (filePath) => {
        console.log('上传完成', filePath)
    }
    error(error: Error) {
        this.uploading = false
        this.startChunkIndex = -1
        this.onError(error)
    }
    onError: FileUploaderOptions['onError'] = (error: Error) => {
        console.log('error', error)
    }
    onUploadProgress(chunkIndex: number, chunkLoaded: number, chunkTotal: number) {
        console.log(
            `当前分片: ${chunkIndex}/${this.chunkCount},分片进度${chunkLoaded}/${chunkTotal}}`
        )
    }

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
        if (file) {
            this.loadFile(file)
        }
    }
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
    readerFile(file: File): Promise<ArrayBuffer> {
        return new Promise((resolve, reject) => {
            if (!file) {
                return reject(new Error('读取文件失败'))
            }
            const reader = new FileReader()
            reader.onload = (e) => {
                resolve(e.target?.result as ArrayBuffer)
            }
            reader.onerror = (e) => {
                reject(e)
            }
            reader.readAsArrayBuffer(file)
        })
    }
    // 开始上传
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

            const arrayBuffer = await this.readerFile(this.file)
            const fileMd5 = this.getMd5(arrayBuffer)
            this.fileMd5 = fileMd5 + '_' + this.fileSize
            const isExistFilePath = await this.checkFileExist()

            if (isExistFilePath) {
                this.success(isExistFilePath)
                return
            }
            const hasChunk = await this.getHasChunk()
            this.startChunkIndex = hasChunk && hasChunk.length ? Math.max(...hasChunk) : -1
            console.log('hasChunk', hasChunk)

            await this.splitChunks()
            await this.mergeChunk()
        } catch (error) {
            this.error(error)
        }
    }

    // 检查上传状态
    getMd5(arrayBuffer: ArrayBuffer): string {
        console.time('SparkMD5')
        const spark = new SparkMD5.ArrayBuffer()
        spark.append(arrayBuffer)
        const hash = spark.end()
        console.timeEnd('SparkMD5')
        return hash
    }
    // 检查文件是否存在,可实现秒传
    async checkFileExist(): Promise<string> {
        /* 检查文件是否存在 */
        const res = await axios.get('/api/admin/common/uploadChunk/CheckFileExist', {
            params: {
                fileMd5: this.fileMd5,
                fileName: this.fileName
            }
        })
        if (res.data.code === 200) {
            return res.data.data
        }
        throw new Error(res.data.message)
    }
    async getHasChunk(): Promise<number[]> {
        const hasChunkRes = await axios.get('/api/admin/common/uploadChunk/HasChunk', {
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
    }

    async splitChunks() {
        for (let index = this.startChunkIndex + 1; index < this.chunkCount; index++) {
            const chunkStart = index * this.chunkSize
            const chunkEnd = Math.min(chunkStart + this.chunkSize, this.file.size)
            const chunk = this.file.slice(chunkStart, chunkEnd)
            await this.uploadChunk(this.fileMd5, index, chunk)
        }
    }
    async uploadChunk(fileMd5: string, index: number, chunk: Blob) {
        try {
            const formData = new FormData()
            formData.append('fileMd5', fileMd5)
            formData.append('chunk', chunk)
            formData.append('chunkSize', String(this.chunkSize))
            formData.append('index', String(index))

            const result = await axios.post('/api/admin/common/uploadChunk/UploadChunk', formData, {
                onUploadProgress: (progressEvent) => {
                    // const percentCompleted = (
                    //     ((this.chunkSize * index + progressEvent.loaded) * 100) /
                    //     this.fileSize
                    // ).toFixed(3)
                    // const loaded = this.chunkSize * index + progressEvent.loaded //TODO progressEvent.loaded体积比文件大，不能直接相加

                    this.onUploadProgress(index, progressEvent.loaded, progressEvent.total)
                }
            })
            chunk = null
            console.log('result', result)

            if (result.data.code === 200) {
                console.log(`分片 ${index + 1}/${this.chunkCount} 上传成功`)
            } else {
                console.error(`分片 ${index + 1}/${this.chunkCount} 上传失败: ${result}`)
                // break
            }
        } catch (error) {
            chunk = null
            console.error(`分片 ${index + 1}/${this.chunkCount} 上传失败: ${error}`)
            this.error(error)
        }
    }
    async mergeChunk() {
        try {
            const res = await axios.post('/api/admin/common/uploadChunk/MergeChunk', {
                fileMd5: this.fileMd5,
                fileName: this.fileName,
                chunkCount: this.chunkCount,
                chunkSize: this.chunkSize
            })
            if (res.data.code === 200) {
                console.log('合并分片成功')
                this.success(res.data.data)
            } else {
                console.log('MergeChunk', res)
                this.error(new Error('合并分片失败'))
            }
        } catch (error) {
            console.error(`合并分片失败: ${error}`)
            this.error(error)
        }
    }
}

/**
 * const uploader = new FileUploader(file, {
  endpoint: "/api/upload",
  concurrency: 4,
  onProgress: (percent, chunkIndex) => {
    console.log(`进度: ${percent}%，当前分片: ${chunkIndex}`);
  },
  onComplete: (fileUrl) => {
    console.log("文件地址:", fileUrl);
  }
});

 */
