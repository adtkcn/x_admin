// import env from "../utils/env";


/**
 * 生成随机数
 * @param {number} min 最小值
 * @param {number} max 最大值
 * */
export function random(min, max) {
	return Math.floor(Math.random() * (max - min)) + min;
}
/**
 * 手机号脱敏
 * @param {string} phone
 * @returns
 */
export function hidePhone(phone) {
	if (phone && phone.length == 11) {
		return phone.substring(0, 3) + "****" + phone.substring(7);
	} else {
		return phone;
	}
}

/**
 * toast提示
 * @param {string} title
 */
export function toast(title) {
	if (title) {
		uni.showToast({
			icon: "none",
			duration: 2000,
			title: title,
		});
	}
}

/**
 * 扫码
 *  @param {boolean} onlyFromCamera 是否只从相机扫码
 */
export function scanCode(onlyFromCamera = true) {
	return new Promise((resolve, reject) => {
		uni.scanCode({
			onlyFromCamera: onlyFromCamera,
			success: (res) => {
				resolve(res.result);
			},
			fail: () => {
				resolve('');
			},
		});
	});
}

export function queryToObj(path) {
	const res = {};
	const search = path.substr(path.indexOf("?") + 1);
	search.split("&").forEach((paramStr) => {
		const arr = paramStr.split("=");
		const key = arr[0];
		const value = arr[1];
		res[key] = value;
	});
	return res;
}

/**
 * 拼接路径
 * @param {string} path
 * @returns
 */
// export function filePath(path) {
//   if (path && path.indexOf("http") != 0) {
//     return env.fileUrl + path;
//   } else if (path) {
//     return path;
//   } else {
//     return "";
//   }
// }

/**
 * 生成参数
 * @param {object} obj
 * @returns
 */
export function generateParams(obj) {
	var keys = Object.keys(obj);
	keys.sort();

	var vals = [];
	keys.forEach((key) => {
		vals.push(`${key}=${obj[key]}`);
	});

	return vals.join("&");
}

/**
 * 跳转页面
 * @param {string} path 路径
 * @param {object} query 参数
 * @returns
 */
export function toPath(path, query = {}) {
	if (!path) {
		toast("没有指定路径");
		return;
	}

	var url = generateParams(query) ? path + "?" + generateParams(query) : path;
	uni.navigateTo({
		url: url,
	});
}

/**
 * alert 提示
 * @export
 * @param {string} content 内容
 * @param {string} title 标题
 * @param {object} options 配置项
 *
 */
export function alert(content, title = "提示", options = {}) {
	return new Promise((resolve, reject) => {
		uni.showModal({
			title: title,
			content: content,
			showCancel: options.showCancel || false,
			...options,
			success: function(res) {
				if (res.confirm) {
					resolve(res.content);
				} else if (res.cancel) {
					reject();
				}
			},
		});
	});
}

/**
 * 浅拷贝一个对象，去掉''、null、undefined字段
 * @param {object} obj
 */
export function clearObjEmpty(obj) {
	var newObj = {};
	for (let key in obj) {
		if (obj[key] !== "" && obj[key] !== null && obj[key] !== undefined) {
			newObj[key] = obj[key];
		}
	}
	return newObj;
}

/**
 * 转换毫秒数为天时分
 * @param {*} inputSeconds
 * @returns
 */
export function convertMillisecondsToTime(inputSeconds) {
	if (!inputSeconds) return;
	const days = Math.floor(inputSeconds / (60 * 60 * 24));
	const hours = Math.floor((inputSeconds % (60 * 60 * 24)) / (60 * 60));
	const minutes = Math.floor((inputSeconds % (60 * 60)) / 60);
	return `${days}天${hours}时${minutes}分`;
}
/**
 * 防抖函数
 * @param {function} fn
 * @param {number} delay 延迟时间
 */
export function debounce(fn, delay) {
	var timer;
	return function() {
		var context = this;
		var args = arguments;
		clearTimeout(timer);
		timer = setTimeout(function() {
			fn.apply(context, args);
		}, delay);
	};
}