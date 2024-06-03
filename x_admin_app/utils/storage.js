/**
 * 设置key
 * @param {string} key 
 * @param {*} data 
 */
export function setLocalStorage(key, data) {
	uni.setStorageSync(key, JSON.stringify(data))
}
/**
 * 获取本地的key
 *  @param {string} key
 * */
export function getLocalStorage(key) {
	var val = uni.getStorageSync(key)
	try {
		if (val) {
			val = JSON.parse(val)
		}
	} catch (e) {
		//TODO handle the exception
	}
	return val
}

/**
 * 删除key
 * @param {string} key 
 **/
export function removeLocalStorage(key) {
	return uni.removeStorageSync(key)
}