/**
 * @description 将文件字节数格式化为可读大小（B/KB/MB/GB/TB）
 * @param bytes 字节数
 * @param decimals 非 B 单位的小数位数，默认 1
 */
export function formatSize(bytes: number, decimals = 1): string {
    if (!bytes || bytes < 0) return '0 B'
    const units = ['B', 'KB', 'MB', 'GB', 'TB']
    const i = Math.min(Math.floor(Math.log(bytes) / Math.log(1024)), units.length - 1)
    return `${(bytes / Math.pow(1024, i)).toFixed(i === 0 ? 0 : decimals)} ${units[i]}`
}

/**
 * @description
 * @param file
 */
export function streamFileDownload(file: any, fileName = '文件名称.zip') {
    const blob = new Blob([file], { type: 'application/octet-stream;charset=UTF-8' })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.style.display = 'none'
    link.href = url
    link.setAttribute('download', fileName)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link) // 下载完成移除元素
    window.URL.revokeObjectURL(url)
}
