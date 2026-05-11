// import CryptoJS from 'crypto-js'
// /**
//  * @word 要加密的内容
//  * @keyWord String  服务器随机返回的关键字
//  *  */
// export function aesEncrypt(word: string, keyWord = 'XwKsGlMcdPMEhR1B') {
//     const key = CryptoJS.enc.Utf8.parse(keyWord)
//     const src = CryptoJS.enc.Utf8.parse(word)
//     const encrypted = CryptoJS.AES.encrypt(src, key, {
//         mode: CryptoJS.mode.ECB,
//         padding: CryptoJS.pad.Pkcs7
//     })
//     return encrypted.toString()
// }

// 引入核心和具体算法，而不是整个库
import CryptoJS from 'crypto-js/core'
import AES from 'crypto-js/aes'
import Utf8 from 'crypto-js/enc-utf8'
// 1. 引入 ECB 模式
import ModeECB from 'crypto-js/mode-ecb'
// 2. 引入 Pkcs7 填充
import PadPkcs7 from 'crypto-js/pad-pkcs7'

export function aesEncrypt(word: string, keyWord = 'XwKsGlMcdPMEhR1B') {
    const key = Utf8.parse(keyWord)
    const src = Utf8.parse(word)
    // 注意：这里直接使用 AES 对象，而不是 CryptoJS.AES
    const encrypted = AES.encrypt(src, key, {
        mode: ModeECB,
        padding: PadPkcs7
    })
    return encrypted.toString()
}
