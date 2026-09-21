export const imageExt = ['png', 'jpg', 'jpeg', 'gif', 'ico', 'bmp', 'webp', 'avif']
export const videoExt = ['mp4', 'avi', 'mov', 'wmv', 'flv', 'rmvb']
export const audioExt = ['mp3', 'wav', 'aac', 'flac']
export const officeExt = ['doc', 'docx', 'xls', 'xlsx', 'ppt', 'pptx', 'pdf', 'txt']
export const fileExt = ['zip', '7z', 'rar']
export const All_EXT = [...imageExt, ...videoExt, ...audioExt, ...fileExt]

// 获取文件类型
export function GetFileType(url: string) {
    // 移除query
    url = url.split('?')[0]
    const ext = url.split('.').pop()?.toLowerCase()
    if (!ext) {
        return 'file'
    }
    if (officeExt.includes(ext)) {
        return 'office'
    }
    if (audioExt.includes(ext)) {
        return 'audio'
    }
    if (imageExt.includes(ext)) {
        return 'image'
    }
    if (videoExt.includes(ext)) {
        return 'video'
    }
    return 'file'
}
// 获取文件类型上传的accept
function getAccept(ext: string[]) {
    return ext
        .map((item) => {
            return `.${item}`
        })
        .join(',')
}
export const FileTabsMap = [
    {
        name: '全部',
        fileType: 'all',
        accept: getAccept(All_EXT),
        ext: []
    },
    {
        name: '图片',
        fileType: 'image',
        accept: getAccept(imageExt),
        ext: imageExt
    },
    {
        name: '视频',
        fileType: 'video',
        accept: getAccept(videoExt),
        ext: videoExt
    },
    {
        name: '音频',
        fileType: 'audio',
        accept: getAccept(audioExt),
        ext: audioExt
    },
    {
        name: '文档',
        fileType: 'office',
        accept: getAccept(officeExt),
        ext: officeExt
    },
    {
        name: '文件',
        fileType: 'file',
        accept: getAccept(fileExt),
        ext: fileExt
    }
]
export const FileExt = {
    all: All_EXT,
    image: imageExt,
    video: videoExt,
    audio: audioExt,
    office: officeExt,
    file: fileExt
}
export const FileAccept = {
    all: getAccept(All_EXT),
    image: getAccept(imageExt),
    video: getAccept(videoExt),
    audio: getAccept(audioExt),
    office: getAccept(officeExt),
    file: getAccept(fileExt)
}
