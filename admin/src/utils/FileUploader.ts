import SparkMD5 from 'spark-md5'
import axios from 'axios'

export interface FileUploaderOptions {
    chunkSize?: number
    fileName?: string
    onSuccess?: (filePath: string) => void
    onError?: (error: Error) => void
}

export default class FileUploader {
    file: File
    fileMd5: string
    fileName: string
    chunkSize: number = 1024 * 1024 // 1MB
    chunkCount: number = 0
    onSuccess: FileUploaderOptions['onSuccess'] = () => {}
    onError: FileUploaderOptions['onError'] = () => {}
    onUploadProgress(chunkIndex, percent) {
        console.log(`当前分片: ${chunkIndex}，进度: ${percent}%`)
    }

    constructor(file: File, options: FileUploaderOptions) {
        this.file = file
        this.fileName = file.name

        if (options?.chunkSize) {
            this.chunkSize = options.chunkSize
        }
        this.chunkCount = Math.ceil(this.file.size / this.chunkSize)
        if (options?.fileName) {
            this.fileName = options.fileName
        }
        if (options?.onSuccess) {
            this.onSuccess = options.onSuccess
        }
        if (options?.onError) {
            this.onError = options.onError
        }
    }

    readerFile(file: File): Promise<ArrayBuffer> {
        return new Promise((resolve, reject) => {
            if (!file) {
                return reject(new Error('文件不存在'))
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
    // 开始/恢复上传
    async start() {
        try {
            const arrayBuffer = await this.readerFile(this.file)
            this.fileMd5 = this.getMd5(arrayBuffer)
            const isExistFilePath = await this.checkFileExist()
            if (isExistFilePath) {
                this.complete(isExistFilePath)
                return
            }
            this.splitChunks()
        } catch (error) {
            this.onError(error)
        }
    }

    // 上传完成
    complete(filePath) {
        /* 合并分片 */
        console.log('complete:', filePath)
        this.onSuccess(filePath)
    }
    // 检查上传状态
    getMd5(arrayBuffer: ArrayBuffer): string {
        const spark = new SparkMD5.ArrayBuffer()
        spark.append(arrayBuffer)
        return spark.end()
    }
    // 检查文件是否存在,可实现秒传
    async checkFileExist(): Promise<string> {
        try {
            /* 检查文件是否存在 */
            const res = await axios.get('/api/admin/common/uploadChunk/CheckFileExist', {
                params: {
                    fileMd5: this.fileMd5,
                    fileName: this.file.name
                }
            })
            console.log('Init', res)

            if (res.data.code === 200 && res.data.data) {
                return res.data.data
            }
            return ''
        } catch (error) {
            return ''
        }
    }
    async splitChunks() {
        for (let index = 0; index < this.chunkCount; index++) {
            const chunkStart = index * this.chunkSize
            const chunkEnd = Math.min(chunkStart + this.chunkSize, this.file.size)
            const chunk = this.file.slice(chunkStart, chunkEnd)
            await this.uploadChunk(this.fileMd5, index, chunk)
        }
        await this.mergeChunk()
    }
    async uploadChunk(fileMd5: string, index: number, chunk: Blob) {
        try {
            const checkResult = await axios.get('/api/admin/common/uploadChunk/CheckChunkExist', {
                params: {
                    index,
                    fileMd5
                }
            })
            console.log('checkResult', checkResult)

            if (checkResult.data.code === 200) {
                console.log(`分片 ${index + 1}/${this.chunkCount} 已存在`)
                return
            } else if (checkResult.data.code === 500) {
                const formData = new FormData()
                formData.append('chunk', chunk)
                formData.append('index', String(index))
                formData.append('fileMd5', fileMd5)
                const result = await axios.post(
                    '/api/admin/common/uploadChunk/UploadChunk',
                    formData,
                    {
                        onUploadProgress: (progressEvent) => {
                            const percentCompleted = Math.round(
                                (progressEvent.loaded * 100) / progressEvent.total
                            )
                            this.onUploadProgress(index, percentCompleted)
                        }
                    }
                )
                console.log('result', result)

                if (result.data.code === 200) {
                    console.log(`分片 ${index + 1}/${this.chunkCount} 上传成功`)
                } else {
                    console.error(`分片 ${index + 1}/${this.chunkCount} 上传失败: ${result}`)
                    // break
                }
            }
        } catch (error) {
            console.error(`分片 ${index + 1}/${this.chunkCount} 上传失败: ${error}`)
            this.onError(error)
        }
    }
    async mergeChunk() {
        try {
            const res = await axios.post('/api/admin/common/uploadChunk/MergeChunk', {
                fileMd5: this.fileMd5,
                fileName: this.file.name,
                chunkCount: this.chunkCount
            })
            if (res.data.code === 200) {
                console.log('合并分片成功')
                this.complete(res.data.data)
            } else {
                console.log('MergeChunk', res)
                this.onError(new Error('合并分片失败'))
            }
        } catch (error) {
            console.error(`合并分片失败: ${error}`)
            this.onError(error)
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
